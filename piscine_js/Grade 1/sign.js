function sign(number) {
    if(number > 0) {
        return 1
    } else if(number === 0) {/* == comparaison et === égualité stricte*/
        return 0
    } else {
        return -1
    }
    
}
function sameSign(number1, number2) {
    if(sign(number1) == sign(number2)) {
        return true
    } else {
        return false
    }

}