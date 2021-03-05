package solution

// TreeNode is definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
    if root == nil {
        return res
    }
    var q []*TreeNode
    q = append(q, root)

    for len(q) > 0 {
        n := len(q)
        average := 0
        for i := 1; i <= n; i++ {
            node := q[0]
            average += node.Val
            q = q[1:]
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        res = append(res, float64(average)/float64(n))
    }
    return res
}
