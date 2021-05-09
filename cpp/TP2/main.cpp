#include "header.h"

int main() {
    int answer;
    while (true) {


        cout << "Bonjour et bienvenu dans les exercices du TP2! :) \nQuel exercice voulez vous voir? (1-6, 0 pour quitter) \n\n";
        cout << "1 - Reference/pointeur: A \n2 - Reference/pointeur: B \n3 - Macro \n4 - Syracus\n5 - Fibonacci\n6 - Identite d'une personne: les classes\n\nVotre reponse: ";
        cin >> answer;

        switch (answer) {

        default: main(); break;

        case 0: break; // Quitter

        case 1: C_2_6b(); break;

        case 2: C_2_6c(); break;

        case 3: macro();  break;

        case 4: syracuse(); break;

        case 5: fibonacciCalcul();  break;

        case 6: Init();  break;


        }
    }
    return 0;

}