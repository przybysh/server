//go:build !prod || full || agt

package wizard

import (
	"github.com/slotopol/server/game"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "AGT", Name: "Wizard"},           // see: https://demo.agtsoftware.com/games/agt/wizard
		{Prov: "AGT", Name: "Around The World"}, // see: https://demo.agtsoftware.com/games/agt/aroundtheworld
	},
	GP: game.GPsel |
		game.GPretrig |
		game.GPscat |
		game.GPwild,
	SX:  5,
	SY:  4,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  0,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	Info.SetupFactory(func() any { return NewGame() }, CalcStat)
}
