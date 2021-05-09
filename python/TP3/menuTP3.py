#SARTORI BOUTY Marie-Laure
#En collaboration avec ALEM Lina

from TP3.exercice1 import rulesExo1
from TP3.exercice2 import rulesExo2
from TP3.exercice3 import rulesExo3
from TP3.exercice4 import rulesExo4
from PUtilitaires import inputInt

def menuexo():
    print("Bienvenu sur le TP3!")
    print("Entrez le numéro de l'exercice à lancer (1-4 ou 0 pour quitter):")
    numero=inputInt("➢")


    if numero == 1:
        rulesExo1()
        menuexo()
    elif numero == 2:
        rulesExo2()
        menuexo()
    elif numero == 3:
        rulesExo3()
        menuexo()
    elif numero == 4:
        rulesExo4()
        menuexo()
    if numero == 0:
        pass



