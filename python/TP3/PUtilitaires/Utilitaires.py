def StartOver(func):
    print("voulez-vous rejouer? (O/N)")
    answer = input()

    if (answer == "O" or answer == "o"):
        func()
    elif (answer == "N" or answer == "n"):
        print("Au revoir! ;)")
        pass
    else:
        print("je n'ai pas compris votre réponse")
        StartOver(func)

#----------------------------------------------------------------------------------
def inputInt(msg,msgerr="❌ ce n'est pas un nombre entier!"):
    while True:
        try:
            return int(input(msg))
        except:
            print(msgerr)

#----------------------------------------------------------------------------------
def inputFloat(msg,msgerr="❌ Valeur incorrecte."):
    while True:
        try:
            return float(input(msg))
        except:
            print(msgerr)
#----------------------------------------------------------------------------------
def inputFloatN(msgerr="❌ Valeur incorrecte."):
    while True:
        try:
            return float(input())
        except:
            print(msgerr)
#----------------------------------------------------------------------------------

def inferior(nbr, func, msgerr="❌ Valeur incorrecte."):
    if nbr <= 0:
        print(msgerr)
        func()