#include "header.h"

int fibonacci(int);

void fibonacciCalcul() {

	cout << " \n -- Suite de Fibonnaci --\nEntrez la valeur de N pour calculer la suite de Fibonacci!";

	int N;
	float n;

	cout << "\n\nEntrez la valeur de N : ";
	cin >> n;

	while (cin.fail()) {//s�curisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> n;
	}

	N = round(n);

	if (N < 0 || N > 150) {//limite de la valeur N

		cout << "\n\n!!! N doit �tre compris entre 0 et 150 inclus! \n\n";
		fibonacciCalcul();
	}

	cout << "Voici la suite de Fibonacci:" << endl;

	for (int i = 0; i <= N; i++) {//le nombre de fois que la site tourne

		cout << fibonacci(i) << endl;
	}
	cout << "\nFin de la suite Fibonacci! \n" << endl;
}

int fibonacci(int n) {//fibonacci en r�cursive

	if (n <= 1) {
		return n;
	}
	else {

		return (fibonacci(n - 1) + fibonacci(n - 2));
	}
}