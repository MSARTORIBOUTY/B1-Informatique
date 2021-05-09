function sameAmount(str, regexe1, regexe2) {
    const reg1 = new RegExp(regexe1, 'g');//objet d'expression rationnelle pour la reconnaissance d'un modèle dans un texte
    const reg2 = new RegExp(regexe2, 'g');//g pour recherche globale dans la string

    const resReg1 = str.match(reg1);
    const resReg2 = str.match(reg2);

    if(resReg1 == null) {
        return false
    }
    if (resReg2 == null) {
        return false
    }
    return resReg1.length === resReg2.length

}