package captainstreasure

// See: https://freeslotshub.com/playtech/captains-treasure/

import (
	"math"

	"github.com/slotopol/server/game"
)

// reels lengths [32, 42, 42, 42, 32], total reshuffles 75866112
// RTP = 83.431(lined) + 7.3986(scatter) = 90.829406%
var Reels91 = game.Reels5x{
	{8, 6, 5, 9, 7, 10, 11, 7, 10, 3, 11, 8, 7, 9, 8, 10, 9, 8, 11, 6, 7, 10, 4, 6, 11, 8, 2, 9, 11, 6, 9, 10},
	{10, 4, 2, 6, 9, 10, 7, 6, 8, 9, 6, 4, 7, 11, 6, 9, 10, 8, 9, 11, 5, 9, 10, 8, 7, 5, 11, 10, 8, 5, 7, 8, 11, 4, 5, 11, 3, 10, 11, 1, 7, 6},
	{9, 2, 8, 6, 11, 9, 7, 5, 6, 1, 10, 7, 8, 5, 10, 9, 11, 10, 5, 9, 6, 11, 8, 3, 7, 11, 6, 10, 4, 3, 5, 8, 11, 5, 4, 6, 9, 7, 8, 10, 4, 7},
	{11, 2, 7, 5, 8, 11, 10, 7, 8, 1, 11, 8, 9, 6, 11, 7, 6, 11, 9, 3, 4, 5, 6, 4, 9, 10, 7, 9, 10, 6, 11, 7, 5, 10, 8, 9, 6, 10, 4, 8, 10, 5},
	{7, 11, 10, 6, 7, 2, 8, 7, 6, 8, 11, 6, 7, 8, 9, 10, 6, 9, 8, 10, 9, 4, 11, 9, 10, 5, 11, 8, 9, 3, 11, 10},
}

// reels lengths [32, 41, 42, 41, 32], total reshuffles 72296448
// RTP = 84.746(lined) + 7.5408(scatter) = 92.287131%
var Reels92 = game.Reels5x{
	{8, 6, 5, 9, 7, 10, 11, 7, 10, 3, 11, 8, 7, 9, 8, 10, 9, 8, 11, 6, 7, 10, 4, 6, 11, 8, 2, 9, 11, 6, 9, 10},
	{10, 9, 8, 6, 10, 9, 5, 6, 7, 8, 9, 5, 8, 4, 9, 11, 10, 2, 4, 7, 11, 5, 8, 3, 1, 5, 9, 6, 10, 4, 11, 5, 7, 10, 11, 7, 6, 8, 7, 6, 11},
	{4, 9, 11, 7, 10, 8, 6, 1, 5, 7, 6, 9, 11, 8, 7, 6, 9, 4, 11, 7, 10, 5, 9, 8, 7, 3, 8, 10, 6, 5, 10, 2, 8, 11, 4, 5, 10, 9, 3, 5, 11, 6},
	{5, 9, 1, 11, 10, 9, 11, 10, 8, 5, 2, 11, 6, 8, 4, 7, 10, 5, 6, 7, 5, 10, 4, 8, 7, 6, 8, 9, 3, 6, 7, 8, 5, 10, 9, 6, 11, 9, 7, 4, 11},
	{7, 11, 10, 6, 7, 2, 8, 7, 6, 8, 11, 6, 7, 8, 9, 10, 6, 9, 8, 10, 9, 4, 11, 9, 10, 5, 11, 8, 9, 3, 11, 10},
}

// reels lengths [32, 40, 40, 40, 32], total reshuffles 65536000
// RTP = 85.561(lined) + 7.8403(scatter) = 93.400847%
var Reels93 = game.Reels5x{
	{8, 6, 5, 9, 7, 10, 11, 7, 10, 3, 11, 8, 7, 9, 8, 10, 9, 8, 11, 6, 7, 10, 4, 6, 11, 8, 2, 9, 11, 6, 9, 10},
	{6, 11, 8, 7, 9, 10, 6, 3, 11, 1, 9, 10, 7, 11, 8, 10, 11, 8, 7, 5, 2, 8, 5, 9, 6, 4, 5, 7, 6, 10, 5, 9, 6, 8, 5, 4, 7, 9, 11, 10},
	{11, 8, 7, 5, 10, 9, 5, 8, 10, 5, 4, 6, 3, 11, 4, 8, 9, 11, 5, 6, 7, 9, 5, 1, 6, 10, 7, 8, 9, 7, 11, 10, 9, 6, 8, 11, 2, 10, 6, 7},
	{9, 5, 7, 9, 10, 7, 6, 10, 11, 5, 8, 9, 7, 11, 1, 4, 6, 7, 8, 4, 11, 9, 3, 8, 5, 11, 10, 5, 2, 8, 7, 10, 6, 5, 8, 9, 6, 10, 11, 6},
	{7, 11, 10, 6, 7, 2, 8, 7, 6, 8, 11, 6, 7, 8, 9, 10, 6, 9, 8, 10, 9, 4, 11, 9, 10, 5, 11, 8, 9, 3, 11, 10},
}

// reels lengths [32, 40, 39, 40, 32], total reshuffles 63897600
// RTP = 86.295(lined) + 7.9206(scatter) = 94.215525%
var Reels94 = game.Reels5x{
	{8, 6, 5, 9, 7, 10, 11, 7, 10, 3, 11, 8, 7, 9, 8, 10, 9, 8, 11, 6, 7, 10, 4, 6, 11, 8, 2, 9, 11, 6, 9, 10},
	{6, 11, 8, 7, 9, 10, 6, 3, 11, 1, 9, 10, 7, 11, 8, 10, 11, 8, 7, 5, 2, 8, 5, 9, 6, 4, 5, 7, 6, 10, 5, 9, 6, 8, 5, 4, 7, 9, 11, 10},
	{6, 7, 5, 11, 6, 8, 9, 7, 5, 8, 4, 11, 3, 7, 8, 1, 10, 6, 7, 5, 10, 11, 6, 7, 9, 11, 2, 9, 10, 4, 9, 8, 6, 9, 10, 8, 5, 11, 10},
	{9, 5, 7, 9, 10, 7, 6, 10, 11, 5, 8, 9, 7, 11, 1, 4, 6, 7, 8, 4, 11, 9, 3, 8, 5, 11, 10, 5, 2, 8, 7, 10, 6, 5, 8, 9, 6, 10, 11, 6},
	{7, 11, 10, 6, 7, 2, 8, 7, 6, 8, 11, 6, 7, 8, 9, 10, 6, 9, 8, 10, 9, 4, 11, 9, 10, 5, 11, 8, 9, 3, 11, 10},
}

// reels lengths [32, 39, 39, 39, 32], total reshuffles 60742656
// RTP = 87.27(lined) + 8.0828(scatter) = 95.353086%
var Reels95 = game.Reels5x{
	{8, 6, 5, 9, 7, 10, 11, 7, 10, 3, 11, 8, 7, 9, 8, 10, 9, 8, 11, 6, 7, 10, 4, 6, 11, 8, 2, 9, 11, 6, 9, 10},
	{6, 7, 5, 11, 6, 8, 9, 7, 5, 8, 4, 11, 3, 7, 8, 1, 10, 6, 7, 5, 10, 11, 6, 7, 9, 11, 2, 9, 10, 4, 9, 8, 6, 9, 10, 8, 5, 11, 10},
	{10, 11, 9, 2, 7, 8, 4, 7, 9, 10, 11, 5, 8, 11, 5, 8, 9, 7, 6, 9, 5, 10, 6, 7, 11, 4, 8, 6, 10, 8, 3, 6, 9, 5, 11, 1, 10, 7, 6},
	{11, 10, 6, 4, 5, 9, 11, 7, 9, 3, 11, 10, 5, 9, 10, 8, 6, 11, 8, 6, 10, 7, 5, 10, 8, 6, 7, 4, 8, 11, 9, 2, 7, 9, 5, 1, 6, 8, 7},
	{7, 11, 10, 6, 7, 2, 8, 7, 6, 8, 11, 6, 7, 8, 9, 10, 6, 9, 8, 10, 9, 4, 11, 9, 10, 5, 11, 8, 9, 3, 11, 10},
}

// reels lengths [32, 38, 39, 38, 32], total reshuffles 57667584
// RTP = 88.42(lined) + 8.2548(scatter) = 96.674653%
var Reels97 = game.Reels5x{
	{8, 6, 5, 9, 7, 10, 11, 7, 10, 3, 11, 8, 7, 9, 8, 10, 9, 8, 11, 6, 7, 10, 4, 6, 11, 8, 2, 9, 11, 6, 9, 10},
	{11, 7, 10, 3, 6, 4, 5, 10, 9, 7, 5, 8, 11, 4, 7, 10, 8, 11, 9, 6, 11, 8, 9, 10, 5, 6, 10, 8, 2, 6, 8, 7, 9, 1, 11, 7, 6, 9},
	{4, 5, 8, 9, 7, 11, 2, 10, 4, 9, 6, 3, 7, 8, 10, 7, 6, 9, 8, 6, 11, 5, 8, 1, 6, 11, 10, 5, 6, 9, 7, 10, 11, 8, 7, 11, 9, 5, 10},
	{7, 11, 9, 10, 11, 4, 5, 8, 1, 6, 7, 10, 11, 8, 5, 6, 9, 10, 2, 7, 4, 6, 8, 9, 3, 5, 7, 10, 11, 6, 7, 9, 8, 11, 9, 6, 8, 10},
	{7, 11, 10, 6, 7, 2, 8, 7, 6, 8, 11, 6, 7, 8, 9, 10, 6, 9, 8, 10, 9, 4, 11, 9, 10, 5, 11, 8, 9, 3, 11, 10},
}

// reels lengths [32, 38, 38, 38, 32], total reshuffles 56188928
// RTP = 89.575(lined) + 8.3417(scatter) = 97.917193%
var Reels98 = game.Reels5x{
	{6, 11, 10, 6, 7, 9, 11, 10, 3, 7, 8, 9, 10, 2, 9, 10, 8, 11, 4, 6, 11, 9, 10, 8, 7, 6, 8, 11, 7, 8, 5, 9},
	{11, 7, 10, 3, 6, 4, 5, 10, 9, 7, 5, 8, 11, 4, 7, 10, 8, 11, 9, 6, 11, 8, 9, 10, 5, 6, 10, 8, 2, 6, 8, 7, 9, 1, 11, 7, 6, 9},
	{6, 4, 9, 11, 10, 7, 3, 4, 6, 11, 7, 8, 10, 9, 11, 5, 9, 8, 7, 2, 11, 6, 8, 10, 9, 6, 8, 7, 10, 5, 9, 1, 10, 5, 6, 11, 8, 7},
	{7, 11, 9, 10, 11, 4, 5, 8, 1, 6, 7, 10, 11, 8, 5, 6, 9, 10, 2, 7, 4, 6, 8, 9, 3, 5, 7, 10, 11, 6, 7, 9, 8, 11, 9, 6, 8, 10},
	{11, 8, 10, 9, 11, 7, 9, 8, 10, 2, 9, 7, 6, 11, 8, 10, 3, 11, 6, 8, 9, 4, 6, 9, 10, 5, 11, 7, 10, 6, 8, 7},
}

// reels lengths [32, 37, 39, 37, 32], total reshuffles 54672384
// RTP = 90.443(lined) + 8.4374(scatter) = 98.880793%
var Reels99 = game.Reels5x{
	{6, 11, 10, 6, 7, 9, 11, 10, 3, 7, 8, 9, 10, 2, 9, 10, 8, 11, 4, 6, 11, 9, 10, 8, 7, 6, 8, 11, 7, 8, 5, 9},
	{7, 5, 4, 7, 10, 6, 11, 8, 9, 10, 8, 3, 11, 1, 8, 10, 6, 11, 9, 5, 6, 7, 8, 11, 7, 9, 6, 2, 9, 11, 10, 9, 8, 6, 7, 4, 10},
	{9, 10, 6, 8, 5, 11, 9, 8, 10, 4, 1, 6, 9, 8, 7, 11, 4, 8, 10, 6, 11, 10, 8, 7, 9, 10, 7, 11, 3, 5, 6, 9, 5, 2, 7, 11, 4, 7, 6},
	{11, 6, 7, 8, 1, 5, 6, 9, 8, 6, 9, 11, 8, 5, 2, 10, 11, 9, 7, 10, 8, 6, 10, 4, 7, 8, 4, 9, 10, 7, 11, 9, 3, 11, 6, 10, 7},
	{11, 8, 10, 9, 11, 7, 9, 8, 10, 2, 9, 7, 6, 11, 8, 10, 3, 11, 6, 8, 9, 4, 6, 9, 10, 5, 11, 7, 10, 6, 8, 7},
}

// reels lengths [32, 37, 38, 37, 32], total reshuffles 53270528
// RTP = 91.111(lined) + 8.5256(scatter) = 99.636996%
var Reels100 = game.Reels5x{
	{6, 11, 10, 6, 7, 9, 11, 10, 3, 7, 8, 9, 10, 2, 9, 10, 8, 11, 4, 6, 11, 9, 10, 8, 7, 6, 8, 11, 7, 8, 5, 9},
	{7, 5, 4, 7, 10, 6, 11, 8, 9, 10, 8, 3, 11, 1, 8, 10, 6, 11, 9, 5, 6, 7, 8, 11, 7, 9, 6, 2, 9, 11, 10, 9, 8, 6, 7, 4, 10},
	{6, 4, 9, 11, 10, 7, 3, 4, 6, 11, 7, 8, 10, 9, 11, 5, 9, 8, 7, 2, 11, 6, 8, 10, 9, 6, 8, 7, 10, 5, 9, 1, 10, 5, 6, 11, 8, 7},
	{11, 6, 7, 8, 1, 5, 6, 9, 8, 6, 9, 11, 8, 5, 2, 10, 11, 9, 7, 10, 8, 6, 10, 4, 7, 8, 4, 9, 10, 7, 11, 9, 3, 11, 6, 10, 7},
	{11, 8, 10, 9, 11, 7, 9, 8, 10, 2, 9, 7, 6, 11, 8, 10, 3, 11, 6, 8, 9, 4, 6, 9, 10, 5, 11, 7, 10, 6, 8, 7},
}

// reels lengths [31, 32, 32, 32, 31], total reshuffles 31490048
// RTP = 101.31(lined) + 10.636(scatter) = 111.944212%
var Reels112 = game.Reels5x{
	{8, 11, 3, 10, 5, 8, 9, 6, 11, 7, 10, 11, 6, 7, 8, 6, 10, 8, 2, 10, 9, 11, 7, 9, 10, 4, 9, 11, 7, 5, 6},
	{9, 10, 6, 11, 7, 6, 11, 9, 8, 10, 4, 8, 11, 6, 7, 10, 8, 2, 9, 11, 7, 1, 10, 7, 9, 10, 5, 6, 3, 8, 5, 11},
	{1, 11, 7, 9, 6, 10, 8, 11, 3, 6, 11, 7, 9, 2, 5, 9, 10, 8, 7, 6, 9, 10, 11, 8, 6, 4, 10, 5, 11, 7, 8, 10},
	{3, 10, 6, 9, 1, 11, 7, 5, 4, 6, 11, 7, 9, 10, 8, 6, 5, 7, 10, 11, 9, 2, 7, 9, 8, 6, 11, 10, 8, 11, 10, 8},
	{7, 8, 10, 11, 7, 8, 11, 6, 10, 8, 11, 10, 6, 5, 9, 8, 10, 7, 6, 9, 11, 10, 9, 3, 7, 6, 2, 11, 5, 9, 4},
}

// Map with available reels.
var reelsmap = map[float64]*game.Reels5x{
	90.829406:  &Reels91,
	92.287131:  &Reels92,
	93.400847:  &Reels93,
	94.215525:  &Reels94,
	95.353086:  &Reels95,
	96.674653:  &Reels97,
	97.917193:  &Reels98,
	98.880793:  &Reels99,
	99.636996:  &Reels100,
	111.944212: &Reels112,
}

func FindReels(mrtp float64) (rtp float64, reels game.Reels) {
	for p, r := range reelsmap {
		if math.Abs(mrtp-p) < math.Abs(mrtp-rtp) {
			rtp, reels = p, r
		}
	}
	return
}

// Lined payment.
var LinePay = [11][5]float64{
	{0, 0, 0, 0, 0},         //  1 wild
	{0, 0, 0, 0, 0},         //  2 scatter
	{2, 10, 100, 500, 5000}, //  3 sabers
	{0, 5, 50, 250, 2500},   //  4 map
	{0, 3, 20, 100, 1000},   //  5 anchor
	{0, 0, 10, 30, 500},     //  6 ace
	{0, 0, 5, 25, 300},      //  7 king
	{0, 0, 5, 20, 200},      //  8 queen
	{0, 0, 5, 20, 200},      //  9 jack
	{0, 0, 5, 15, 100},      // 10 ten
	{0, 0, 5, 15, 100},      // 11 nine
}

// Scatters payment.
var ScatPay = [5]float64{0, 1, 5, 10, 100} //  2 suitcase

type Game struct {
	game.Slot5x3 `yaml:",inline"`
}

func NewGame(rtp float64) *Game {
	return &Game{
		Slot5x3: game.Slot5x3{
			RTP: rtp,
			SBL: game.MakeBitNum(9),
			Bet: 1,
		},
	}
}

const wild, scat = 1, 2

var bl = game.Lineset5x{
	{2, 2, 2, 2, 2}, // 1
	{1, 1, 1, 1, 1}, // 2
	{3, 3, 3, 3, 3}, // 3
	{1, 2, 3, 2, 1}, // 4
	{3, 2, 1, 2, 3}, // 5
	{1, 1, 2, 1, 1}, // 6
	{3, 3, 2, 3, 3}, // 7
	{2, 1, 1, 1, 2}, // 8
	{2, 3, 3, 3, 2}, // 9
}

func (g *Game) Scanner(screen game.Screen, wins *game.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen game.Screen, wins *game.Wins) {
	for li := g.SBL.Next(0); li != 0; li = g.SBL.Next(li) {
		var line = bl.Line(li)

		var syml, numl = screen.Pos(1, line), 1
		var mw float64 = 1 // mult wild
		for x := 2; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild {
				mw = 2
			} else if sx != syml {
				break
			}
			numl++
		}

		if pay := LinePay[syml-1][numl-1]; pay > 0 {
			*wins = append(*wins, game.WinItem{
				Pay:  g.Bet * pay,
				Mult: mw,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		}

		if numl < 5 {
			var symr, numr = screen.Pos(5, line), 1
			var mw float64 = 1 // mult wild
			for x := 4; x > numl; x-- {
				var sx = screen.Pos(x, line)
				if sx == wild {
					mw = 2
				} else if sx != symr {
					break
				}
				numr++
			}

			if pay := LinePay[symr-1][numr-1]; pay > 0 {
				*wins = append(*wins, game.WinItem{
					Pay:  g.Bet * pay,
					Mult: mw,
					Sym:  symr,
					Num:  numr,
					Line: li,
					XY:   line.CopyR(numr),
				})
			}
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen game.Screen, wins *game.Wins) {
	if count := screen.ScatNum(scat); count >= 2 {
		var pay = ScatPay[count-1]
		*wins = append(*wins, game.WinItem{
			Pay:  g.Bet * float64(g.SBL.Num()) * pay,
			Mult: 1,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
		})
	}
}

func (g *Game) Spin(screen game.Screen) {
	var _, reels = FindReels(g.RTP)
	screen.Spin(reels)
}

func (g *Game) SetLines(sbl game.Bitset) error {
	var mask game.Bitset = (1<<len(bl) - 1) << 1
	if sbl == 0 {
		return game.ErrNoLineset
	}
	if sbl&^mask != 0 {
		return game.ErrLinesetOut
	}
	if g.FreeSpins() > 0 {
		return game.ErrNoFeature
	}
	g.SBL = sbl
	return nil
}
