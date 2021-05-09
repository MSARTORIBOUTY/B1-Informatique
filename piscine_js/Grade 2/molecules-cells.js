function RNA(n) {

    let str ="";

    for(let i = 0; i < n.length; i++) {
        if(n[i] === "A") {
            str += "U"

        } else if(n[i] == "T") {
            str += "A"

        } else if(n[i] == "G") {
            str += "C"

        } else {
            str += "G"
        }
        
    }
    return str

}

function DNA(n) {
    let str ="";

    for(let i = 0; i < n.length; i++) {
        if(n[i] === "A") {
            str += "T"

        } else if(n[i] == "U") {
            str += "A"

        } else if(n[i] == "G") {
            str += "C"

        } else  {
            str += "G"

        } 
        
    }
    return str

}