package structures

type TreeNode[T comparable] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func Ints2TreeNode(slice []int) *TreeNode[int] {
	n := len(slice)
	if n == 0 {
		return nil
	}

	root := &TreeNode[int]{
		Val: slice[0],
	}

	queue := make([]*TreeNode[int], 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n {
			node.Left = &TreeNode[int]{Val: slice[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n {
			node.Right = &TreeNode[int]{Val: slice[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
