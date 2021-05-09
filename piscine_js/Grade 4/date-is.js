function isValid(date) {

    if(isNaN(date)) {//NaN : is not a number. Donc si date est NaN,
        return false
    }
    if(date == 0) {//si il n'y a tout simplement pas de date
        return false
    }
    
    return true
}

function isAfter(date, dateToCompare) {
    if(date > dateToCompare) {//si la date est après la date à comparé,
        return true
    }
    return false
}

function isBefore(date, dateToCompare) {
    if(date < dateToCompare) {//si la date est avant la date à comparé,
        return true
    }
    return false

}

function isFuture(date) {
    var today = new Date();


    if(date > today) {//si la date est après aujourd'hui, alors c'est le future
        return true
    }
    return false
}

function isPast(date) {
    var today = new Date();
    
    if(date == 0) {
        return false 
    }
    if(date < today) {//si la date est avant aujourd'hui, alors c'est le passé
        return true
    }
    return false

}

