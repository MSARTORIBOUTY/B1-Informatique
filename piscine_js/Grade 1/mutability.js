const clone1 = Object.assign({}, person);/*copy*/
const clone2 = Object.assign({}, person);

const samePerson = person;

samePerson.age = 78;
samePerson.country = 'FR';