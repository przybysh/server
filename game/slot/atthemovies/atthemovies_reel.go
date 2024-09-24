package atthemovies

import (
	slot "github.com/slotopol/server/game/slot"
)

// reels lengths [36, 37, 36, 37, 36], total reshuffles 63872064
// symbols: 75.457(lined) + 10.499(scatter) = 85.955963%
// free spins 2576664, q = 0.040341, sq = 1/(1-q) = 1.042037
// free games frequency: 1/202.8
// RTP = rtp(sym) + q*sq*2*rtp(sym) = 85.956 + 0.040341*179.14 = 93.182595%
var Reels93 = slot.Reels5x{
	{4, 8, 7, 6, 5, 7, 4, 5, 6, 4, 8, 3, 6, 4, 7, 5, 6, 8, 3, 7, 10, 8, 7, 6, 8, 5, 6, 7, 8, 6, 5, 1, 8, 5, 2, 7},
	{4, 8, 6, 4, 5, 6, 9, 8, 2, 7, 8, 4, 7, 5, 8, 7, 5, 8, 6, 4, 5, 8, 6, 3, 7, 6, 8, 10, 7, 6, 3, 5, 7, 1, 5, 6, 7},
	{8, 4, 6, 7, 4, 6, 2, 7, 3, 8, 7, 5, 8, 6, 10, 5, 6, 8, 7, 4, 6, 7, 3, 8, 5, 1, 6, 5, 4, 8, 7, 5, 8, 6, 5, 7},
	{4, 7, 8, 6, 7, 8, 6, 5, 3, 7, 2, 5, 1, 8, 4, 7, 5, 4, 7, 9, 6, 5, 8, 6, 5, 4, 6, 8, 10, 3, 7, 6, 8, 7, 5, 6, 8},
	{4, 6, 8, 4, 7, 6, 4, 7, 8, 3, 7, 5, 2, 8, 5, 7, 8, 1, 7, 6, 8, 7, 5, 6, 10, 7, 6, 3, 5, 4, 6, 5, 8, 6, 5, 8},
}

// reels lengths [31, 35, 34, 35, 31], total reshuffles 40025650
// symbols: 72.319(lined) + 12.352(scatter) = 84.671307%
// free spins 2119824, q = 0.052962, sq = 1/(1-q) = 1.055923
// free games frequency: 1/154.84
// RTP = rtp(sym) + q*sq*2*rtp(sym) = 84.671 + 0.052962*178.81 = 94.141528%
var Reels94 = slot.Reels5x{
	{5, 2, 7, 8, 5, 6, 8, 7, 3, 5, 6, 7, 8, 3, 6, 7, 5, 6, 10, 8, 7, 4, 8, 7, 5, 8, 6, 7, 1, 8, 4},
	{5, 6, 8, 1, 6, 8, 7, 3, 9, 7, 4, 6, 8, 4, 6, 7, 5, 10, 7, 8, 5, 7, 4, 8, 5, 3, 7, 5, 4, 6, 2, 8, 7, 6, 8},
	{8, 6, 3, 5, 8, 4, 6, 1, 7, 6, 8, 7, 5, 8, 7, 6, 8, 7, 4, 8, 5, 2, 7, 6, 4, 7, 5, 6, 3, 5, 8, 7, 4, 10},
	{2, 7, 6, 10, 4, 8, 3, 7, 5, 6, 8, 1, 5, 7, 6, 8, 4, 5, 8, 7, 9, 3, 7, 4, 8, 6, 7, 5, 8, 7, 6, 5, 4, 6, 8},
	{7, 8, 3, 7, 8, 5, 6, 8, 4, 7, 8, 6, 5, 7, 8, 5, 7, 10, 6, 3, 8, 4, 7, 6, 8, 5, 2, 7, 6, 5, 1},
}

// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 74.5(lined) + 11.802(scatter) = 86.301737%
// free spins 2231280, q = 0.049109, sq = 1/(1-q) = 1.051645
// free games frequency: 1/166.88
// RTP = rtp(sym) + q*sq*2*rtp(sym) = 86.302 + 0.049109*181.52 = 95.215851%
var Reels95 = slot.Reels5x{
	{7, 8, 5, 6, 7, 5, 3, 8, 6, 5, 8, 7, 5, 1, 7, 4, 8, 6, 5, 7, 4, 6, 10, 8, 7, 4, 6, 3, 8, 7, 2, 6, 8, 4},
	{5, 2, 7, 4, 6, 3, 8, 7, 5, 6, 10, 8, 3, 6, 8, 5, 7, 6, 8, 7, 6, 8, 7, 4, 6, 8, 5, 4, 6, 1, 7, 9, 8, 7},
	{8, 6, 3, 5, 8, 4, 6, 1, 7, 6, 8, 7, 5, 8, 7, 6, 8, 7, 4, 8, 5, 2, 7, 6, 4, 7, 5, 6, 3, 5, 8, 7, 4, 10},
	{7, 6, 8, 5, 6, 7, 4, 8, 7, 6, 5, 7, 3, 8, 4, 6, 8, 7, 5, 8, 2, 5, 7, 8, 6, 1, 8, 6, 7, 9, 3, 6, 10, 4},
	{5, 8, 6, 1, 5, 7, 8, 6, 10, 5, 7, 6, 5, 7, 6, 4, 7, 5, 8, 4, 7, 8, 4, 7, 3, 8, 6, 7, 3, 8, 2, 6, 8, 4},
}

// reels lengths [34, 35, 34, 35, 34], total reshuffles 48147400
// symbols: 76.744(lined) + 11.572(scatter) = 88.316046%
// free spins 2287008, q = 0.0475, sq = 1/(1-q) = 1.049869
// free games frequency: 1/172.48
// RTP = rtp(sym) + q*sq*2*rtp(sym) = 88.316 + 0.0475*185.44 = 97.124497%
var Reels97 = slot.Reels5x{
	{8, 6, 7, 4, 8, 6, 1, 8, 5, 2, 7, 8, 5, 7, 6, 5, 3, 6, 5, 7, 4, 8, 6, 7, 8, 6, 4, 8, 7, 10, 5, 6, 7, 3},
	{4, 8, 6, 7, 2, 8, 6, 5, 7, 8, 4, 6, 5, 7, 6, 8, 4, 10, 8, 7, 6, 8, 1, 5, 3, 7, 6, 5, 8, 7, 9, 3, 6, 5, 7},
	{7, 8, 5, 6, 8, 7, 5, 6, 7, 5, 3, 7, 5, 6, 8, 4, 7, 8, 6, 3, 8, 6, 4, 8, 7, 6, 1, 8, 6, 7, 2, 5, 10, 4},
	{5, 7, 6, 8, 7, 6, 8, 2, 6, 7, 1, 5, 4, 8, 10, 5, 6, 7, 5, 8, 3, 6, 4, 8, 7, 5, 6, 7, 4, 6, 8, 9, 3, 7, 8},
	{1, 6, 5, 8, 7, 2, 8, 6, 5, 4, 7, 10, 8, 3, 5, 7, 4, 8, 6, 7, 8, 6, 5, 7, 8, 6, 5, 8, 6, 7, 3, 6, 4, 7},
}

// reels lengths [33, 34, 33, 34, 33], total reshuffles 41543172
// symbols: 78.344(lined) + 12.171(scatter) = 90.514056%
// free spins 2148660, q = 0.051721, sq = 1/(1-q) = 1.054542
// free games frequency: 1/158.52
// RTP = rtp(sym) + q*sq*2*rtp(sym) = 90.514 + 0.051721*190.9 = 100.387712%
var Reels100 = slot.Reels5x{
	{8, 7, 3, 6, 5, 7, 8, 1, 6, 4, 7, 8, 2, 6, 5, 8, 7, 6, 8, 5, 6, 7, 10, 8, 4, 7, 6, 8, 3, 7, 6, 4, 5},
	{5, 2, 7, 4, 6, 3, 8, 7, 5, 6, 10, 8, 3, 6, 8, 5, 7, 6, 8, 7, 6, 8, 7, 4, 6, 8, 5, 4, 6, 1, 7, 9, 8, 7},
	{7, 6, 4, 5, 8, 7, 6, 8, 5, 6, 3, 8, 6, 1, 7, 8, 4, 7, 8, 2, 7, 6, 8, 7, 5, 8, 6, 7, 10, 4, 6, 3, 5},
	{7, 6, 8, 5, 6, 7, 4, 8, 7, 6, 5, 7, 3, 8, 4, 6, 8, 7, 5, 8, 2, 5, 7, 8, 6, 1, 8, 6, 7, 9, 3, 6, 10, 4},
	{5, 6, 7, 8, 6, 7, 8, 6, 7, 5, 8, 3, 6, 2, 5, 7, 3, 8, 6, 4, 7, 6, 4, 8, 5, 10, 6, 8, 7, 4, 8, 1, 7},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	93.182595:  &Reels93,
	94.141528:  &Reels94,
	95.215851:  &Reels95,
	97.124497:  &Reels97,
	100.387712: &Reels100,
}
