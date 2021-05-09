package reloaded

type TreeNode struct {
	right *TreeNode
	left  *TreeNode
	Data  string
}

//On ajoute les enfants à partir de root
func BTreeInsertData(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: elem}
	}
	if root != nil {
		if root.Data < elem {
			root.right = BTreeInsertData(root.right, elem)
		} else {
			root.left = BTreeInsertData(root.left, elem)
		}
	}
	return root
}

//On cherche si un enfant ou le root est équivalent à un élément
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.left, elem)
	}
	if elem > root.Data {
		return BTreeSearchItem(root.right, elem)
	}
	return root
}

//On met dans l'ordre chaque chiffre de l'arborescence
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	BTreeApplyInorder(root.left, f)
	f(root.Data)
	BTreeApplyInorder(root.right, f)

}

//Supprimer un chiffre de l'arborescence
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if root == node {
		if root.right != nil { //Si l'enfant de droite existe
			min := BTreeMin(root.right) //On prend l'enfant de gauche de l'enfant de droite du root
			root.Data = min.Data        //Et le root prend l'identité du plus petit chiffre de droite
			node = min
		} else if root.left != nil { //même système
			max := BTreeMax(root.left)
			root.Data = max.Data
			node = max
		} else {
			return nil
		}
	}
	root.left = BTreeDeleteNode(root.left, node) //Faire la boucle pour chaque enfants
	root.right = BTreeDeleteNode(root.right, node)
	return root
}

func BTreeMin(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	if root.left == nil {
		return root
	}
	return BTreeMin(root.left)
}

func BTreeMax(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	if root.right == nil {
		return root
	}
	return BTreeMax(root.right)
}
