package main

import (
	"github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pkg/pnp"
	"github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pkg/pnp/engine/tview"
	"github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pkg/pnpdev"
)

func main() {
	game := pnp.New(pnpdev.NewRubyist())
	game.Run(engine.New())
}
