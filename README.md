# fancyprint
A fancyprint library

![Demo](https://raw.githubusercontent.com/nohupped/fancyprint/master/demo.gif)

##### Install:

`go get -u github.com/nohupped/fancyprint`

##### Use:

* All functions accept type `[]byte`. Cast the string to `[]byte()` before passing it to the function.
* Works only on terminals supporting [`Ansi Escape Sequences.`](https://en.wikipedia.org/wiki/ANSI_escape_code "Ansi Escape Codes - Wikipedia")
* Works for single line texts only. Multi-line texts are not supported.

Eg. code that prints similar to the gif above:
```
package main

import (
	"github.com/nohupped/fancyprint"
	"fmt"
)

var str []byte = []byte("Sphinx of black quartz, judge my vow")

func main() {
	fancyprint.SetSpinnerRotation(3)
	fancyprint.SetJumpDelay(10)
	fancyprint.SetSpinnerRotationDelay(10)
	fancyprint.PrintAndEraseWithSpinner(str, 1, "*")
	fmt.Println("\n\n")

	fmt.Printf(string(str))
	fancyprint.EraseFromCurrentCursorPositionWithSpinner(len(str))
	fmt.Println("\n\n")

	fancyprint.PrintEach(str)
	fmt.Println("\n\n")

	fancyprint.SetJumpDelay(500)
	fmt.Println("Actual string:")
	fmt.Println(string(str))
	fancyprint.PrintAtRandom(str)
	fmt.Println("\n\n\n")


}

```
