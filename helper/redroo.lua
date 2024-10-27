local path = arg[0]:match("(.*[/\\])")
dofile(path.."lib/reelgen.lua")

local symset = {
	1, --  1 wild
	1, --  2 scatter
	5, --  3 redroo
	4, --  4 girl
	4, --  5 boy
	4, --  6 dog
	4, --  7 parrot
	4, --  8 ace
	4, --  9 king
	4, -- 10 queen
	4, -- 11 jack
	4, -- 12 ten
	5, -- 13 nine
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,
	{ 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,}, --  1
	{ 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  2
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  3
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  4
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  5
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  6
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  7
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  8
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, --  9
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, -- 10
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, -- 11
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, -- 12
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,}, -- 13
}

math.randomseed(os.time())
local reel, iter = makereel(symset, neighbours)
printreel(reel, iter)
