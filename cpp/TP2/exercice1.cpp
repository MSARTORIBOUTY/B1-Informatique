#include "header.h"

typedef double* ptrDouble; 

void constructeur(ptrDouble& dd, const unsigned& taille) {// cr�ation d'un pointeur
    cout << "\n   ---CONSTRUCTION!--- \n";
    dd = new double[taille] {0};//initialisation de ces valeurs � 0 avec son type
}

void afficher(const ptrDouble dd, const unsigned& index) {//affichage de svaleurs du pointeur

    cout << "\n   ---AFFICHAGE---" << endl;

    if (dd != nullptr) {//si il est diff�rent de null
        for (int i = 0; i < index; i++) {
            cout << "Valeur du tableau[ " << i << " ] : " << dd[i] << endl;
        }
    }
    else {//si il est null et donc n'existe pas
        cout << "ERROR: Le tableau n'existe pas\n" << endl;
    }
    

}

void modifier(const ptrDouble dd, const unsigned& index, const double& valeur) {//modification des vlaeurs du pointeur

    cout << "\n   ---MODIFICATION---" << endl;

        dd[index] = valeur;

    
}

void destructeur(ptrDouble& dd) {

    cout << "\n   ---DESTRUCTION---" << endl;

    delete[] dd;//on a utiliser new pour lui assigner une case et l� on le d�truit
    dd = nullptr;//il faut donc lib�rer la cas. Si le pointeur est initialiser � 0, aucune case ne lui est assign�
    //dd = NULL;

}

const double get(const ptrDouble& dd, const unsigned& index) {//on retourne une valeur
    cout << "\n   ---GET---" << endl;
    return dd[index];
}

double& get(ptrDouble& dd, const unsigned& index) {
    cout << "\n   ---GET---" << endl;
    return dd[index];
}

void C_2_6b() {

    ptrDouble d1 = nullptr;//d1 est un pointeur double que l'on initialise � 0
    int t1 = 5;//variable qui poss�de la valeur 5, taille du tableau
    constructeur(d1, t1);
    afficher(d1, t1);
    modifier(d1, 2, 3.13589985);
    afficher(d1, t1);
    cout << get(d1, 2) << endl;
    get(d1, 2) = 62.1;
    cout << get(d1, 2) << endl;
    afficher(d1, t1);
    destructeur(d1);
    afficher(d1, t1);
}