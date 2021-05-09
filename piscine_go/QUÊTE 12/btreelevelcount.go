package piscine

type TreeNode struct {
	Right, Left, Parent *TreeNode
	Data                string
}

func BTreeLevelCount(root *TreeNode) int {

	CountLeft := 0
	CountRight := 0

	if root == nil {
		return 0
	}
	if root != nil {
		CountRight = BTreeLevelCount(root.Right)
		CountRight++
		CountLeft = BTreeLevelCount(root.Left)
		CountLeft++
	}
	if CountLeft > CountRight {
		return CountLeft
	}
	return CountRight
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil { //Les enfants n'ont aucune identité avant qu'on les définisse donc le root est égale à nil. ce n'est pas root en lui même
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root

}
