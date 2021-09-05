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
	var max int = 5

    correctNumber = randGen(max)

	fmt.Println("Welcome to alloy guess game!!")

	for guess != correctNumber {
		
		fmt.Printf("\n Guess a number between 1- %d: ", max)
		fmt.Scanln(&guess)

		if guess == correctNumber {

			score ++

			fmt.Printf("%d is the correct number. You'r score is %d", guess, score)

			for guess == correctNumber {

				correctNumber = randGen(max)
				fmt.Scanln(&guess)

				check := score % 3 

				if check == 0 {
					stage ++
					max ++
					fmt.Printf("\n Good job, Your in stage %d", stage )
				}
				
			}
			
		}
	}

}

func randGen (max int) int {
	rand.Seed(time.Now().UnixNano())
    min := 1
    maximum := max
	return rand.Intn(maximum - min + 1) + min
}

