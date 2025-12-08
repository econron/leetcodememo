package main

// フロイドの循環検出法を使用した解法
// https://leetcode.com/problems/linked-list-cycle-ii/description/
type ListNode struct {
    Val int
    Next *ListNode
 }

 func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    fast := head
    slow := head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            // headと出会った地点それぞれから1ステップずつ進むと出会える
            start := head
            for start != slow {
                start = start.Next
                slow = slow.Next
            }
            return slow
        }
    }
    return nil
}

// num of islands
// gridを回しながら島が見つかったら上下左右を探索して海にする
// mainのforで1つ1つ探索していく。0なら探索しない。1なら探索して上下左右を探索して海にする。
func numIslands(grid [][]byte) int {
    rows := len(grid)
    columns := len(grid[0])
    islandCount := 0
    for r := 0; r < rows; r++ {
        for c := 0; c < columns; c++ {
            if grid[r][c] == '1' {
                islandCount++
                dfs(grid, r, c)
            }
        }
    }
    return islandCount
}

func dfs(grid [][]byte, r,c int) {
    rows := len(grid)
    columns := len(grid[0])
    if r < 0 || c < 0 || r >= rows || c >= columns {
        return
    }
    if grid[r][c] == '0' {
        return
    }
    grid[r][c] = '0'
    dfs(grid, r+1, c) // 下
    dfs(grid, r-1, c) // 上
    dfs(grid, r, c+1) // 右
    dfs(grid, r, c-1) // 左
}