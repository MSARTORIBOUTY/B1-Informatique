import sys
sys.path.append('..')

from PUtilitaires import StartOver
from PUtilitaires import inputInt
from PUtilitaires import inputFloat
from PUtilitaires import inferior

def rulesExo4():

    print("Bonjour et bienvenu dans l'exo4!")                                                                           #Hello and welcome in the fourth exercise!
    print("---------------------------FONCTION FACTORIELLE----------------------------------")                          #Factorial function
    print("après avoir choisit la valeur maximal de n ainsi que celle de x,")                                           #after choosing the maximum value of n as well as that of x
    print("l'algorithme vous calculera la fonction e^x et sin(x)")                                                      #The algorithm will calculate the e^x and sin(x)'s function
    print("A vous de jouer! ;)")                                                                                        #Here you go! ;)
    print("----------------------------------------------------------------------------------")
    inputN()

def inputN():


        Nmax=inputInt("entrez la valeur Nmax:")                                                                         #enter Nmax value
        inferior(Nmax, inputN)
        x= inputFloat("entrez la valeur de x:")                                                                         #enter x value
        inferior(x, inputN)

        functions(Nmax, x)


def functions(Nmax, x):

    e = 0
    f=0
    for n in range(0, Nmax+1):                                                                                          #n prend les valeurs de 0 à Nmax // n takes values from 0 to Nmax
        e += x**n/fact(n)                                                                                               #fonction e^x  // e^x's function

    for m in range(0, Nmax+1):                                                                                          #m prend les valeurs de 0 à Nmax // m takes values from 0 to Nmax
        f += ((-1)**m) * (x**(2*m+1))/fact(2*m+1)


    print("fonction e^x:", round(e,5))                                                                                  #afficher le résultat // display the result
    print("fonction sin(x):", round(f, 5))

    StartOver(inputN)


def fact(n):

    if n < 2:                                                                                                           #factoriel de 0 et 1 = 1 // factorial of 0 and 1 = 1
        return 1
    else:                                                                                                               #autrement //otherwhise
        return n * fact(n - 1)                                                                                          #factorielle de n = n * factorielle de n-1



