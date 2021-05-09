#include "header.h"


int maxfact(int);
int fact(int);

void factmax(){//donne la factorielle la plus proche d'un nombre donné /gives the closest factorial to a given number
    int k;
    float floatk;
    int n;

    cout << "Calcul de n tel que n! <= k\n \nEntrez la valeur de k: ";
    cin >> floatk;

    while (cin.fail()) {
        cin.clear();
        cin.ignore(1, '\n');
        cout << "\nIl me faut un nombre entier:";
        cin >> floatk;
    }

    k = round(floatk);

    if (k < 1) {

        cout << "k doit etre superieur ou egal a 1\n" << endl;
        factmax();
    }

    cout << "Resultat n! <= k : " << maxfact(k) << "! <= " << k << "\n" << endl;

}

int fact(int n){//calcul factoriel / factorial calculation

    int factorielle=1;
    int i;
    for(i=1;i<=n;i++)
        factorielle*=i;
    return factorielle;
}

int maxfact(int k){// définit la factorielle la plus proche / define the closest factorial

    int n=0;
    while(fact(n)<=k)
    {
        n++;
    }
    return n-1;
}
