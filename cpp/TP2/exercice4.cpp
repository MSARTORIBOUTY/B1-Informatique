#include "header.h"


void syracuse() {

	cout << "\nSuite de Syracuse, entrez la valeur de N pour obtenir la suite\n" << endl;

	float n;
	int N;
	int count =0;


	cout << "Entrez la valeur de N :";
	cin >> n;


	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> n;
	}

	N = round(n);

	if (N < 1 || N > 50) {

		cout << "\n!!!N doit etre compris entre 1 et 50 inclus\nRecommencez!" << endl;
		syracuse();
	}

	cout << "U0 = " << N << endl;

	while (N > 1) {//tant que la suite ne converge pas à 1 on fait la boucle

		if (N % 2 == 0) {//si le résultat est paire

			N /= 2;

		}
		else {//si il ne l'est pas
			N = 3*N + 1;
			
		}
		count += 1;
		cout << "U" << count << " = " << N << endl;

	}
	cout << "\nLa suite de Syracus avec U0 = " << round(n) << " se termine a U" << count << " = " << N << "\n\n" << endl;

}
