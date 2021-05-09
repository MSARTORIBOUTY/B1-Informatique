import sys
sys.path.append('..')

from PUtilitaires import StartOver


import random

def rulesExo1():

    print("Bonjour et bienvenu dans l'exercice 1! :)")                                                                  #Hello and welcome to the first exercise! :)
    print("---------------------------CHIFFRE ALEATOIRE---------------------------")                                    #Random number
    print("vous allez devoir rentréer des chiffres afin de créer une intervalle")                                       #you will need to enter numbers to create an interval
    print("évitez les lettres !")                                                                                       #do not enter letters!
    print("l'algorithme va vous donner un nombre alléatoire dans cette dernière ")                                      #the algorithm will give you a random number in it
    print("vous aurez la possibilité de choisir si ce sera un entier ou un réel")                                       #you will have the possibility to choose between a real or a whole number
    print("A vous de jouer! ;)")                                                                                        #Here you go!
    print("-----------------------------------------------------------------------")
    Interval()



def Interval():


    m = input("entrez la valeur initial:")                                                                              #enter the initial value
    n = input("entrez la valeur final:")                                                                                #enter the last value



    if(n.isdecimal()== False or m.isdecimal() == False):                                                                #si n et m ne sont pas des chiffres entiers // if n and m are note whole numbers
        try:
            m=round(float(m))                                                                                           #on arrondit leur valeur avec round en les castant en float // round their value and cast them in float
            n=round(float(n))
        except ValueError:                                                                                              #si c'est pas possible (parce que ce sont des lettres) // if it's not possible (because they are letters)
            print("il semblerait que vous n'ayez pas rentré un chiffre...")                                             #it seems that you did not enter a number
            Interval()

    m= int(m)                                                                                                           #on cast m et n en int //m and n becomes interger
    n=int(n)

    if n < m:                                                                                                           #si la valeur final est plus petite que l'initial // if the final value is bigger than the initial value
        print("Alors... Dans la logique des choses... initial < final.")                                                #So... in the logic of things... initial < final
        Interval()
    else:
        lance_de(m, n)


def lance_de(m, n):
    print("Maintenant il est temps de choisir si vous voulez que je vous donne un réel ou un entier! ^^")               #Now, it's time to choose if you want me to give you a whole or a real number
    print("R pour 'réel' et E pour 'entier' ^^")                                                                        #R for real and E for whole number
    answer = input()

    if (answer == 'R' or answer == 'r') :
        nbr = random.uniform(m, n)                                                                                      #valeur random d'un réel// random value for a real number
        print("voici la valeur que j'ai choisi pour vous :",round(nbr, 2))                                              #this is the value I have choosen for you

    elif (answer == 'E' or answer == 'e') :
        nbr = random.randint(m, n)                                                                                      #valeur random d'un entier // random value of a integer
        print("voici la valeur que j'ai choisi pour vous :",nbr)
    else:
        print(".... N'étais-je pas assez clair? >:(")                                                                   #Was I not clear enough?
        lance_de(m, n)


    StartOver(Interval)




