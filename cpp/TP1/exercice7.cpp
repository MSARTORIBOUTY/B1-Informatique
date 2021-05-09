#include "header.h"
#include <cmath>


void integrales() {// calcul des intégrales / integral calculation

    cout << "\nEntrez les valeurs de a, b et le pas pour calculer l'air sous la courbe d'equation y = x^2\n" << endl;

    double a;
    double b;
    double c;
    double p;
    double integrale{};

    cout << "Entrez la valeur de a: ";
    cin >> a;
    //gestion erreur
    while (cin.fail()) {
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nIl me faut un nombre entier:";
        cin >> a;
    }

    cout << "Entrez la valeur de b: ";
    cin >> b;
    //gestion erreur
    while (cin.fail()) {
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nIl me faut un nombre entier:";
        cin >> b;
    }

    cout << "Entrez la valeur de p: ";
    cin >> p;
    //gestion erreur
    while (cin.fail()) {
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nIl me faut un nombre entier:";
        cin >> p;
    }

    //gestion erreur si les input ne respectent pas certaines conditions
    if (b <= a || p <= 0 || a < 0 || b < 1) {
        cout << "Valeur incorrect! \nVous devez respecter certaines conditions:\nb > a, p > 0, a >= 0 et b >= 1\n\n  ";
        integrales();
    }

    c=a;

    while(a < b) { //calcul de l'intégrale

        integrale += (pow((a+p),3)/3) - (pow(a,3)/3);
        a += p;
    };

    cout << "\nCalcul de l'integrale de la fonction y = x^2 \n avec " << c << " < x < " << b << " et p = " << p << endl;
    cout << "Resultat : " << round(integrale*100)/100 << "\n" << endl;

}
