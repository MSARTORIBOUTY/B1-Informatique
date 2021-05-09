import sys
sys.path.append('..')

from PUtilitaires import StartOver
from PUtilitaires import inputInt
from PUtilitaires import inputFloatN
from PUtilitaires import inferior



def rulesExo2():
    """rulesExo2: règles de l'exercice 2"""

    print("Bonjour et bienvenu dans l'exercice 2! :)")                                                                  #Hello and welcome to the second exercise
    print("---------------------------TABLEAU MULTIDIMENSSIONEL---------------------------")                            #multidimensional array
    print("vous allez devoir créer un tableau et y ajouter des valeurs!")                                               #you will need to create an array and add values to it
    print("d'abord définissez le nombres de colonnes et de lignes qu'il doit comporter")                                #first define the number of columns and rows it should contain
    print("puis ajoutez-y vos valeurs!")                                                                                #then, add your values
    print("seul des chiffres positifs supérieur à 0 sont acceptés pour créer le tableau!")                              #only positive numbers > 0 are allowed to create the array
    print("A vous de jouer! ;)")                                                                                        #Here you go! :)
    print("-----------------------------------------------------------------------")
    Elem()


def Elem():


    colonnes=inputInt("entrez le nombre de colonnes que doit avoir votre tableau:")                                     #enter the number of columns your array should have
    inferior(colonnes, Elem)
    lignes = inputInt("entrez le nombre de dimensions:")                                                                #enter the number of dimensions
    inferior(lignes, Elem)


    ArrayValue(colonnes, lignes)

def ArrayValue(colonnes, lignes):


    array = []
    print("Maintenant, ajoutez des valeurs à votre tableau : ")                                                         #Now add values to your table
    print("appuyez sur 'enter' après chacun d'eux)")                                                                    #press enter after each values
    for elem in range(lignes):                                                                                          #elem prend l'index des lignes // elem tooks the lines' index
        a = []                                                                                                          #On crée autant de tableau qu'il y a de ligne // Create as many arrays as there are rows
        for item in range(colonnes):                                                                                    #item pour l'index des colonnes //item for columns' index
            a.append(inputFloatN())                                                                                     #on ajoute la valeur au tableau // we add the value

        array.append(a)                                                                                                 #On ajoute les lignes avec leurs valeurs dans le tableau // we add lines with their values in the array


    print("Voici votre Tableau:", array, end=" ")                                                                       #On remplace le saut à la ligne par un espace // Replace line jumping with space
    print()
    StartOver(Elem)



