//remove the 2 first characters.
function cutFirst(str) {
    let result = str.substring(2);
    return result

    //on peu aussi faire str.slice(2)

}

//remove the 2 last characters.
function cutLast(str) {
    str = str.slice(0, -2); 
    return str
}


//remove the 2 first characters and 2 last characters.
function cutFirstLast(str) {
    str = str.slice(2, -2); 
    return str
}

//return the string only keeping the 2 first characters.
function keepFirst(str) {
    str = str.slice(0, 2); 
    return str
}

//return the string only keeping the 2 last characters.
function keepLast(str) {
    str = str.slice(-2); 
    return str
}

//keep 2 first characters and 2 last characters.
function keepFirstLast(str) {
    if(str.length > 3){
        str = keepFirst(str) + keepLast(str);
        return str
    }
    return str

}

