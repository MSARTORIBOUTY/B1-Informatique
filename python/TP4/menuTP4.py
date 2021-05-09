#SARTORI BOUTY Marie-Laure
#En collaboration avec ALEM Lina

from TP4.exercice1 import ruleExo1
from TP4.exercice2 import ruleExo2
from TP4.exercice3 import ruleExo3
from TP4.exercice4 import Calculatrice
from TP4.exercice5 import ruleExo5


def menuexo():
    print("Bienvenu sur le TP4!")
    print("Entrez le numéro de l'exercice à lancer (1-5 ou 'e' pour quitter):")

    numero=input()


    if numero == '1':
        ruleExo1()
        menuexo()
    elif numero == '2':
        ruleExo2()
        menuexo()
    elif numero == '3':
        ruleExo3()
        menuexo()
        menuexo()
    elif numero == '4':
        Calculatrice()
        menuexo()
    elif numero == '5':
        ruleExo5()
        menuexo()


    elif numero == 'e' or numero == 'E' :
        pass
    else:
        print("je n'ai pas compris votre demande :(")
        menuexo()



menuexo()