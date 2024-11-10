package main

import (
	"github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pkg/pnp"
	"github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pkg/pnp/engine/tview"
)

func main() {
	game := pnp.New()
	game.Run(engine.New())
}
