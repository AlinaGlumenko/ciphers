package view

import (
	"bufio"
	"fmt"
	"os"

	"../controller"
	strs "../model/strings"
)

var answer string
var text string

func Startloop() {
	fmt.Println(Greetings)
	consoleReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(br)
		fmt.Print(StartInfo)
		fmt.Println(br)
		fmt.Println(Colors[Yellow] + "Your choice: ")
		fmt.Scanln(&answer)
		fmt.Print(Colors[Reset])

		if should, errNum := controller.ShouldContinue(answer); !should {
			fmt.Println(br)
			fmt.Println(Errors[errNum])
			continue
		}

		fmt.Println(Colors[Yellow] + "Your text: ")
		text, _ = consoleReader.ReadString('\n')
		fmt.Print(Colors[Reset])

		fmt.Println(br)
		if res, isErr := controller.Do(text, answer); isErr {
			fmt.Println(Errors[int(strs.PROC_IMPOS)])
		} else {
			fmt.Println(Colors[Purple] + "Result: ")
			fmt.Print(res)
			fmt.Print(Colors[Reset] + "\n")
		}

	}
}
