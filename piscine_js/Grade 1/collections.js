
    /*arrToSet from Array to Set
    arrToStr from Array to String
    setToArr from Set to Array
    setToStr from Set to String
    strToArr from String to Array
    strToSet from String to Set
    mapToObj from Map to Object
    objToArr from Object to Array
    objToMap from Object to Map
    arrToObj from Array to Object
    strToObj from String to Object*/
function arrToSet(arr) {
    var set = new Set;
    for( let contenu of arr) {
        set.add(contenu)
    }
    return set
}
function arrToStr(arr) {
    var str = "";
    for(let contenu of arr) {
        str += contenu
    }
    return str
}

function setToArr(set) {
    var arr = Array.from(set);
    return arr
}

function setToStr(set) {
    var str= "";
    for (let contenu of set) {
        str += contenu;
    } 
    return str
}
function strToArr(str) {
    var arr = Array.from(str);
    return arr
}
function strToSet(str) {
    var set = new Set;
    for(let i = 0; i < str.length; i++) {
        set.add(str[i])
    }
    return set
}
function mapToObj(map) {
    var obj = {};
    for (var [key, values] of map) {
        obj[key] = values;
    } 
    return obj
}
function objToArr(obj) {
    let arr = [];
    for (let item in obj) {
        /* push() ajoute un ou plusieurs éléments à la fin d'un tableau 
        et retourne la nouvelle taille du tableau */
        arr.push(obj[item]); 
    }
    return arr
}

function objToMap(obj) {
    var map = new Map();
    for (let contenu in obj) {
        map.set(contenu, obj[contenu]) // Définit la valeur d'une clé pour Map
    }
    return map
}
function arrToObj(arr) {
    var obj = {};
    for(let i = 0; i < arr.length; i++) {
        obj[i] = arr[i];
    }
    return obj
}
function strToObj(str) {
    var obj = {};
    for(let i = 0; i < str.length; i++) {
        obj[i] = str[i];
    }
    return obj
}

function superTypeOf(variable) {
    switch (typeof variable) {
        case "number":
          return "Number";

        case "string":
          return "String";

        case "boolean":
          return "Boolean";

        case "undefined":
          return "undefined";

        case "object":
    // L'opérateur instanceof permet de tester si un objet possède la propriété voulue.
          if (variable instanceof Array) {
            return "Array";
          }

          if (variable instanceof Set) {
            return "Set";
          }

          if (variable instanceof Map) {
            return "Map";
          }

          if (variable === null) {
            return "null";
          }

          return "Object";

        case "function":
          return "Function";
        
        default: // si rien ne correspond au cas du dessus
          return typeof variable
    }  

}