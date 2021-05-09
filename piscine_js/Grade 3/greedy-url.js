function getURL(data) {
    var result = /(https?:\/\/[a-zA-Z0-9]+[^\s]{2,})/g;
    return data.match(result);
}

function greedyQuery(data) {
    var get = getURL(data);
    var arr = [];
    var result = /([^=]*[=]){3,}/g;//pour séparer dès que un caractère est entouré de =

    for (let item of get) {//ajout de l'élément dans le tableau. Item = ensemble de la variable matchable 
        if (item.match(result) !== null) {
          arr.push(item);
      }
    }
      return arr;
    }

function notSoGreedy(data) {
        let get = getURL(data);
        let arr = [];
        var result = /[=]/g;
        for (let item of get) {
          if ( item.match(result) !== null &&//si y a un match
            item.match(result).length >= 2 &&//et que la longueur est sup ou = à 2 paramètres
            item.match(result).length <= 3) //et inf ou égale à 3
            {
            arr.push(item);
          }
        }
        return arr;

}