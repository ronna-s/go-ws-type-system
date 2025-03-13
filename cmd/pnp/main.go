package main

import (
	"github.com/ronna-s/go-ws-type-system/pkg/pnp"
	"github.com/ronna-s/go-ws-type-system/pkg/pnp/engine/tview"
)

func main() {
	game := pnp.New()
	pm := ProductManager{Game: game}
	game.Players = append(game.Players, pm)
	game.Run(tview.New())
}

type ProductManager struct {
	Game *pnp.Game
}

func (p ProductManager) Options(g *pnp.Game) []pnp.Option {
	return []pnp.Option{
		{
			Description: "Pay wages",
			OnSelect: func() pnp.Outcome {
				p.Game.Coins -= len(p.Game.Players)
				return "Wages paid"
			},
		},
	}
}

func (p ProductManager) String() string {
	return "Sir Tan Lee Knot"
}

func (p ProductManager) AsciiArt() string {
	return `
 O
/|\
/ \`
}

func (p ProductManager) Alive() bool {
	return p.Game.Coins > 0
}
