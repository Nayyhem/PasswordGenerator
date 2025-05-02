package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
)

const (
	// Les ensembles de caractères disponibles pour le mot de passe
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // Majuscules
	lowerChars  = "abcdefghijklmnopqrstuvwxyz" // Minuscules
	digitChars  = "0123456789"                 // Chiffres
	symbolChars = "!@#$%&*()-_=+[]{}<>?/|"     // Symboles
)

// generate génère un mot de passe aléatoire en fonction de la longueur et des options spécifiées.
func generate(length int, useUpperChars, useLowerChars, useDigitChars, useSymbolChars bool) string {
	charset := "" // Initialisation de l'ensemble des caractères vides

	// Ajoute les caractères majuscules si l'option useUpperChars est activée
	if useUpperChars {
		charset += upperChars
	}
	// Ajoute les caractères minuscules si l'option useLowerChars est activée
	if useLowerChars {
		charset += lowerChars
	}
	// Ajoute les chiffres si l'option useDigitChars est activée
	if useDigitChars {
		charset += digitChars
	}
	// Ajoute les symboles si l'option useSymbolChars est activée
	if useSymbolChars {
		charset += symbolChars
	}
	// Si aucun ensemble de caractères n'a été sélectionné, arrête le programme avec une erreur
	if charset == "" {
		log.Fatal("Erreur : aucune catégorie de caractère selectionner.")
	}

	password := ""

	// Génère chaque caractère du mot de passe en tirant un caractère aléatoire de charset
	for i := 0; i < length; i++ {
		char, err := randomChar(charset) // Appel à randomChar pour obtenir un caractère
		if err != nil {
			log.Fatal(err)
		}
		password += char // Ajoute le caractère généré au mot de passe
	}

	return password
}

// randomChar génère un caractère aléatoire à partir de l'ensemble de caractères spécifié.
func randomChar(charset string) (string, error) {
	randomByte := make([]byte, 1)   // Crée un tableau pour stocker un octet aléatoire
	_, err := rand.Read(randomByte) // Lit un octet aléatoire cryptographiquement sécurisé
	if err != nil {
		return "", err
	}
	index := int(randomByte[0]) % len(charset) // Calcule un index dans le charset
	return string(charset[index]), nil         // Retourne le caractère correspondant à cet index
}

func main() {
	// Définition des options en ligne de commande
	length := flag.Int("length", 15, "Longueur du mot de passe")
	useLower := flag.Bool("lower", true, "Inclure des minuscules")
	useUpper := flag.Bool("upper", true, "Inclure des majuscules")
	useDigits := flag.Bool("digits", true, "Inclure des chiffres")
	useSymbols := flag.Bool("symbols", true, "Inclure des symboles")
	flag.Parse() // Analyse les arguments de ligne de commande

	// Génération du mot de passe en fonction des options spécifiées
	password := generate(*length, *useLower, *useUpper, *useDigits, *useSymbols)

	// Affichage du mot de passe généré
	fmt.Println("Mot de passe généré :", password)
}
