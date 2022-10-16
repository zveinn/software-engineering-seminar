package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// memory()
	files()
}

func memory() {

	var XX = make([]int, 10, 10)

	// var XX [10]int

	fmt.Println("LENGTH(no. of items):", len(XX), " // CAPACITY(max items):", cap(XX), " \n DATA:", XX)
	// XX[3] = 4
	// fmt.Println(XX)
	XX = append(XX, 10)

	fmt.Println("===========================")
	fmt.Println("LENGTH(no. of items):", len(XX), " // CAPACITY(max items):", cap(XX), " \n DATA:", XX)

	XX = XX[:15]
	fmt.Println("===========================")
	fmt.Println("LENGTH(no. of items):", len(XX), " // CAPACITY(max items):", cap(XX), " \n DATA:", XX)

	// XX[9] = 2
}

func files() {

	file, err := os.Create("file1")
	if err != nil {
		fmt.Println(err)
		return
	}

	stat, _ := file.Stat()
	fmt.Println("NEW FILE:", stat.Name(), "// SIZE:", stat.Size(), "// MOD TIME:", stat.ModTime())

	helloworld := "hello world\n"
	var totalBytesWritten int

	numberOfBytesWrittenToFile, _ := file.WriteString(helloworld)

	totalBytesWritten += numberOfBytesWrittenToFile

	numberOfBytesWrittenToFile, _ = file.WriteString(helloworld)

	totalBytesWritten += numberOfBytesWrittenToFile

	fmt.Println("WROTE", totalBytesWritten, "BYTES TO FILE >> DATA:", helloworld+helloworld)

	// GO BACK TO THE START OF THE FILE
	file.Seek(0, 0)
	// READ THE FULL FILE
	fileBytes, _ := io.ReadAll(file)
	fmt.Println("FILE BYTES:", fileBytes)

	numberOfBytesWrittenToFile, _ = file.Write([]byte{104, 104, 104, 13})

	totalBytesWritten += numberOfBytesWrittenToFile

	fmt.Println("WROTE", totalBytesWritten, "BYTES TO FILE")

	file.Seek(0, 0)
	fileBytes, _ = io.ReadAll(file)
	fmt.Println("FILE BYTES:", fileBytes)

	for i := 0; i < 1000; i++ {
		OPEN_FILE_AND_WRITE_STRING(file, "HELLO FUCKING WORLD "+strconv.Itoa(i)+"\n")
	}

	file.Close()
}

func OPEN_FILE_AND_WRITE_STRING(file *os.File, data string) {
	_, err := file.WriteString(data)
	if err != nil {
		log.Println(err)
		return
	}
}
