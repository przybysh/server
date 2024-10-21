package trolls

import (
	slot "github.com/slotopol/server/game/slot"
)

// Lined payment.
var LinePay = [14][5]float64{
	{0, 3, 25, 100, 750},      //  1 troll1
	{0, 0, 25, 100, 500},      //  2 troll2
	{0, 0, 15, 100, 500},      //  3 troll3
	{0, 0, 10, 75, 250},       //  4 troll4
	{0, 0, 10, 75, 250},       //  5 troll5
	{0, 0, 10, 50, 200},       //  6 troll6
	{0, 0, 5, 50, 150},        //  7 ace
	{0, 0, 5, 25, 125},        //  8 king
	{0, 0, 5, 25, 125},        //  9 queen
	{0, 0, 5, 25, 125},        // 10 jack
	{0, 2, 5, 25, 100},        // 11 ten
	{0, 10, 250, 2500, 10000}, // 12 wild
	{0, 0, 0, 0, 0},           // 13 golden
	{0, 0, 0, 0, 0},           // 14 scatter
}

// Scatters payment.
var ScatPay = [5]float64{0, 2, 5, 25, 500} // 14 scatter

// Scatter freespins table
var ScatFreespin = [5]int{0, 0, 10, 20, 30} // 14 scatter

const (
	jid = 1 // jackpot ID
)

// Jackpot win combinations.
var Jackpot = [14][5]int{
	{0, 0, 0, 0, 0}, //  1 troll1
	{0, 0, 0, 0, 0}, //  2 troll2
	{0, 0, 0, 0, 0}, //  3 troll3
	{0, 0, 0, 0, 0}, //  4 troll4
	{0, 0, 0, 0, 0}, //  5 troll5
	{0, 0, 0, 0, 0}, //  6 troll6
	{0, 0, 0, 0, 0}, //  7 ace
	{0, 0, 0, 0, 0}, //  8 king
	{0, 0, 0, 0, 0}, //  9 queen
	{0, 0, 0, 0, 0}, // 10 jack
	{0, 0, 0, 0, 0}, // 11 ten
	{0, 0, 0, 0, 0}, // 12 wild
	{0, 0, 0, 0, 0}, // 13 golden
	{0, 0, 0, 0, 0}, // 14 scatter
}

// Bet lines
var BetLines = slot.BetLinesNetEnt5x3[:20]

type Game struct {
	slot.Slot5x3 `yaml:",inline"`
	// free spin number
	FS int `json:"fs,omitempty" yaml:"fs,omitempty" xml:"fs,omitempty"`
}

func NewGame() *Game {
	return &Game{
		Slot5x3: slot.Slot5x3{
			Sel: slot.MakeBitNum(len(BetLines), 1),
			Bet: 1,
		},
		FS: 0,
	}
}

const wild1, wild2, scat = 12, 13, 14

func (g *Game) Scanner(screen slot.Screen, wins *slot.Wins) {
	g.ScanLined(screen, wins)
	g.ScanScatters(screen, wins)
}

// Lined symbols calculation.
func (g *Game) ScanLined(screen slot.Screen, wins *slot.Wins) {
	for li := g.Sel.Next(0); li != -1; li = g.Sel.Next(li) {
		var line = BetLines[li-1]

		var mw float64 = 1 // mult wild
		var numw, numl slot.Pos = 0, 5
		var syml slot.Sym
		var x slot.Pos
		for x = 1; x <= 5; x++ {
			var sx = screen.Pos(x, line)
			if sx == wild1 {
				if syml == 0 {
					numw = x
				}
				if mw < 4 {
					mw = 2
				}
			} else if sx == wild2 {
				if syml == 0 {
					numw = x
				}
				mw = 4
			} else if syml == 0 && sx != scat {
				syml = sx
			} else if sx != syml {
				numl = x - 1
				break
			}
		}

		var payw, payl float64
		if numw > 0 {
			payw = LinePay[wild1-1][numw-1]
		}
		if numl > 0 && syml > 0 {
			payl = LinePay[syml-1][numl-1]
		}
		if payl*mw > payw {
			var mm float64 = 1 // mult mode
			if g.FS > 0 {
				mm = 3
			}
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * payl,
				Mult: mw * mm,
				Sym:  syml,
				Num:  numl,
				Line: li,
				XY:   line.CopyL(numl),
			})
		} else if payw > 0 {
			var mm float64 = 1 // mult mode
			if g.FS > 0 {
				mm = 3
			}
			*wins = append(*wins, slot.WinItem{
				Pay:  g.Bet * payw,
				Mult: mm, // no multiplyer on line by double symbol
				Sym:  wild1,
				Num:  numw,
				Line: li,
				XY:   line.CopyL(numw),
				Jack: Jackpot[wild1-1][numw-1],
			})
		}
	}
}

// Scatters calculation.
func (g *Game) ScanScatters(screen slot.Screen, wins *slot.Wins) {
	if count := screen.ScatNum(scat); count >= 2 {
		var mm float64 = 1 // mult mode
		if g.FS > 0 {
			mm = 3
		}
		var pay, fs = ScatPay[count-1], ScatFreespin[count-1]
		*wins = append(*wins, slot.WinItem{
			Pay:  g.Bet * float64(g.Sel.Num()) * pay,
			Mult: mm,
			Sym:  scat,
			Num:  count,
			XY:   screen.ScatPos(scat),
			Free: fs,
		})
	}
}

func (g *Game) Spin(screen slot.Screen, mrtp float64) {
	var reels, _ = slot.FindReels(ReelsMap, mrtp)
	screen.Spin(reels)
}

func (g *Game) Apply(screen slot.Screen, wins slot.Wins) {
	if g.FS > 0 {
		g.Gain += wins.Gain()
	} else {
		g.Gain = wins.Gain()
	}

	if g.FS > 0 {
		g.FS--
	}
	for _, wi := range wins {
		if wi.Free > 0 {
			g.FS += wi.Free
		}
	}
}

func (g *Game) FreeSpins() int {
	return g.FS
}

func (g *Game) SetSel(sel slot.Bitset) error {
	return g.SetSelNum(sel, len(BetLines))
}
