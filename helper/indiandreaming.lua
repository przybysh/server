local path = arg[0]:match("(.*[/\\])")
dofile(path.."lib/reelgen.lua")

local symset = {
	1, --  1 wild
	1, --  2 scatter
	5, --  3 cash catcher
	5, --  4 man
	5, --  5 woman
	5, --  6 guy
	5, --  7 bull
	5, --  8 hatchet
	5, --  9 ace
	6, -- 10 king
	6, -- 11 queen
	6, -- 12 jack
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,
	{ 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, }, --  1
	{ 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, }, --  2
	{ 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, }, --  3
	{ 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, }, --  4
	{ 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, }, --  5
	{ 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, }, --  6
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, }, --  7
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, }, --  8
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, }, --  9
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, }, -- 10
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, }, -- 11
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, }, -- 12
}

math.randomseed(os.time())
local reel, iter = makereel(symset, neighbours)
printreel(reel, iter)
