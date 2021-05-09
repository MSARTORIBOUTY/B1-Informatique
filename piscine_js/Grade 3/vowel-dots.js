var vowels = /[aeiouAEIOU]/g;///.../g permet de faire une comparaison globale

const vowelDots = (str) => 

    str.replace(vowels, '$&.');// $ pour ajouter à la fin et & car c'est statique sinon c'est à la fin de chaque voyelles

