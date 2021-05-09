from random import randint

def rulesExo5():
    print("Bonjouer et bienvenu dans l'exercice 5! :)")                                                     #Hello and welcome to the fifth exercise! :)
    print("---------------Le jeu des allumettes-------------------")                                        #the game of matches
    print("La règle est simple")                                                                            #The rule is simple
    print("ne prenez pas la dernière allumette! :D")                                                        #don't take the last match! :D
    print("A vous de jouer! ;)")                                                                            #Here you go! ;)
    print("_____________________________________________________________")
    nbrMatches()

def nbrMatches():
    print("Avant de commencer, avec combien d'allumettes voulez-vous jouer?")                               #Before we start, how many matches do you want to play with?
    nbrA= input()

    if(nbrA.isdecimal() == False or int(nbrA) == 0):                                                        #si la réponse n'est pas un nombre entier positif > 0 // if the answer is not a positive whole number > 0
        print("un nombre positif supérieur à 0 dans la limite du raisonnable :)")                           #a prositive number greater than 0 within the reasonable limit
        nbrMatches()                                                                                        #On relance la fonction // start over the function
    else:                                                                                                   #sinon.. // otherwhise...
        nbrA = int(nbrA)
        WhoBegins(nbrA)

def WhoBegins(nbrA):
    print("maintenant que l'on a toutes nos allumettes")                                                    #now that we have all our matches
    print("déterminons qui commence!")                                                                      #we will determine who will start
    print("Voulez-vous que je choisisse? (O/N)")                                                            #shall I pick out?
    answer = input()

    if(answer == "o" or answer == "O"):                                                                     #si la réponse est oui // if the answer is yes
        random = randint(1, 2)                                                                              #Albert choisit un nombre entre 1 et 2 // Albert choose a number between 1 or 2
        if (random == 1):                                                                                   #si c'est 1 // if it's 1
            print("vous commencez! ^^")                                                                     #You start! ^^
            Joueur(nbrA)

        elif (random == 2):                                                                                 #si c'est 2 // if it's 2
            print("je commence! ^^")                                                                        #I start
            Albert(nbrA)


    elif(answer == "N" or answer == "n"):                                                                   #si la réponse est non // if the answer is no
        print("dans ce cas dites moi,")                                                                     #In that case tell me
        print("1 - Vous")                                                                                   #1 - you
        print("2 - Albert l'ordinateur")                                                                    #2 - Albert the computer
        result = input()

        if(result == '1'):                                                                                  #si c'est vous // if it's you
            print("D'accord! Allez-y")                                                                      #ok! Go ahead
            Joueur(nbrA)

        elif(result == '2'):                                                                                #Si c'est Albert // if it's Albert
            print("Très bien! j'espère que vous ne serez pas froissé si je vous bat! ^^")                   #Very well! I hope you will not be offended if I beat you! ^^
            Albert(nbrA)
        else:                                                                                               #si la reponse n'est pas 1 ou 2 // if the answer is not 1 or 2
            print("je n'ai pas compris")                                                                    #I didn't understand
            WhoBegins(nbrA)

    else:                                                                                                   #Si la réponse n'est ni oui ni non // if the answer is not yes or no
        print("je n'ai pas compris")                                                                        #I didn't understand
        WhoBegins(nbrA)                                                                                     #start over

def Joueur(nbrA):                                                                                           #la fonction du joueur // the player's function

    while (nbrA > 0):                                                                                       #tant qu'il y a des allumettes // as long as there are matches
        print("combient voulez vous enlever d'allumettes?")
        remove = int(input("choisissez un nombre entre 1 et 3 compris:"))                                   #choose a number between 1 to 3

        while(remove > nbrA or remove < 1 or remove > 3):                                                   #Si la réponse n'est pas 1, 2 ou 3 //if the answer is not 1, 2 or 3
            print("Oops! il y a une erreur, ressayez")                                                      #oops! there's an error, try again
            remove= int(input())

        nbrA-=remove                                                                                        #on enlève le nombre choisie au nombre d'allumettes // We removed the chosen number from the number of matches


        if(nbrA == 0):                                                                                      #si il n'y a plus d'allumettes // if there's no more matches
            print("vous avez perdu! dommage :)")                                                            #you lost! too bad :)
            StartOver()
        else:                                                                                               #Sinon... // Otherwhise...

            print("il reste", nbrA, "allumettes")                                                           #there are .... matches
            print("A mon tour de jouer!")                                                                   #my turn!
            Albert(nbrA)


def Albert(nbrA):                                                                                           #La fonction de l'IA // AI's function

    while (nbrA > 0):                                                                                       #tant qu'il y a des allumettes // as long as there are matches

        if (nbrA == 0):                                                                                     #si il n'y a plus d'allumettes // if there's no more matches
            print("vous avez perdu! dommage :)")                                                            #you lost! too bad :)
            StartOver()
        else:                                                                                               #Sinon... // Otherwhise...

            nbr = nbrA % 4                                                                                  #le reste de la division du nombre d'allumettes et 4 // the rest of the division of the number of matches and 4

            if (nbr == 0):                                                                                  #Si le reste est égale à 0 // if the rest is 0
                take = 3                                                                                    #Albert en prend 3 //Albert take 3 of them
            elif (nbr == 3):                                                                                #Si le reste est égale à 3 // if the rest is 3
                take = 2                                                                                    #Albert en prend 2 //Albert take 2 of them
            elif (nbr == 2):                                                                                #Si le reste est égale à 2 // if the rest is 2
                take = 1                                                                                    #Albert en prend 2 //Albert take 2 of them
            else:                                                                                           #Sinon... // Otherwhise...
                take = 1                                                                                    #Albert en prend 1 //Albert take 1 of them

            print("j'en prend", take)                                                                       #I take
            nbrA -= take                                                                                    #on enlève le nombre choisie au nombre d'allumettes // We removed the chosen number from the number of matches

            if (nbrA == 0):                                                                                 #si il n'y a plus d'allumettes // if there's no more matches
                print("Ha! On dirait bien que vous avez gagné, Félicitation! \^0^/")                        #Ha! Looks like you won, Congratulations! \^0^/
                StartOver()
            else:
                print("il reste", nbrA, "allumettes")                                                       #there are .... matches
                Joueur(nbrA)



def StartOver():
    print("voulez-vous rejouer? (O/N)")
    answer = input()

    if (answer == "O" or answer == "o"):
        nbrMatches()
    elif (answer == "N" or answer == "n"):
        pass
    else:
        print("je n'ai pas compris votre réponse")
        nbrMatches()

