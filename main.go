package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	myData := []byte("This is written file\n")
	err := ioutil.WriteFile("newfile.data", myData, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully added to file")

	data, err := ioutil.ReadFile("newfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	//adding additional data in existing file
	f, err := os.OpenFile("newfile.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	myNewData := []byte("New data that wasn't there originally\n")
	if _, err := f.Write(myNewData); err != nil {
		panic(err)
	}

	// reading file again after changes
	data, err = ioutil.ReadFile("newfile.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}
