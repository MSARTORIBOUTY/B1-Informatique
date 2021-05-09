
/*la fonction retourne le premier élément*/
const first = (n) => {
    return n[0]
}

/*la fonction retourne le dernier élément*/
const last = (m) => {
    return m[m.length - 1]
}

/*la fonction retourne le dernier et le premier élément, dans cet ordre*/
const kiss = (o) => {
    return [o[o.length - 1], o[0]]
}