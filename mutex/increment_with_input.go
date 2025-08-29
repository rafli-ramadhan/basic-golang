package mutex

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IncrementInput() {
	counter := 1
	scanner := bufio.NewScanner(os.Stdin)

	for counter <= 5 {
		fmt.Print("Type 'Yes': ")
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if strings.EqualFold(input, "Yes") {
				fmt.Println(counter)
				counter++
			} else {
				fmt.Println("Please type 'Yes' to continue.")
			}
		} else {
			fmt.Println("Error reading input.")
			break
		}
	}
}
