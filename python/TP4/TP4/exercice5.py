import sys
sys.path.append('..')

from PUtilitaires import StartOver
from PUtilitaires import inputInt
import math

def ruleExo5():                                                                                                         #fonction qui active la classe AFormeGeometrique // activate AFormeGeometrique class

    create = AFormeGeometrique()
    create.getNom()
    StartOver(ruleExo5)

class iFormeGeometrique():                                                                                              #classe interface // Interface class
    def __init__(self):
        pass
    def perimetre(self):
        pass
    def affichage(self):
        pass
    def surface(self):
        pass


class AFormeGeometrique(iFormeGeometrique):                                                                             #classe abstraite // abstract class

    def __init__(self):
        super(iFormeGeometrique, self).__init__()
        pass

    def getNom(self):                                                                                                   #active la classe avec l'input // activate class with the input
        print("Cercle = 1 | Rectangle/Carré = 2")
        nbr = inputInt("Quelle forme voulez-vous afficher ? ")
        if nbr == 1:
            name = Cercle()
            name.affichage()
        elif nbr == 2:
            name = Rectangle()
            name.surface()
        else:
            print("Vous devez entrez 1 ou 2")
            ruleExo5()



class Cercle(AFormeGeometrique):                                                                                        #classe cercle avec tous ses éléments // circle class
    def __init__(self):
        super(AFormeGeometrique, self).__init__()
        self.rayon = inputInt("Quel est la valeur de son rayon ? ")
        self.name = input("Nom du Cercle: ")

    def perimetre(self):
        return round(2 * math.pi * self.rayon, 2)

    def surface(self):
        return round(math.pi * (self.rayon ** 2), 2)

    def affichage(self):
        print(self.name, "est un cercle de rayon égal à", self.rayon)
        print("Son périmètre est de : ", Cercle.perimetre(self))
        print("Et sa surface est : ", Cercle.surface(self))

class Rectangle(AFormeGeometrique):                                                                                     #classe rectangle avec tous ses éléments // rectangle class
    def __init__(self):
        super(AFormeGeometrique, self).__init__()
        self.longueur = inputInt("Quelle est sa longueur ? ")
        self.largeur = inputInt("Quelle est sa largeur ? ")


    def perimetre(self):
        return round(2 * (self.largeur + self.longueur), 2)

    def surface(self):
        if self.longueur == self.largeur:
            carre = Carre()
            carre.affichage(self.longueur)
        else:
            self.surfacerec = round(self.longueur * self.largeur, 2)
            Rectangle.affichage(self)

    def affichage(self):
        self.nameR = input("Nom du Rectangle:")
        print(self.nameR, " est un rectangle de largeur égale à", self.largeur, "et de longueur égale à", self.longueur)
        print("Son périmètre est de : ", Rectangle.perimetre(self))
        print("Et sa surface est : ", self.surfacerec)

class Carre(Rectangle):                                                                                                 ##classe carré avec tous ses éléments // square class
    def __init__(self):
        super(Rectangle, self).__init__()
        self.name = input("Nom du carré: ")

    def affichage(self, longueur):
        print(self.name, " est un carré de côté égale à ", longueur)
        print("Son aire est donc :", longueur*longueur)
