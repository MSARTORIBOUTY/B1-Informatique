#include "CarteY2.h"
#include "header.h"

void CarteM();
void CarteMag();
void CarteP();

namespace Carte_YUGIOH3 {

	//Classe monstre-----------------------------------------------------------------------------------------------------------------
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

	//Class MagiePiege, Magique et Piege-----------------------------------------------------------------------------------------------------------------------------------------------------------------
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
		cout << "			" << getType() << Carte_YUGIOH3::Carte_MagiePiege::afficher() << endl;
		return string();
	}

	string Carte_Piege::afficher()
	{
		cout << "			" << getType() << Carte_YUGIOH3::Carte_MagiePiege::afficher() << endl;
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


void CarteY3() {
	cout << "\nCartes Yu-Gi-Ho!\n" << endl;

	int answer;
	cout << "\nVoulez-vous une carte Magie (1),  Piege (2), ou Monstre (3) ? ";
	cin >> answer;

	while (cin.fail()) {//sécurisation
		cin.clear();
		cin.ignore(1, '\n');
		cout << "\nEntrez un entier:";
		cin >> answer;
	}

	if (answer == 1) {
		CarteMag();
	}
	else if (answer == 2) {
		CarteP();
	}
	else if (answer == 3) {
		CarteM();
	}
	else {
		cout << "Je n'ai pas compris votre reponse...\n" << endl;
	}

}

//____________________________________________________________________________________

void CarteM() {

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

		Carte_YUGIOH3::Carte_Monstre c1(400, 800, "Lorsque vous declarez une attaque, vous pouvez \nactiver l'effet de cette carte. Si vous le faite, \ncette carte perd 500 points d'ATK et de DEF (definitivement)\net le monstre qu'elle combat perd 1500points d'ATK et de DEF \njusqu'a la End Phase ", 8, "Dragon Ultime de Lumiere", "RGBT-FR091", "Dragon", "Synchro", Carte_YUGIOH3::LUMIERE);
		c1.afficher();
	}
	else if (answer == 2) {
		Carte_YUGIOH3::Carte_Monstre c2(1600, 1100, "Durant votre Main Phase, vous pouvez\nInvoquer Specialement 1 monstre Peluchimal depuis\nvotre main. Cet effet ne peut être utilise qu'une \nfois tant que cette carte est face recto sur le \nTerrain. Si cette carte est envoyee au Cimetiere comme \n Materiel de Fusion pour l'Invocation d'un Monstre Frourreur : \nvous pouvez piocher 2 cartes, puis defaussez 1 carte.\n\nVous ne pouvez utiliser cet effet de Pingouin Peluchimal qu'une fois par tour.", 4, "PINGUINO FLUFFAL", "FUEN-FR015", "Elfe", "Effet", Carte_YUGIOH3::EAU);
		c2.afficher();
	}
	else if (answer == 3) {
		Carte_YUGIOH3::Carte_Monstre c3(200, 100, "Si cette carte est Invoquee Normalement ou retournee face recto:\nvous pouvez bannir 1 monstre 'Seigneur Lumiere'\ndepuis votre main ou Cimetiere: bannissez\n1 carte sur le Terrain. Une fois par\ntour, si un autre effet de votre monstre\n'Seigneur Lumiere' est active: envoyez les 3\ncartes du dessus de votre Deck au Cimetiere", 2, "Ryko le Combattant, Seigneur Lumiere du Crepuscule", "MP18-FR053", "Bete", "Effet", Carte_YUGIOH3::TENEBRE);
		c3.afficher();
	}
	else if (answer == 4) {
		Carte_YUGIOH3::Carte_Monstre c4(800, 1, "Durant votre Main Phase, vous pouvez\nInvoquer Normalement 1 monstre Calice du Monde en plus\nde votre Invocation/Pose Normale.Au debut de la\nDamage Step, si cette carte combat un monstre de \npointe par cette carte : vous pouvez detruire son monstre ", 1, "Imduk le Dragon du Calice du Monde", "COTD-FR047", "Dragon", "Lien", Carte_YUGIOH3::VENT);
		c4.afficher();
	}
	else if (answer == 5) {
		Carte_YUGIOH3::Carte_Monstre c5(1700, 1400, "Lorsque cette carte detruit un monstre de l'adversaire\nau combat: vous pouvez ajouter 1 monstre Bete\ndemax. Niveau 4 depuis votre Deck a votre main.", 4, "Diable de Desmanie", "MP18-FR191", "Bete", "Effet", Carte_YUGIOH3::TERRE);
		c5.afficher();
		cout << "\n-----------------------------------------------------------------------" << endl;
		cout << "\n			---CARTE MONSTRE ULTRA RARE---" << endl;
		Carte_YUGIOH3::Carte_Monstre c6(2600, 1900, "2 monstres de niveau 9\nDurant votre tour, lorsque vous activez: une carte\nou un effet (Effet Rapide): vous pouvez cibler:\n1 carte du meme type (Monstre, Magie, Piege) dans votre\nCimetiere: attachez le a cette carte comme Materiel.\nDurant le tour de votre adversaire, lorsqu'une carte\nou un effet est active (Effet Rapide): vous pouvez detacher un Materiel\ndu meme type(Monstre, Magie, Piege) de cette carte, et si vous le faites,\nannulez l'activation, et si vous faites ca, detruisez la carte.\nVous ne pouvez utiliser chaque effet de 'Bete-Arbre Sacree, Hyperyton'\nqu'une fois par tout", 9, "Bete-Arbre Sacree, Hyperyton", "BLVO-FR047", "Plante", "Xyz", Carte_YUGIOH3::TERRE);
		c6.afficher();
	}
	else if (answer == 6) {

		Carte_YUGIOH3::Carte_Monstre c7(500, 200, "Si vous contrôlez un monstre\n   Cloche de Feu face recto, excepté Origine de la Neo Cloche\nde Feu, et que votre adversaire a 3 cartes ou moins\ndans son Cimetiere, vous pouvez Invoquer Specialement\ncette carte depuis votre main. ", 2, "Origine de la Neo Cloche de Feu", "HA04-FR031", "Pyro", "Syntoniseur", Carte_YUGIOH3::FEU);
		c7.afficher();

	}
	else {
		cout << "\nJe n'ai pas compris votre demande....\n" << endl;
	}
	
	cout << "\n-----------------------------------------------------------------------" << endl;

}

//________________________________________________________________________________

void CarteMag() {

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

		Carte_YUGIOH3::Carte_Magique c1("Equipable uniquement a un monstre 'Inzektor'. Il gagne\n1000 ATK. Lorsque le monstre equipe declare une attaque, votre\nadversaire ne peut pas activer de Carte Magie/Piege", "Hache Inzektor - Zektahawk", "ORCS-FR000", Carte_YUGIOH3::Carte_Magique::EQUIPEMENT);
		c1.afficher();

	}
	else if (answer == 2) {

		Carte_YUGIOH3::Carte_Magique c2("Chaque fois qu'un ou plusieurs monstres\nsamourais sont invoques, placez 1 compteur\nBushido surcette carte. Les monstres controles \npar votre adversaire perdent 100 ATK pour chaque \ncompteur Bushido sur cette carte.", "Temple des six", "MAGO-FR146", Carte_YUGIOH3::Carte_Magique::TERRAIN);
		c2.afficher();

	}
	else if (answer == 3) {

		Carte_YUGIOH3::Carte_Magique c3("Activez 1 de ces effets:\n-->Ce tour, votre adversaire ne peut activer ni de cartes\nni d'effets en reponse de l'activation des\neffets de votre monstre 'Croisedia'\n-->Apres calcul des dommages, si votre monstre lien 'Croisedia' detruit un monstre\nde l'adversaire au combat : piochez un nombre de cartes egale\na la Classification Lien de votre monstre.\nVous ne pouvez activer qu'un 'Testament Croisedia' par tour.", "Testament Croisedia", "MP20-FR78", Carte_YUGIOH3::Carte_Magique::JEURAPIDE);
		c3.afficher();

	}
	else if (answer == 4) {

		Carte_YUGIOH3::Carte_Magique c4("Chaque fois qu'eun Carte Magie est activee, gagnez 500 LP\nimmediatement apres sa resolution.", "Absorption de Magie", "INCH-FR053", Carte_YUGIOH3::Carte_Magique::CONTINUE);
		c4.afficher();

	}
	else if (answer == 5) {

		Carte_YUGIOH3::Carte_Magique c5("Cette carte est utilisee pour Invoquer\nRituellement Miroir du Demon. Vous devez\naussi Sacrifier des monstres depuis le Terrain\nou votre main dont le Niveau total est egal a minimum 6.", "Rituel du Miroir Bestial", "PRC1-FR002", Carte_YUGIOH3::Carte_Magique::RITUEL);
		c5.afficher();

	}
	else if (answer == 7) {

		Carte_YUGIOH3::Carte_Magique c6("Envoyez 1 monstre 'Orcust' ou 'Heritage du Monde' depuis votre\nmain ou Terrain face recto au Cimetiere; piochez 2 cartes.\nVous ne pouvez activer qu'1 'Retour Orcustre' par tour.", "Retour Orcustre", "GFTP-FR119", Carte_YUGIOH3::Carte_Magique::NORMAL);
		c6.afficher();

	}
	else {
		cout << "\nje n'ai pas compris votre demande...\n" << endl;
	}
	
	cout << "\n-----------------------------------------------------------------------" << endl;

}

//________________________________________________________

void CarteP() {

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

		Carte_YUGIOH3::Carte_Piege c1("Payez 1000 Life Points. Annulez l'invocation \nd'un monstre du meme Type qu'un monstre sur le\nTerrain et detruisez-le.", "Selection de l'arche", "TDGS-FR093", Carte_YUGIOH3::Carte_Piege::CONTRE);
		c1.afficher();
	}
	else if (answer == 2) {
		Carte_YUGIOH3::Carte_Piege c2("Reduisez le Niveau de tous les monstres face \nrecto controles par votre adversaire de 1.", "Champ Funebre Stygien", "SOVR-FR078", Carte_YUGIOH3::Carte_Piege::CONTINUE);
		c2.afficher();

	}
	else if (answer == 3) {
		Carte_YUGIOH3::Carte_Piege c3("Si vous controlez un monstre EAU: Ciblez 1 Magie\ndans le Cimetiere de votre adversaire ; banissez-la,\net si vous le faites, annulez tous\nles effets de Magie sur le Terrain de votre adversaire\nle reste de ce tour.", "Hurlement Blanc", "MP20-FR143", Carte_YUGIOH3::Carte_Piege::NORMAL);
		c3.afficher();


	}
	else {
		cout << "\nje n'ai pas compris votre demande...\n" << endl;
	}
	
	cout << "\n-----------------------------------------------------------------------" << endl;
}