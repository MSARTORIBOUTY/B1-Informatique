package main

import (
	"fmt"

	student ".."
)

//Atoi

/*func main() {
	fmt.Println(student.Atoi(""))
	fmt.Println(student.Atoi("-"))
	fmt.Println(student.Atoi("--123"))
	fmt.Println(student.Atoi("1"))
	fmt.Println(student.Atoi("-3"))
	fmt.Println(student.Atoi("8292"))
	fmt.Println(student.Atoi("9223372036854775807"))
	fmt.Println(student.Atoi("-9223372036854775808"))
}*/

//-----------------------------------------------------------------------------------------------

//Recursive power

/*func main() {
	//arg1 := -7
	//arg2 := -2
	arg1 := -8
	arg2 := -7
	//arg1 := 4
	//arg2 := 8
	//arg1 := 1
	//arg2 := 3
	//arg1 := -1
	//arg2 := 1
	//arg1 := -6
	//arg2 := 5
	fmt.Println(student.RecursivePower(arg1, arg2))
}*/

//-----------------------------------------------------------------------

//PrintCombN

/*func main() {
	student.PrintCombN(1)
	student.PrintCombN(2)
	student.PrintCombN(3)
	student.PrintCombN(4)
	student.PrintCombN(5)
	student.PrintCombN(6)
	student.PrintCombN(7)
	student.PrintCombN(8)
	student.PrintCombN(9)
}*/

//-----------------------------------------------------------------

//PrintNbrBase

/*func main() {
	student.PrintNbrBase(919617, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	z01.PrintRune('\n')
	student.PrintNbrBase(-861737, "Zone01Zone01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "-ab")
	z01.PrintRune('\n')
	student.PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')
}*/

//---------------------------------------------------------------

//AtoiBase

/*func main() {
	fmt.Println(student.AtoiBase("bcbbbbaab", " abc"))
	fmt.Println(student.AtoiBase("0001", "01"))
	fmt.Println(student.AtoiBase("00", "01"))
	fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))
	fmt.Println(student.AtoiBase("125", "0123456789"))
	fmt.Println(student.AtoiBase("uoi", "choumi"))
	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
}*/

//------------------------------------------------------------------

//SplitWhitSpaces

/*func main() {

	fmt.Println(student.SplitWhiteSpaces("The earliest foundations of what would become computer science predate the invention of the modern digital computer"))
	fmt.Println(student.SplitWhiteSpaces("Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"))
	fmt.Println(student.SplitWhiteSpaces("aiding in computations such as multiplication and division ."))
	fmt.Println(student.SplitWhiteSpaces("Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."))
	fmt.Println(student.SplitWhiteSpaces("Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"))
	fmt.Println(student.SplitWhiteSpaces("In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"))
}*/

//---------------------------------------------------------------------------------

//Split

/*func main() {
	//s := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished, and charset = |=choumi=|"
	//s := "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator, and charset = !==!"
	//s := "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator, and charset = AFJ"
	//s := "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer, and charset = <<==123==>>"
	fmt.Println(student.Split(s, "<<== ==>"))
}*/

//---------------------------------------------------------------------

//ConvertBase

/*func main() {
	//result := student.ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	//result := student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	//result := student.ConvertBase("256850", "0123456789", "01")
	//result := student.ConvertBase("uuhoumo", "choumi", "Zone01")
	result := student.ConvertBase("683241", "0123456789", "0123456789")
	fmt.Println(result)
}*/

//-------------------------------------------------------------------------------

//AdvancedSortWordArr

/*func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result, student.Compare)

	fmt.Println(result)
}*/

//---------------------------------------------------------------------------------

//ActiveBits

/*func main() {
	//nbits := student.ActiveBits(15)
	//nbits := student.ActiveBits(17)
	//nbits := student.ActiveBits(4)
	//nbits := student.ActiveBits(11)
	//nbits := student.ActiveBits(9)
	//nbits := student.ActiveBits(12)
	nbits := student.ActiveBits(2)
	fmt.Println(nbits)
}*/

//------------------------------------------------------------------------------------

//SortListInsert

/*func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *student.NodeI

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = student.SortListInsert(link, -2)
	link = student.SortListInsert(link, 2)
	PrintList(link)
}*/

//------------------------------------------------------------------------------------------------

//SortedListMerge

/*func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *student.NodeI
	var link2 *student.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(student.SortedListMerge(link2, link))
}*/

//--------------------------------------------------------------------------------------------

//ListRemoveIf

/*func PrintList(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &student.List{}
	link2 := &student.List{}

	fmt.Println("----normal state----")
	student.ListPushBack(link2, 1)
	PrintList(link2)
	student.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "Hello")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "There")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "How")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "are")
	student.ListPushBack(link, "you")
	student.ListPushBack(link, 1)
	PrintList(link)

	student.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}*/

//----------------------------------------------------------------------------------------------

//BTreeTransplant

/*func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "1")
	replacement := &student.TreeNode{Data: "3"}
	root = student.BTreeTransplant(root, node, replacement)
	fmt.Println(root, node)
	student.BTreeApplyInorder(root, fmt.Println)
}*/

//---------------------------------------------------------------------------------------------

//BTreeApplyByLevel

/*func main() {
	root := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root, "07")
	student.BTreeInsertData(root, "12")
	student.BTreeInsertData(root, "10")
	student.BTreeInsertData(root, "05")
	student.BTreeInsertData(root, "02")
	student.BTreeInsertData(root, "03")
	student.BTreeApplyByLevel(root, fmt.Print)
}*/

//-----------------------------------------------------------------------------------------------------

//BTreeDeleteNode

func main() {
	root := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root, "07")
	student.BTreeInsertData(root, "12")
	student.BTreeInsertData(root, "10")
	student.BTreeInsertData(root, "05")
	student.BTreeInsertData(root, "02")
	student.BTreeInsertData(root, "03")
	node := student.BTreeSearchItem(root, "02")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root, fmt.Println)
}
