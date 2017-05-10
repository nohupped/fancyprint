package fancyprint_test
import (
	"fancyprint"
	"fmt"
	"testing"
)

var str []byte = []byte("Sphinx of black quartz, judge my vow")

func TestPrintAndErase(t *testing.T) {
	fancyprint.SetSpinnerRotation(3)
	fancyprint.SetJumpDelay(10)
	fancyprint.SetSpinnerRotationDelay(10)
	fancyprint.PrintAndEraseWithSpinner(str, 1, "*")
	fmt.Println("\n\n")
}

func TestEraseFromCurrentCursorPosition(t *testing.T) {
	fmt.Printf(string(str))
//	fancyprint.SetSpinnerRotation(10)
//	fancyprint.SetJumpDelay(100)
	fancyprint.EraseFromCurrentCursorPositionWithSpinner(len(str))
	fmt.Println("\n\n")
}

func TestPrintEach(t *testing.T) {
	fancyprint.PrintEach(str)
	fmt.Println("\n\n")
}

func TestPrintAtRandom(t *testing.T) {
	fancyprint.SetJumpDelay(500)
	fmt.Println("Actual string:")
	fmt.Println(string(str))
	fancyprint.PrintAtRandom(str)
	fmt.Println("\n\n\n")
}