package db

import (
	"fmt"
	"os"
)

const fileName = ""
const filePath = ""

func main() {
	openFile("db.json")
	/*read()
	file,err:=os.ReadFile("db.json")
	if err != nil {
		panic(err)
	}
	//rewrites the database
	write(file, getJsonString())*/
}

func getJsonString() string {
	return ""
}

func CreateAndWriteWithConstants() {
	createAndWrite(filePath, fileName)
}

func createAndWrite(directory string, txt string) {
	file, err := os.Create(directory)
	PanicIfError(err)
	defer file.Close()
	file.WriteString(txt)

}

func openFile(directory string) {
	file, err := os.Open(directory)
	PanicIfError(err)
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(data)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
func create(directory string) {
	file, err := os.Create(directory)
	if err != nil {
		panic(err)
	}
	write(file, "")
}

func write(file *os.File, txt string) {
	length, err := io.WriteString(file, txt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Length is %d", length)

}

func read() {
	file, ferr := os.Open("db.csv")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}*/
