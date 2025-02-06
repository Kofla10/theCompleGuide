package writetext

import (
	"os"
)

func WriteFile(write string) {

	os.WriteFile("user.text", []byte(write), 0664)
}
