package fancyprint

import (
	"fmt"
	"time"
	"math/rand"
)
var spinnerRotation int = 5

// SetSpinnerRotation will set how many times the spinner has to rotate when
// the function is called each time.
func SetSpinnerRotation(i int) {
	spinnerRotation = i
}

var spinnerRotationDelay time.Duration = time.Millisecond * 100

// SetSpinnerRotationDelay will set the delay between each spinner movement in MilliSeconds.
// Default value, if not manually set would be 100 MilliSeconds.
func SetSpinnerRotationDelay(i int) {
	spinnerRotationDelay = time.Duration(i) * time.Millisecond
}

var jumpdelay time.Duration = time.Millisecond * 10

// SetJumpDelay will set how long in MilliSeconds the cursor has to stay in the current position
// before jumping to the next position. Default value if not manually set would be 10 MilliSeconds.
func SetJumpDelay(i int) {
	jumpdelay = time.Millisecond * time.Duration(i)
}

// PrintAndEraseWithSpinner will print the series of byte array as strings, wait for "showTextFor"
// number of seconds and starts erasing the text backwards with a spinner, masks it
// with "endwithmask" string. The cursor movement is ensured only for single line printing,
// and if there are new line characters that cause the cursor to jump to new line, this function
// will not move the cursor to previous lines.
func PrintAndEraseWithSpinner(s []byte, showTextFor int, endwithmask string) {
	if len(endwithmask) > 1 {
		endwithmask = "*"
	}
	fmt.Printf("%s",s)
	time.Sleep(time.Second * time.Duration(showTextFor))
	for i := 0; i < len(s); i++ {
		if i == 0 {
			//fmt.Printf("\033[%dD%s", 1, "*")
			fmt.Printf("\033[%dD%s", 1," ")
			PrintSpinnerOnOffset(1, endwithmask)
		} else {
			//fmt.Printf("\033[%dD%s", 2, "*")
			//fmt.Printf("\033[%dD%s", 2, "*")
			fmt.Printf("\033[%dD%s", 2," ")
			PrintSpinnerOnOffset(1, endwithmask)
		}
		//time.Sleep(time.Millisecond * 1)
		time.Sleep(jumpdelay)
	}

	fmt.Printf("%c[2K", 27);
	fmt.Printf("\n")



}

// EraseFromCurrentCursorPositionWithSpinner will start to erase backwards from the current cursor
// position until it reaches the start of line. An offset needs to be provided.
func EraseFromCurrentCursorPositionWithSpinner(offset int) {
	for i := 0; i < offset; i++ {
		if i == 0 {
			//fmt.Printf("\033[%dD%s", 1, "*")
			fmt.Printf("\033[%dD%s", 1," ")
			PrintSpinnerOnOffset(1, " ")
		} else {
			//fmt.Printf("\033[%dD%s", 2, "*")
			//fmt.Printf("\033[%dD%s", 2, "*")
			fmt.Printf("\033[%dD%s", 2," ")
			PrintSpinnerOnOffset(1, " ")
		}
		time.Sleep(jumpdelay)
	}
}

// PrintEach will individually print each byte as type string to STDOUT honoring the delay
// set with SetJumpDelay function.
func PrintEach(s []byte)  {
	for _, i := range s {
		PrintSpinnerOnOffset(1, string(i))
		fmt.Printf("\033[%dC%s", -1, " ")
		time.Sleep(jumpdelay)
	}
}

// PrintAtRandom will print a single line byte array as string at random offsets.
func PrintAtRandom(s []byte) {
	perm := rand.Perm(len(s)+1)
	for _, shuffledIndex := range perm {
		if shuffledIndex > 1 {
			fmt.Printf("\033[%dC", shuffledIndex - 1)
			fmt.Printf(string(s[shuffledIndex - 1]))
			fmt.Printf("\033[%dD", shuffledIndex)
			time.Sleep(jumpdelay)
		} else if shuffledIndex == 1 {
			fmt.Printf("\033[%dC", shuffledIndex - 2)
			fmt.Printf(string(s[shuffledIndex - 1]))
			fmt.Printf("\033[%dD", shuffledIndex)
			time.Sleep(jumpdelay)
		}

	}

}

// PrintSpinnerOnOffset will print a spinner character and jumps to the mentioned position offset.
func PrintSpinnerOnOffset(pos int, endmask string) {
	rotate := []string{"-", "\\", "|", "/"}
	for i := spinnerRotation; i >= 0; i-- {
		for _, j := range rotate {
			fmt.Printf("\033[%dD%s", pos, j)
			time.Sleep(spinnerRotationDelay)
		}
	}

	fmt.Printf("\033[%dD%s", 1, endmask)
}

// PrintSpinnerWithCarriageReturn will go to the start of line with \r and print a spinner.
func PrintSpinnerWithCarriageReturn() {
	rotate := []string{"-", "\\", "|", "/"}
	for i := spinnerRotation; i >= 0; i-- {
		for _, j := range rotate {
			fmt.Printf("\r %s", j)
			time.Sleep(spinnerRotationDelay)
		}
	}
}
