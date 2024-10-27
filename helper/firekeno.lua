local path = arg[0]:match("(.*[/\\])")
dofile(path.."lib/keno.lua")

-- Fire Keno
local paytable = {
--        0      1      2      3      4      5      6      7      8      9     10  WIN BALLS
    {     0,     0,     0,     0,     0,     0,     0,     0,     0,     0,     0,}, --  1 sel
    {     1,     0,     6,     0,     0,     0,     0,     0,     0,     0,     0,}, --  2 sel
    {     1,     0,     1,    26,     0,     0,     0,     0,     0,     0,     0,}, --  3 sel
    {     1,     0,     0,     7,   100,     0,     0,     0,     0,     0,     0,}, --  4 sel
    {     1,     0,     0,     1,    15,   666,     0,     0,     0,     0,     0,}, --  5 sel
    {     2,     0,     0,     1,     2,    25,  2500,     0,     0,     0,     0,}, --  6 sel
    {     3,     0,     0,     0,     3,    10,   100, 10000,     0,     0,     0,}, --  7 sel
    {     5,     0,     0,     0,     1,     4,    46,   666, 25000,     0,     0,}, --  8 sel
    {    10,     0,     0,     0,     0,     3,    10,   100,  1000, 50000,     0,}, --  9 sel
    {    15,     0,     0,     0,     0,     1,     3,    25,   666,  1000,100000,}, -- 10 sel
}

kenoprobtable(10, 6)
kenortp(paytable, 6)
