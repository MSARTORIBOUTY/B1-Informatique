#include "header.h"


int BinaryToDecimal() {//conversion binaire en décimal / transform binary to decimal

    int decimalNumber = 0;
    string n;
    int i = 1;
    int y; 

    cout << "\nSaisissez un nombre binaire : ";
    cin >> n;

    for (int j = 0; j < n.size(); j++)//gestion de l'erreur si on entre un chiffre décimal / if it's not a binary number
        if (n[j] > '1') {
            cout << "Malheureusement, ce n'est pas un nombre binaire :/ \n" << endl;
            BinaryToDecimal();

        }

    y = stoi(n);//cast en int / transform string (n) into int


    while (y) {//conversion 
        int lastDigit = y % 10;
        y = y / 10;
        decimalNumber += lastDigit * i ;
        i = i * 2;
   }
    cout << "La forme decimal de " << n << " est " << decimalNumber << "\n" << endl;

    return 0;

}

int DecimalToBinary() {//conversion décimal en binaire / transform decimal to binary
    long long m;
    int j = 0;
    int binaryNumber[100];
    int y;

    cout << "\nSaisissez un nombre decimal : ";
    cin >> m;

    while (cin.fail()) {//gestion des erreurs
        cin.clear();
        cin.ignore(1, '\n');//supprime la valeur précédente
        cout << "\nIl me faut un nombre entier:";
        cin >> m;
    }

    y = m;

    while (m > 0) {
            binaryNumber[j] = m % 2;
            m = m / 2;
            j++;
   }
   cout<<"La forme binaire de " << y << " est ";
   for (int p = j - 1; p >= 0; p--) {
        cout << binaryNumber[p];
   }

   cout<< "\n" << endl;
    return 0;

}

void choose() {//réponse du choix de la conversion 

    int answer;

    cout << "\n\nVoulez-vous convertir:\n 1- un nombre binaire en decimal \n 2- un nombre decimal en binaire \n\nEntrez votre choix : ";
    cin >> answer;

    while (cin.fail()) {//gestion des erreurs
        cin.clear();
        cin.ignore(1, '\n');//supprime la valeur précédente
        cout << "\n (1) ou (2):";
        cin >> answer;
    }

    if(answer == 1) {
        BinaryToDecimal();

    } else if (answer == 2) {
        DecimalToBinary();

    } else {
        cout << "Je ne comprends pas votre reponse :/" << endl;
        main();

    }

}
