function getAcceleration(objet) {/* f est un objet donnée avec toutes les valeurs sauf l'accélération qu'on doit calculer*/
    if(typeof(objet.f) == 'number' && typeof(objet.m) == 'number') {
        return objet.f / objet.m
    } else if(typeof(objet.Δv) == 'number' && typeof(objet.Δt) == 'number') {
        return objet.Δv / objet.Δt
    } else if(typeof(objet.d) == 'number' && typeof(objet.t) == 'number') {
        return (2 * objet.d) / (objet.t * objet.t)
    } else {
        return "impossible"
    }
}