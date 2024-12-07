//go:build !prod || full || agt

package valkyrie

import (
	"github.com/slotopol/server/game"
)

var Info = game.GameInfo{
	Aliases: []game.GameAlias{
		{Prov: "AGT", Name: "Valkyrie"},
	},
	GP: game.GPsel |
		game.GPfghas |
		game.GPscat |
		game.GPwild |
		game.GPbsym,
	SX:  5,
	SY:  3,
	SN:  len(LinePay),
	LN:  len(BetLines),
	BN:  0,
	RTP: game.MakeRtpList(ReelsMap),
}

func init() {
	Info.SetupFactory(func() any { return NewGame() }, CalcStatReg)
}
