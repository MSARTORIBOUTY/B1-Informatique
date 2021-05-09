def StartOver(func):
    print("voulez-vous rejouer? (O/N)")
    answer = input()

    if (answer == "O" or answer == "o"):
        func()
    elif (answer == "N" or answer == "n"):
        print("Au revoir! ;)")
        return
    else:
        print("je n'ai pas compris votre réponse")
        func()

#----------------------------------------------------------------------------------
def inputInt(msg,msgerr="❌ Veuillez saisir un nombre entier."):
    while True:
        try:
            return int(input(msg))
        except:
            print(msgerr)

#----------------------------------------------------------------------------------
def inputFloat(msg,msgerr="❌ Veuillez saisir un nombre entier ou décimal."):
    while True:
        try:
            return float(input(msg))
        except:
            print(msgerr)

#----------------------------------------------------------------------------------
def YorN(answer, func, modif, aff):
    if (answer == "O" or answer == "o"):
        StartOver(func)
    elif (answer == "N" or answer == "n"):
        modif()
        aff()
    else:
        print("je n'ai pas compris votre réponse")
        StartOver(func)

#-----------------------------------------------------------------------------------
def validInput(item, mynotes):
    while True:

            for i in range(0 <= item <= 20):
                return i

            else:
                print("📚 Désolée, mais une note doit être comprise entre 0 et 20 📚 \n")
                mynotes()

#------------------------------------------------------------------------------------
def Answer(answer, liste):

    while True:

        if answer == 0 :
            print("d'accord! :p")
            break
        elif answer == 1 or answer == 2 or answer == 3 or answer == 4 or answer == 5 or answer == 6 or answer == 7 or answer == 8:
            print("---------------------------CONSIGNES DE MODIFICATIONS----------------------------------------------")
            print(
                "choix possibles des informations:", liste)
            print(
                "si vous n'entrez pas exactement l'une de ces valeurs lorsque l'information à modifier'vous sera demandé, une nouvelle ligne à la classe sera ajoutée")
            print(
                "il vous sera cependant possible de la supprimer! ^^ de ce fait il est également possible d'ajouter de nouvelles informations!")
            print(
                "-----------------------------------------------------------------------------------------------------")
            print()
            return answer
        else:
            print("(⊙_⊙)？...Je n'ai pas compris votre réponse")
            answer = inputInt("➢")


