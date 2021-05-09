def rulesExo1():
    print("Bonjour et bienvenu dans l'exercice 1! :)")                                                  #Hello and welcome in the first exercise
    print("Cet exercice consiste à entrer un chiffre de votre choix pour peut que ce soit un entier")   #This exercise consist of entering a number of your choice. It must be an integer
    print("l'algorithme ce chargera de calculer son factorielle!")                                      #The algorithm will load to calculate his factorial
    print("A vous de jouer! ;)")                                                                        #Here you go!
    print("__________________________________________________________________________________________")
    factorial()

def factorial():

    print("saisissez un entier n :")                                #enter a whole number n
    n= input()



    if (n.isdecimal() == False or int(n) == 0):
        print("ce n'est pas un nombre entier, recommencez.")        #It's not a whole number, start over
        factorial()

    while True :                                                    #si il n'y a pas d'erreur // if there's no error
        factoriel = 1                                               #le factoriel prend la valeur de 1 au début // the factorial takes the value of 1 at the begining
        for i in range(1, int(n) + 1):                              #i prend la valeur de 1 à n + 1 // i takes the value of 1 to n+1
            factoriel *= i                                          #on multiplie le factoriel à i //we multiply the factorial by
        print("la factorielle de", n, "est:", factoriel)              #the factorial of n is
        break                                                       #on sort de la boucle //we stop

    else:
        print("êtes-vous sûr que c'est un nombre entier?")          #are you sure it's a whole number?
        factorial()                                                 #Start over

#----------------------------------------------------------------------------------------------------------------------------

    print("voulez-vous recommencer? (O/N)")
    result=input()
    if(result == "O" or result == "o"):
            factorial()
    elif(result == "N" or result == "n"):
            pass
    else:
            print("je n'ai pas compris")
            fin()