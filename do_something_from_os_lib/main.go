package main

import (
	"fmt"
	"log"
	"os"
)

func Write_go_code_using_go(dirName string) {

	err := os.Mkdir(dirName, 0750)
	if err != nil && os.IsExist(err) {
		log.Fatal(err)
	}
	filePath := fmt.Sprintf("%s/hello.go", dirName)

	data := []byte("package main\n\n import \"fmt\" \n\n func main() {\n\tfmt.Println(\"Hello, World!\")\n}\n")
	err = os.WriteFile(filePath, data, 0660)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	tempFile := os.TempDir()
	fmt.Println(tempFile)

	Write_go_code_using_go("test")

}
