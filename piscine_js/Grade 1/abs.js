function isPositive(number) {
    if(number > 0) {
        return true
    } else {
        return false
    }
}
function abs(number) {
    if(number < 0 ) {
        return -number;
    }else if(number == 0) {/*c'est une égualité faible donc ==, car il va comparer deux chiffres*/
        return 0;
    } else {
        return number;
    }
}