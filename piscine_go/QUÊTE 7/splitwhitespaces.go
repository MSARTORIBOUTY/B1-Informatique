package piscine

func SplitWhiteSpaces(s string) []string {

	lettres := ""        //on créé une vraiable vide de type 'chaîne de charactère'
	var tableau []string // on créé un tableau vide

	for i, char := range s { //On cherche s dans le range pour obtenir tout ce qu'il contient et le lister
		if char != ' ' && char != '\t' && char != '\n' { //si le charactère est != d'un espace, tabulation et saut à la ligne (donc si c'est une lettre, une ponctuation, un chiffre,...)
			lettres += string(s[i]) // on ajoute le caractère s[i] à la variable lettres

		} else if lettres != "" { //si c'est un espace, une tabulation ou un saut à la ligne
			tableau = append(tableau, lettres) //On ajoute (=append) le mot contenue dans lettres au tableau
			lettres = ""                       //on réinitialise enuite lettres pour qu'il redevienne une chaîne de caractère vide pour créer un nouveau mot
		}
	}

	tableau = append(tableau, lettres) //Pour bien ajouter le dernier mot de s au tableau sinon le dernier mot sera manquant car il sera juste ajouter à lettres et le programme sera terminer avant l'ajouter au tableau
	return tableau
}
