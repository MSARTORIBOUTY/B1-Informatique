package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args
	if len(args) < 1 {
		z01.PrintRune('\n')
	}

	str := ""
	for i, val := range args[1:] {
		if i+1 == len(args[1:]) { //Si i fait toute la longueur de l'args
			str += string(val) //On ajoute le caractère à la string vide
		} else {
			str += val + " " //on met un espace après l'argument. Donc entre chaque args
			//Ici val équivaut à l'argument et non au mot ou à chaque lettres
		}
	}
	strVowels := ""
	for _, char := range str { //Chaque caractère de la string qui continent chaques lettres séparement pas comme le premier for
		if Vowels(string(char)) { //Si le charactère est une voyelle
			strVowels += string(char) //On ajoute cette voyelle à la string vide de voyelle
		}

	}
	LenStrVowels := 0
	for range strVowels { //On balaye toutes les voyelles une par une de la string en question
		LenStrVowels++
	}

	ArrVowels := make([]int, LenStrVowels) //On place chaques voyelles quand un tableau
	j := 0
	for i, char := range str {
		if Vowels(string(char)) { //Si les char de la string sont des voyelles
			ArrVowels[j] = i //si l'index i est équivalent à un index du tableau de voyelles
			j++              //On incrémente j

		}
	}
	strRune := []rune(str)

	for i := range strVowels {
		strRune[ArrVowels[LenStrVowels-1-i]] = rune(strVowels[i])
		//*note
	}
	Print(string(strRune)) //On peut pas print avec PrintRune un tableau de rune

	/**note: StrVowels[i] détermine les voyelles à l'index i dans le tableau. Et on donne cette valeur à la voyelle contenue dans
	le tableau de rune qui possède les caractères des arguments. En gros il va permettre d'inverser les voyelles là où il y en a dans l'argument.*/

}
func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
func Vowels(s string) bool {

	if s == "a" || s == "A" {
		return true
	}
	if s == "e" || s == "E" {
		return true

	}
	if s == "i" || s == "I" {
		return true
	}
	if s == "o" || s == "O" {
		return true

	}
	if s == "u" || s == "U" {
		return true
	}
	return false
}

/*args := os.Args[1:]

vowels := []rune{}
arr := make([]int, len(args))
str := ""
tableStr := []rune(str)

if len(args) < 1 { //Si la longueur de l'argument est inférieur à 1 carcatère
	z01.PrintRune('\n') //Print uniquement un retour à la ligne
	return
}

lenargs := 0
for _, char := range args {
	str += string(char)
	lenargs++
}

for i, lettre := range tableStr {
	if Vowels(string(lettre)) {
		vowels = append(vowels, lettre)
		arr = append(arr, i)
	}
}
//fmt.Println(vowels)
//fmt.Println(arr)

lenVowels := 0
for range vowels {
	lenVowels++
}

for i, char := range arr {
	tableStr[char] = vowels[lenVowels-i-1]
}
for _, charStr := range tableStr {
	z01.PrintRune(charStr)
}*/
