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
	var guess int 
	var correctNumber int 

    correctNumber = randGen()

	for guess != correctNumber {
		
		fmt.Println("Guess a number between 1-5: ")
		fmt.Scanln(&guess)

		if guess == correctNumber {

			score ++
			fmt.Printf("%d is the correct number. You'r score is %d", guess, score)

			for guess == correctNumber {

				correctNumber = randGen()
				fmt.Scanln(&guess)
				
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

