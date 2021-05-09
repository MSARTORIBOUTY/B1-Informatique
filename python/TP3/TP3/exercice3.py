import sys
sys.path.append('..')

from PUtilitaires import StartOver
from PUtilitaires import inputInt
from PUtilitaires import inferior


def rulesExo3():

    print("Bonjour et bienvenu à l'exercice 3! :)")                                                                     ##hello and welcome to the third exercise! :)
    print("-----------------------FIBONACCI---------------------")
    print("Le but de l'algorithme et d'afficher la suite de fibonacci,")                                                #The purpose of the algorithm and display the fibonacci suite
    print("jusqu'au terme où vous voulez qu'il s'arrête")                                                               #until the end you want it to stop
    print("Attention! seul les nombres entiers positifs, sont autorisés")                                               #Attention! only positive integers, are allowed
    print("A vous de jour! ;)")                                                                                         #Here you go! ;)
    print("------------------------------------------------------")
    fibonacciElem()

def fibonacciElem():

    print("entrez le nombre de fois que vous voulez lancer la suite:")                                                  #enter the number of times you want to launch the suite
    nbr = inputInt("➢")
    inferior(nbr, fibonacciElem)


    for i in range(nbr):                                                                                                #i prend toutes les valeurs de 0 à nbr // i takes all values from 0 to nbr
        print(fibonacci(i))

    StartOver(fibonacciElem)




def fibonacci(n):

    if n <=1 :
        return n
    else:
        return (fibonacci(n-1) + fibonacci(n-2))                                                                        #on prend les deux valeurs précédentes de la suite // we take the two previous values of the suite

