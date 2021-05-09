#include "header.h"
#include "Carte.h"

namespace DomainModel2 {//tout est contenu dans cette classe (évite de devoir déterminer les chemins pour toutes les fonction mais également parce que ostream ne fonctionnait pas sinon

	int countCard;//variable qui compte le nombre de carte créé, définit à l'extérieur pour ne pas qu'elle se reset

	unsigned Carte::GetNbCreation() {//compteur du nombre de cartes créés
		countCard++;
		cout << "\n				[Carte numero : " << countCard << "]"<< endl;
		return 0;
	}
	Carte& Carte::operator=(Carte& carte)//si les cartes sont identiques
	{    
		if (this == &carte) return *this;
		_couleur = carte._couleur;
		_valeur = carte._valeur;
		return *this;
	}

	bool Carte::operator==(Carte& val) const//si la valeur des cartes sont strictement identiques
	{
		if (this->_couleur == val._couleur && this->_valeur == val._valeur) {
			return true;
		} else {
			return false;
		}

	}

	bool Carte::operator!=(Carte& val) const//si la valeur des cartes sont différentes
	{
		if (this->_couleur != val._couleur && this->_valeur != val._valeur) {
			return true;
		}
		else {
			return false;
		}

	}

	void Carte::setType(Couleur elem) {//on affecte la couleur à la carte
		this->_couleur = elem;
	}

	void Carte::setValeur(const string& valeur) {//on affecte la valeur à la carte
		this->_valeur = valeur;
	}

	ostream& operator<<(ostream& elem, Carte& val)//on surcharge l'operateur cout, et on affiche les cartes
	{
		string icone[4] = { "PIQUE", "COEUR", "CARREAU", "TREFLE"};

		elem << "\n			couleur de la carte: " << icone[val._couleur] << "\n			Valeur de la carte: " << val._valeur;
		return elem;
	}

	Carte::Carte(Couleur color, string val) ://constructeur
		_couleur(color),
		_valeur(val)
	{};

	Carte::Carte(Carte& copy) ://copy contructeur
		Carte(copy._couleur, copy._valeur)
	{}

	Carte::~Carte()//destructeur
	{
		cout << __func__ << endl;
	}
}
void Carte2() {

	cout << "\n   ---Jeu de carte numero 2---\n" << endl;

	cout << "\n-----------------------------------------------------------------------" << endl;

	DomainModel2::Carte c1(DomainModel2::Couleur::PIQUE, "As");
	c1.GetNbCreation();
	cout << "\n				---CARTE 1---\n";
	cout << c1 << endl;

	DomainModel2::Carte c2(c1);
	c2.GetNbCreation();
	cout << "\n				---CARTE 2 : AVANT---\n";
	cout << c2 << endl;
	c2.setType(DomainModel2::Couleur::TREFLE);
	c2.setValeur("Queen");
	cout << "\n				---CARTE 2 : APRES---\n";
	cout  << c2 << endl;
	DomainModel2::Carte c3(DomainModel2::Couleur::CARREAU, "valet");
	c3.GetNbCreation();
	cout << "\n				---CARTE 3 : AVANT---\n";
	cout << c3 << endl;
	c3.setType(DomainModel2::Couleur::PIQUE);
	c3.setValeur("As");
	cout << "\n				---CARTE 3 : APRES---\n";
	cout << c3 << endl;

	cout << "\n-----------------------------------------------------------------------" << endl;

	if (c1 != c2) {
		cout << "\n\n-->is ok :)\nLa premiere et deuxieme cartes sont bien differentes :)" << endl;
	}
	else {
		cout << "\n\n-->problem bug\nIls s'emblerait que certaines cartes soient identiques" << endl;
		
	}

	cout << "\n-----------------------------------------------------------------------" << endl;

	if (c1 == c3) {
		cout << "\n\n-->is ok :)\nLa premiere et la troisieme cartes sont bien identiques :)" << endl;
	}
	else {
		cout << "\n\n-->problem bug\nIls s'emblerait que certaines cartes soient differentes" << endl;

	}

}