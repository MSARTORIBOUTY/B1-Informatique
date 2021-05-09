def introduction():
    print("Bonjour et bienvenu dans l'exercice 2! :)")                                  #Hello and welcome to this second exercise
    print("Il s'agit tout simplement d'une calculatrice toute simple!")                 #It's just a simple calculator
    print("seul les chiffres sont autorisés!")                                          #Only digits are allowed
    print("A Vous de jouer! ;)")                                                        #Here you go! :)
    print("____________________________________________________________________")
    calculator()

def calculator():

    print("que voulez-vous faire?")                                                     #What do you want to do?
    print("1 - additioner")                                                             #1 - add
    print("2 - soustraire")                                                             #2 - subtract
    print("3 - multiplier")                                                             #3 - multiply
    print("4 - diviser")                                                                #4 - divide
    choose = input("entrez l'option choisie:")                                          #enter the option you want

    if (choose.isdecimal() == False or int(choose) > 4 or int(choose) == 0) :           #si l'option entrée n'est pas un entier // if the option entered is not an integer
        print("je n'ai pas compris votre demande")                                      #I did not understand your request
        calculator()                                                                    #Start over

    elif(choose == '1' or '2' or '3' or '4'):                                           #Si l'option est 1, 2, 3, ou 4 //if the option is 1, 2, 3, or 4

        calculation(choose)


def calculation(choose):
    nbr1 = input("entrez le premier nombre:")                                           #enter the first number

    try:                                                                                #Si c'est possible...// If it's possible...
        nbr1 = float(nbr1)                                                              #conversion de l'input en float // conversion input to float
    except:                                                                             #...Si non // ...If not
        print("ce n'est pas un nombre!")                                                #It's not a number
        calculation(choose)                                                             #Recommencer (on garde en mémoire l'option choisie) // Start over (We keep in memory the choosen option)

    nbr2 = input("entrez le deuxième nombre:")                                          #Enter the second number
    try:                                                                                #Si c'est possible...// If it's possible...
        nbr2 = float(nbr2)                                                              #conversion de l'input en float // conversion input to float
    except:                                                                             #...Si non // ...If not
        print("ce n'est pas un nombre!")                                                #It's not a number
        calculation(choose)                                                             #Recommencer (on garde en mémoire l'option choisie) // Start over (We keep in memory the choosen option)

    if (choose == '1'):                                                                 #On définit la bonne fonction de calcul avec l'option choisie
        add(nbr1, nbr2)                                                                 #We define the correct calculation function with the chosen option

    elif (choose == '2'):
        subtract(nbr1, nbr2)

    elif (choose == '3'):
        multiply(nbr1, nbr2)

    elif (choose == '4'):
        if(nbr2 == 0):
            print("la valeur du deuxième nombre est invalide, il doit êtes différent de 0") #invalide value of the second number. It must be different to 0
            calculation(choose)
        divide(nbr1, nbr2)


#fonction de calcul // Calculation function
def add(x, y):
    print(x, "+", y, "=", round(x+y, 2))
    StartOver()                                                                         #demander si on fait un autre calcul // at the end, the algorithm ask if we want to do another calculation

def subtract(x, y):
    print(x, "-", y, "=", round(x-y, 2))
    StartOver()

def multiply(x, y):
    print(x, "*", y, "=", round(x*y, 2))
    StartOver()

def divide(x, y):
    print(x, "/", y, "=", round(x/y, 2))
    StartOver()


def StartOver():
    print("voulez-vous calculer autre chose? (O/N)")                                     #Do you want to calculate anything else?
    answer = input()

    if (answer == "O" or answer == "o"):
        calculator()
    elif (answer == "N" or answer == "n"):
        pass
    else:
        print("je n'ai pas compris votre réponse")
        calculator()