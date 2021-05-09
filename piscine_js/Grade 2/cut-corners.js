//arrondit selon les règles
function round(num) {
    let arrondi = 0;
    let flag = false;

    if(num < 0) {
        num = -num;
        flag = true;
    }

    let decimal = num % 1//permet de prendre la valeur après la virgule
    
    if(decimal >= 0.5) {
        num = num + 1;
        arrondi = num - num % 1;
        
    } 
    if(decimal < 0.5) {
        arrondi = num - num % 1;
    }
    if(flag) {
        arrondi = -arrondi;
    }
    return arrondi

}

//enlève décimal et arrondi au supérieur 
function ceil(num) {


    let decimal = num % 1;
    if(decimal) {
        if(num > 0) {
            num = (num + 1) - (num % 1);
        }
        if(num < 0) {
            num = - num;
            num = num - num%1;
            num = - num;
        }

    }
    return num
}

//enlève les décimals et rrondi à l'inférieur
function floor(num) {

    let decimal = num % 1;
    if(num > 0) {
        if (decimal) {
            num = num - num % 1;//on lui enlève la décimale
        }
    }
    if(num < 0) {
        if(decimal) {
            num = -num;//on transforme num en positif
            num = num + 1 - num%1;//on enlève la décimale en rajoutant 1
            num = -num;//on lui rend sa négativité
        }
    }
    return num

}

//enlève les décimales
function trunc(num) {
    let flag = false;


    if(num < 0) {
        num = -num;
        flag = true;
    }
    num = num - num % 1;
    if(flag) {
        num = -num;
    }
    return num


}