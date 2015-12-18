% Enter your code here. Read input from STDIN. Print output to STDOUT
% Your class should be named solution

-module(first).
-export([main/0]).

main() ->
	{ok, [N]} = io:fread("", "~d"),
	Map = cMap(N, maps:new()),
    {ok, [T]} = io:fread("", "~d"),
    solveT(T, Map).

cMap(0, M) ->
	M;
cMap(T, M) ->
    {ok, [Key]} = io:fread("", "~s"),
    {ok, [_]} = io:fread("", "~s"),
    {ok, [Val]} = io:fread("", "~s"),
    cMap(T-1, maps:put(Key, Val, M)).

solveT(0, _) ->
    ok;
solveT(T, M) ->
    {ok, [N]} = io:fread("", "~d"),
    solveOne(N, M),
    solveT(T-1, M).

solveOne(0, _) ->
    io:format("~n");
solveOne(N, M) ->
    {ok, [Key]} = io:fread("", "~s"),
    io:format(maps:get(Key, M)),
    io:format(" "),
    solveOne(N-1, M).
