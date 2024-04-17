package errorHandler

import (
	"fmt"
	"os"
)

func Handler(err any) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
