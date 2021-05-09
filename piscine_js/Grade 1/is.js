is.num = (n) => typeof n === 'number';
is.nan = (n) => Number.isNaN(n);
is.str = (n) => typeof n === 'string';
is.bool = (n) => typeof n === 'boolean';
is.undef = (n) => typeof n === 'undefined';
is.def = (n) => !is.undef(n);/*différent de indéfinit car defined n'a pas de type*/
is.arr = (n) => Array.isArray(n);/*déterminer n comme type tableau*/
is.obj = (n) => typeof n === 'object' && n !== null && !Array.isArray(n);/*quand il n'est pas null et que ce n'est pas un tableau*/
is.fun = (n) => typeof n === 'function';/*c'est du type fonction*/
is.truthy = (n) =>  true && n;/*quand il est vrai*/
is.falsy = (n) => !n;/*il est différent de vrai*/