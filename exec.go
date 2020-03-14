package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

//Example command execution and capturing output
func main() {

	os.RemoveAll("temp")
	err := os.MkdirAll("temp", 0755)
	if err != nil {
		fmt.Printf("Error in creating dir %s\n", err)
		return
	}

	sFile, err := os.Create("temp/text")
	if err != nil {
		fmt.Printf("unable to create temp/text file")
		return
	}
	defer sFile.Close()

	cmd := exec.Command("ls", "-ltr")

	var stderr bytes.Buffer
	//Writing to both stdout and file
	cmd.Stdout = io.MultiWriter(sFile, os.Stdout)
	//Capture stderr to a buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Finished writing to log file")
	fmt.Printf("stderr %s", string(stderr.Bytes()))

}
