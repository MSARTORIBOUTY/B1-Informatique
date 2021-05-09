function isFriday(date) {
    let day = date.getDay();

    if(day === 5) {
        return true
    }
    return false

}

function isWeekend(date) {
    let day = date.getDay();

    if(day === 6 || day === 7) {
        return true
    }
    return false
}

function isLeapYear(date) {
    let year = date.getFullYear();
    let LeapYear = year % 4;

    if (LeapYear==0) {
        return true;
    } 
    return false

}

function isLastDayOfMonth(date) {

    //date actuelle comparé avec le dernier jour du mois
    if (date.getDate() == new Date(date.getFullYear(),date.getMonth()+1,0).getDate()) {
        return true
    }
    return false

}