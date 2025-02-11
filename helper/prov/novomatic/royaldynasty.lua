local path = arg[0]:match("(.*[/\\])")
dofile(path.."../../lib/reelgen.lua")

local symset = {
	1, --  1 oldking  5000
	2, --  2 lord     500
	3, --  3 gryphon  300
	3, --  4 leon     300
	3, --  5 money    200
	3, --  6 shield   200
	3, --  7 ace      150
	3, --  8 king     150
	3, --  9 queen    100
	3, -- 10 jack     100
	3, -- 11 ten      100
	4, -- 12 nine     100
	1, -- 13 princess
	1, -- 14 prince
}

local neighbours = {
	--1, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13,14,
	{ 2, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 2, 2,}, --  1 oldking
	{ 1, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,}, --  2 lord
	{ 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,}, --  3 gryphon
	{ 1, 1, 1, 2, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,}, --  4 leon
	{ 1, 1, 1, 1, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0,}, --  5 money
	{ 1, 1, 1, 1, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0,}, --  6 shield
	{ 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0,}, --  7 ace
	{ 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0,}, --  8 king
	{ 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0,}, --  9 queen
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0,}, -- 10 jack
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0,}, -- 11 ten
	{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0,}, -- 12 nine
	{ 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2,}, -- 13 princess
	{ 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2,}, -- 14 prince
}

math.randomseed(os.time())
printreel(makereel(symset, neighbours))
