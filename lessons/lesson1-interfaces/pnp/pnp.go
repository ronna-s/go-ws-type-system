package pnp

import (
	_ "embed"
)

type (
	// Game represents a Platforms and Programmers Game
	// where a band of player will attempt to take on production
	Game struct {
		Players       []Player
		Prod          ProductionState
		BandName      string
		Coins         int
		CurrentPlayer int
	}

	Option interface {
		// String returns a description of the option (e.g. "Drink coffee and ponder about life").
		String() string
		// Selected notifies the option that it was selected, and returns a description of what happened to PRODUCTION.
		Selected() string
	}

	// Player represents a P&P player
	Player interface {
		Options(g *Game) []Option
		AsciiArt() string
		Alive() bool
	}

	// Engine represents the game's user interface rendering engine
	Engine interface {
		Start()
		RenderGame(g *Game)
		SelectOption(g *Game, player Player, cb func())
		GameOver()
		GameWon()
		PizzaDelivery(fn func())
		RenderActivity(desc string, fn func())
		Welcome(fn func(string))
	}
)

// New returns a new P&P game
func New(players ...Player) *Game {
	g := Game{Players: players, Prod: NewProduction(), Coins: 10}
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

// MainLoop kicks off the next players round
func (g *Game) MainLoop(e Engine) {
	if g.Coins == 0 {
		e.GameOver()
		return
	}
	p := g.Players[g.CurrentPlayer]
	e.RenderGame(g)
	e.SelectOption(g, p, func() {
		g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
		g.MainLoop(e)
	})
}
