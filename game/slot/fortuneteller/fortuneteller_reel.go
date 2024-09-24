package fortuneteller

import (
	slot "github.com/slotopol/server/game/slot"
)

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 83.217(lined) + 0(scatter) = 83.216850%
// cards bonuses: frequency 1/181.31, rtp = 22.449700%
// RTP = 83.217(sym) + 22.45(cards) = 105.666550%
// *regular reels calculations*
// reels lengths [35, 35, 35, 35, 35], total reshuffles 52521875
// symbols: 39.928(lined) + 13.688(scatter) = 53.616431%
// free spins 4412475, q = 0.084012
// free games frequency: 1/181.31
// cards bonuses: frequency 1/181.31, rtp = 22.449700%
// RTP = 53.616(sym) + 22.45(cards) + 0.084012*105.67(fg) = 84.943404%
var Reels85 = slot.Reels5x{
	{4, 8, 12, 4, 10, 11, 5, 4, 12, 8, 10, 3, 12, 11, 7, 5, 9, 10, 8, 11, 7, 6, 9, 7, 6, 8, 5, 10, 9, 12, 2, 6, 9, 1, 11},
	{9, 3, 10, 5, 8, 10, 7, 8, 12, 9, 8, 11, 12, 4, 1, 5, 6, 12, 8, 11, 10, 2, 11, 6, 7, 9, 4, 12, 5, 9, 11, 6, 7, 10, 4},
	{12, 9, 5, 10, 1, 6, 9, 12, 11, 8, 7, 6, 10, 8, 3, 12, 4, 11, 12, 5, 9, 8, 2, 7, 11, 9, 5, 4, 8, 7, 10, 11, 6, 4, 10},
	{11, 8, 5, 11, 9, 8, 10, 11, 12, 5, 4, 8, 12, 9, 7, 10, 6, 8, 2, 12, 4, 6, 3, 9, 10, 7, 5, 6, 1, 11, 9, 12, 4, 7, 10},
	{8, 1, 12, 7, 4, 10, 3, 5, 8, 11, 12, 7, 10, 5, 8, 12, 9, 11, 12, 9, 11, 5, 10, 2, 11, 6, 7, 4, 9, 8, 6, 9, 4, 6, 10},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [34, 35, 35, 35, 34], total reshuffles 49563500
// symbols: 84.226(lined) + 0(scatter) = 84.225781%
// cards bonuses: frequency 1/175.36, rtp = 23.210969%
// RTP = 84.226(sym) + 23.211(cards) = 107.436750%
// *regular reels calculations*
// reels lengths [34, 35, 35, 35, 34], total reshuffles 49563500
// symbols: 40.111(lined) + 14.011(scatter) = 54.121680%
// free spins 4305960, q = 0.086878
// free games frequency: 1/175.36
// cards bonuses: frequency 1/175.36, rtp = 23.210969%
// RTP = 54.122(sym) + 23.211(cards) + 0.086878*107.44(fg) = 86.666500%
var Reels87 = slot.Reels5x{
	{4, 6, 10, 8, 2, 11, 5, 8, 6, 9, 7, 12, 9, 7, 11, 12, 4, 9, 10, 3, 12, 11, 5, 9, 10, 1, 4, 11, 7, 6, 8, 10, 5, 12},
	{9, 3, 10, 5, 8, 10, 7, 8, 12, 9, 8, 11, 12, 4, 1, 5, 6, 12, 8, 11, 10, 2, 11, 6, 7, 9, 4, 12, 5, 9, 11, 6, 7, 10, 4},
	{12, 9, 5, 10, 1, 6, 9, 12, 11, 8, 7, 6, 10, 8, 3, 12, 4, 11, 12, 5, 9, 8, 2, 7, 11, 9, 5, 4, 8, 7, 10, 11, 6, 4, 10},
	{11, 8, 5, 11, 9, 8, 10, 11, 12, 5, 4, 8, 12, 9, 7, 10, 6, 8, 2, 12, 4, 6, 3, 9, 10, 7, 5, 6, 1, 11, 9, 12, 4, 7, 10},
	{10, 9, 11, 3, 6, 8, 12, 4, 7, 9, 12, 8, 7, 5, 4, 10, 2, 12, 4, 6, 9, 11, 7, 10, 9, 5, 10, 12, 8, 6, 11, 1, 5, 11},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [34, 34, 35, 34, 34], total reshuffles 46771760
// symbols: 86.095(lined) + 0(scatter) = 86.095082%
// cards bonuses: frequency 1/169.65, rtp = 23.992532%
// RTP = 86.095(sym) + 23.993(cards) = 110.087614%
// *regular reels calculations*
// reels lengths [34, 34, 35, 34, 34], total reshuffles 46771760
// symbols: 40.714(lined) + 14.339(scatter) = 55.053417%
// free spins 4201065, q = 0.089821
// free games frequency: 1/169.65
// cards bonuses: frequency 1/169.65, rtp = 23.992532%
// RTP = 55.053(sym) + 23.993(cards) + 0.089821*110.09(fg) = 88.934078%
var Reels89 = slot.Reels5x{
	{4, 6, 10, 8, 2, 11, 5, 8, 6, 9, 7, 12, 9, 7, 11, 12, 4, 9, 10, 3, 12, 11, 5, 9, 10, 1, 4, 11, 7, 6, 8, 10, 5, 12},
	{10, 5, 6, 10, 11, 6, 12, 2, 9, 7, 10, 9, 12, 5, 4, 8, 11, 4, 8, 7, 12, 10, 6, 7, 9, 12, 11, 3, 4, 8, 5, 1, 9, 11},
	{12, 9, 5, 10, 1, 6, 9, 12, 11, 8, 7, 6, 10, 8, 3, 12, 4, 11, 12, 5, 9, 8, 2, 7, 11, 9, 5, 4, 8, 7, 10, 11, 6, 4, 10},
	{9, 5, 2, 8, 4, 12, 7, 4, 6, 11, 12, 10, 6, 9, 1, 7, 12, 5, 10, 6, 11, 8, 4, 10, 11, 7, 9, 11, 5, 10, 9, 12, 3, 8},
	{10, 9, 11, 3, 6, 8, 12, 4, 7, 9, 12, 8, 7, 5, 4, 10, 2, 12, 4, 6, 9, 11, 7, 10, 9, 5, 10, 12, 8, 6, 11, 1, 5, 11},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 87.48(lined) + 0(scatter) = 87.479507%
// cards bonuses: frequency 1/166.88, rtp = 24.391009%
// RTP = 87.48(sym) + 24.391(cards) = 111.870516%
// *regular reels calculations*
// reels lengths [34, 34, 34, 34, 34], total reshuffles 45435424
// symbols: 41.257(lined) + 14.505(scatter) = 55.761755%
// free spins 4149225, q = 0.091321
// free games frequency: 1/166.88
// cards bonuses: frequency 1/166.88, rtp = 24.391009%
// RTP = 55.762(sym) + 24.391(cards) + 0.091321*111.87(fg) = 90.368933%
var Reels90 = slot.Reels5x{
	{4, 6, 10, 8, 2, 11, 5, 8, 6, 9, 7, 12, 9, 7, 11, 12, 4, 9, 10, 3, 12, 11, 5, 9, 10, 1, 4, 11, 7, 6, 8, 10, 5, 12},
	{10, 5, 6, 10, 11, 6, 12, 2, 9, 7, 10, 9, 12, 5, 4, 8, 11, 4, 8, 7, 12, 10, 6, 7, 9, 12, 11, 3, 4, 8, 5, 1, 9, 11},
	{5, 10, 11, 12, 10, 6, 11, 2, 8, 10, 1, 9, 11, 7, 9, 4, 6, 12, 8, 10, 3, 4, 12, 7, 9, 12, 5, 9, 8, 5, 4, 7, 6, 11},
	{9, 5, 2, 8, 4, 12, 7, 4, 6, 11, 12, 10, 6, 9, 1, 7, 12, 5, 10, 6, 11, 8, 4, 10, 11, 7, 9, 11, 5, 10, 9, 12, 3, 8},
	{10, 9, 11, 3, 6, 8, 12, 4, 7, 9, 12, 8, 7, 5, 4, 10, 2, 12, 4, 6, 9, 11, 7, 10, 9, 5, 10, 12, 8, 6, 11, 1, 5, 11},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [34, 34, 33, 34, 34], total reshuffles 44099088
// symbols: 88.227(lined) + 0(scatter) = 88.226688%
// cards bonuses: frequency 1/164.04, rtp = 24.813637%
// RTP = 88.227(sym) + 24.814(cards) = 113.040325%
// *regular reels calculations*
// reels lengths [34, 34, 33, 34, 34], total reshuffles 44099088
// symbols: 41.325(lined) + 14.681(scatter) = 56.005303%
// free spins 4097385, q = 0.092913
// free games frequency: 1/164.04
// cards bonuses: frequency 1/164.04, rtp = 24.813637%
// RTP = 56.005(sym) + 24.814(cards) + 0.092913*113.04(fg) = 91.321872%
var Reels91 = slot.Reels5x{
	{4, 6, 10, 8, 2, 11, 5, 8, 6, 9, 7, 12, 9, 7, 11, 12, 4, 9, 10, 3, 12, 11, 5, 9, 10, 1, 4, 11, 7, 6, 8, 10, 5, 12},
	{10, 5, 6, 10, 11, 6, 12, 2, 9, 7, 10, 9, 12, 5, 4, 8, 11, 4, 8, 7, 12, 10, 6, 7, 9, 12, 11, 3, 4, 8, 5, 1, 9, 11},
	{11, 9, 10, 12, 3, 9, 6, 1, 7, 10, 8, 4, 11, 2, 10, 6, 4, 5, 11, 12, 8, 4, 7, 9, 8, 10, 12, 11, 5, 6, 12, 7, 5},
	{9, 5, 2, 8, 4, 12, 7, 4, 6, 11, 12, 10, 6, 9, 1, 7, 12, 5, 10, 6, 11, 8, 4, 10, 11, 7, 9, 11, 5, 10, 9, 12, 3, 8},
	{10, 9, 11, 3, 6, 8, 12, 4, 7, 9, 12, 8, 7, 5, 4, 10, 2, 12, 4, 6, 9, 11, 7, 10, 9, 5, 10, 12, 8, 6, 11, 1, 5, 11},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [33, 34, 34, 34, 33], total reshuffles 42802056
// symbols: 88.527(lined) + 0(scatter) = 88.526892%
// cards bonuses: frequency 1/161.25, rtp = 25.242048%
// RTP = 88.527(sym) + 25.242(cards) = 113.768940%
// *regular reels calculations*
// reels lengths [33, 34, 34, 34, 33], total reshuffles 42802056
// symbols: 41.402(lined) + 14.858(scatter) = 56.259671%
// free spins 4045950, q = 0.094527
// free games frequency: 1/161.25
// cards bonuses: frequency 1/161.25, rtp = 25.242048%
// RTP = 56.26(sym) + 25.242(cards) + 0.094527*113.77(fg) = 92.255956%
var Reels92 = slot.Reels5x{
	{11, 8, 4, 12, 3, 11, 10, 7, 4, 12, 5, 8, 10, 7, 11, 12, 2, 10, 9, 11, 5, 6, 7, 9, 8, 12, 6, 9, 5, 6, 4, 10, 1},
	{10, 5, 6, 10, 11, 6, 12, 2, 9, 7, 10, 9, 12, 5, 4, 8, 11, 4, 8, 7, 12, 10, 6, 7, 9, 12, 11, 3, 4, 8, 5, 1, 9, 11},
	{5, 10, 11, 12, 10, 6, 11, 2, 8, 10, 1, 9, 11, 7, 9, 4, 6, 12, 8, 10, 3, 4, 12, 7, 9, 12, 5, 9, 8, 5, 4, 7, 6, 11},
	{9, 5, 2, 8, 4, 12, 7, 4, 6, 11, 12, 10, 6, 9, 1, 7, 12, 5, 10, 6, 11, 8, 4, 10, 11, 7, 9, 11, 5, 10, 9, 12, 3, 8},
	{11, 7, 8, 3, 5, 10, 9, 6, 4, 8, 5, 6, 10, 4, 12, 9, 11, 4, 6, 1, 8, 10, 12, 11, 7, 12, 9, 11, 7, 10, 5, 2, 12},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [34, 33, 33, 33, 34], total reshuffles 41543172
// symbols: 90.119(lined) + 0(scatter) = 90.118889%
// cards bonuses: frequency 1/158.52, rtp = 25.676283%
// RTP = 90.119(sym) + 25.676(cards) = 115.795172%
// *regular reels calculations*
// reels lengths [34, 33, 33, 33, 34], total reshuffles 41543172
// symbols: 41.874(lined) + 15.037(scatter) = 56.911280%
// free spins 3994920, q = 0.096163
// free games frequency: 1/158.52
// cards bonuses: frequency 1/158.52, rtp = 25.676283%
// RTP = 56.911(sym) + 25.676(cards) + 0.096163*115.8(fg) = 93.722785%
var Reels94 = slot.Reels5x{
	{4, 6, 10, 8, 2, 11, 5, 8, 6, 9, 7, 12, 9, 7, 11, 12, 4, 9, 10, 3, 12, 11, 5, 9, 10, 1, 4, 11, 7, 6, 8, 10, 5, 12},
	{5, 9, 12, 4, 10, 12, 9, 2, 10, 8, 6, 11, 8, 4, 12, 11, 7, 3, 12, 11, 7, 5, 11, 9, 10, 1, 5, 6, 7, 10, 8, 4, 6},
	{11, 9, 10, 12, 3, 9, 6, 1, 7, 10, 8, 4, 11, 2, 10, 6, 4, 5, 11, 12, 8, 4, 7, 9, 8, 10, 12, 11, 5, 6, 12, 7, 5},
	{12, 11, 6, 4, 2, 6, 9, 11, 10, 8, 3, 12, 4, 8, 9, 4, 12, 1, 9, 7, 10, 5, 11, 8, 7, 10, 12, 7, 11, 5, 6, 10, 5},
	{10, 9, 11, 3, 6, 8, 12, 4, 7, 9, 12, 8, 7, 5, 4, 10, 2, 12, 4, 6, 9, 11, 7, 10, 9, 5, 10, 12, 8, 6, 11, 1, 5, 11},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [33, 33, 34, 33, 33], total reshuffles 40321314
// symbols: 90.505(lined) + 0(scatter) = 90.504962%
// cards bonuses: frequency 1/155.85, rtp = 26.116382%
// RTP = 90.505(sym) + 26.116(cards) = 116.621344%
// *regular reels calculations*
// reels lengths [33, 33, 34, 33, 33], total reshuffles 40321314
// symbols: 41.994(lined) + 15.217(scatter) = 57.211367%
// free spins 3944295, q = 0.097822
// free games frequency: 1/155.85
// cards bonuses: frequency 1/155.85, rtp = 26.116382%
// RTP = 57.211(sym) + 26.116(cards) + 0.097822*116.62(fg) = 94.735834%
var Reels95 = slot.Reels5x{
	{11, 8, 4, 12, 3, 11, 10, 7, 4, 12, 5, 8, 10, 7, 11, 12, 2, 10, 9, 11, 5, 6, 7, 9, 8, 12, 6, 9, 5, 6, 4, 10, 1},
	{5, 9, 12, 4, 10, 12, 9, 2, 10, 8, 6, 11, 8, 4, 12, 11, 7, 3, 12, 11, 7, 5, 11, 9, 10, 1, 5, 6, 7, 10, 8, 4, 6},
	{5, 10, 11, 12, 10, 6, 11, 2, 8, 10, 1, 9, 11, 7, 9, 4, 6, 12, 8, 10, 3, 4, 12, 7, 9, 12, 5, 9, 8, 5, 4, 7, 6, 11},
	{12, 11, 6, 4, 2, 6, 9, 11, 10, 8, 3, 12, 4, 8, 9, 4, 12, 1, 9, 7, 10, 5, 11, 8, 7, 10, 12, 7, 11, 5, 6, 10, 5},
	{11, 7, 8, 3, 5, 10, 9, 6, 4, 8, 5, 6, 10, 4, 12, 9, 11, 4, 6, 1, 8, 10, 12, 11, 7, 12, 9, 11, 7, 10, 5, 2, 12},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [33, 33, 33, 33, 33], total reshuffles 39135393
// symbols: 91.971(lined) + 0(scatter) = 91.970560%
// cards bonuses: frequency 1/153.24, rtp = 26.562384%
// RTP = 91.971(sym) + 26.562(cards) = 118.532943%
// *regular reels calculations*
// reels lengths [33, 33, 33, 33, 33], total reshuffles 39135393
// symbols: 42.54(lined) + 15.399(scatter) = 57.939127%
// free spins 3894075, q = 0.099503
// free games frequency: 1/153.24
// cards bonuses: frequency 1/153.24, rtp = 26.562384%
// RTP = 57.939(sym) + 26.562(cards) + 0.099503*118.53(fg) = 96.295851%
var Reels96 = slot.Reels5x{
	{11, 8, 4, 12, 3, 11, 10, 7, 4, 12, 5, 8, 10, 7, 11, 12, 2, 10, 9, 11, 5, 6, 7, 9, 8, 12, 6, 9, 5, 6, 4, 10, 1},
	{5, 9, 12, 4, 10, 12, 9, 2, 10, 8, 6, 11, 8, 4, 12, 11, 7, 3, 12, 11, 7, 5, 11, 9, 10, 1, 5, 6, 7, 10, 8, 4, 6},
	{11, 9, 10, 12, 3, 9, 6, 1, 7, 10, 8, 4, 11, 2, 10, 6, 4, 5, 11, 12, 8, 4, 7, 9, 8, 10, 12, 11, 5, 6, 12, 7, 5},
	{12, 11, 6, 4, 2, 6, 9, 11, 10, 8, 3, 12, 4, 8, 9, 4, 12, 1, 9, 7, 10, 5, 11, 8, 7, 10, 12, 7, 11, 5, 6, 10, 5},
	{11, 7, 8, 3, 5, 10, 9, 6, 4, 8, 5, 6, 10, 4, 12, 9, 11, 4, 6, 1, 8, 10, 12, 11, 7, 12, 9, 11, 7, 10, 5, 2, 12},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [33, 33, 32, 33, 33], total reshuffles 37949472
// symbols: 93.696(lined) + 0(scatter) = 93.696323%
// cards bonuses: frequency 1/150.55, rtp = 27.036261%
// RTP = 93.696(sym) + 27.036(cards) = 120.732584%
// *regular reels calculations*
// reels lengths [33, 33, 32, 33, 33], total reshuffles 37949472
// symbols: 43.17(lined) + 15.592(scatter) = 58.761964%
// free spins 3843855, q = 0.101289
// free games frequency: 1/150.55
// cards bonuses: frequency 1/150.55, rtp = 27.036261%
// RTP = 58.762(sym) + 27.036(cards) + 0.10129*120.73(fg) = 98.027078%
var Reels98 = slot.Reels5x{
	{11, 8, 4, 12, 3, 11, 10, 7, 4, 12, 5, 8, 10, 7, 11, 12, 2, 10, 9, 11, 5, 6, 7, 9, 8, 12, 6, 9, 5, 6, 4, 10, 1},
	{5, 9, 12, 4, 10, 12, 9, 2, 10, 8, 6, 11, 8, 4, 12, 11, 7, 3, 12, 11, 7, 5, 11, 9, 10, 1, 5, 6, 7, 10, 8, 4, 6},
	{11, 5, 10, 4, 1, 12, 7, 10, 9, 12, 8, 6, 5, 7, 11, 12, 7, 3, 4, 9, 11, 6, 8, 4, 9, 8, 11, 12, 6, 2, 5, 10},
	{12, 11, 6, 4, 2, 6, 9, 11, 10, 8, 3, 12, 4, 8, 9, 4, 12, 1, 9, 7, 10, 5, 11, 8, 7, 10, 12, 7, 11, 5, 6, 10, 5},
	{11, 7, 8, 3, 5, 10, 9, 6, 4, 8, 5, 6, 10, 4, 12, 9, 11, 4, 6, 1, 8, 10, 12, 11, 7, 12, 9, 11, 7, 10, 5, 2, 12},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [32, 33, 33, 33, 32], total reshuffles 36799488
// symbols: 94.104(lined) + 0(scatter) = 94.103932%
// cards bonuses: frequency 1/147.92, rtp = 27.516801%
// RTP = 94.104(sym) + 27.517(cards) = 121.620733%
// *regular reels calculations*
// reels lengths [32, 33, 33, 33, 32], total reshuffles 36799488
// symbols: 43.285(lined) + 15.787(scatter) = 59.071259%
// free spins 3794040, q = 0.103100
// free games frequency: 1/147.92
// cards bonuses: frequency 1/147.92, rtp = 27.516801%
// RTP = 59.071(sym) + 27.517(cards) + 0.1031*121.62(fg) = 99.127201%
var Reels99 = slot.Reels5x{
	{8, 11, 12, 8, 9, 7, 12, 5, 10, 1, 5, 7, 4, 6, 11, 12, 9, 3, 8, 4, 11, 6, 7, 9, 4, 6, 10, 5, 11, 2, 10, 12},
	{5, 9, 12, 4, 10, 12, 9, 2, 10, 8, 6, 11, 8, 4, 12, 11, 7, 3, 12, 11, 7, 5, 11, 9, 10, 1, 5, 6, 7, 10, 8, 4, 6},
	{11, 9, 10, 12, 3, 9, 6, 1, 7, 10, 8, 4, 11, 2, 10, 6, 4, 5, 11, 12, 8, 4, 7, 9, 8, 10, 12, 11, 5, 6, 12, 7, 5},
	{12, 11, 6, 4, 2, 6, 9, 11, 10, 8, 3, 12, 4, 8, 9, 4, 12, 1, 9, 7, 10, 5, 11, 8, 7, 10, 12, 7, 11, 5, 6, 10, 5},
	{12, 4, 9, 6, 7, 3, 10, 9, 11, 5, 8, 10, 11, 5, 4, 7, 8, 6, 12, 11, 2, 6, 7, 12, 5, 8, 12, 1, 9, 11, 4, 10},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [31, 31, 32, 31, 31], total reshuffles 29552672
// symbols: 105.81(lined) + 0(scatter) = 105.810229%
// cards bonuses: frequency 1/130.55, rtp = 31.177841%
// RTP = 105.81(sym) + 31.178(cards) = 136.988070%
// *regular reels calculations*
// reels lengths [31, 31, 32, 31, 31], total reshuffles 29552672
// symbols: 47.745(lined) + 17.243(scatter) = 64.987284%
// free spins 3455055, q = 0.116912
// free games frequency: 1/130.55
// cards bonuses: frequency 1/130.55, rtp = 31.177841%
// RTP = 64.987(sym) + 31.178(cards) + 0.11691*136.99(fg) = 112.180642%
var Reels112 = slot.Reels5x{
	{8, 9, 6, 4, 5, 8, 7, 11, 12, 5, 6, 9, 10, 7, 4, 8, 10, 12, 11, 6, 4, 12, 1, 7, 12, 11, 2, 10, 9, 5, 3},
	{10, 8, 6, 12, 8, 4, 9, 7, 12, 5, 4, 7, 3, 11, 5, 8, 12, 1, 10, 6, 11, 2, 4, 11, 6, 10, 9, 7, 12, 9, 5},
	{11, 7, 6, 8, 12, 9, 7, 11, 12, 5, 11, 6, 4, 3, 10, 12, 8, 1, 9, 4, 12, 5, 8, 10, 4, 9, 2, 5, 11, 10, 7, 6},
	{2, 5, 12, 9, 7, 8, 11, 4, 12, 11, 6, 12, 4, 5, 10, 8, 7, 5, 1, 8, 4, 10, 3, 7, 10, 9, 6, 11, 12, 6, 9},
	{6, 4, 9, 8, 10, 5, 12, 7, 11, 4, 12, 3, 6, 12, 9, 11, 5, 10, 8, 6, 12, 10, 8, 11, 7, 4, 1, 7, 9, 2, 5},
}

// *bonus games calculations*
// total = 64, E = 40.703125
// *bonus reels calculations*
// reels lengths [30, 30, 30, 30, 30], total reshuffles 24300000
// symbols: 118.13(lined) + 0(scatter) = 118.131790%
// cards bonuses: frequency 1/116.82, rtp = 34.841875%
// RTP = 118.13(sym) + 34.842(cards) = 152.973665%
// *regular reels calculations*
// reels lengths [30, 30, 30, 30, 30], total reshuffles 24300000
// symbols: 52.418(lined) + 18.66(scatter) = 71.077901%
// free spins 3177225, q = 0.130750
// free games frequency: 1/116.82
// cards bonuses: frequency 1/116.82, rtp = 34.841875%
// RTP = 71.078(sym) + 34.842(cards) + 0.13075*152.97(fg) = 125.921083%
var Reels126 = slot.Reels5x{
	{9, 4, 5, 2, 7, 10, 11, 12, 4, 6, 1, 4, 12, 7, 6, 11, 5, 9, 10, 8, 6, 3, 7, 10, 8, 11, 9, 12, 8, 5},
	{11, 12, 6, 9, 7, 8, 9, 10, 7, 6, 12, 10, 11, 4, 9, 10, 5, 6, 2, 8, 4, 3, 5, 12, 7, 4, 11, 5, 1, 8},
	{10, 6, 12, 3, 4, 8, 6, 5, 7, 11, 1, 7, 6, 8, 5, 7, 11, 2, 8, 10, 4, 5, 9, 12, 10, 9, 11, 12, 4, 9},
	{8, 5, 11, 10, 1, 11, 7, 6, 5, 4, 11, 9, 3, 12, 4, 10, 7, 2, 8, 9, 4, 8, 9, 12, 6, 7, 12, 5, 6, 10},
	{10, 4, 1, 12, 11, 8, 6, 10, 5, 9, 6, 11, 7, 3, 8, 5, 6, 8, 4, 9, 12, 7, 5, 9, 10, 2, 12, 11, 4, 7},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	84.943404:  &Reels85,
	86.666500:  &Reels87,
	88.934078:  &Reels89,
	90.368933:  &Reels90,
	91.321872:  &Reels91,
	92.255956:  &Reels92,
	93.722785:  &Reels94,
	94.735834:  &Reels95,
	96.295851:  &Reels96,
	98.027078:  &Reels98,
	99.127201:  &Reels99,
	112.180642: &Reels112,
	125.921083: &Reels126,
}
