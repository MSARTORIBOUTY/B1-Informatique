#include "header.h"
#include <time.h>


void randomNum() {//génère un nombre aléatoire / generate a random number

    int random;
    int nbr;
    float nbrInterval;
    int v2;
    int answer;
    float nbrAnswer;
    int countNbr = 0;
    srand((unsigned) time(0));

    cout << "\nEntrez un nombre entier qui definira votre intervalle :";
    cin >> nbrInterval;

    while (cin.fail()) {
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nIl me faut un nombre entier:";
        cin >> nbrInterval;
    }

    nbr = round(nbrInterval);

    for (int index = -nbr; index <= nbr; index++) {//génère un nombre aléatoire... / generate a random number
        v2 = -nbr + rand() % nbr;
    }

    random = rand()%2+1;//...puis un deuxième entre 1 et 2 / and another between 1 and 2

    if (random == 1) {//si il tombe sur 1, il multiplie par -1 le nombre mystère / if it's 1, then it multiply the mysterynumber by -1

        v2 *= -1;
    } else {//Sinon il n'y touche pas // otherwhise it doesn't touch it 
        v2 = v2;
    }


    cout << "\nVous devez trouver un nombre compris entre " << -nbr << " et " << nbr << "\nEssayez de deviner!\nNe rentrez que des entiers! jouez le jeu ;)" << endl;
    cout << "--> ";
    cin >> nbrAnswer;// variable pour trouver le nombre mystère / variable to find the mystery number 

    while (cin.fail()) {
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nIl me faut un nombre entier:";
        cin >> nbrAnswer;
    }

    answer = round(nbrAnswer);

    while (answer != v2) {//si le chiffre est différent / if the input is diferent
        if (answer < v2) {//si il est plus petit / if it is smaller

            cout << "trop petit!\n->";
            cin >> nbrAnswer;
            answer = round(nbrAnswer);
            countNbr += 1;

        } else if(answer > v2) {//si il est plus grand / if it is bigger

            cout << "trop grand!\n->";
            cin >> nbrAnswer;
            answer = round(nbrAnswer);
            countNbr += 1;
        }

    }

    if (answer == v2) {
        cout << "Bravo! le chiffre mystere etait bien " << v2 << "! \n Vous avez gagne(e)! \^o^/ \n Vous avez reussis en " << countNbr << " coup(s)\n\n";



    }


}
