package main

import (
	"github.com/ronna-s/go-ws-type-system/pkg/pnp"
	"github.com/ronna-s/go-ws-type-system/pkg/pnp/engine/tview"
	"github.com/ronna-s/go-ws-type-system/pkg/pnpdev"
)

func main() {
	game := pnp.New(pnpdev.NewNopLivable(pnpdev.NewMinion("Bob")))
	game.Run(tview.New())
}
