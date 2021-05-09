function nasa(n) {
    let str = "";

    //si n est divisible par 3 on ajout NA
    for(let i = 1;i<=n;i++){

        if(i%3 ===0 && i%5 ===0){
            str += "NASA" + " "

        }else if(i%3 === 0 && i%5 != 0){
            str += "NA" + " "

        }else if(i%3 != 0 && i%5 === 0){
            str += "SA" + " " 

        }else{
            str += String(i) + " "
        }   
    }
    str = str.slice(0, str.length-1)//pour enlever le denrier espace à la fin de la string
    return str


}