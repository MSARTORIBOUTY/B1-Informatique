function letterSpaceNumber(str) {
    var result = str.match(/. \d((?=\W)|$)/g);// |$ avant ou après charactère, ?= c'est pour pas prendre en compte les (), sinon le reste c'est pour sélectionner les char qu'on veut

    if (result) {
       return result;//on retourne le résultat...
   }
   return []//... en tant que tableau
}