package pnpdev

import (
	pnp2 "github.com/ronna-s/go-design-workshop/lessons/lesson1-interfaces/pkg/pnp"
)

// Minion represents a minion P&P player
// The zero value is a dead minion player.
type Minion struct {
}

type Option struct {
	Fn          func() string
	Description string
}

func (o Option) String() string {
	return o.Description
}
func (o Option) Selected() string {
	return o.Fn()
}

func (m Minion) Options(g *pnp2.Game) []pnp2.Option {
	opts := []pnp2.Option{
		Option{
			Description: "Create a bug",
			Fn: func() string {
				return g.Prod.Upset()
			},
		}}
	if g.Coins > 0 {
		opts = append(opts, Option{
			Description: "Buy and eat a banana (1 gold coin)",
			Fn: func() string {
				g.Coins -= 1
				return g.Prod.NoImpact()
			},
		})
	}
	return opts
}

// NewMinion returns a minion
func NewMinion() *Minion {
	return &Minion{}
}

// AsciiArt returns the minion's ascii-art
func (m Minion) AsciiArt() string {
	return minionArt
}

func (m Minion) String() string {
	return "Minion"
}

type Rubyist struct {
	Dead bool
}

func (r *Rubyist) Options(g *pnp2.Game) []pnp2.Option {
	return []pnp2.Option{
		Option{
			Description: "Dark magic",
			Fn: func() string {
				if g.Prod == pnp2.Legacy {
					return g.Prod.CalmDown()
				}
				return g.Prod.Upset()
			},
		},
	}
}
func (r *Rubyist) Alive() bool {
	return !r.Dead
}

func NewRubyist() *Rubyist {
	return &Rubyist{}
}

func (r *Rubyist) AsciiArt() string {
	return rubyistArt
}

func (r *Rubyist) String() string {
	return "Rubyist"
}
