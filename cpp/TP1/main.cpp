#include "header.h"

int main() {
    int answer;
    while (true) {

        cout << "Bonjour et bienvenu dans les exercices du TP1! :) \nQuel exercice voulez vous voir? (1-8, 0 pour quitter) \n\n";
        cout << "1 - Saisie et boucle \n2 - Somme 20 premiers entiers et factorielle \n3 - Conversion binaire et decimal \n4 - Nombre Mystere\n5 - Jeu du 421\n6 - Jeu des allumettes\n7 - Les integrales\n8 - Factorielle maximale\n\nVotre reponse: ";
        cin >> answer;

        switch (answer) {

        default: break;

        case 0: return 0; break; // Quitter

        case 1: Debut(); break;

        case 2: addition(); break;

        case 3: choose(); break;

        case 4: randomNum(); break;

        case 5: jeu421(); break;

        case 6: allumettes(); break;

        case 7: integrales(); break;

        case 8: factmax(); break;
        }
    }
    return 0;

}
