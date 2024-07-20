package answer

import (
	"fmt"
	"math/rand"
)

func RandomNumber() {
	var jeu = 1000
	randomizeNumber := rand.Intn(jeu)
	var userInput int
	var maxi = jeu
	var mini = 0
	// fmt.Println(randomizeNumber)
	fmt.Printf("Le but est de deviné le chiffre entre 0 et %d\n", jeu)
	for {
		if _, err := fmt.Scanln(&userInput); err != nil {
			fmt.Println("Veuillez entrer un nombre entier valide.")
			continue
		}
		if userInput > randomizeNumber {
			maxi = userInput
			fmt.Printf("Votre valeur est supérieur à la valeur aléatoire, vous etes donc entre %d et %d\n", mini, maxi)
		} else if randomizeNumber > userInput {
			mini = userInput
			fmt.Printf("Votre valeur est inférieur à la valeur aléatoire, vous etes donc entre %d et %d\n", mini, maxi)
		} else {
			fmt.Printf("Bravo vous avez trouvé la valeur qui est %d\n", randomizeNumber)
			break
		}
	}

}
