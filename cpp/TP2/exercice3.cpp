#include "header.h"

//macro de débogage, si il n'y a aucune erreur dans le code
#if DEBUG == 0
#define TRACE() {\
cout << "Aucune erreur! tout est OK dans le fichier: "<< __FILE__ << " jusqu'a la fin du code ligne " << __LINE__ << "\n";\
}
#else//si il y a une erreur
#define TRACE() {\
std::cerr << "Erreur dans le fichier : " << __FILE__ << ", a la ligne : " << __LINE__<< "\n";\
std::cerr.flush();\
}
#endif

//macro de la valeur la plus grande, addition, soustraction, multiplication et division
#define MAXI(a, b) ((a) > (b) ? (a) : (b)) //si a > b on retourne a sinon b
#define ADD(a, b) ((a) + (b)) 
#define SOUS(a, b) ((a) - (b)) 
#define MULT(a, b) ((a) * (b)) 
#define DIV(a, b) ((a) / (b)) 

void macro() {

	int x;
	int y;

	cout << "\nEntrez deux entiers positifs, la Macro se chargera de vous indiquer le plus grand et de faires quelques calculs\n avec ces derniers ainsi que d'afficher le nom de l'exercice\net la ligne ou se termine le code" << endl;

	cout << "\nEntrez deux entiers!\nEntier a: ";
	cin >> x;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> x;
	}

	cout << "Entier b: ";
	cin >> y;

	while (cin.fail()) {
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> y;
	}

	if (x < 0 || y < 0) {
		cout << "\na et b doivent être supérieur à 0\n" << endl;
		macro();
	}

	//les variables prennent la valeur du resultat de la Macro
	int result = MAXI(x, y);
	int add = ADD(x, y);
	unsigned sous = SOUS(x, y);
	unsigned mult = MULT(x, y);
	unsigned div = DIV(x, y);
	

	//Affichage
	cout << "\nLe chiffre le plus grand entre " << x << " et " << y << " est : " << result << "\n" << endl;
	cout << "\nAddition : " << x << " + " << y << " = " << add << "\n" << endl;

	if (x > y) {//condition car la macro fait une erreur si x < y 
		cout << "\nSoustraction : " << x << " - " << y << " = " << sous << "\n" << endl;
	}
	
	cout << "\nMultiplication : " << x << " * " << y << " = " << mult << "\n" << endl;
	cout << "\nDivision : " << x << " / " << y << " = " << div << "\n" << endl;
	TRACE();
	cout << "\n"<<endl;

}