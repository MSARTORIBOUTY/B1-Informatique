package reloaded

type NodeI struct {
	Next *NodeI
	Data int
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {

	n := &NodeI{}
	n.Data = data_ref
	n.Next = nil

	if l == nil || l.Data >= n.Data {
		n.Next = l //si la liste est vide ça retourna forcément rien.On met cette ligne car on prend les deux conditions. Si l == nil et qu'on retourn n, il retourne la valeur de data_ref
		//mais si l>data_ref, n.Next prend la valeur de la liste
		return n //n prend la valeur du data_ref
	} else {
		temp := l                                         //il sert à prendre les données de la liste, ça aide pour situer le chiffre.
		for temp.Next != nil && temp.Next.Data < n.Data { //Si la suite de la liste est =/= nil et que cette suite est inf au chiffre que l'on veut insérer (data-ref)
			temp = temp.Next //Alors la liste prend la valeur de la suite de la liste, pour la décaller et ainsi insérer le data-ref
		}
		n.Next = temp.Next //Les chiffres de la list sup au data_ref son ajoutés
		temp.Next = n      //Affiche le chiffre dans la liste
	}
	return l //affiche la liste

}
