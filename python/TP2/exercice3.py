from random import randint

def ruleExo3():
    print("Bonjour et bienvenu dans l'exercice 3 ! :)")                                                         #hello and welcome to the third exercise! :)
    print("Les règles du jeu sont simple!")                                                                     #the rules are simple
    print("après avoir défini un interval, par exemple [-10, 10]")                                              #after you have choosen the interval, for exemple [-10, 10]
    print("je choisirais un nombre dans ce dernier et vous devrez le trouver!")                                 #I will choose a whole number in it and you will have to find it
    print("plutôt simple non? ^^")                                                                              #It's simple, no? ^^
    print("A vous de jouer! ;)")                                                                                #Here you go! ;)
    print("_________________________________________________________________________________________")
    interval()

def interval():
    print("avant de commencer je vais vous demander de me donner un chiffre entier et positif, un seul!")       #Before to begin, I will ask you to give me a positive integer, only one!
    num = input()

    if(num.isdecimal() == False or int(num) == 0):                                                              #Si les condition du nombre entier ne sont pas respectés // if the conditions of the integer arent respected
        print("Non, j'ai dit entier ET positif! Egalement différent de 0")                                      #No, I said a positive integer different from 0
        interval()                                                                                              #On recommence // start over
    else:
        num = int(num)
        exercice3(num)

def exercice3(num):
    s=randint(-num, num)                                                                                        #On définit l'interval avec le chiffre entré // we define the interval with the input
    print("Le nombre à trouver est compris entre -" + str(num) + " et", num)                                    #the number to be found is included between
    print("entrez un nombre:")                                                                                  #enter a number
    n=input()

    try:                                                                                                        #Si on ne peut pas le convertir en int... // if it can't be converted to int...
        n=int(n)
        StayInTheInterval(n, num)                                                                               #On vérifie que le chiffre reste bien dans l'interval // Check that the number remains in the interval
        count = 1
    except:                                                                                                     #...On relance la fonction avec les mêmes paramètres // ...Restart the function with the same parameters
        print("ce n'est pas un nombre")                                                                         #it's not a number
        exercice3(num)


    while (s != n):                                                                                             #Tant que le nombre entier est différent du nombre choisit par l'IA // while the number is different from the one choosen by the AI
        if int(n) > s:                                                                                          #Si le nombre entier est > à celui choisit // if the number is bigger thant the one selected
            print("Trop grand")                                                                                 #too big
        else:                                                                                                   #dans l'autre cas // In another case
            print("Trop petit")                                                                                 #too small
        n= input()                                                                                              #Prend nouveau chiffre // take another number

        try:                                                                                                    #On fait la même chause que plus haut // We do the same thing as above
            n = int(n)
            StayInTheInterval(n, num)
            count += 1                                                                                          #On compte le nombre de fois où on a besoins d'entrer un nombre jusqu'à trouver le bon // We count how many time we need to enter a number until we find the good one
        except:
            print("ce n'est pas un chiffre!")
            StartOverButKeep(num)


    print("Bravo! Le chiffre est bien", s, ", vous avez gagné en ", count, "coup(s)")
    StartOver()

def StayInTheInterval(n, num):                                                                                  #Fonction qui vérifie si on reste bien dans l'interval // function that makes sure we stay in the interval
    if (n < -num or n > num):
        print("Vous avez entrez un chiffre en dehors de l'interval....")                                        #You entered a number outside the interval...
        StartOverButKeep(num)

def StartOverButKeep(num):                                                                                      #function that ask if we want to start over again with the same parameters

    print("voulez-vous recommencer la partie en conservant le même interval? (O/N)")                            #Do you want to play again by keeping the same interval?
    answer = input()

    if (answer == "O" or answer == "o"):
        exercice3(num)
    elif (answer == "N" or answer == "n"):
        StartOver()
    else:
        print("je n'ai pas compris votre réponse")                                                              #I didn't understand your answer
        StartOverButKeep(num)


def StartOver():
    print("voulez-vous faire une nouvelle partie? (O/N)")                                                       #Do you want to play again?
    answer = input()

    if (answer == "O" or answer == "o"):
        interval()
    elif (answer == "N" or answer == "n"):
        pass
    else:
        print("je n'ai pas compris votre réponse")                                                              #I didn't understand your answer
        interval()