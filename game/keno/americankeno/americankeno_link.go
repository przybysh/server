//go:build !prod || full || keno || aristocrat

package americankeno

import (
	"context"

	"github.com/slotopol/server/game"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "Aristocrat", Name: "American Keno"},
	},
	GP:  0,
	SX:  80,
	SY:  0,
	SN:  0,
	LN:  0,
	BN:  0,
	RTP: []float64{89.250235},
}

func init() {
	Info.SetupFactory(func() any { return NewGame() }, func(ctx context.Context, mrtp float64) float64 {
		return Paytable.CalcStat(ctx)
	})
}
