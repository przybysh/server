local path = arg[0]:match("(.*[/\\])")
dofile(path.."lib/keno.lua")

-- Keno Centurion
local paytable = {
--       0     1     2     3     4     5     6     7     8     9    10  WIN BALLS
    {    0,    0,    0,    0,    0,    0,    0,    0,    0,    0,    0,}, --  1 sel
    {    0,    1,   10,    0,    0,    0,    0,    0,    0,    0,    0,}, --  2 sel
    {    0,    0,    5,   20,    0,    0,    0,    0,    0,    0,    0,}, --  3 sel
    {    0,    0,    2,   10,   40,    0,    0,    0,    0,    0,    0,}, --  4 sel
    {    0,    0,    1,    5,   20,   75,    0,    0,    0,    0,    0,}, --  5 sel
    {    0,    0,    1,    3,    5,   40,  150,    0,    0,    0,    0,}, --  6 sel
    {    0,    0,    0,    3,    5,   15,   80,  300,    0,    0,    0,}, --  7 sel
    {    0,    0,    0,    1,    5,   15,   24,  150,  500,    0,    0,}, --  8 sel
    {    0,    0,    0,    1,    3,    5,   30,   80,  300, 1000,    0,}, --  9 sel
    {    0,    0,    0,    0,    3,    5,   15,   50,  180,  500, 2000,}, -- 10 sel
}

kenoprobtable(10, 6)
kenortp(paytable, 6)
