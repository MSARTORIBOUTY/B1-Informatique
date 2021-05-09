def rulesExo8():
    print("Bonjouer et bienvenu dans l'exercice 8! :)")                                                                 #Hello and welcome to the eith exercise! :)
    print("__________________________________________LES INTEGRALES_____________________________________________")
    print("Cet exercice consiste à calculer l'air sous la courbe d'équation 'y=x^2'")                                   #This exercise involves calculating the air below the equation curve
    print("je vous demanderais pour cela de rentrer la valeur de a, b et le pas (p)")                                   #I would ask you to enter the value of a, b and step (p)
    print("les lettres sont bien évidemment interdites!")                                                               #letters are obviously forbidden!
    print("pour les valeurs de a et b, seuls des entiers sont autorisés avec une valeur supérieur à 0 et a < b")        #for values of a and b, only integers are allowed with a value greater than 0 and a < b
    print("A vous de jouer! ;)")                                                                                        #Ready? Here you go! ;)
    print("_____________________________________________________________________________________________________")
    integral()

def integral():

    print("entre la valeur de a:")                                                                                      #enter the value of a
    a= input()

    print("entrez la valeur de b:")                                                                                     #enter the value of b
    b=input()


    try:
        a= int(a)                                                                                                       #si a et b ne sont pas des entiers, qu'il sont négatifs et que b < 0
        b=int(b)                                                                                                        #if a and b are not whole numbers, they are negative and b < 0
        if( a < 0 and b <= 0):
            print("il semblerait que vous n'ayez pas compris les consignes")                                            #it seems that you didn't understand the rules
            rulesExo8()                                                                                                 #retour à la fonction des règles // go back to rules' function

    except:
        print("il semblerait que vous n'ayez pas compris les consignes")                                                #it seems that you didn't understand the rules
        rulesExo8()


    if(a > b):
        print(" a doit être inférieur à b!")                                                                            #a must be inferior to b
        integral()

    print("entrez la valeur de p:")                                                                                     #enter the value of p
    p= input()

    try:
        p= float(p)
        if( p <= 0):
            print("il semblerait que vous n'ayez pas compris les consignes")                                            #it seems that you didn't understand the rules
            rulesExo8()

    except:
        print("il semblerait que vous n'ayez pas compris les consignes")                                                #it seems that you didn't understand the rules
        rulesExo8()


    calculate(a, b, p)






def calculate(a, b, p):
    c = a                                                                                                               #c prend la valeur de a //c takes the value of a
    integrale = float(0)
    while a < b :                                                                                                       #pour faire l'interval [a,b] // to do the interval [a,b]
        integrale += float(((a+p)**3/3) - ((a**3)/3))                                                                   #on prend la primitive de x^2, on enlève la valeur de a pour avoir juste le rectangle du pas // we take the primitive of x^2 and we remove the value of a to just keep the value of the rectangle of the step
        a += p                                                                                                          #on ajoute la valeur p à a pour le rectangle suivant // we add the p value to a for the next rectangle

    print("Calcul de l’intégrale de la fonction y = x^2 ")                                                              #calculation of the integral of the y=x^2 function
    print("avec", c, "< x <", b," et p =", p)                                                                           #with a < x < b and p =
    print("résultat:", round(integrale, 1))                                                                             #result
    StartOver()


def StartOver():
    print("voulez-vous rejouer? (O/N)")
    answer = input()

    if (answer == "O" or answer == "o"):
        integral()
    elif (answer == "N" or answer == "n"):
        pass
    else:
        print("je n'ai pas compris votre réponse")
        integral()