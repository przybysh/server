//go:build !prod || full || playtech

package links

import (
	"context"

	"github.com/slotopol/server/game/deserttreasure"
	"github.com/spf13/pflag"
)

func init() {
	var gi = GameInfo{
		Aliases: []GameAlias{
			{"deserttreasure", "Desert Treasure"},
		},
		Provider: "Playtech",
		ScrnX:    5,
		ScrnY:    3,
	}
	GameList = append(GameList, gi)

	for _, ga := range gi.Aliases {
		ScanIters = append(ScanIters, func(flags *pflag.FlagSet, ctx context.Context) {
			if is, _ := flags.GetBool(ga.ID); is {
				var rn, _ = flags.GetString("reels")
				if rn == "bon" {
					deserttreasure.CalcStatBon(ctx)
				} else {
					deserttreasure.CalcStatReg(ctx, rn)
				}
			}
		})
		GameFactory[ga.ID] = func(rd string) any {
			if _, ok := deserttreasure.ReelsMap[rd]; ok {
				return deserttreasure.NewGame(rd)
			}
			return nil
		}
	}
}
