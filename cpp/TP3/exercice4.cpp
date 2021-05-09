#include "header.h"
#include "CarteY.h"

void CarteMagie();
void CartePiege();

namespace Carte_YUGIOH2 {


	//On retourne les types directements en fonction du contructeur appelé pour construire la carte
	string Carte_Piege::getType()
	{
		return "[CARTE PIEGE]";
	}

	string Carte_Magique::getType()
	{
		return "[CARTE MAGIE]";
	}

	//on affiche les cartes d'après les contructeurs des classes filles
	string Carte_Magique::afficher()
	{
		cout << "			" << getType() << Carte_YUGIOH2::Carte_MagiePiege::afficher() << endl;
		return string();
	}

	string Carte_Piege::afficher()
	{
		cout << "			" << getType() << Carte_YUGIOH2::Carte_MagiePiege::afficher() << endl;
		return string();
	}

	//on affiche les reste des éléments avec la fonction afficher de la classe mère
	string Carte_MagiePiege::afficher()
	{

		string icone[7] = { "EQUIPEMENT", "TERRAIN", "JEU RAPIDE", "CONTRE", "CONTINU", "RITUEL", "NORMAL" };

		cout << "\n\n			Nom:" << getNomCarte() << "\n\n			Numero: " << getNumeroCarte() << "\n\n			Icone: [" << icone[getIcone()] << "]\n\nDescription: " << getDescription() << endl;
		return string();
	}

	//constructeur des classes filles d'après le cnstructeur de la classe Mère
	Carte_Magique::Carte_Magique(string desc, string nom, string num, Icone icone) :
		Carte_MagiePiege(desc, nom, num, icone)
	{}

	Carte_Piege::Carte_Piege(string desc, string nom, string num, Icone icone) :
		Carte_MagiePiege(desc, nom, num, icone)
	{}

	//Constructeur classe Mère, on initialise pas les variables car ce constructeur sert de passerelle 
	Carte_MagiePiege::Carte_MagiePiege(string desc, string nom, string num, Icone icone)
	{
		this->a_description = desc;
		this->a_nomcarte = nom;
		this->a_numerocarte = num;
		this->a_icone = icone;

	}

}

void CarteY2() {

	cout << "\nMagie et Piege Yu-Gi-Ho!\n" << endl;

	int answer;
	cout << "\nVoulez-vous une carte Magie (1) ou Piege (2) ? ";
	cin >> answer;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> answer;
	}


	if (answer == 1) {
		CarteMagie();
	}
	else if (answer == 2) {
		CartePiege();
	}

}

void CarteMagie() {

	int answer;
	cout << "\nVoulez-vous une carte Magie EQUIPEMENT (1), TERRAIN (2) , JEU RAPIDE (3),  CONTINUE (4), RITUEL (5) ou NORMAL (6) ? ";
	cin >> answer;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> answer;
	}
	
	cout << "\n-----------------------------------------------------------------------" << endl;
	cout << "\n			---CARTE MAGIE---" << endl;

	if (answer == 1) {

		Carte_YUGIOH2::Carte_Magique c1("Equipable uniquement a un monstre 'Inzektor'. Il gagne\n1000 ATK. Lorsque le monstre equipe declare une attaque, votre\nadversaire ne peut pas activer de Carte Magie/Piege", "Hache Inzektor - Zektahawk", "ORCS-FR000", Carte_YUGIOH2::Carte_Magique::EQUIPEMENT);
		c1.afficher();

	}
	else if (answer == 2) {

		Carte_YUGIOH2::Carte_Magique c2("Chaque fois qu'un ou plusieurs monstres\nsamourais sont invoques, placez 1 compteur\nBushido surcette carte. Les monstres controles \npar votre adversaire perdent 100 ATK pour chaque \ncompteur Bushido sur cette carte.","Temple des six", "MAGO-FR146", Carte_YUGIOH2::Carte_Magique::TERRAIN);
		c2.afficher();

	}
	else if (answer == 3) {

		Carte_YUGIOH2::Carte_Magique c3("Activez 1 de ces effets:\n-->Ce tour, votre adversaire ne peut activer ni de cartes\nni d'effets en reponse de l'activation des\neffets de votre monstre 'Croisedia'\n-->Apres calcul des dommages, si votre monstre lien 'Croisedia' detruit un monstre\nde l'adversaire au combat : piochez un nombre de cartes egale\na la Classification Lien de votre monstre.\nVous ne pouvez activer qu'un 'Testament Croisedia' par tour.", "Testament Croisedia", "MP20-FR78", Carte_YUGIOH2::Carte_Magique::JEURAPIDE);
		c3.afficher();

	}
	else if (answer == 4) {

		Carte_YUGIOH2::Carte_Magique c4("Chaque fois qu'eun Carte Magie est activee, gagnez 500 LP\nimmediatement apres sa resolution.","Absorption de Magie", "INCH-FR053", Carte_YUGIOH2::Carte_Magique::CONTINUE);
		c4.afficher();

	}
	else if (answer == 5) {

		Carte_YUGIOH2::Carte_Magique c5("Cette carte est utilisee pour Invoquer\nRituellement Miroir du Demon. Vous devez\naussi Sacrifier des monstres depuis le Terrain\nou votre main dont le Niveau total est egal a minimum 6.","Rituel du Miroir Bestial", "PRC1-FR002", Carte_YUGIOH2::Carte_Magique::RITUEL);		
		c5.afficher();

	}
	else if (answer == 7) {

		Carte_YUGIOH2::Carte_Magique c6("Envoyez 1 monstre 'Orcust' ou 'Heritage du Monde' depuis votre\nmain ou Terrain face recto au Cimetiere; piochez 2 cartes.\nVous ne pouvez activer qu'1 'Retour Orcustre' par tour.", "Retour Orcustre", "GFTP-FR119", Carte_YUGIOH2::Carte_Magique::NORMAL);
		c6.afficher();

	}
	else {
		cout << "\nje n'ai pas compris votre demande...\n" << endl;
	}
	
	cout << "\n-----------------------------------------------------------------------" << endl;

}

void CartePiege() {

	int answer;
	cout << "\nVoulez-vous une carte Magie CONTRE (1), CONTINU (2), ou NORMAL (3) ? ";
	cin >> answer;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> answer;
	}

	cout << "\n-----------------------------------------------------------------------" << endl;
	cout << "\n			---CARTE PIEGE---" << endl;

	if (answer == 1) {

		Carte_YUGIOH2::Carte_Piege c1("Payez 1000 Life Points. Annulez l'invocation \nd'un monstre du meme Type qu'un monstre sur le\nTerrain et detruisez-le.", "Selection de l'arche", "TDGS-FR093", Carte_YUGIOH2::Carte_Piege::CONTRE);
		c1.afficher();

	}
	else if (answer == 2) {

		Carte_YUGIOH2::Carte_Piege c2("Reduisez le Niveau de tous les monstres face \nrecto controles par votre adversaire de 1.","Champ Funebre Stygien", "SOVR-FR078", Carte_YUGIOH2::Carte_Piege::CONTINUE);
		c2.afficher();


	}
	else if (answer == 3) {

		Carte_YUGIOH2::Carte_Piege c3("Si vous controlez un monstre EAU: Ciblez 1 Magie\ndans le Cimetiere de votre adversaire ; banissez-la,\net si vous le faites, annulez tous\nles effets de Magie sur le Terrain de votre adversaire\nle reste de ce tour.", "Hurlement Blanc", "MP20-FR143", Carte_YUGIOH2::Carte_Piege::NORMAL);
		c3.afficher();


	}
	else {
		cout << "\nje n'ai pas compris votre demande...\n" << endl;
	}
	cout << "\n-----------------------------------------------------------------------" << endl;
}