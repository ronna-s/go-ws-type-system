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

	Option interface {
		// String returns a description of the option (e.g. "Drink coffee and ponder about life").
		String() string
		// Selected notifies the option that it was selected, and returns a description of the outcome (e.g. "You feel more awake, but PRODUCTION is unhappy").
		Selected() Outcome
	}

	// Player represents a P&P player
	Player interface {
		Options(g *Game) []Option
		AsciiArt() string
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

// New returns a new P&P game
func New(players ...Player) *Game {
	g := Game{Players: players, Prod: NewProduction(), Coins: 10}
	return &g
}

type Livable interface {
	Alive() bool
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

func alive(p Player) bool {
	if livable, ok := p.(Livable); ok {
		return livable.Alive()
	}
	return true
}

func AllPlayersAreDead(players []Player) bool {
	for _, p := range players {
		if alive(p) {
			return false
		}
	}
	return true
}

type IsMinion interface {
	IsMinion() bool
}

func AllPlayersAreMinions(players []Player) bool {
	for _, p := range players {
		if minion, ok := p.(IsMinion); ok {
			if !minion.IsMinion() {
				return false
			}
		}
	}
	return true
}

// MainLoop kicks off the next players round
func (g *Game) MainLoop(e Engine) {
	if AllPlayersAreDead(g.Players) {
		e.GameOver()
		return
	}
	//if AllPlayersAreMinions(g.Players) {
	//	e.GameOver()
	//	return
	//}
	if g.Coins == 0 {
		e.GameOver()
		return
	}
	g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
	e.RenderGame(g)
	e.SelectOption(g, g.Players[g.CurrentPlayer], func(selected Option) {
		outcome := selected.Selected()
		e.RenderOutcome(outcome, func() {
			g.MainLoop(e)
		})
	})

}
