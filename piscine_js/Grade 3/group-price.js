function groupPrice(str) {
    var result = /(USD)\d*\.\d*|[$]\d*\.\d*/g;
  
    var regex1 = /\d*(?=\.)/;
    var regex2 = /\d*$/;
    var match = str.match(result);
    var arr = [];
    
    if (match !== null) {
      for (let item of match) {
          let arrayquifonctionne = []
        arrayquifonctionne.push(item);
        arrayquifonctionne.push(regex1.exec(item)[0]);
        arrayquifonctionne.push(regex2.exec(item)[0]);
        arr.push(arrayquifonctionne);
      }
    }
    return arr;
  }