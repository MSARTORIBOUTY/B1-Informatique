
/*let est limité à la fonction, on ne pourra pas utiliser la varaible déclaré avec
dans une autre fonction*/
const concatStr = (n, m) => {
    let string = String(n) + String(m)/*on les mets bout à bout (pas d'addition) et on cast en string pour que, même si ce sont des chiffres ils se mettent bout à bout*/
    return string
}