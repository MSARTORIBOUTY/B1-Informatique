function pyramid(str, nbr) {
    let result = "";
    let largeur = (2 * nbr) -1;
    let milieu = Math.floor(largeur / 2);

    for(let i= 0; i < nbr; i++) {
        for(let j = 0; j < largeur; j++) {
            if(j < milieu-i) {
                result += ' ';
            } else if( j > milieu+i) {
                
            }else if( i === 1) {
                result += str;
            }else {
                result += str;
            }
        }
        if(i === nbr) {
           return result
        }
        result += '\n'

    }
    result = result.slice(0, result.length - 1)
    return result
}
module.exports = pyramid;