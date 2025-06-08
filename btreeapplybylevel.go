package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	type nodeWithLevel struct {
		node  *TreeNode
		level int
	}

	queue := []nodeWithLevel{{node: root, level: 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		f(curr.node.Data)

		if curr.node.Left != nil {
			queue = append(queue, nodeWithLevel{node: curr.node.Left, level: curr.level + 1})
		}
		if curr.node.Right != nil {
			queue = append(queue, nodeWithLevel{node: curr.node.Right, level: curr.level + 1})
		}
	}
}
