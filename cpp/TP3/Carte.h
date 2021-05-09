#ifndef _CARTE_H_
#define _CARTE_H_
#include "header.h"


namespace DomainModel {  

	class Carte {
		
	public:
		
		enum Couleur {PIQUE, COEUR, CARREAU, TREFLE};

		Carte(Couleur, string);
		Carte(Carte&);
		~Carte();


		void setType(Couleur);
		void setValeur(const string&);
		void afficher() const;
		bool equal(Carte&) const;
		void affecter(Carte&);


	private:
		string _valeur;
		Couleur _couleur;
	};

};

namespace DomainModel2 {
	enum Couleur { PIQUE, COEUR, CARREAU, TREFLE };

	class Carte {

	public:

		

		Carte(Couleur, string);
		Carte(Carte&);
		~Carte();


		void setType(Couleur);
		void setValeur(const string&);


		Carte& operator=(Carte&);
		bool operator==(Carte&) const;
		bool operator!=(Carte&) const;
		unsigned GetNbCreation();
		friend ostream& operator<<(ostream&, Carte&);

	private:
		string _valeur;
		Couleur _couleur;
	};

};
#endif // !_CARTE_H_
