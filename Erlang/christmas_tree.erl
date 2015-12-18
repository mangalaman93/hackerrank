% O(n) algorith, can be improved further

-module(christmas_tree).
-export([main/0]).

main() ->
    {ok, [T]} = io:fread("", "~d"),
    testCase(T).

testCase(0) ->
    ok;
testCase(T) ->
    {ok, [N]} = io:fread("", "~d"),
    Tree = constructTree(N-1, array:new(N, {default, []})),
    {ok, [Q]} = io:fread("", "~d"),
    readAndFindLCAs(Q, Tree, array:new(N, {default, false})),
    testCase(T-1).

constructTree(0, Tree) ->
    Tree;
constructTree(N, Tree) ->
    {ok, [A, B]} = io:fread("", "~d ~d"),
    T1 = array:set(A, [B|array:get(A, Tree)], Tree),
    T2 = array:set(B, [A|array:get(B, Tree)], T1),
    constructTree(N-1, T2).

readAndFindLCAs(0, _, _) ->
    ok;
readAndFindLCAs(N, NbrTree, ColorTree) ->
    {ok, [A]} = io:fread("", "~s"),
    {ok, [B]} = io:fread("", "~d"),
    Val = case A of
        "+" ->
            true;
        "-" ->
            false
    end,
    NewColorTree = array:set(B, Val, ColorTree),
    io:format("~p~n", [findLCA(-1, 0, NbrTree, NewColorTree)]),
    readAndFindLCAs(N-1, NbrTree, NewColorTree).

% returns -1 if no colored node is present
% otherwise returns lca of the tree rooted at RootNode
findLCA(LastNode, RootNode, NbrTree, ColorTree) ->
    case array:get(RootNode, ColorTree) of
        true ->
            RootNode;
        false ->
            case lists:map(fun(E) -> findLCA(RootNode, E, NbrTree, ColorTree) end,
                    lists:delete(LastNode, array:get(RootNode, NbrTree))) of
                [] ->
                    -1;
                LCAs ->
                    checkLCAs(LCAs, -1, RootNode)
            end
    end.

checkLCAs([], Node, _) ->
    Node;
checkLCAs([-1|T], N, RootNode) ->
    checkLCAs(T, N, RootNode);
checkLCAs([X|T], -1, RootNode) ->
    checkLCAs(T, X, RootNode);
checkLCAs([_X|_T], _N, RootNode) ->
    RootNode.
