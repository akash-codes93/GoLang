package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	p := fmt.Printf
	pl := fmt.Println

	var newFile *os.File
	p("%T\n", newFile)

	var err error

	newFile, err = os.Create("a.txt")

	if err != nil {
		log.Fatal(err) // log is thread safe and it prints the timestamp
	}

	pl(newFile)

	err = os.Truncate("a.txt", 0)

	if err != nil {
		log.Fatal(err) // log is thread safe and it prints the timestamp
	}

	newFile.Close()

	// file, err := os.Open("a.txt")
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	pl(file)

	file.Close()

	var fileInfo os.FileInfo

	fileInfo, err = os.Stat("a.txt")
	pl("File Name:", fileInfo.Name())
	pl("Size in bytes:", fileInfo.Size)
	pl("Last modified:", fileInfo.ModTime())
	pl("Is Directory: ", fileInfo.IsDir())
	pl("Permissions:", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")

	if err != nil {
		// pl(err)
		if os.IsNotExist(err) {
			pl("File does not exists")
		}
	}

	oldPath := "a.txt"
	newPath := "aa.txt"

	err = os.Rename(oldPath, newPath)

	if err != nil {
		pl(err)
	}

	err = os.Remove("aa.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}

	myfile, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer myfile.Close() // close the file in case of panic

	byteSlice := []byte("I learn Golang!")
	byteWritten, err := myfile.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	pl(byteWritten)

	bs := []byte("Go programming is cool !!!!")
	err = ioutil.WriteFile("c.txt", bs, 0644)

	file1, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	defer file1.Close()

	bufferedWriter := bufio.NewWriter(file1)
	bs1 := []byte{97, 98, 99}

	bytesWritten, err := bufferedWriter.Write(bs1)

	log.Printf("Bytes written to buffer (not cool) %d", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("\nJust some things")
	unflushedBufferedSize := bufferedWriter.Buffered()

	log.Printf("%d\n", unflushedBufferedSize)
	bufferedWriter.Flush()

	// read entire file
	mainFile, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}

	defer mainFile.Close()

	data, err := ioutil.ReadAll(mainFile)
	if err != nil {
		log.Fatal(err)
	}

	p("Data as string: %s\n", data)
	pl("len of data read: %d\n", len(data))
	// ioutil han
	data1, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	p("data read using ioutil: %s\n", data1)

	// read file line by line
	delimetedFile, err := os.Open("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer delimetedFile.Close()

	scanner := bufio.NewScanner(delimetedFile)

	scanner.Split(bufio.ScanWords) // scan the file word by word
	// scanner.Split(bufio.ScanRunes) // cahr by char
	// success := scanner.Scan()

	// if !success {
	// 	err = scanner.Err()
	// 	if err == nil {
	// 		log.Println("end of file reached")
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }

	// pl("First line found: ", scanner.Text())

	for scanner.Scan() {
		pl(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	newscanner := bufio.NewScanner(os.Stdin) // reading from the terminal
	newscanner.Scan()
	text1 := newscanner.Text()
	pl(text1)

	 // reading the input continously until a specific string is scanned
	 for newscanner.Scan() {
        text1 = newscanner.Text()
        fmt.Println("You entered:", text1)
        if text1 == "exit" {
            fmt.Println("Exiting the scanning ...")
            break
        }
    }
 
    // error handling
    if err := newscanner.Err(); err != nil {
        log.Println(err)
    }
}
