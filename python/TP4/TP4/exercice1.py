import sys
sys.path.append('..')

from PUtilitaires import StartOver
from PUtilitaires import inputInt

def ruleExo1():
    print("Bonjour et bienvenu dans l'exercice 1! :)")                                                                  #Hello and welcome in the first exercise! :)
    print("------------------------------------EDITER UN FICHIER----------------------------")
    print("Creer, lisez, modifiez un fichier!")                                                                         #create, modify and read a file
    print("----------------------------------------------------------------------------------")
    Choose()

def Choose():
    print("voulez-vous un fichier texte (1) ou binaire (2)?")                                                           #do you want a text file or a binary file
    answer = inputInt("➢")

    if answer ==1 :                                                                                                     #1 est un fichier texte // 1 is a text file
        editTxt()
    elif answer == 2:                                                                                                   #2 est un fichier binaire // 2 is a binary file
        editBin()
    else:
        print("Réponse incorrecte! :(")                                                                                 #incorrect answer
        Choose()


def editBin():
    print("Comment voulez-vous appeler le fichier bin?")                                                                #How do you want to call your bin file?
    Bin=input()
    Bin = Bin + ".bin"                                                                                                  #On ajoute .bin au nom //we add .bin to the name

    try:
        file = open(Bin, "xb")                                                                                          #On essaye d'ouvrir le fichier pour vérifier si il existe déjà //We try to open the file to check if it already exist
        bin(Bin, file)                                                                                                  #si il n'existe pas on continue // if it do not exist we continue

    except :
        print("il semblerait que le fichier existe déjà!")                                                              #it seems that the file already exists
        existe(Bin)                                                                                                     #Il y a une étape supplémentaire // There is an additional step

def editTxt():                                                                                                          #for the text file it's exactly the same as binary, the only things that chnage are the permissions
    print("Comment voulez-vous appeler le fichier texte?")                                                              #permissions need to add a b for binary but that's all
    Txt=input()                                                                                                         #Pour les fichiers textes, c'est exactement la même choses, seul les permissions changent. On ajoute un b pour binaire
    Txt = Txt + ".txt"

    try:
        file = open(Txt, "x")
        txt(Txt, file)

    except :
        print("il semblerait que le fichier existe déjà!")
        existeT(Txt)

def txt(Txt, file):
    a = input("a:")
    b = input("b:")


    file.write(a)
    file.write("\n")
    file.write(b)
    file.write("\n")
    file.close()

    with open(Txt, 'r') as file:
        for ligne in file:
            print(ligne)

    StartOver(Choose)

def bin(Bin, file):
    a = inputInt("a:")                                                                                                  #on prend les deux éléments à ajouter dans le fichier
    b = inputInt("b:")                                                                                                  #we take two elements to add in the file

    arr=[a, b]                                                                                                          #On les met dans un tableau // we add them in an array
    file.write(bytearray(arr))                                                                                          #On cast le tableau en byte et on l'ajoute au fichier // we cast the array in byte and we add it in the file
    file.close()

    with open(Bin, 'rb') as file:                                                                                       #On lit le fichier // Read the file
        print(list(file.read()))

    StartOver(Choose)

def existe(Bin):                                                                                                        #Si le fichier existe // if the file already exist
    print("Changer le nom (1), Ajouter (2) ou écraser (3) le fichier?")                                                 #Change the name, add or overwrite?
    answer=inputInt("➢")
    if answer ==1 :                                                                                                     #Changer le nom --> retour à la fonction où on nomme le fichier //Change the name --> go back to the function where we call the file
        editBin()
    elif answer == 2:                                                                                                   #On ajoute au fichier // add to the file
        file = open(Bin, "ab")
        bin(Bin, file)
    elif answer == 3:
        file = open(Bin, "wb")                                                                                          #On écrase le fichier // we overwrite
        bin(Bin, file)
    else:
        print("Réponse incorrecte! :(")                                                                                 #incorrect answer
        existe(Bin)

def existeT(Txt):
    print("Changer le nom (1), Ajouter (2) ou écraser (3) le fichier?")
    answer=inputInt("➢")
    if answer ==1 :
        editTxt()
    elif answer == 2:
        file = open(Txt, "a")
        txt(Txt, file)
    elif answer == 3:
        file = open(Txt, "w")
        txt(Txt, file)
    else:
        print("Réponse incorrecte! :(")
        existeT(Txt)