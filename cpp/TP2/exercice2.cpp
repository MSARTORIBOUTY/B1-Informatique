#include "header.h"
#include <memory>

typedef std::shared_ptr<double> ptrStdDouble;
using ptrDouble = shared_ptr<double[]>;


void constructeur(ptrDouble& dd, const unsigned& taille) {//construction of a smart pointer
    cout << "\n   ---CONSTRUCTION!--- \n";
    dd = shared_ptr<double[]>(new double[taille]{0}) ; // construction d'un pointeur intelligent, on l'initialise à 0
}

void afficher(const ptrDouble dd, const unsigned& index) { // display smart pointer's value
    cout << "\n   ---AFFICHAGE---" << endl;
    
    if (dd != nullptr) {//si le pointeur n'est pas null, on affiche toutes les valeurs du pointeur intelligent
        for (int i = 0; i < index; i++) {
            cout << "Valeur du tableau[ " << i << " ] : " << dd.get()[i]  << endl;
        }

    }
    else {//sinon le tableau n'existe pas
        cout << "ERROR: Le tableau n'existe pas\n" << endl;
    }


}

void modifier(const ptrDouble dd, const unsigned& index, const double& valeur) {//change value of the smart pointer
    cout << "\n   ---MODIFICATION---" << endl;//on modifie la valeur du pointeur intelligent à un emplacement précis

    dd.get()[index] = valeur;
}

void destructeur(ptrDouble& dd) {

    cout << "\n   ---DESTRUCTION---" << endl;

    dd.reset();//on a utiliser new pour lui assigner une case et là on le détruit
    dd = nullptr;//il faut donc libérer la cas. Si le pointeur est initialiser à 0, aucune case ne lui est assigné
    
    

}

const double get(const ptrDouble& dd, const unsigned& index) {//get return the value of the smart pointeur
    cout << "\n   ---GET---" << endl;//affiche la valeur à un index précis
    return dd.get()[index];
}

double& get(ptrDouble& dd, const unsigned& index) {//get return the value of the smart pointeur
    cout << "\n   ---GET---" << endl;//affiche la valeur à un index précis
    return dd.get()[index];
}

void C_2_6c() {

    ptrDouble d1 = nullptr;//d1 est un pointeur double que l'on initialise à 0
    int t1 = 5;//variable qui possède la valeur 5
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