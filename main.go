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

// bfs的にmergeTreeを解く
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 * Val int
 * Left *TreeNode
 * Right *TreeNode
 * }
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func mergeTreesbyBFS(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    // 例外処理: どちらかが空なら、もう片方を返すだけで終わり
    if root1 == nil {
        return root2
    }
    if root2 == nil {
        return root1
    }

    // キューの定義 (スライスとして定義します)
    // root1をベースにして、そこにroot2の値を足していく戦略です
    queue1 := []*TreeNode{root1}
    queue2 := []*TreeNode{root2}

    for len(queue1) > 0 {
        // Pop (先頭を取り出す)
        node1 := queue1[0]
        queue1 = queue1[1:]
        
        node2 := queue2[0]
        queue2 = queue2[1:]

        // 値を足し合わせる (node1を更新)
        node1.Val += node2.Val

        // --- 左の子ノードの処理 ---
        if node1.Left == nil && node2.Left != nil {
            // node1になく、node2にある -> node2のサブツリーをそのまま移植
            node1.Left = node2.Left
        } else if node1.Left != nil && node2.Left != nil {
            // 両方にある -> 加算処理が必要なのでキューに追加して後で処理
            queue1 = append(queue1, node1.Left)
            queue2 = append(queue2, node2.Left)
        }
        // (注: node1にあってnode2にない場合は、何もしなくてOK)
        // つまるところnode1に寄せていく

        // --- 右の子ノードの処理 ---
        if node1.Right == nil && node2.Right != nil {
            node1.Right = node2.Right
        } else if node1.Right != nil && node2.Right != nil {
            queue1 = append(queue1, node1.Right)
            queue2 = append(queue2, node2.Right)
        }
    }

    return root1
}

func mergeTreesbyDFS(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil {
        return root2
    }
    if root2 == nil {
        return root1
    }
    root1.Val += root2.Val
    root1.Left = mergeTreesbyDFS(root1.Left, root2.Left)
    root1.Right = mergeTreesbyDFS(root1.Right, root2.Right)
    
    return root1
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    // ベースケース：両方の葉がもうない
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }
    newSum := targetSum - root.Val
    return hasPathSum(root.Left, newSum) || hasPathSum(root.Right, newSum)
}

type Minheap struct {
    nodes []int
}

func (h *Minheap) parentIndex(i int) int {
    return (i - 1)/2
}

func (h *Minheap) leftIndex(i int) int {
    return 2*i + 1
}

func (h *Minheap) rightIndex(i int) int {
    return 2*i + 2
}

func (h *Minheap) Add(num int) {
    h.nodes = append(h.nodes, num)
    h.upHeap(len(h.nodes)-1)
}

func (h *Minheap) upHeap(i int) {
    for i > 0 {
        p := h.parentIndex(i)
        if h.nodes[p] > h.nodes[i] {
            h.nodes[i], h.nodes[p] = h.nodes[p], h.nodes[i]
            i = p
        } else {
            break
        }
    }
}

func (h *Minheap) Pop(num int) int {
    h.nodes = append(h.nodes, num)
    h.nodes[0], h.nodes[len(h.nodes)-1] = h.nodes[len(h.nodes)-1], h.nodes[0]
    h.downHeap(0)
    return h.nodes[0]
}

func (h *Minheap) downHeap(i int) {
    for i <= len(h.nodes) - 1 {
        left := h.leftIndex(i)
        right := h.rightIndex(i)
        smallest := i
        if right < len(h.nodes) && h.nodes[right] < h.nodes[i] {
            smallest = right
        }
        if left < len(h.nodes) && h.nodes[left] < h.nodes[i] {
            smallest = left
        }
        if smallest == i {
            break
        }
        h.nodes[i], h.nodes[smallest] = h.nodes[smallest], h.nodes[i]
        i = smallest
    }
}

func twoSum(nums []int, target int) []int {
    visited := make(map[int]int)
    for i,num := range nums {
        diff := target - num
        if idx, ok := visited[diff]; ok {
            return []int{idx,i}
        }
        visited[num] = i
    }
    return nil
}

func main() {
	fmt.Println(minSteps(5, 17))
}