def rulesExo7():
    print("Bonjour et bienvenu dans l'exercice 7! :)")                                                                  #Hello and welcome to the seventh exercise
    print("---------------------------------------LES SUITES-----------------------------------------------------")     #arithmetic suites
    print("Cet exercice consiste à calculer une suite: Un = a+ b/U(n-1)")                                               #this exercise consists to calculate an arthmetic suites: Un = a + b/U(n-1)
    print("vous déterminerez la valeur de U0 ainsi que Nmax, la limite et le chiffre a et b qui forment la suite")      #you will determine the value of U0 as well as Nmax, the limit and the number a and b that form the following
    print("les réels sont autorisés ainsi que les nombres négatifs")                                                    #the real ones are authorized as well as the negative numbers
    print("A vous de jouer! ;)")                                                                                        #Here you go! ;)
    data()

def data():
    print("déterminez la valeur de a et b")                                                                             #determine the value of a and b
    a= input("a:")

    try:
        a = float(a)
    except:                                                                                                             #Si a n'est pas un nombre // if a is not a digit
        print("un chiffre s'il vous plaît!")                                                                            #a number please
        data()

    b= input("b:")

    try:
        b = float(b)
    except:                                                                                                             #de même pour b // same for b
        print("un chiffre s'il vous plaît!")
        data()

    print("maintenant entrez la valeur de U0")                                                                          #Now enter the value of U0
    U0= input("U0:")

    try:
        U0 = float(U0)
    except:                                                                                                             #on vérifie pour U0 // we do the same thing for U0
        print("un chiffre s'il vous plaît!")
        data()

    print("maintenant entrez la valeur de Nmax")                                                                        #Now enter the value of Nmax
    Nmax= input("Nmax:")

    try:
        Nmax = int(Nmax)
    except:                                                                                                             #Rebelote // we do the same
        print("un chiffre s'il vous plaît!")
        data()

    print("et l'epsilon")
    eps= input("epsilon:")

    try:
        eps = float(eps)
    except:                                                                                                             #Pareil // and again
        print("un chiffre s'il vous plaît!")
        data()

    suite(a, b, U0, Nmax, eps)

def suite(a, b, U0, Nmax, eps):
    U=U0                                                                                                                #U prend la valeur initiale //U takes the initial value
    count = 0
    option=0
    Un = a+(b/U)                                                                                                        #ici Un est équivalent à U(n+1), c'est surtout pour déclarer la variable // here, Un is equal to U(n+1), it's mostly to declare the variable
    e = abs(Un-U)                                                                                                       #e : valeur absolue de U(n-1) - Un //e: absolute value equal to : U(n-1) - Un

    for n in range(0, Nmax):                                                                                            #n prendra la valeur de 0 à Nmax inclus // n will take the 0 value until Nmax included
        while (e > eps):                                                                                                #tant que e > epsilon//as long as e > epsilon
            Un = U                                                                                                      #Un prend la valeur précédente (pour l'epsilon) // Un take the previous value (for the epsilon)
            U = a+(b/Un)                                                                                                #ici : Un is Un-1 value and U becomes U(n+1) value
            e = abs(U-Un)                                                                                               #On actualise la valeur de e // the value of e is discounted
            count +=1                                                                                                   #count is the number of iteration

        if eps > 0:
            result = U-Un
            if ( result >= 0 and result < eps):                                                                         #0 <= result < eps
                option = 1
            if(U >= 0 and U < Un and Un < Nmax):                                                                        #0 <= U(n-1) < Un < Nmax
                option = 2

    print("Suite Un="+str(a)+"+"+str(b)+"/(Un-1) avec:")                                                                #On affiche les résultats // display the results
    print("U0=", U0)
    print("Nmax=", Nmax)
    print("epsilon=", eps)

    if(option == 1):                                                                                                    #and if the suite is converging....
        print("0 <= Un - U(n-1) <",eps)
        print("la suite est donc convergente")
    elif(option == 2):                                                                                                  #...or diverging
        print("0 <= Un < U(n-1) <", Nmax)
        print("la suite est donc divergente")

    print("Résulat", round(U, 5), "pour un nombre d'itération n=", count)
    StartOver()

def StartOver():
    print("voulez-vous rejouer? (O/N)")
    answer = input()

    if (answer == "O" or answer == "o"):
        data()
    elif (answer == "N" or answer == "n"):
        pass
    else:
        print("je n'ai pas compris votre réponse")
        data()