package tools

import (
	"fmt"
	"os"
)

func Confirm(kind, data string) error {
	fmt.Printf("[%s]: `%s`, are you sure? (y/N)", kind, data)
	confirmData := make([]byte, 2)
	os.Stdin.Read(confirmData)
	if confirmData[0] != 'y' && confirmData[0] != 'Y' {
		return fmt.Errorf("confirm got negative answer")
	}
	return nil
}
