#include "header.h"

void Result(int, int, float, unsigned long long, int);
void facto(int, int, float);


void addition() {//après avoir rentré le nombre de chiffres que l'on veut additionner, on additionne des chiffres
    int c{};
    float nbr;
    float somme{};
    int HowMany;

    cout << "\nAdditionner des chiffres, et entrez un chiffre dont vous voulez la factorielle\n" << endl; 
   

    cout << "Combien de chiffres voulez-vous additionner? (dans une limite de 50 chiffres et 2 chiffres minimum)" <<endl;
    cin >> HowMany;

    while (cin.fail()) {//gestion des erreurs
        cin.clear();
        cin.ignore(1, '\n');//supprime la valeur précédente
        cout << "\nIl me faut un nombre entier:";
        cin >> HowMany;
    }
    cout << "\n" << endl;

    if(HowMany < 2 || HowMany > 50) {
        cout << "Valeur invalide!!" << endl;
        addition();
    }

    cout << "Entrez " << HowMany << " valeurs les unes apres les autres:" << endl;

    while (c < HowMany){//on fait la somme des chiffres
        cin >> nbr;
        somme += nbr;
        c += 1;
    }

    facto(c, HowMany, somme);
 



}

void facto(int c, int HowMany, float somme){//factoriel d'un chiffre
    unsigned long long Fact = 1;
    float nbrFactoF;
    int nbrFacto;

    cout << "\nQuel est le chiffre dont vous voulez la factorielle (20 maximum) : ";
    cin >> nbrFactoF;

    while (cin.fail()) {
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nEntrez un nombre entier:";
        cin >> nbrFactoF;
    }

    nbrFacto = round(nbrFactoF);//on arrondit pour sécuriser. Si on entre un float, le resultat est faussé 

    if (nbrFacto < 0 || nbrFacto > 20) {
        cout << "Valeur invalide!!" << endl;
        facto(c, HowMany, somme);
    }

    for (int i = 1; i <= nbrFacto; ++i) {//calcul factoriel
        Fact *= i;

    }

    Result(c, HowMany, somme, Fact, nbrFacto);


}

void Result(int c, int HowMany, float somme, unsigned long long Fact, int nbrFacto) {

    if (c == HowMany) {
        cout << "\nla somme de ces " << HowMany << " chiffres est: " << somme << endl;
        cout << "\nla factorielle de " << nbrFacto << " est: " << Fact << "\n\n" << endl;
    }

}

