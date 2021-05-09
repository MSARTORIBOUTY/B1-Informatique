package piscine

type TreeNode struct {
	Right, Left, Parent *TreeNode
	Data                string
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	}
	return true
}
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil { //Les enfants n'ont aucune identité avant qu'on les définisse donc le root est égale à nil. ce n'est pas root en lui même
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root //On doit définir le root enfant comme étant le parent
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root

}
