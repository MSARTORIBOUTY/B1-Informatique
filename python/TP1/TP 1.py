#SARTORI BOUTY Marie-Laure
# en colaboration avec ALEM Lina 

#déclarer deux variables // report two variables
def exercice1():
        print("veuillez choisir un chiffre, essayez 65:")#please, choose a digit, try 65

        try:#essayer de faire la suite... // try to do the following...
                x=int(input())#conversion de l'input en int // convert input to int
        except:#...si il y a une erreur // ...if there's an error
                print("Vous êtes sûr que c'est un chiffre?")# are you sure it's a digit?
                exercice1() #recommencer // start over
                
        print("veuillez choisir une lettre, essayez 'A':")#please, choose a letter, try 'A'
        y=input()
        if(y.isdigit()):#si y est un chiffre // if y is a digit
                print("j'ai dit une lettre!")#"I said a letter!"
                exercice1()#Recommencer // start over

        if (len(y) > 1):#si y contient plus d'une lettre // if y has more than 1 letter
                print("j'ai dit UNE lettre! :)")#"I said A letter! :)"
                exercice1()#Recommencer // start Over

        try:#si il n'y a aucune erreur détecté, on fait la suite // if there's no error, we move on
                x=chr(x)#on converti les lettres et chiffres avec l'ASCII
                y=ord(y)# we convert digits and letters into ASCII
        except:#si y a une erreur // if there's an error
                print("Oops! Il y a due avoir une erreur quelque part, êtes-vous sûr d'avoir bien suivie les consignes?")
                exercice1()#Recommencer // start over


        print('x:', x,',', 'y:', y)#On affiche le résultat // display the result


        print("voulez-vous recommencer? (O/N)")
        result=input()
        if(result == "O" or result == "o"):
                exercice1()
        elif(result == "N" or result == "n"):
                fin()
        else:
                fin()


#faire un calcul avec deux variable // do a calculation with two variables
def exercice2():
    x=input("entrez la valeur de x:")#enter x value
    if(x.isdecimal() == False):#si x n'est pas une décimal // if x is not a decimal
            print("x et y doivent être des entiers, recommencez.")#x and y must be whole number, start over
            exercice2()
    y=input("entrez la valeur de y:")#enter y value
    if(y.isdecimal() == False):#si y n'est pas une décimal // if y is not a decimal
            print("x et y doivent être des entiers, recommencez.")
            exercice2()
    print('(x + y)*2 =', (int(x)+int(y))*2)#afficher le resultat // display the result


    print("voulez-vous recommencer? (O/N)")
    result=input()
    if(result == "O" or result == "o"):
            exercice2()
    elif(result == "N" or result == "n"):
            fin()
    else:
            fin()



#division de nombres réels et entiers //divide two real numbers and two whole numbers
def exercice3():
    
    print('Entrez la valeur de 2 réels. Commençons par le premier:')#enter the value of 2 real digits, let's start with the first
    x=float(input())#float est utiliser pour les nombres réels // float is used for real numbers
    print('et le deuxième:')#and the second
    y=float(input())
    if(y == 0):#si y est égale  à 0 // if y is equal to 0
            print("valeur invalide")#invalid value
            exercice3()#Recommencer // start over

    
    print('Entrez la valeur de 2 entiers:')#enter the value of two whole numbers
    a = input()
    if(a.isdecimal() == False):#si a n'est pas une décimal // if a is not a decimal
            print("a et b doivent être des entiers, recommencez.")#a and b must be whole number, start over
            exercice3()
    else:
        a=int(a)
    print('et le deuxième:')#and the second
    b = input()
    if(b == 0):
            print("valeur invalide")
            exercice3()
    if(b.isdecimal() == False):#si b n'est pas une décimal // if b is not a decimal
            print("a et b doivent être des entiers, recommencez.")#a and b must be whole number, start over
            exercice3()
    else:
            b=int(b)
    
    print('x/y =', x/y)#afficher le resultat
    print('a/b =', a/b)#display the result


    print("voulez-vous recommencer? (O/N)")
    result=input()
    if(result == "O" or result == "o"):
            exercice3()
    elif(result == "N" or result == "n"):
            fin()
    else:
            fin()
            

#convertir en ascii une lettre ou un chiffre // Converto to ascii a letter or a digit
def exercice4():
    print("voulez vous rentrer une lettre (l) ou un entier (e)?:")#do you want to enter a letter or a whole number
    choix= input()

    if (choix == 'l'):#si le choix est une lettre // if the choice is a letter
        print("entrez une lettre:")#please enter a letter
        phrase=input()
        if(phrase.isdigit() == True):#si c'est un nombre // if it's a digit 
                print("ce n'est pas une lettre.")#it's not a letter
                exercice4()#recommencer // start over
        if(len(phrase) > 1):#si il y a plus d'une lettre // if there's more than one letter
                print("entrez seulement une seule lettre!")#enter only one letter
                exercice4()#recommencer // start over
        else:
                print("conversion en ascii:", ord(phrase))#si tout est bon on affiche le résultat // if everything is okay, we display the result


    if(choix == 'e'):
        print("entrez un chiffre:")#please enter a digit
        chiffre=input()
        if (chiffre.isdecimal() == False):#si c'est pas un entier // if it's not a whole number
                print("vous n'avez pas rentré un nombre entier.")#You did not enter a whole number
                exercice4()
        print("conversion en ascii:", chr(int(chiffre)))#afficher le résultat // display the result
        pass


    print("voulez-vous recommencer? (O/N)")
    result=input()
    if(result == "O" or result == "o"):
            exercice4()
    elif(result == "N" or result == "n"):
            fin()
    else:
            fin()
            


#donner des conditions en fonction le l'input // give the condition based on the answer        
def exercice5():
    print("Voulez-vous jouer?")#do you want to play?
    reponse=input()

    if(reponse == 'O' or reponse == 'o'):
        print("c'est parti!")#let's go!
        exercice5()
    elif(reponse =='N' or reponse == 'n'):
        print("tant pis!")#too bad
        pass
    else:
            print("je n'ai pas compris votre réponse.")#I did not understand your answer
            exercice5()
            

    print("voulez-vous recommencer? (O/N)")
    result=input()
    if(result == "O" or result == "o"):
            exercice5()
    elif(result == "N" or result == "n"):
            fin()
    else:
            fin()

            

#calculer le ln d'un nombre
def exercice6():
    from math import log#import de la bibliothèque math // import math's library
    print("entrez un entier:")#enter a whole number
    entier=input()
    if(entier.isdecimal() == False): #si ce n'est pas un décimal // if it's not a decimal
        print("ce n'est pas un entier.")#it's not a whole number
        exercice6()#recommencer // start over
    print("logarithme népérien de", int(entier), ":", log(int(entier)))#afficher le résltat // display the result


    print("voulez-vous recommencer? (O/N)")
    result=input()
    if(result == "O" or result == "o"):
            exercice6()
    elif(result == "N" or result == "n"):
            fin()
    else:
            fin()

  
#compter le nombre d'entier positif entrés
#count the number of positive integers entered
def exercice7():
    print("entrez des entiers:")#enter whole number
    entier=input()
    
    if(entier.isdecimal() == False):#si ce n'est pas un nombre décimal // if it's not a decimal number
            print("ce n'est pas un nombre entier") #it's not a whole number
            exercice7()
            
    count = 0#pour compter le nombre de chiffre postif que nous allons entrer
             #to count how many positive number we have

    while(int(entier) > 0):#si le nombre ets > 0 //if the number is > 0

        entier = input()#on réécrit un chiffre // we write another number
        try:
                if(entier.isdecimal() == False):#si c'est pas un nombre décimal // if it's not a decimal number
                        if(int(entier) < 0):#si le chiffre est < 0 // if the number is <0
                              count +=1
                              print("nombre d'entiers positifs entrés:", count)
                              
                        else:
                              print("ce n'est pas un nombre entier")#it's not a whole number
                              exercice7()
                       
                else:#si tout va bien, on continue // if everything is okay, we continue
                        entier = int(entier)
                        count += 1
                        
        except:#si ce n'est pas un nombre // if it's not a number
                print("êtes-vous sûr d'avoir rentré un entier?")#Are you sure it was a whole number
                exercice7()

        
    
    print("voulez-vous recommencer? (O/N)")
    result=input()
    if(result == "O" or result == "o"):
            exercice7()
    elif(result == "N" or result == "n"):
            fin()
    else:
            fin()

    
def fin() :
        
        print("retour au menu?(O/N)")#"go back to menu?"
        answer=input()#take the answer and do the action it means
        if(answer == "O" or answer == "o"):
            menu()
        elif(answer =="N" or answer =="n"):
            pass
        else:
                print("je n'ai pas compris votre réponse")
                fin()

def menu():
    print("Entrez le numéro de l'exercice à lancer (1-7):")
    numero=int(input())

    if(numero == 1):
        exercice1()
    if(numero == 2):
        exercice2()
    if(numero == 3):
        exercice3()
    if(numero == 4):
        exercice4()
    if(numero == 5):
        exercice5()
    if(numero == 6):
        exercice6()
    if(numero == 7):
        exercice7()

menu()
