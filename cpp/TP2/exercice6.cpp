#include "Header.h"

enum Sexe { INCONNU = 0, MASCULIN = 1, FEMININ = 2 };



//champs à manipuler
struct Personne {
	int numero;
	int age;
	string ville;
	string pays;
	string prenom;
	string nom;
	string metier;
	Sexe sexe;

};


//foncion de destruction
void Destruction(Personne* pointer) {
	cout << "\nVue que le crime commis ne meritait pas la peine capitale, nous l'avons laissé vivre paisiblement...\nMort de ";
	cout << pointer->prenom << " " << pointer->nom << " a l'age de " << pointer->age << " ans.\n" << endl;
	delete pointer;
	pointer = nullptr;
}

//fonction de l'âge random du décès
void vieillir(Personne* pointer) { 
	srand(time(NULL));

	pointer->age += rand() % 50 + 1; 
	Destruction(pointer);
}
//fonction affichage
void Affichage(Personne* pointer) {

	cout << "\n   ---CARTE D'IDENTITE---\n";
	if (pointer->sexe == MASCULIN) {
		cout << "\nCreation d'un homme";
	}
	else if (pointer->sexe == FEMININ) {
		cout << "\nCreation d'une femme";
	}
	else if (pointer->sexe == INCONNU) {
		cout << "\nCreation de";
	}

	cout << " : " << pointer->prenom << " " << pointer->nom << endl;
	cout << " " << " Age :  " << pointer->age << " ans" << endl;
	cout << " " << " Vivant a " << pointer->ville << ", " << pointer->pays << endl;
	cout << " " << " Metier: " << pointer->metier << endl;
	cout << " " << " Contactable au : " << pointer->numero << endl;
	cout << " " << "\n\n" << endl;

	vieillir(pointer);


}

// fonction détermination des informations
void Structure(Personne* pointer){

	int DeterSexe;

	cout << "\n   ---ENTREZ LES INFORMATIONS DE LA PERSONNE---\n";

	cout << "\nNom : ";
	cin >> pointer->nom;
	cout << "Prenom : ";
	cin >> pointer->prenom;
	cout << "Age : ";
	cin >> pointer->age;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> pointer->age;
	}

	cout << "Ville : ";
	cin >> pointer->ville;
	cout << "Pays : ";
	cin >> pointer->pays;
	cout << "Metier : ";
	cin >> pointer->metier;
	cout << "Numero de telephone : ";
	cin >> pointer->numero;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> pointer->numero;
	}

	cout << "Sexe (0 = Inconnu) (1 = Masculin) (2 = Feminin) : ";
	cin >> DeterSexe;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "Sexe (0 = Inconnu) (1 = Masculin) (2 = Feminin) : ";
		cin >> DeterSexe;
	}

	if (DeterSexe != 0 || DeterSexe != 1 || DeterSexe != 2) {
		DeterSexe = 0;
	}

	switch (DeterSexe) {
	case MASCULIN: {
		pointer->sexe = MASCULIN;
		break;
	}
	case INCONNU: {
		pointer->sexe = INCONNU;
		break;
	}
	case FEMININ: {
		pointer->sexe = FEMININ;
		break;
	}
	}
	Affichage(pointer);

}


//fonction de construction
void Init() {
	cout << "\nBienvenu dans l'exercice 6, votre mission si vous l'acceptez, est de rentrer les informations sur\n l'identite du potentiel voleur de la derniere rillette de canard qui se\n trouvait encore dans le frigo, pas plus tard qu'hier\nBonne chance!\n";
	cout << "\n   ---CONSTRUCTION---" << endl;
	Personne* pointer = new Personne;
	Structure(pointer);

}
