function addWeek(date) {
    //let TwoWeeks = 1000 * 60 * 60 * 24 * 14;//pour transformer une semaine en 14j, on multiplie le nombre de ms, s, m, h et j.
    //let WeekDays = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "saturday", "sunday"]

    //return WeekDays[date.getDay() * TwoWeeks/7]



    let day0 = new Date('0001-01-01')//la date 1 qui donne le premier j de la semaine
    function diffDates(day_one, day_two) {
        //on soustrait le jour donné en argument et le jour 1 
        return ((day_one-day_two) / (60 * 60 * 24 * 1000))/7//divisé par les ms d'une journée, divisé par 7j (nbr j dans une semaine)
    }    
    //obtenir le jour de la semaine
    function getWeekDay(date) {
        let days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
        //on obtient le jour de la date donné en argument dans la toute première fonction;
        return days[date.getDay()];
    }
    let diff = diffDates(date,day0)%2 // on prend le reste de la div avec 2 (pr 2 semaines)
    if(diff%2===0||diff<1){//0 étant la première semaine
        return getWeekDay(date)// on retourne le tableau avec le bon jour 
    }
    else {
        let day = getWeekDay(date)
        return 'second'+day//sinon on rajoute second avant
    }
    
}


function timeTravel({date,hour,minute,second}){
    date.setHours(hour)
    date.setMinutes(minute)
    date.setSeconds(second)
    return date
}