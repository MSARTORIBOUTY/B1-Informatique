//répéter une string n fois
//avec retour à la ligne chaque fois

function triangle(str, nbr) {

    let result = "";

    for(let i= 1; i <= nbr; i++) {//on print la string jusqu'à ce que i == nbr pour avoir le nbr de fois voulue la str
        for(let j = 0; j < i; j++) {//le nbr de fois qu'on print la string sur une même ligne. si i = 2, j printera 2 fois la string sur la même ligne
            result += str
        }
        //on doit avoir un retour à la ligne à la fin
        //sans cela on a rien et result +='\n' rajoute 2 '\n'. donc on retourne la string pour en avoir un seul
        if(i === nbr) {
           return result
        }
        result += '\n'//à la fin de chaque ligne on rajoute un '\n'

    }
    return result
}