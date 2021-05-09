#include "header.h"
#include <time.h>


void Player(int, int);


void Albert(int nbr, int win) {//Fonction de l'ordinateur qui choisit combien d'allumettes enlever en fonction de combien il en reste / computer's function who choose how many matches it must remove 
    int take;
    int reste;


    while (nbr > 0 && win == 0) {

            reste = nbr % 4;

            if(reste == 0) {
                take = 3;

            } else if( reste ==3) {
                take = 2;

            } else {
                take = 1;

            }

            cout << "J'en prends " << take << endl;
            nbr -= take;

            if(nbr == 0) {

                cout << "Zut! J'ai perdu! \nFelicitaiton pour votre victoire! ^o^/ \n\n" << endl;
                win =1;
                main();

            } else {
                cout << "il reste " << nbr << " allumette(s)" << endl;
                Player(nbr, win);
            }



    }

}

void Player(int nbr, int win) {//Fonction du joueur / player's function

    int removeA;
    float removeP;

    while(nbr > 0 && win == 0) {

        cout << "Combien voulez-vous enlever d'allumettes?" << endl;//how many matches do you want to remove
        cin >> removeP;

        while (cin.fail()) {
            cin.clear();
            cin.ignore(1, '\n');
            cout << "\nVous ne pouvez enlever que 1, 2 ou 3 allumettes:";
            cin >> removeP;
        }

        removeA = round(removeP);

        if (removeA < 1 || removeA > 3) {

            cout << "vous n'avez le droit de n'enlever que 1, 2 ou 3 allumettes! Ce n'est pas bien d'essayer de tricher ;)\n" << endl;
            Player(nbr, win);
        }
        

        nbr -= removeA;

        if(nbr == 0) {
            cout << "Vous avez perdu! Dommage!\n\n" << endl;
            win =1;
            main();


        } else {
            cout<< "Il reste " << nbr << " allumettes" << endl;
            Albert(nbr, win);

        }
    }

}


void allumettes() {//définit le nombre d'allumettes mises en jeu / define with how many matches the player will play
    cout << "\n -- jeu des allumettes --  \nDeterminer le nombre d'allumettes, retire 1, 2 ou 3 allumettes quand c'est a votre tour de jouer! \nAttention! Vous prenez la derniere, vous perdez ;)\n\n";

    int nbr;
    float nbrArrondi;
    int randm;
    int win = 0;
    srand((unsigned) time(0));

    cout << "Avec combien d'allumettes voulez-vous jouer?" << endl;
    cin >> nbrArrondi;

    while (cin.fail()) {
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nIl me faut un nombre entier:";
        cin >> nbrArrondi;
    }

    nbr = round(nbrArrondi);

    if (nbr < 10 || nbr > 50) {//entre 10 et 50 allumettes possible / between 10 to 50 matches

        cout << "vous avez le choix entre 10 à 50 allumettes ;)\n" << endl;
        allumettes();
    }

    cout << "\nNous avons donc " << nbr << " allumettes\n" << endl;

    randm = rand()%2+1; //chiffre random pour savoir qui commence / a random number to decide who will begin

    if(randm == 1) {//1 - le joueur commence / 1 - player begins
        cout << "Vous commencez!" << endl;
        Player(nbr, win);

    } else if (randm == 2) { //2 - l'ordi commence / 2 - computer begins
        cout << "Je commence!" << endl;
        Albert(nbr, win);

    }

}


