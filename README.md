# go-code-problem-day3

## 題目描述

給定 Binary Tree 的根節點 root, 找出此 Binary Tree 的最大寬度。

Binary Tree 的寬度定義為: 在某個 tree level 中，最左邊的節點至最右邊的節點的距離
若中間包含空節點 (null node) 則必須將空節點也納入距離的計算
Constraints:

節點值均為整數
最大寬度不會超過 32-bit integer 的可表示範圍
Example:

Input: binary tree =

    Level 0         9 (root)
                   / \
    Level 1       8   7
                 /   / \
    Level 2     6   5   4
                     \
    Level 3           3

Output: 4
Explanation: Level 2 的節點為 [6, null, 5, 4]，其寬度為 4 且是整個 Binary Tree 中最寬的

## 題目分析

給定一個 一個BinaryTree root 想要找到最大的width

首先可以使用 BFT(Breadth First Transerval)來對 整個Tree作遍歷

因為每次都是同一層先找 因此每次計算 每個層的width 然後只紀錄最大的那個

跑完整個tree 結果就能出來了

## refDoc

[BFT](https://ithelp.ithome.com.tw/articles/10204111)