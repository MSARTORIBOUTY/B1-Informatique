#ifndef _CARTEY2_H_
#define _CARTEY2_H_
#include "header.h"


//.h de l'exercice 5

namespace Carte_YUGIOH3 {

    class ICarte_YUGIOH {//classe virtuelle pure, elle ne contien pas de constructeur parce que l'on en a pas besoins, elle prend celui par défaut

    public:
        virtual ~ICarte_YUGIOH() {};

        virtual string afficher() = 0;
        virtual string getDescription() = 0;
        virtual string getNomCarte() = 0;
        virtual string getNumeroCarte() = 0;
        virtual string getType() = 0;
    };

    enum Attribut { TENEBRE, TERRE, FEU, LUMIERE, EAU, VENT };

    class Carte_Monstre : public ICarte_YUGIOH {//classe qui hérite de la classe interface

    public:

        

        Carte_Monstre(int, int, string, int, string, string, string, string, Attribut);
        ~Carte_Monstre() {};


        string afficher();
        int getATK() { return this->a_ATK; };
        Attribut getAttribut()  { return this->a_attribut; };
        int getDEF() { return this->a_DEF; };
        string getDescription() { return this->a_description; };
        int getNiveau() { return this->a_niveau; };
        string getNomCarte() { return this->a_nomcarte; };
        string getNumeroCarte() { return this->a_numerocarte; };
        string getType() { return this->a_typeMonstre + "/" +  this->a_typeCarte ; };



    private:

        int a_ATK, a_DEF, a_niveau;
        string a_description, a_nomcarte, a_numerocarte, a_typeMonstre, a_typeCarte;

        Attribut a_attribut;
    };


    class Carte_MagiePiege : public ICarte_YUGIOH {//herite de la classe interface

    public:
        enum Icone { EQUIPEMENT, TERRAIN, JEURAPIDE, CONTRE, CONTINUE, RITUEL, NORMAL };

        Carte_MagiePiege(string, string, string, Icone);
        virtual ~Carte_MagiePiege() {};

        virtual string afficher();
        virtual string getDescription() { return this->a_description; };
        virtual Icone getIcone() { return this->a_icone; };
        virtual string getNomCarte() { return this->a_nomcarte; };
        virtual string getNumeroCarte() { return this->a_numerocarte; };
        virtual string getType() = 0;

    protected:

        string a_description, a_nomcarte, a_numerocarte;

        Icone a_icone;

    };

    class Carte_Magique : public Carte_MagiePiege{//herite de la classe MagiePiege

    public:

        Carte_Magique(string, string, string, Icone);
        ~Carte_Magique() {};

        string afficher();
        string getType();

    };

    class Carte_Piege : public Carte_MagiePiege {//herite de la classe MagiePiege

    public:

        Carte_Piege(string, string, string, Icone);
        ~Carte_Piege() {};

        string afficher();
        string getType();

    };

}

#endif