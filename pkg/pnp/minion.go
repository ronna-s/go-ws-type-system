package pnp

import _ "embed"

func NewMinion(name string) Minion {
	return Minion{}
}

type Minion struct {
	Name string
}

//go:embed resources/minion.txt
var minionArt string

func (m Minion) AsciiArt() string {
	return minionArt
}

func (m Minion) Options(g *Game) []Option {
	var options []Option
	if g.Coins > 0 {
		options = append(options, Option{
			Description: "Buy a banana and eat it (1 gold coin)",
			OnSelect: func() Outcome {
				g.Coins--
				return "You ate a banana"
			},
		})
	}
	options = append(options, Option{
		Description: "Add a bug to the code",
		OnSelect: func() Outcome {
			g.Prod.Upset()
			return "PRODUCTION is upset"
		},
	})
	return options
}

type IsMinion interface {
	isMinion() bool
}

func (m Minion) isMinion() bool {
	return true
}

func (m Minion) String() string {
	return "Minion"
}
