package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	watcher *fsnotify.Watcher
	stopCh  chan struct{}
)

func fileEventWatcher(dirPath string) (retErr error) {

	fmt.Println("Starting File Event Watcher")

	// Uncomment this to simulate error in fileEventWatcher
	return fmt.Errorf("error in starting watcher")

	retErr = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				fmt.Println("Error adding watch for dir:", err)
			}
		}
		return nil
	})

	if retErr != nil {
		fmt.Println("Error setting up file watches:", retErr)
		return retErr
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				// fmt error and return
				// ok is False only when channel is closed
				return fmt.Errorf("event channel closed")

			}

			fmt.Println("Event:", event)
			return fmt.Errorf("error in watch events")

		case err, ok := <-watcher.Errors:
			if !ok {
				// ok is False only when channel is closed
				return fmt.Errorf("error channel closed")
			}
			fmt.Println("Error:", err)
			return err
		case <-stopCh:
			fmt.Println("Stopping File Event Watcher")
			return nil
		}
	}
}

func create() (err error) {
	// create a directory named tstDir in the current directory
	_, filename, _, _ := runtime.Caller(0)
	dirPath := filepath.Join(filepath.Dir(filename), "tstDir")
	//dirPath := filepath.Join(os.TempDir(), "tstDir")

	//err := os.Mkdir(dirPath, 0755)
	//if err != nil {
	//	fmt.Println("Error creating directory:", err)
	//	return err
	//}
	//defer os.RemoveAll(dirPath)

	stopCh = make(chan struct{})
	//defer close(stopCh)

	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		return err
	}
	//defer watcher.Close()

	startFEV := true
	if startFEV {
		go func() (err error) {
			if err = fileEventWatcher(dirPath); err != nil {
				fmt.Println("Stopping File Event Watcher on err")
				fmt.Println("error in func is: ", err)
				// If fileEventWatcher throws error and we also close stopCh here
				// then we will hit a panic in the main function when we try to close the
				// stopCh channel again.
				// This is due to the fact that the stopCh channel is already closed in the
				// defer statement in the main function.
				// So closing stopCh channel here will result in a panic.
				// Hence, close stopCh channel only in the defer statement in the main function.
				//close(stopCh)
				return err
			}
			return nil
		}()

		defer func() {
			fmt.Println("Error in defer: ", err)
			if err != nil {
				fmt.Println("Stopping File Event Watcher-1")
				fmt.Println("Number of goroutines 1:", runtime.NumGoroutine())
				close(stopCh)
			}
		}()

	}

	// Simulate some error conditions
	simulateError := false
	if simulateError {
		fmt.Println("Simulating an error...")
		time.Sleep(2 * time.Second)
		// We encountered an error, so we'll stop the goroutine.
		fmt.Println("Stopping the fileEventWatcher due to an error.")
		close(stopCh)
		return err // This will also close the 'stopCh' channel.
	}

	fmt.Println("Number of goroutines 2", runtime.NumGoroutine())
	// In case of no errors, the main function continues executing here.
	fmt.Println("Main function is returning error")
	time.Sleep(10 * time.Second)

	return fmt.Errorf("error")
	// Do some other work or wait indefinitely, so the program doesn't terminate immediately.
	fmt.Println("Main function is running...")
	time.Sleep(30 * time.Second)
	return nil
}

func main() {
	fmt.Println("Number of goroutines 0", runtime.NumGoroutine())
	err := create()
	if err != nil {
		fmt.Println("Error:", err)
		// Check which go routines are still running.
		fmt.Println("Number of goroutines post err:", runtime.NumGoroutine())
		//return
	}
	// Check which go routines are still running.
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())

}
