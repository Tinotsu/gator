package config

import(
	"fmt"
	"os"
)

func HandleError (err error) {
	if err != nil {
		fmt.Print("\n⚠️ error: ", err, "\n")
		os.Exit(1)
	}
}
