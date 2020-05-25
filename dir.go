package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
        var m map[string]string
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	m = make(map[string]string)
	for _, file := range files {
		fmt.Println(file.Name())
                content, err := ioutil.ReadFile(file.Name())
    	        if err != nil {
		     log.Fatal(err)
          	}
                m[file.Name()] = string(content)
	        fmt.Printf("File contents: %s", string(content))
	}
	fmt.Printf("map: %v", m)
}
