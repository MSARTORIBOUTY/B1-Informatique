package piscine

type TreeNode struct {
	Right, Left, Parent *TreeNode
	Data                string
}

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
	f(root.Data)
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
