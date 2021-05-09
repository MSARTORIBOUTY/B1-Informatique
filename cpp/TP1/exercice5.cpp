#include "header.h"
#include <string>
#include <time.h>
#include <stdio.h>
#include <stdlib.h>


void jeu421() {

    cout << " \n\n-- le jeu du 421 -- \nTrois coups, trois des, essayez d'obtenir la combinaison 421\n\n";

    int de1 = 0;
    int de2 = 0;
    int de3 = 0;
    int coups = 0;
    vector<int> de(3, 0);
    int arr[3] = {0, 0, 0};

    srand(time(NULL));//remise à zero du nombre random / reset random number 
    const int MAX = 6, MIN = 1;

        while (coups <= 3)// tant que le nombre de coup est inf à 3 
        {
            coups += 1;
            de1 = (rand() % (MAX - MIN + 1)) + MIN;//chaque dé prend une valeur comprise entre 1 et 6
            de2 = (rand() % (MAX - MIN + 1)) + MIN;//each die takes a value between 1 to 6
            de3 = (rand() % (MAX - MIN + 1)) + MIN;

            de[0] = de1;//on met cette valeur dans un tableau / we put this value in an array
            de[1] = de2;
            de[2] = de3;

            cout << "\nVoici votre jeu : " << de1 << de2 << de3 << endl;

            cout << "Quelle(s) valeur(s) voulez-vous garder ? " << endl;

            for(int i = 0; i < 3; i++){// si les valeurs du tableau sont equivalent à 1, 2, 4/ if the array values are equal to 1, 2 or 4

                if (de[i] == 1) {
                    if (arr[2] == 1) {
                        cout << "J'ai deja un 1..." << endl;// already has this number
                    }else {
                        cout << "Je conserve le 1" << endl;//keep this number
                        arr[2] = 1;//chaque nombre a sa place définit dans un tableau / each number has his own place in the array
                    }
                }

                if (de[i] == 2) {
                    if (arr[1] == 2) {
                            cout << "J'ai deja un 2..." << endl;
                    }else {
                        cout << "Je conserve le 2" << endl;
                        arr[1] = 2;

                    }
                }

                if (de[i] == 4) {
                    if (arr[0] == 4) {
                        cout << "J'ai deja un 4..." << endl;
                    }else {
                        cout << "Je conserve le 4" << endl;
                        arr[0] = 4;

                    }
                }

                if(de1 == de2 == de3){
                cout << "WOW ! Impressionnant ! \^o^/ " << endl;

                }
                if(arr[2] == 1 && arr[1] == 2 && arr[0] == 4){//si on obtient un 421 / if we have a 421
                    cout << arr[0] << arr[1] << arr[2] << endl;
                    cout << "BRAVO !! Vous avez fait un 421 en " << coups << " coups ! \^o^/" << endl;
                    break;

                }

            }
            if(arr[2] == 1 && arr[1] == 2 && arr[0] == 4){
                cout << "le jeu est termine!" << endl;
                break;
            }else if (coups == 3 ) {
                cout << "\nOh non ! Vous avez perdu car vous avez atteint le nombre de lances maximum...\n" << endl;
                break;


        }

    }
}
