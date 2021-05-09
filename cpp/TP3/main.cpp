#include "header.h"

int main() {
    int answer;
    while (true) {


        cout << "Bonjour et bienvenu dans les exercices du TP3! :) \nQuel exercice voulez vous voir? (1-6, 0 pour quitter) \n\n";
        cout << "1 - Jeu de carte : A \n2 - Jeu de carte : B \n3 - Jeu de carte YU-GI-OH (1) : MONSTRE \n4 - Jeu de carte YU-GI-OH (2) : MAGIE et PIEGE\n5 - Jeu de carte YU-GI-OH (3) : GENERALISE\n\nVotre reponse: ";
        cin >> answer;

        switch (answer) {

        default: main(); break;

        case 0:   break; // Quitter

        case 1: Carte1(); break;

        case 2: Carte2(); break;

        case 3: CarteY1(); break;

        case 4: CarteY2(); break;

        case 5: CarteY3(); break;

        }
    }
    return 0;

}