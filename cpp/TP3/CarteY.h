#ifndef _CARTEY_H_
#define _CARTEY_H_
#include "header.h"

//1------------------------------------------------------------------------------------------------------------------------
namespace Carte_YUGIOH {

	class Carte_Monstre {

	public:

		enum Attribut {TENEBRE, TERRE, FEU, LUMIERE, EAU, VENT};

		Carte_Monstre(int, int, string, int, string, string, string, string, Attribut);
		~Carte_Monstre() { cout << __func__ << endl; };

		string afficher();
		int getATK()const { return this->a_ATK; };
		Attribut getAttribut() const { return this->a_attribut; };
		int getDEF()const { return this->a_DEF; };
		string getDescription()const { return this->a_description; };
		int getNiveau()const { return this->a_niveau; };
		string getNomCarte()const { return this->a_nomcarte; };
		string getNumeroCarte()const { return this->a_numerocarte; };
		string getType()const { return this->a_typeMonstre + "/" + this->a_typeCarte; };


	private:

		int a_ATK, a_DEF, a_niveau;
		string a_description, a_nomcarte, a_numerocarte,a_typeMonstre,a_typeCarte ;

		Attribut a_attribut;
	};

}

//2------------------------------------------------------------------------------------------------------------------------
namespace Carte_YUGIOH2 {

	class Carte_MagiePiege {

	public:
		enum Icone {EQUIPEMENT, TERRAIN, JEURAPIDE, CONTRE, CONTINUE, RITUEL, NORMAL};

		Carte_MagiePiege(string, string, string, Icone);
		virtual ~Carte_MagiePiege() { cout << __func__ << endl; };

		virtual string afficher();
		virtual string getDescription() const { return this->a_description; };
		virtual Icone getIcone()const { return this->a_icone; };
		virtual string getNomCarte()const { return this->a_nomcarte; };
		virtual string getNumeroCarte()const { return this->a_numerocarte; };
		virtual string getType()=0;

	protected:
		
		string a_description, a_nomcarte, a_numerocarte;

		Icone a_icone;

	};

	class Carte_Magique: public Carte_MagiePiege {

	public:

		Carte_Magique(string, string, string, Icone) ;
		~Carte_Magique() { cout << __func__ << endl; };

		string afficher();
		string getType();

	};

	class Carte_Piege: public Carte_MagiePiege {

	public:

		Carte_Piege(string, string, string, Icone) ;
		~Carte_Piege() { cout << __func__ << endl; };

		string afficher();
		string getType();

	};

}


#endif