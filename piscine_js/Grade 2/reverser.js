function reverse(obj) {
    //let result="";
    let result = [];

    for(let i = obj.length - 1; i >= 0; i--) {//on balaye le contenue de la variable à l'envers
        if(typeof(obj)== 'string') {
            result += obj[i];//on ajoute le caractère de la string à la suite du résultat
            
        }
        
        if(Array.isArray(obj)) {
            result.push(obj[i]);//on ajoute chaque valeur du tableau sépparément
            
        }
        
        
    }
    return result

}