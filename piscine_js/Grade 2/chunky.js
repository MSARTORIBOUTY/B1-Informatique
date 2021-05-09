function chunk(arrI, sep) {
    let result= [];


    for(let i = 0; i < arrI.length; i += sep) {//on incrémente selon la longueur
        result.push(arrI.slice(i, i + sep))//incrémente l'index juqu'à la longueur sep où on veut couper le tableau en 2
   
    }
    
    return result
}