package pnpdev

import "github.com/ronna-s/go-ws-type-system/pkg/pnp"

type Option struct {
	Fn          func() pnp.Outcome
	Description string
}

func (o Option) String() string {
	return o.Description
}

func (o Option) Selected() pnp.Outcome {
	return o.Fn()
}

func NewMinion(name string) Minion {
	return Minion{}
}

type Minion struct {
	Name string
}

func (m Minion) AsciiArt() string {
	return minionArt
}

func (m Minion) IsMinion() bool {
	return true
}

func (m Minion) Options(g *pnp.Game) []pnp.Option {
	var options []pnp.Option
	if g.Coins > 0 {
		options = append(options, Option{
			Description: "Buy a banana and eat it (1 gold coin)",
			Fn: func() pnp.Outcome {
				g.Coins--
				return "You ate a banana"
			},
		})
	}
	options = append(options, Option{
		Description: "Add a bug to the code",
		Fn: func() pnp.Outcome {
			g.Prod.React(false)
			return "PRODUCTION is upset"
		},
	})
	return options
}

func (m Minion) String() string {
	return "Minion"
}

type NopLivable struct {
	Minion
}

func (m NopLivable) Alive() bool {
	return true
}

func NewNopLivable(m Minion) NopLivable {
	return NopLivable{m}
}

var _ pnp.Player = NewNopLivable(Minion{})
