// La méthode slice() renvoie une copie d'une partie d'un array dans un nouveau 
// sélectionné du début à la fin (par les index). Le array original ne sera pas modifié.
//start et end sont des intervalles

const slice = (contenu, start, end) => {

    let result = [];

    if (start === undefined) {//si la valur par défaut est indéfinit, on lui donbne la valeur de 0
        start = 0;
    }

    if (start < 0) { 
        start = contenu.length + start // on soustrait la valeur de start à la longueur de la string
    }

    if (end === undefined) {
        end = contenu.length
    }

    if (end < 0) {
        end = contenu.length + end
    }

    for (let i = start; i < end; i++) { // s'il y a une fin
        result.push(contenu[i]) // ajoute un nouvel élément à la fin du tableau
    }

    if (typeof contenu === 'string') {
        // crée et renvoie une nouvelle chaîne de caractères en concaténant tous les éléments d'un tableau
        return result.join('') 
    }

    return result
}
