from tkinter import *
import math


def Calculatrice():
    calculator = Calculator()
    calculator.calculator()                                                                                             #On active la fonction calculator de la classe Calculator // Activate calculator function from the class Calculator

class Calculator():
    erase = False                                                                                                       #erase est un bool avec Faux comme valeur par défaut // erase is a bool with False as default value

    def __init__(self):                                                                                                 #Fonction qui construit la classe // Function who create the class
        "constructeur"

    def calculator(self):
        "création des widgets et on les lies à leur fonction"

        self.root= Tk()                                                                                                 #fenêtre principale de l'application // principal window of the app
        self.root.title("calculatrice")
        self.display= Entry(self.root, width=13, bg="lavender", bd=2, font="courier 20")                                #fenêtre de la calculatrice // window of the calculator
        self.display.grid()                                                                                             #grid sert à l'affichage
        self.frame = Frame(self.root, relief="raised", border=0, bg="lavender")                                         #frame regroupe les composants
        self.frame.grid(sticky="nsew")

        self.buttons = [["QUITTER", 0, 0, 5, 1, 12, 4],                                                                 #Tous les bouttons sont contenus dans un tableau multidimenssionnel //all buttons containing in a multidimensional array
                        ["(", 1, 0, 1, 1, 6, 4],                                                                        #name, row, column, columnspan, rowspan, width, height
                        [")", 1, 1, 1, 1, 6, 4],                                                                        #nom, ligne, colonnes, columnspan (fusion de colonnes), rowspan(fusion de colonnes), largeur, hauteur
                        ["x²", 1, 2, 1, 1, 6, 4],
                        ["C", 1, 3, 1, 1, 6, 4],
                        ["CE", 1, 4, 1, 1, 6, 4],
                        ["7", 2, 0, 1, 1, 6, 4],
                        ["8", 2, 1, 1, 1, 6, 4],
                        ["9", 2, 2, 1, 1, 6, 4],
                        ["/", 2, 3, 1, 1, 6, 4],
                        ['√', 2, 4, 1, 1, 6, 4],
                        ["4", 3, 0, 1, 1, 6, 4],
                        ["5", 3, 1, 1, 1, 6, 4],
                        ["6", 3, 2, 1, 1, 6, 4],
                        ["x", 3, 3, 1, 1, 6, 4],
                        ["%", 3, 4, 1, 1, 6, 4],
                        ["1", 4, 0, 1, 1, 6, 4],
                        ["2", 4, 1, 1, 1, 6, 4],
                        ["3", 4, 2, 1, 1, 6, 4],
                        ["-", 4, 3, 1, 1, 6, 4],
                        ["=", 4, 4, 1, 2, 6, 8],
                        ["+/-", 5, 0, 1, 1, 6, 4],
                        ["0", 5, 1, 1, 1, 6, 4],
                        [".", 5, 2, 1, 1, 6, 4],
                        ["+", 5, 3, 1, 1, 6, 4]]

        #création des bouttons // creation of buttons

        for item in self.buttons:                                                                                       #pour chaque tableau dans le self.button //fo each array in the self.button array
            self.button = Button(self.frame, text=item[0], highlightbackground="white", height=item[6], width=item[5], relief="raised",
                                 border=1, command=lambda arg=item[0]: self.action(arg))                                #on affiche le bouton à partir des paramètres du frame, puis on le personnalises avec les données du tableau en indiquant le positionnement
                                                                                                                        #display the button from the parameters of the frame, then personalize it with the data of the table indicating the positioning
            self.button.grid(row=item[1], column=item[2], columnspan=item[3], rowspan=item[4], sticky="nsew")           #nsew permet l'étirement pour remplir la cellule

        self.root.mainloop()                                                                                            #lance la fenêtre // launches the window

    def action(self, signe):                                                                                            #action lié aux bouttons spéciaux // action link to special buttons
        signesSpeciaux=  ["QUITTER",
                           "C",
                           "CE",
                           "x²",
                           "√",
                           "%",
                           "=",
                           "+/-"]

        if signe not in signesSpeciaux:
            if signe in ["+", "-", "x", "/", "."]:                                                                      #pour ces boutons non contenu dans le tableau signesSpeciaux for these buttons not contained in the signesSpeciaux array
                Calculator.erase = False                                                                                #ils gardent l'option effacé inactif // they keep the erased option inactive
            if Calculator.erase == True and signe not in ["+", "-", "x", "/", "."]:                                     #si ils éffacent et ne font pas partie de ces bouttons // if they erase and are not part of these buttons
                self.display.delete(0, 'end')                                                                           #on efface tout // erase everything
                Calculator.erase = False                                                                                #et erase repasse à False // erase tkaes False value again
            self.display.insert('end', signe)                                                                           #on insert le signe à la fin des éléments (surtout pour le quitter) // insert the sign at the end of elements

        elif signe == signesSpeciaux[0]: #le bouton quitter // quit button
            self.root.destroy()                                                                                         #on détruit la classe // destroy the class

        elif signe == signesSpeciaux[1]: #le bouton supp tout // C button
            self.display.delete(0, 'end')                                                                               #supprime tout //delete everything

        elif signe == signesSpeciaux[2]:  # le bouton sup 1 char // CE button
            afficher = self.display.get()[:-1]                                                                          #afficher prend tout les élément afficher sauf le dernier // afficher takes all display item except the last one
            self.display.delete(0, 'end')                                                                               #supp tout //delete everything
            self.display.insert('end', afficher)                                                                        #on insert à l'écran 'afficher' // display on the screen 'afficher'

        elif signe == signesSpeciaux[3]: #le bouton puissance 2 // x² button
            try:
                value = float(self.display.get())                                                                       #on récupère les éléments afficher // we retrieve the display elements
                self.display.delete(0, 'end')                                                                           #supp tout // erase everything

                if str(value).endswith('.0'):                                                                           #Si le résultat finit avec un .0 // if the result end with a .0
                    self.display.insert('end', int(value*value))                                                        #on cast en int pour pas avoir la virgule et on affiche le résultat// transform into int to not have the comma and display the result
                else:
                    self.display.insert('end', value*value)                                                             #sinon on affiche le résultat // otherwhise we display the result

                Calculator.erase =True                                                                                  #erase passe à True pour passer à un autre calcul//erase value becomes True, to switch to another calculation
            except:
                self.display.delete(0, 'end')                                                                           #si ce n'est pas possible de faire tout ça, on supprime ce qu'il y a 'afficher à l'écran //if it's not possible to do that, we delete what is display on the screen

        elif signe == signesSpeciaux[4]: #le bouton racine carrée                                                       #même principe que la puissance, seul le calcul change
            try:
                value = float(self.display.get())                                                                       #same principle as power, only the calculation changes
                self.display.delete(0, 'end')

                if str(math.sqrt(value)).endswith('.0'):
                    self.display.insert('end', int(math.sqrt(value)))
                else:
                    self.display.insert('end', math.sqrt(value))

                Calculator.erase =True
            except:
                self.display.delete(0, 'end')

        elif signe == signesSpeciaux[5]: #le boutton %
            value = self.display.get()                                                                                  #On récupère la valeur affiché à l'écrant // takes the value display on the screen
            for i in value:                                                                                             #pour chaque éléments de value // for each items of value
                try:
                    float(i)                                                                                            #cast en float // transform into float
                except :                                                                                                #si ce n'est pas possible// if it's not possible
                    try:
                        value = ("".join(value)).split(i)                                                               #on sépare chaque i de value // we separate each i of value
                        if i == 'x':                                                                                    #si i est le signe de la multiplication //if i is the sign of multiplication
                            i = '*'
                        percentage=float(value[0])*float(value[1])/100                                                  #on calcul le reste // we calculate the rest
                        result = eval(str(value[0]) + i + str(percentage))
                        if str(result).endswith('.0'):                                                                  #si le resultat est un nombre entier on enlève la virgule
                            result = int(result)                                                                        #if the result is a whole number, we delete the comma
                        self.display.delete(0, 'end')
                        self.display.insert("end", result)                                                              #On affiche le résultat // display the result
                        Calculator.erase = True
                    except :
                        self.display.delete(0, 'end')

        elif signe == signesSpeciaux[6]:# le boutton =
            try:
                if "x" in self.display.get():                                                                           #Si le calcul est une multiplication // if the calculation is a multiplication
                    value = eval(self.display.get().replace('x', '*'))                                                  #On remplace X par le signe de la multiplication // we replace 'X' by multiplicaiton sign
                else:
                    value = eval(self.display.get())                                                                    #autrement on fait siplement le calcul // otherwhise we do the calculation
                if str(value).endswith('.0'):
                    value = int(value)
                self.display.delete(0, 'end')
                self.display.insert("end", value)
                Calculator.erase = True
            except:
                self.display.delete(0, 'end')

        elif signe == signesSpeciaux[7]:#le boutton +/-
            try:
                value = float(self.display.get())                                                                       #on prend la valeur afficher // we takes the display value
                self.display.delete(0, 'end')                                                                           #supp tout // delete everything
                self.display.insert("end", value*-1)                                                                    #on prend la valeur et on la multiplie par -1 // takes the value and multiply it by -1
                Calculator.erase = True
            except :
                self.display.delete(0, 'end')
