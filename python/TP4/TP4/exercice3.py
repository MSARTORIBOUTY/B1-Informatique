import TP4.exercice2
TP4.exercice2.personne()

import sys
sys.path.append('..')

from PUtilitaires import Answer
from PUtilitaires import YorN
from PUtilitaires import inputInt
from PUtilitaires import inputFloat
from PUtilitaires import validInput
from PUtilitaires import StartOver
from statistics import mean


def ruleExo3():
    print("Bonjour et bienvenu dans l'exercice 3! :)")                                                                  #Hello and welcome to the third exercise!
    print("-----------------------HERITAGE----------------------")
    print("ajouter, modifier des notes dans une classe qui hérite\n de la classe personne dans l'exercice d'avant ")    #add, edit notes in a class that inherits the person class in the previous exercise
    print("------------------------------------------------------")
    heritage()

def heritage():                                                                                                         #fonction qui appel la classe et ces fonction dans un certain ordre
    etudiant= Etudiant()                                                                                                #function who call the class and her functions in a specific order
    etudiant.__init__()
    etudiant.aff()

    question(etudiant.modif, etudiant.aff)
    StartOver(heritage)

def question(modif, aff):

    print("Êtes-vous satisfait des données?(O/N)")                                                                      #Are you satisfied with the data
    answer=input("➢")

    YorN(answer, heritage, modif, aff)

class Etudiant(TP4.exercice2.personne):
    global notes
    notes={}

    def __init__(self):                                                                                                 #on rappel les éléments de la classe dont Etudiant hérite, et on ajouter les éléments propre à la classe Etudiante
        TP4.exercice2.personne.__init__(self)
        notes['maths']= 15                                                                                              #recall the elements of the class that Student inherits, and add the elements specific to the Student class
        notes['français'] = 12
        notes['svt'] = 16
        notes['physique-chimie'] = 13
        notes['anglais'] = 14
        notes['sport'] = 9
        notes['histoire'] = 12
        notes['moyenne générale'] = 13

    def modif(self):                                                                                                    #modifie les éléments de la classe Etudiante, calcul les moyennes et les ajoutes au tableau
        print("combien d'éléments voulez-vous modifier? (1-8, 0 pour quitter)")                                         #modifies the elements of the Student class, calculates the averages and adds them to the table
        answer= inputInt("➢")
        liste= ['maths', 'français', 'svt', 'physique-chimie', 'anglais', 'sport', 'histoire']
        Answer(answer, liste)


        for i in range(answer):
            print("entrez l'information à modifier':")
            key = input("➢")
            print("entrez la valeur de l'information':")
            value = inputFloat("➢")
            mynotes = self.modif
            validInput(value, mynotes)


            try:
                notes[key] = [notes[key]]
                notes[key].append(value)
                print()
                print(notes)
                print()

                array = notes[key]
                moyenne_key = round(mean(array), 2)
                print(key, ":", moyenne_key)

                notes[key] = moyenne_key

                count = 0
                a = 0
                for i in notes:
                    count += 1
                    a += notes[i]
                print("La moyenne générale est ", round(a / count, 2), "/ 20")
                print()


            except:
                print("attention! il ne vous est possible que de changer les notes!")
                print("vous revenez aux valeurs par défaut en cas d'erreur! ")
                print("faites plus attention la prochaine fois")
                print()
                self.modif()


    def aff(self):                                                                                                      #affiche les éléments de la classe et calcule la moyenne générale en fonction de ce qu'elle contient
        TP4.exercice2.personne.aff(self)                                                                                #displays the elements of the class and calculates the overall average based on what it contains
        for clef, valeur in notes.items():
            print("➢",clef, ":", valeur)

        count = 0
        a = 0
        for i in notes:
            count += 1
            a += notes[i]
        print("La moyenne générale est ", round(a/count, 2), "/ 20")

