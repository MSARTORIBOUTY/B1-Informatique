function get(src, path) {
    let paths = path.split('.');//on sépare la chaine de caractère src à partir du mot que contient path pour juste avoir la suite
    var current = src; 

    for (let i = 0; i < paths.length; ++i) {
        if (current[paths[i]] == undefined) {
          return undefined;
        } else {
          current = current[paths[i]];//on donne a current la valeur de la clé
        }
      }
      return current;
}