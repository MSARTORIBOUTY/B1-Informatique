export function getArchitects(people) {
    let all = [];

    //first array --> architects corresponding <a> tag
    let architects = [];
    architects = people.getElementsByTagName("a");

    //second array --> non architects people
    let non_architects = [];
    non_architects = people.getElementsByTagName("span")

    //all 
    all.push(architects);
    all.push(non_architects)

    return all;
}


export function getClassical(people) {
    let all2 = [];

    //fist array --> architects belonging to the classical class
    let classical = [];
    classical = people.getElemntsByClassName("classical");


    //second array --> non-classical architects
    let non_classical = [];
    non_classical = people.querySelectorAll("a:not(.classical)");
    



    all2.push(classical);
    all2.push(non_classical);


    return all2
}



export function getActive(people) {
    let all3 = [];

    //first array --> classical architects active in their class
    let active = [];
    let active = people.getElemntsByClassName("true");

    //second array --> non-active classical architects 
    let non_active = [];    
    let non_active = people.querySelectorAll("calssical:not(.true)");


    all3.push(active);
    all3.push(non_active);

    return all3

}

export function getBonannoPisano(active) {
    //array contain --> HTML element of our architect, id = BonannoPisano
    let statue = classical.getElemntsById("BonannoPisano");
    let active = [];
    active.push(statue)

    //array --> remaining active classical architects
}