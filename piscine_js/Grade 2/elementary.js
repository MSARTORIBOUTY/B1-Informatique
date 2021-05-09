const multiply = (a, b) => {

    let result = 0;
    if (b > 0) {
      for (let i = 0; i < b; i++) {
        result += a; // permet d'incrémenter le résultat par lui-même
      }
  
    } else {
      for (let i = b; i < 0; i++) {
        result -= a; // permet de décrémenter le résultat par lui-même
      }
    }
    return result;
  }
  
  const divide = (a, b) => {
    let result = 0
    
    // on implémente grâce à un bool = s'il y a encore des chiffres, on continue la division
    let count = false
  
    // chaque possibilité des dividende et diviseur :
    if (a < 0 && b > 0) {
        count = true 
        a = -a
    }
  
    if (a == b) {
        return 1
    }
  
    if (a > 0 && b < 0) {
        count = true
        b = -b
    }
  
    if (a < 0 && b < 0) {
        a = -a
        b = -b
    }
  
    if (a < b) {
        return 0
    }
  
    if (b == 0) { // la division est impossible
        return
    }
  
    while (a >= b) { 
        a = a - b
        result++
    }
  
    if (count) { // on remet le signe 
        result = -result
    }
  
    return result
  
  }
  
  const modulo = (a, b) => {
  
      return a - multiply(divide(a, b), b);
  }