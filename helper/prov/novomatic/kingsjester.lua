local path = arg[0]:match("(.*[/\\])")
dofile(path.."../../lib/reelgen.lua")

local symset = {
	1, --  1 double      1000
	2, --  2 jester      500
	1, --  3 funny king  500
	1, --  4 funny queen 500
	2, --  5 cards       350
	3, --  6 bandura     250
	3, --  7 pan flute   250
	3, --  8 ace         125
	3, --  9 king        125
	4, -- 10 queen       100
	4, -- 11 jack        100
	4, -- 12 ten         100
	4, -- 13 nine        100
	1, -- 14 scatter
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,14,
	{ 2, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 2,}, --  1 double
	{ 2, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 2,}, --  2 jester
	{ 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,}, --  3 funny king
	{ 1, 1, 1, 2, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,}, --  4 funny queen
	{ 1, 1, 1, 1, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0,}, --  5 cards
	{ 1, 1, 1, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0,}, --  6 bandura
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0,}, --  7 pan flute
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0,}, --  8 ace
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0,}, --  9 king
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0,}, -- 10 queen
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0,}, -- 11 jack
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0,}, -- 12 ten
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0,}, -- 13 nine
	{ 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,}, -- 14 scatter
}

math.randomseed(os.time())
printreel(makereel(symset, neighbours))
