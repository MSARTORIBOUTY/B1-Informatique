#include "header.h"
#include "Carte.h"


namespace DomainModel {

    void Carte::afficher() const {//affiche la carte


        string icone[4] = { "PIQUE", "COEUR", "CARREAU", "TREFLE" };


        cout << "\n         Symbole de la carte : " << icone[this->_couleur] << "\n         Valeur de la carte : " << this->_valeur << endl;
    }

    void Carte::setType(Couleur elem) {//détermine le type de la carte
        this->_couleur = elem;
    }

    void Carte::setValeur(const string& valeur) {//détermine la valeur de la carte
        this->_valeur = valeur;
    }

    void Carte::affecter(Carte& ptr) {//on affecte à la carte les élémeents d'une autre carte
        this->_couleur = ptr._couleur;
        this->_valeur = ptr._valeur;

    }

    bool Carte::equal(Carte& carte) const {//on vérifie si deux cartes sont identiques

        if (this->_couleur == carte._couleur && this->_valeur == carte._valeur) {
            return true;
        }
        else {
            return false;
        }

    }

    //on construit la cartes en initialisant des variables
    Carte::Carte(Couleur color, string val) :
        _couleur(color),
        _valeur(val)
    {};

    //on copie une carte d'après une carte existante
    Carte::Carte(Carte& copy) :
        Carte(copy._couleur, copy._valeur)
    {}

    //destructeur de la carte
    Carte::~Carte()
    {
        cout << __func__ << endl;
    }
}
void Carte1() {
	cout << "\nJeu de carte numero 1\n" << endl;

    cout << "\n-----------------------------------------------------------------------" << endl;

    DomainModel::Carte c1(DomainModel::Carte::Couleur::PIQUE, "As");
    cout << "\n             ---CARTE 1---\n";
    c1.afficher();
    DomainModel::Carte c2(c1);
    cout << "\n             ---CARTE 2 : AVANT---\n";
    c2.afficher();
    c2.setType(DomainModel::Carte::Couleur::TREFLE);
    c2.setValeur("Queen");
    cout << "\n             ---CARTE 2 : APRES---\n";
    c2.afficher();
    DomainModel::Carte c3(DomainModel::Carte::Couleur::PIQUE, "2");
    c2.affecter(c3);
    cout << "\n         ---CARTE 2 : DERNIERE EDITION---\n";
    c2.afficher();
    cout << "\n             ---CARTE 3---\n";
    c3.afficher();

    cout << "\n-----------------------------------------------------------------------" << endl;

    if (c1.equal(c2)) {
        cout << "\n\n-->is ok :)\nla premiere et deuxieme cartes sont equivalente :)" << endl;
    }
    else {
        cerr << " \n\n-->problem bug\nIl s'emblerait que certaines cartes soient differentes" << endl;
        cout << "\n             ---CARTE 1 : DERNIERE EDITION---\n";
        c1.afficher();
        cout << "\n             ---CARTE 2 : DERNIERE EDITION---\n";
        c2.afficher();
    }

    cout << "\n-----------------------------------------------------------------------" << endl;
    
}