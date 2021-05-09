import sys
sys.path.append('..')


from PUtilitaires import inputInt
from PUtilitaires import YorN
from PUtilitaires import Answer
from PUtilitaires import StartOver


def ruleExo2():
    print("Bienvenu dans l'exercice 2! :)")                                                                             #Welcome to the second erercise
    print("--------------------------CLASSE------------------------")
    print("créer, modifier, supprimer des élements d'une classe")                                                       #Create, modify, delete elements from a class
    print("A vous de jouer!")                                                                                           #Here you go!
    print("---------------------------------------------------------")

    script()

def script():
    human=personne()                                                                                                    #la variable "humain" prend les valeurs de la classe personne // the "human" variable takes the values of the person class
    print()
    human.aff()                                                                                                         #on active la fonction afficher (aff) // enable the display function (aff)

    question(human.modif, human.aff)                                                                                    #On passe en argument les fonctions modifier et afficher // pass in arguments, modif (modify) and display (aff) functions

def question(modif, aff):

    print("Êtes-vous satisfait des données?(O/N)")                                                                      #Are you satisfied with the data
    answer=input("➢")

    YorN(answer, script, modif, aff)                                                                                    #On prend la réponse dans une autre fonction qui va afficher une réponse en fonction de la donnée//takes the answer in another function that we print the answer based on the variable "answer" this function is an import from PUtilitaires




class personne:                                                                                                         #configuration de la classe // configuration of the class
    global Personne                                                                                                     #déclaration de la variable en global pour être reconnu par toutes les autres ofnctions dans la classe//declared a variable in global, that way all functions in the same class will reconize the variable
    Personne={}                                                                                                         #la variable est un dictionnaire //the varaible is a dictionnary

    def __init__(self):                                                                                                 #La fonction init crée la classe avec tous les éléments dedans // init function create the class with all informations in it
        Personne["nom"] = "BOUTY"
        Personne["prenom"] = "Marie-Laure"
        Personne["âge"] = "23"
        Personne["nationalité"] = "Française"
        Personne["adresse"] = "1267 boulevard Jean Guigues"
        Personne["ville"] = "Pertuis"
        Personne["pays"] = "France"
        Personne["numéro de tel"] = "0652530912"

    def modif(self):                                                                                                    #fonction qui modifie la classe // modify the class
        print("combien d'éléments voulez-vous modifier? (1-8, 0 pour quitter)")                                         #How many elements do you want to modify ( 1 to 8 and 0 to quit
        answer= inputInt("➢")
        liste=['nom', 'prenom', 'âge', 'nationalité', 'adresse', 'ville', 'pays', 'numéro de tel']
        Answer(answer, liste)                                                                                           #La fonction Answer va verifier si la valeur de la variable est correct par rapport à ce qui est demandé // Answer function will verify if the answer's value is correct based on what is requested


        for i in range(answer):                                                                                         #i prend toutes les valeurs possibles entre 1 et la valeur de answer pour changer autant d'éléments que voulue // i takes all values between 1 and answer's value to change as many elements as wanted
            print("entrez l'information à modifier':")                                                                  #Enter the information to modify
            key = input("➢")
            print("entrez la valeur de l'information':")                                                                #enter the value of the information
            value = input("➢")
            Personne.update({key : value})                                                                              #On ajoute les changements au dictionnaire // we add changed to the dictionnary

        print("Voulez-vous supprimer un élement de votre liste ? Oui(1)/Non(2) ")                                       #do you want to delete an element from your list?
        answer = inputInt("➢")

        if answer == 1:
            key = input("Saisissez parfaitement le nom de l'information que vous voulez supprimer : ")                  #Write the name of the information that you would like to delete
            try:
                Personne.pop(key)                                                                                       #on enlève la clef en question du dictionnaire // we remove the key from the dictionnary
                print(Personne)                                                                                         #On affiche la classe // display the class
            except:
                print("je n'ai pas compris votre demande")                                                              #I didn't understand your answer
                StartOver(script)

        elif answer == 2:                                                                                               #si on ne veut rien supprimer // if we don't want to delete something
            pass                                                                                                        #Alors on ne fait rien//then we do nothing
        else:
            print("(⊙_⊙)？...je n'ai pas compris votre réponse...")                                                     #I didn't understand your answer


    def aff(self):                                                                                                      #affiche la classe // display the class
        print("Voici les renseignements sur la personne:")                                                              #here are the informations about the person
        for clef, valeur in Personne.items():                                                                           #affiche chaque clef avec sa valeur dans le dictionnaire // display each key with his value in the dictionnary
            print("➢",clef, ":", valeur)

