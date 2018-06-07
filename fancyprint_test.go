package fancyprint_test

import (
	"fmt"
	"github.com/nohupped/fancyprint"
	"testing"
)

var str []byte = []byte("Sphinx of black quartz, judge my vow")


func TestPrintAndErase(t *testing.T) {
	// fancyprint.ClearMask = true
	fancyprint.SetSpinnerRotation(3)
	fancyprint.SetJumpDelay(10)
	fancyprint.SetSpinnerRotationDelay(30)
	fancyprint.PrintAndEraseWithSpinner(str, 1, "*")
}

func TestEraseFromCurrentCursorPosition(t *testing.T) {
	fmt.Printf(string(str))
	//	fancyprint.SetSpinnerRotation(10)
	//	fancyprint.SetJumpDelay(100)
	fancyprint.EraseFromCurrentCursorPositionWithSpinner(len(str))
}

func TestPrintEach(t *testing.T) {
	fancyprint.PrintEach(str)
}

func TestPrintAtRandom(t *testing.T) {
	fancyprint.SetJumpDelay(500)
	fmt.Println("Actual string:")
	fmt.Println(string(str))
	fancyprint.PrintAtRandom(str)
}
