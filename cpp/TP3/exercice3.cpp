#include "header.h"
#include "CarteY.h"

namespace Carte_YUGIOH {

	//on affiche les éléments
	string Carte_Monstre::afficher()
	{

		string attribut[6] = { "TENEBRE", "TERRE", "FEU", "LUMIERE", "EAU", "VENT" };

		cout << "Nom:" << getNomCarte() << "			Attribut: " << attribut[getAttribut()] << "\n						Niveau :" << getNiveau() << "\nNumero: " << getNumeroCarte() << "\n\n[" << getType() << "]\nDescription: " << getDescription() << "\n---------------------------------------------------------------------------\n						ATK/" << getATK() << " Defense/" << getDEF() << "\n" << endl;
		return string();
	}


	//on construit la carte 
	Carte_Monstre::Carte_Monstre(int ATK, int DEF, string desc, int niveau, string nom, string numero, string typeMonstre, string typeCarte, Attribut att) :
		a_ATK(ATK),
		a_DEF(DEF),
		a_description(desc),
		a_niveau(niveau),
		a_nomcarte(nom),
		a_numerocarte(numero),
		a_typeMonstre(typeMonstre),
		a_typeCarte(typeCarte),
		a_attribut(att)
	{}

}

void CarteY1() {
	cout << "\nMonstre Yu-Gi-Ho!\n" << endl;


	int answer;
	cout << "\nVoulez-vous une carte Monstre LUMIERE (1), EAU (2), TENEBRE (3), VENT (4), TERRE (5) ou FEU (6)? ";
	cin >> answer;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> answer;
	}
	
	cout << "\n-----------------------------------------------------------------------" << endl;
	cout << "\n			---CARTE MONSTRE---" << endl;


	if (answer == 1) {

		Carte_YUGIOH::Carte_Monstre c1(400, 800, "Lorsque vous declarez une attaque, vous pouvez \nactiver l'effet de cette carte. Si vous le faite, \ncette carte perd 500 points d'ATK et de DEF (definitivement)\net le monstre qu'elle combat perd 1500points d'ATK et de DEF \njusqu'a la End Phase ", 8, "Dragon Ultime de Lumiere", "RGBT-FR091", "Dragon", "Synchro", Carte_YUGIOH::Carte_Monstre::LUMIERE);
		c1.afficher();
	}
	else if (answer == 2) {
		Carte_YUGIOH::Carte_Monstre c2(1600, 1100, "Durant votre Main Phase, vous pouvez\nInvoquer Specialement 1 monstre Peluchimal depuis\nvotre main. Cet effet ne peut être utilise qu'une \nfois tant que cette carte est face recto sur le \nTerrain. Si cette carte est envoyee au Cimetiere comme \n Materiel de Fusion pour l'Invocation d'un Monstre Frourreur : \nvous pouvez piocher 2 cartes, puis defaussez 1 carte.\n\nVous ne pouvez utiliser cet effet de Pingouin Peluchimal qu'une fois par tour.",4, "PINGUINO FLUFFAL", "FUEN-FR015", "Elfe", "Effet", Carte_YUGIOH::Carte_Monstre::EAU);
		c2.afficher();
	}
	else if (answer==3){
		Carte_YUGIOH::Carte_Monstre c3(200, 100, "Si cette carte est Invoquee Normalement ou retournee face recto:\nvous pouvez bannir 1 monstre 'Seigneur Lumiere'\ndepuis votre main ou Cimetiere: bannissez\n1 carte sur le Terrain. Une fois par\ntour, si un autre effet de votre monstre\n'Seigneur Lumiere' est active: envoyez les 3\ncartes du dessus de votre Deck au Cimetiere", 2, "Ryko le Combattant, Seigneur Lumiere du Crepuscule", "MP18-FR053", "Bete", "Effet", Carte_YUGIOH::Carte_Monstre::TENEBRE);
		c3.afficher();
	}
	else if (answer == 4) {
		Carte_YUGIOH::Carte_Monstre c4(800, 1, "Durant votre Main Phase, vous pouvez\nInvoquer Normalement 1 monstre Calice du Monde en plus\nde votre Invocation/Pose Normale.Au debut de la\nDamage Step, si cette carte combat un monstre de \npointe par cette carte : vous pouvez detruire son monstre ",1, "Imduk le Dragon du Calice du Monde", "COTD-FR047", "Dragon", "Lien", Carte_YUGIOH::Carte_Monstre::VENT);
		c4.afficher();
	}
	else if (answer == 5) {
		Carte_YUGIOH::Carte_Monstre c5(1700, 1400, "Lorsque cette carte detruit un monstre de l'adversaire\nau combat: vous pouvez ajouter 1 monstre Bete\ndemax. Niveau 4 depuis votre Deck a votre main.", 4, "Diable de Desmanie", "MP18-FR191", "Bete", "Effet", Carte_YUGIOH::Carte_Monstre::TERRE);
		c5.afficher();
		cout << "\n-----------------------------------------------------------------------" << endl;
		cout << "\n			---CARTE MONSTRE ULTRA RARE---" << endl;
		Carte_YUGIOH::Carte_Monstre c6(2600,1900, "2 monstres de niveau 9\nDurant votre tour, lorsque vous activez: une carte\nou un effet (Effet Rapide): vous pouvez cibler:\n1 carte du meme type (Monstre, Magie, Piege) dans votre\nCimetiere: attachez le a cette carte comme Materiel.\nDurant le tour de votre adversaire, lorsqu'une carte\nou un effet est active (Effet Rapide): vous pouvez detacher un Materiel\ndu meme type(Monstre, Magie, Piege) de cette carte, et si vous le faites,\nannulez l'activation, et si vous faites ca, detruisez la carte.\nVous ne pouvez utiliser chaque effet de 'Bete-Arbre Sacree, Hyperyton'\nqu'une fois par tout", 9, "Bete-Arbre Sacree, Hyperyton", "BLVO-FR047", "Plante", "Xyz", Carte_YUGIOH::Carte_Monstre::TERRE);
		c6.afficher();

	}
	else if (answer == 6) {

		Carte_YUGIOH::Carte_Monstre c7(500, 200, "Si vous contrôlez un monstre\n   Cloche de Feu face recto, excepté Origine de la Neo Cloche\nde Feu, et que votre adversaire a 3 cartes ou moins\ndans son Cimetiere, vous pouvez Invoquer Specialement\ncette carte depuis votre main. ",2, "Origine de la Neo Cloche de Feu", "HA04-FR031", "Pyro", "Syntoniseur", Carte_YUGIOH::Carte_Monstre::FEU);
		c7.afficher();

	}
	else {
		cout << "\nJe n'ai pas compris votre demande....\n" << endl;
	}
	
	cout << "\n-----------------------------------------------------------------------" << endl;

}