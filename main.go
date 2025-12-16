package main

import "fmt"

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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    // うさぎと亀によるサイクルの終了エッジケースを先に排除する
    if head == nil || head.Next == nil {
        return false
    }
    // kameが3にいたらusagiは2つ先にいる
    kame := head
    usagi := head.Next.Next
    // もしusagiがこの時点でnilならforループを抜けてfalseになる
    for kame != nil && usagi != nil {
        // kameの値とusagiの値が等しい時点でサイクルが存在することを示す
        if kame.Val == usagi.Val {
            return true
        }
        kame = kame.Next
        usagi = usagi.Next.Next
    }
    return false
    
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

// bfsの特訓
// 与えられた変数N、到達したい変数Mがあるとする（N < M）。操作は+1,-1,*2のみとする。
// 最短でMに到達する手数を返す関数を作る。

// N < M については確定しているとして内部処理では考慮しないことにする

type Item struct {
	Value int
	Steps int
}
func minSteps(N, M int) int {
	queue := []Item{{Value: N, Steps: 0}}
	visited := make(map[int]bool)
	visited[N] = true
	for len(queue) > 0 {
		// 先頭要素を取得
		item := queue[0]
		// 先頭要素を削除
		queue = queue[1:]
		// 取得した値がMと一致していたらその時のステップ数を返す
		if item.Value == M {
			return item.Steps
		}
		// 取得した値がMより大きい場合は探索を打ち切る
		if item.Value > M {
			continue
		}
		// 取得した値がMより小さい場合は探索を続ける
		if item.Value < M {
			nextValues := []int{
				item.Value + 1,
				item.Value - 1,
				item.Value * 2,
			}
			for _, nextValue := range nextValues {
				if _, ok := visited[nextValue]; !ok {
					visited[nextValue] = true
					queue = append(queue, Item{Value: nextValue, Steps: item.Steps + 1})
				}
			}
		}
	}
	return -1
}



func main() {
	fmt.Println(minSteps(5, 17))
}