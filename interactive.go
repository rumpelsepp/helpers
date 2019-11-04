package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Confirm(message string) bool {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/N]: ", message)

	res, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.ToLower(strings.TrimSpace(res)) == "y"
}
