function ionOut(str) {
    var result = /([a-z]*t(?=(ion)))/g; // match le t qui précède le ion dans toutes les lettres
  
    if (str.match(result) != null) {//si il y a bien un match

      return str.match(result);//tu retourne la string par rapport au result 
    }

    return [];

  }