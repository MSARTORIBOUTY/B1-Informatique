function firstDayWeek(week, year) {
    let first = new Date(year + '-01-01' ); // année en premier
    if (week != 1) { // si ce n'est pas la première semaine, on y retourne
        first.setDate(first.getDate() + ((week-1) * 7) - first.getDay() + 1) 
        // pour obtenir le lundi de la première semaine en avançant et reculant dans le temps par mois
    }
    // on inverse le résultat (qui est en iso)
    // on les divise au niveau des - (car pas tous en ont)
    // on rajoute donc les - 
    // on retire le Z et le T grâce à leurs index 0 et 10
    return first.toISOString().slice(0, 10).split('-').reverse().join('-')
    //iso donne la date en format yyyy-mm-dd T hh:mm:ss Z
    //.slice permet de retirer T et Z. 0 = Z et 10 = T sont comme des index
    //on sépare avec - et non / 
    //et on met à l'envers pour avoir l'année en premier
}