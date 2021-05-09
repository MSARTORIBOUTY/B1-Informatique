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

//On remplace l'élément correspondant au node de root par rplc
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}

	if root.Data > node.Data {
		root.left = BTreeTransplant(root.left, node, rplc)

	} else if root.Data < node.Data {
		root.right = BTreeTransplant(root.right, node, rplc)

	} else if root.Data == node.Data {
		root.Data = rplc.Data
		return root
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

/*note = on ne peut jamais dire : root != elem, si je comprends bien car sinon
ça fait une erreur au niveau de la mémoire. De même que un if dans un if ou un if dans u for
je pense que ça bloque. Ce sont des suppositions, on en est pas sûr*/
