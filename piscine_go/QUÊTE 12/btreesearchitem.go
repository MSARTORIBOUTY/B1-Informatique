package piscine

type TreeNode struct {
	Right, Left, Parent *TreeNode
	Data                string
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	}
	return root
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
