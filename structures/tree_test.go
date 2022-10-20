package structures

import "testing"

func TestIntTree(t *testing.T) {
	root := Ints2TreeNode([]int{1, 2, 3, 4, 5, 6})
	t.Log(root)
}

func TestStringTree(t *testing.T) {
	root := &TreeNode[string]{}
	root.Val = "wf"
	root.Left = &TreeNode[string]{
		Val: "sd",
	}
	root.Right = &TreeNode[string]{
		Val: "sisi",
	}
	t.Log(root)

	t.Log(root.Val > root.Left.Val)
}
