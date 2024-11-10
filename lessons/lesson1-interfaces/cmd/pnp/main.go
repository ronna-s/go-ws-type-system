package main

import (
	"github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pnp"
	"github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pnp/engine/tview"
	"github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pnpdev"
)

func main() {
	game := pnp.New(pnpdev.NewMinion(), pnpdev.NewRubyist())
	game.Run(engine.New())
}
