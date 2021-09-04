package main
import "fmt"


func main () {
	gamePlay()
}

func gamePlay () {

	var score int
	var guess int
	var correctNumber int = 3

	

	for guess != correctNumber {
		fmt.Println("Guess a number between 1-5: ")
		fmt.Scanln(&guess)

		if guess == correctNumber {

			score ++
			fmt.Printf("%d is the correct number. You'r score is %d", guess, score)
			
		}
	}

	

}