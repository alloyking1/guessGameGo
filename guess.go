package main
import (
	"fmt"
	"math/rand"
	"time"
)


func main () {
	gamePlay()
}

func gamePlay () {

	var score int
	var stage int = 1
	var guess int 
	var correctNumber int 

    correctNumber = randGen()

	fmt.Println("Welcome to alloy guess game!!")

	for guess != correctNumber {
		
		fmt.Println("Guess a number between 1-5: ")
		fmt.Scanln(&guess)

		if guess == correctNumber {

			score ++

			fmt.Printf("%d is the correct number. You'r score is %d", guess, score)

			for guess == correctNumber {

				correctNumber = randGen()
				fmt.Scanln(&guess)

				check := score % 3 

				if check == 0 {
					stage ++
					fmt.Printf(" Good job, Your in stage %d", stage )
				}
				
			}
			
		}
	}

}

func randGen () int {
	rand.Seed(time.Now().UnixNano())
    min := 1
    max := 5
	return rand.Intn(max - min + 1) + min
}

