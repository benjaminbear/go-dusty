package inputparse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskForYesNo(forceInput bool) (bool, error) {
	var yesNo string

	for {
		fmt.Print(" (y/n): ")

		num, err := fmt.Scanf("%s\n", &yesNo)
		if err != nil && num != 0 {
			return false, err
		}

		switch strings.ToLower(yesNo) {
		case "y":
			fallthrough
		case "yes":
			return true, nil
		case "n":
			fallthrough
		case "no":
			return false, nil
		default:
			if !forceInput {
				return false, nil
			}
			fmt.Print("Invalid input. Please try again")
		}
	}
}

func AskForWord(forceInput bool) (string, error) {
	var word string

	for {
		fmt.Print(": ")

		num, err := fmt.Scanf("%s\n", &word)
		if err != nil && num != 0 {
			return "", err
		}

		if word == "" && forceInput {
			fmt.Print("Invalid input. Please try again")
		} else {
			return word, nil
		}
	}
}

func AskForInteger(forceInput bool) (int, error) {
	var input int
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(": ")

		num, err := fmt.Fscanf(stdin, "%d\n", &input)
		if err != nil && num != 0 {
			return input, err
		}

		if input == 0 && forceInput {
			// we have to flush stdin buffer, otherwise it will err on every "char"
			_, _ = stdin.ReadString('\n')
			fmt.Print("Invalid input. Please try again")
		} else {
			return input, nil
		}
	}
}
