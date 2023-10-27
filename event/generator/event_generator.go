package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func createFiles(dirPath string) {
	for i := 1; i <= 3; i++ {
		fileName := fmt.Sprintf("file%d.txt", i)
		filePath := filepath.Join(dirPath, fileName)
		data := []byte(fmt.Sprintf("This is file %d", i))
		err := ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		fmt.Println("Created file:", fileName)
	}

	// Wait for a while before modifying and deleting the files
	time.Sleep(3 * time.Second)

	for i := 1; i <= 3; i++ {
		fileName := fmt.Sprintf("file%d.txt", i)
		filePath := filepath.Join(dirPath, fileName)
		data := []byte(fmt.Sprintf("Modified content for file %d", i))
		err := ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			fmt.Println("Error modifying file:", err)
			return
		}
		fmt.Println("Modified file:", fileName)
	}

	// Wait for a while before deleting the files
	time.Sleep(3 * time.Second)

	for i := 1; i <= 3; i++ {
		fileName := fmt.Sprintf("file%d.txt", i)
		filePath := filepath.Join(dirPath, fileName)
		err := os.Remove(filePath)
		if err != nil {
			fmt.Println("Error deleting file:", err)
			return
		}
		fmt.Println("Deleted file:", fileName)
	}
}

func main() {
	//tmpDir := os.TempDir()
	//dirPath := filepath.Join(tmpDir, "testdir")
	// Create the test directory
	//err := os.Mkdir(dirPath, 0755)
	//if err != nil {
	//	fmt.Println("Error creating directory:", err)
	//	return
	//}
	//defer os.RemoveAll(dirPath)

	// Get the executable file path for the running program
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// Get the current directory path
	currentDir := filepath.Dir(executablePath)

	//dirPath := filepath.Join(filepath.Dir(filename), "tstDir")
	dirPath := filepath.Join(currentDir, "tstDir")

	// Run the file event generator to create, modify, and delete files
	createFiles(dirPath)
}
