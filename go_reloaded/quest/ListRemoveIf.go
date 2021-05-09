package reloaded

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

/*func ListPushBack(link *List, data int) {
	n := &NodeL{Data: data}

	if link == nil {
		return
	}
	iterator := link
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return
}*/

//Insérer char dans la liste
func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = node
		l.Tail = node

		return
	}
	l.Tail.Next = node
	l.Tail = node

}
func ListRemoveIf(l *List, data_ref interface{}) {

	item := l.Head

	if l.Head == nil { //si la tête de liste est vide, la liste est vide
		return //Du coup on retourne rien
	}

	//Il sert à supp le data_ref en tête de liste. Sans lui il reste affiché
	if l.Head.Data == data_ref { //Si le chiffre == le chiffre qu'on veut supp
		l.Head = l.Head.Next //le chiffre prend la valeur du chiffre suivant

	}

	//supp autre data_ref dans la liste
	for item.Next != nil { //Si la liste n'est pas nul
		if item.Next.Data == data_ref { //Si le chiffre == au chiffre que l'on veut supp
			item.Next = item.Next.Next //le chiffre prend la valeur du suivant
			ListRemoveIf(l, data_ref)  //On réitère le truc jusqu'à ce qu'il n'y ait plus de data_ref
			return
		}

		item = item.Next
	}
}
