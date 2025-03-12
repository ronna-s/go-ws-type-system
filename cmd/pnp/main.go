package main

import (
	"github.com/ronna-s/go-ws-type-system/pkg/pnp"
	"github.com/ronna-s/go-ws-type-system/pkg/pnp/engine/tview"
)

func main() {
	game := pnp.New()
	game.Run(engine.New())
}
