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

//afficher l'arbre par niveau
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	f(root.Data)//Niveau 1, le root principal, la racine donc on l'affiche simplement avec f
	BTreeApplyByLevel(root.left, f)//On appel la fonction pour lire le niveau de gauche à droite
	BTreeApplyByLevel(root.right, f)
}
