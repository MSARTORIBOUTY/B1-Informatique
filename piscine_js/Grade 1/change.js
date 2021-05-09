/*const sourceObject = {
  num: 42,
  bool: true,
  str: 'some text',
  log: console.log,
}*/


function get(n) {
    return sourceObject[n]
}

function set(n, m) {/*prend une clé et une valeur*/
    sourceObject[n] = m;/*on met à jours là valeur*/
    return sourceObject[n]/*on retourne l'update*/
}