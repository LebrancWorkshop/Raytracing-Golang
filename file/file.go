package file

import (
	"fmt"
	"os"
)

func WriteFile(fileName string, content string) {
	file, err := os.Create(fileName + ".ppm")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(content)

	fmt.Println("[SUCCESS] Writing File Success")
}