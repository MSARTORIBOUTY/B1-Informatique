#SARTORI BOUTY Marie-Laure
#En collaboration avec ALEM Lina


from exercice1 import rulesExo1
from exercice2 import introduction
from exercice3 import ruleExo3
from exercice4 import rulesExo4
from exercice5 import rulesExo5
from exercice6 import rulesExo6
from exercice7 import rulesExo7
from exercice8 import rulesExo8

def menuexo():
    print("Bienvenu sur le TP2!")
    print("Entrez le numéro de l'exercice à lancer (1-8):")
    numero=int(input())

    if numero == 1:     #clear
        rulesExo1()
        fin()
    if numero == 2:
        introduction()  #clear
        fin()
    if numero == 3:
        ruleExo3()      #clear
        fin()
    if numero == 4:
        rulesExo4()
        fin()
    if numero == 5:     #clear
        rulesExo5()
        fin()
    if numero == 6:     #clear
        rulesExo6()
        fin()
    if numero == 7:     #clear
        rulesExo7()
        fin()
    if numero == 8:
        rulesExo8()
        fin()


def fin():
    print("retour au menu?(O/N)")  # "go back to menu?"
    answer = input()  # take the answer and do the action it means
    if (answer == "O" or answer == "o"):
        menuexo()
    elif (answer == "N" or answer == "n"):
        pass
    else:
        print("je n'ai pas compris votre réponse")
        fin()


menuexo()