def rulesExo6():
    print("Bonjour et bienvenu dans l'exercice 6! :)")                                                                  #hello and welcome in the sixth exercise! :)
    print("_______________________CALCULER LA MOYENNE_______________________________________")                          #calculate the average
    print("Cet algorithme consiste à calculer la moyenne de chiffres que vous allez entrer")                            #This algorithm consists of calculating the average number you will enter
    print("attention! Si vous entrez une lettre vous devrez tout recommencer... ce serait dommage! ;)")                 #Attention If you enter a letter you will have to start all over again... it would be a shame
    print("les réels, les chiffres négatifs et les entiers sont autorisés")                                             #real, negative and integers are allowed
    print("A vous de jouer! ;)")                                                                                        #Here you go! ;)
    numberOfDigits()

def numberOfDigits():
    print("_____________________________________________________________________________________________________")
    print("avant de vous faire commencer, indiquez moi le nombre de chiffres que vous voulez entrer.")                  #before to begin, tell me the number of numbers you want to enter
    num = input()
    if(num.isdecimal() == False or num == '0'):                                                                         #si ce n'est pas un nombre entier > 0 // if it's not a whole number > 0
        print("un nombre entier et supérieur à 0 s'il vous plaît!")
        numberOfDigits()
    else:
        num = int(num)
        nombres(num)

def nombres(num):
    arr = []
    count = 0
    print("entrez les chiffres un à un:")                                                                               #enter the numbers one by one

    while (count < num) :                                                                                               #tant que le nombre de chiffre entré est < au nombre de chiffre désiré // as long as the number of digits entered is < of the number of digits desired
        nbr = input()
        count +=1                                                                                                       #compte le nombre de chiffres entrés // it will count how many number the player will enter

        try:
            nbr = (float(nbr))
            arr.append(nbr)                                                                                             #On ajoute les nombres dans un tableau // we add the number in an array

        except:
            print("Il semblerait que vous n'ayez pas compris les règles....")                                           #It seems you didn’t understand the rules
            print("dans ce cas je vous invite à les relires :)")                                                        #in this case I invite you to reread them
            print("_______________________________________________________________________________________________")
            rules()

    average(arr, num)


def average(arr, num):
    print("nous allons maintenant calculer la moyenne des chiffres que vous avez entrés")                               #we will now calculate the average of the numbers you entered

    somme = sum(arr)                                                                                                    #On fait la somme de tous les nombres entrés // add of all the numbers in the array
    average = somme / num                                                                                               #on divise la somme par le nombre de chiffres // divide the add by the number of digits
    print("nous avons une moyenne de:", round(average, 2))                                                              #We have an average of:
    StartOver()

def StartOver():
    print("voulez-vous rejouer? (O/N)")
    answer = input()

    if (answer == "O" or answer == "o"):
        nombres()
    elif (answer == "N" or answer == "n"):
        pass
    else:
        print("je n'ai pas compris votre réponse")
        nombres()
