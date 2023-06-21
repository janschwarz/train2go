package numberguess

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Game() {
	min, max := 1, 100
	rand.Seed(time.Now().UnixNano())
	numberToGuess := rand.Intn(max-min) + min
	playing := true

	fmt.Println(fmt.Sprintf("Errate meine Zahl. Sie liegt zwischen %d und %d.", min, max))
	fmt.Println("Wie lautet dein Tipp?")

	attempts := 0
	for playing {
		attempts++
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		if guess > numberToGuess {
			fmt.Println(fmt.Sprintf("Dein Tipp ist %d, aber meine Zahl ist kleiner.", guess))
			fmt.Println("Wie lautet dein nächster Tipp?")
		} else if guess < numberToGuess {
			fmt.Println(fmt.Sprintf("Dein Tipp ist %d, aber meine Zahl ist größer.", guess))
			fmt.Println("Wie lautet dein nächster Tipp?")
		} else {
			fmt.Println(fmt.Sprintf("Dein Tipp %d war richtig! Du hast nur %d Versuche gebraucht. Glückwunsch!", guess, attempts))
			playing = false
		}
	}

}
