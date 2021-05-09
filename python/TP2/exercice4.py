import random

def rulesExo4():
    print("Bonjour et bienvenu dans l'exercice 4 ! :)")                                                                             #Hello and welcome in the fourth exercise
    print("-----------------------------------------------------421-----------------------------------------------------------")
    print("Le jeu du 421 est très simple")                                                                                          #The 421 game is very simple
    print("vous êtes en posséssion de 3 dés")                                                                                       #you have 3 diced
    print("Si l'un des 3 tombe sur la face '4', '2' ou '1' vous aurez le choix de le garder ou non")                                #If one of the 3 falls on the face '4', '2' or '1' you will have the choice to keep it or not
    print("Exemple: vous lancés les dés et obtenez '5 2 4' vous retenez 2 et 4 en notant 23 soit, leur position respective")        #Example: you roll the dice and get '5 2 4' you hold 2 and 4 noting 23 or, their respective position
    print("si vous choisissez de le garder alors lors de vos prochains lancés, les dés seront gardés en mémoire et seul le ou les dés non retenus changeront de valeur")   #if you choose to keep it then on your next roll, the dice will be kept in memory and only the unaccounted dice will change value                      #
    print("mais ceux retenus, s'afficheront normalement après le lancé de dé")                                                      #but those retained, will normally appear after the roll roll
    print("il est donc important de marquer leur position respectif pour les garder en mémoire au prochain tour")                   #it is therefore important to mark their respective positions to keep them in mind in the next round
    print("si vous dite '0' pour ne rien retenir, les 3 dés se réinitialiseront et vous reprendrez votre partie de 0")              #if you say '0' to retain nothing, the 3 dice will reset and you will resume your part of 0
    print("le but est donc d'obtenir une combinaison formant '421' dans un maximum de 3 tours afin de remporté la victoire")        #the goal being to obtain a combination forming '421' ina maximum of 3 laps in order to win
    print("dans le cas contraire....")                                                                                              #otherwise
    print("..... Vous avez perdu! :D")                                                                                              #it's a Game Over for you :D
    print("Prêt? A vous de jouer! ")                                                                                                #Ready? Here you go!
    print("Bonne chance! ;)")                                                                                                       #Good luck!
    print("________________________________________________________________________________________________________________")

    jeu421()


def jeu421():

    nbrLances = 1                                                                                                       #Nombre de fois où on lance les dés //number of times we roll the dice

    array = [1, 2, 3, 4, 5, 6]                                                                                          #toutes les valeurs libre du dé //all the free value of the dice
    de1 = random.choice(array)                                                                                          #les dés choisissent une valeur libre dans le tableau
    de2 = random.choice(array)                                                                                          #dice choose a free value of the array
    de3 = random.choice(array)


    while nbrLances <= 3:                                                                                               #on a le droit à trois lancés // we get three throws
        print("vous obtenez les dès suivants:")                                                                         #you get the following
        print("[", de1,"] [", de2, "] [", de3, "]")
        nbrLances += 1

        if ((de1 == 4 and de2 == 2 and de3 == 1) or (de1 == 2 and de2 == 4 and de3 == 1) or (                           #toutes les combinaisons gagnantes possible // all possible winning combinations
                de1 == 1 and de3 == 2 and de3 == 4)
                or (de1 == 1 and de2 == 4 and de3 == 2) or (de1 == 2 and de2 == 1 and de3 == 4) or (
                        de1 == 4 and de2 == 1 and de3 == 2)):
            print("Félicitation! Vous avez gagné au bout de", nbrLances, "coups !")                                     #Congratulations! you win!
            StartOver()


        try:
            choix1 = int(input("Que souhaitez-vous conserver (1 / 2 / 3 / Aucun (0)) ? "))                              #What do you want to keep?

            if choix1 not in (0, 1, 2, 3, 12, 21, 13, 31, 23, 32):                                                      #tout les choix possible // all possible choice
                print("Non! Un chiffre compris entre 1 et 3")                                                           #No! it's not a number between 1 and 3!
                jeu421()

            if choix1 == 1:
                print("Vous conservez ", de1)                                                                           #you keep:
                if de1 in array:
                    array.pop(de1)                                                                                      #on enlève la valeur du dé du tableau // we remove de dice's value from the array
                    de2 = random.choice(array)                                                                          #on relance les dés en conservant le dé gardé // roll the dice again, keeping the dice
                    de3 = random.choice(array)
                    continue

            if choix1 == 2:
                print("Vous conservez ", de2)
                if de2 in array:
                    array.pop(de2)
                    de1 = random.choice(array)
                    de3 = random.choice(array)
                    continue

            if choix1 == 3:
                print("Vous conservez ", de3)
                if de3 in array:
                    array.pop(de3)
                    de1 = random.choice(array)
                    de2 = random.choice(array)
                    continue


            if choix1 == 12 or choix1 == 21:
                print("Vous conservez ", de1, de2)
                if de1 and de2 in array:
                    array.pop(de1 and de2)
                    de3 = random.choice(array)
                    continue

            if choix1 == 23 or choix1 == 32:
                print("Vous conservez ", de2, de3)
                if de2 and de3 in array:
                    array.pop(de2 and de3)
                    de1 = random.choice(array)
                    continue

            if choix1 == 13 or choix1 == 31:
                print("Vous conservez ", de1, de3)
                if de1 and de3 in array:
                    array.pop(de1 and de3)
                    de2 = random.choice(array)
                    continue

            if choix1 == 0:
                print("Recommençons.")                                                                                  #let's start over
                de1 = random.choice(array)
                de2 = random.choice(array)
                de3 = random.choice(array)

        except ValueError:
            print("Non! Un chiffre compris entre 1 et 3")                                                               #No! it's not a number between 1 and 3!
            jeu421()

    else:
        print("vous obtenez les dès suivants:")                                                                         ##on a le droit à trois lancés // we get three throws
        print("[", de1, "] [", de2, "] [", de3, "]")
        print("Il semblerait que vous avez perdu! Dommage! :D")                                                         #Looks like you lost! Too bad! :D
        StartOver()






#-------------------------------------------------------------------------------------------------

def StartOver():
    print("voulez-vous rejouer? (O/N)")
    answer = input()

    if (answer == "O" or answer == "o"):
        jeu421()
    elif (answer == "N" or answer == "n"):
        pass
    else:
        print("je n'ai pas compris votre réponse")
        jeu421()