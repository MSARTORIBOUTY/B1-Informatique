#include "header.h"


int ReturnNbr() {//compare deux réels et dit quel est le plus grand des deux ou si ils sont équivalents 
    float nbr1;
    float nbr2;

    std::cout << "\nEntrez le premier Reel: ";
    std::cin >> nbr1;
    while (cin.fail()) {//Si on entre une string au lieu d'un chiffre
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nEntrez un nombre Reel:";
        cin >> nbr1;
    }
    std::cout << "Entrez le deuxieme Reel: ";
    std::cin >> nbr2;
    while (cin.fail()) {//gestion des erreurs
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nle deuxième aussi doit etre un nombre Reel:";
        cin >> nbr2;

    }



    if(nbr1 < nbr2) {
        std::cout << "\nle plus grand nombre des deux est: " << nbr2 << std::endl;
        return 0;

    } else if (nbr1 > nbr2) {
        std::cout << "\nle plus grand nombre des deux est: " << nbr1 << std::endl;
        return 0;
    } else {
        std::cout << "\nles deux nombres sont strictement identiques: " << nbr1 << " = " << nbr2 << std::endl;
        return 0;
    }

}

int Modulo() {
    int nbr;

    std::cout << "Etrez un nombre entier:";
    std::cin >> nbr;
    while (cin.fail()) {//gestion des erreurs
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nEntrez un nombre entier:" ;
        cin >> nbr;
    }

    if (nbr%2==0) {
        std::cout << "votre nombre " << nbr << " est paire" << std::endl;
        return 0;
    } else {
        std::cout << "votre nombre " << nbr << " n'est pas paire\n" << std::endl;
        return 0;
    }
    
}

void Debut() {//cette fonction prend un nom comme input qu'il va print dans une phrase // this function takes a name as an input and display it in a sentence 

    cout <<"\n\nApres avoir donner votre nom, lalgorithme vous donnera le chiffre le plus grand \n entre deux reels que vous aurez entre ainsi que si votre chiffre est paire ou impaire\n"<< endl;

    string name;

    std::cout << "Comment vous appelez-vous?" << std::endl;
    std::cin >> name;
    std::cout << "Bonjour " << name << ", bienvenu dans le premier exercice du C++! :)" << std::endl;
    ReturnNbr();//appel deux autres fonctions
    Modulo();// call two other functions

}

