function repeat(str, num) { //repeat est un loop, donc on utilise while)
    //return str.repeat(num)
    //dans les fichiers test, si num == 0, str ne print rien, donc num est > 0. num indique le nbr de fois que str est répété.
    //on doit donc ajouter str à une nouvelle variable de type str, qui rajouetras à chaque fois le str. 
    //on décrémente à chaque tour pour avoir le bon nbr de fois, la str répété
    let result = "";

    while(num > 0) {
        result = result += str;
        num--
    }
    return result
}
