package pnp

import (
	_ "embed"
)

type (
	// Game represents a Platforms and Programmers Game
	// where a band of player will attempt to take on production
	Game struct {
		Players        []Player
		Prod           ProductionState
		BandName       string
		Coins          int
		CurrentPlayer  int
		ProductManager string
	}

	Outcome string

	BasePlayer interface {
		Options(g *Game) []Option
		AsciiArt() string
	}

	CanDie interface {
		Alive() bool
	}

	// Player represents a P&P player
	Player interface {
		BasePlayer
		CanDie
	}

	// Engine represents the game's user interface rendering engine
	Engine interface {
		Welcome(cb func(bandName string))
		Start()
		GameOver()
		GameWon()
		RenderGame(g *Game)
		SelectOption(g *Game, player Player, cb func(selected Option))
		RenderOutcome(outcome Outcome, cb func())
		PizzaDelivery(cb func())
	}
)

func NopImortal(player BasePlayer) Player {
	return nopImortal{BasePlayer: player}
}

type nopImortal struct {
	BasePlayer
}

func (player nopImortal) Alive() bool {
	if mortal, ok := player.BasePlayer.(CanDie); ok {
		return mortal.Alive()
	}
	return true
}

func (player nopImortal) isMinion() bool {
	if mortal, ok := player.BasePlayer.(IsMinion); ok {
		return mortal.isMinion()
	}
	return true
}

// New returns a new P&P game
func New(players ...Player) *Game {
	// plugin the minion

	g := Game{Players: append(players, NopImortal(NewMinion("Bob"))), Prod: NewProduction(), Coins: 10}
	return &g
}

// Run starts a new game
func (g *Game) Run(e Engine) {
	g.Welcome(e, func() {
		g.MainLoop(e)
	})
	e.Start()
}
func (g *Game) Welcome(e Engine, fn func()) {
	e.Welcome(func(bandName string) {
		g.BandName = bandName
		fn()
	})
}

func allPlayersAreMinions(players []Player) bool {
	for _, p := range players {
		if !p.Alive() {
			continue
		}
		if m, ok := p.(IsMinion); ok {
			if !m.isMinion() {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func allPlayersAreDead(players []Player) bool {
	for _, p := range players {
		if p.Alive() {
			return false
		}
	}
	return true
}

// MainLoop kicks off the next players round
func (g *Game) MainLoop(e Engine) {
	if allPlayersAreMinions(g.Players) {
		e.GameOver()
		return
	}

	if allPlayersAreDead(g.Players) {
		e.GameOver()
		return
	}

	for !g.Players[g.CurrentPlayer].Alive() {
		g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
	}
	// check if all players are dead and end the game (if they are).
	// check if the current player is dead and skip it.

	e.RenderGame(g)
	e.SelectOption(g, g.Players[g.CurrentPlayer], func(selected Option) {
		outcome := selected.Selected()
		e.RenderOutcome(outcome, func() {
			g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			g.MainLoop(e)
		})
	})
}

type Option struct {
	OnSelect    func() Outcome
	Description string
}

func (o Option) String() string {
	return o.Description
}

func (o Option) Selected() Outcome {
	return o.OnSelect()
}
