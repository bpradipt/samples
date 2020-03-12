package main

import (
	"fmt"
	"os/exec"
        "os"
)

func main() {
	cmd := exec.Command("/bin/echo", "howare")
        out, err := cmd.CombinedOutput()
        if err != nil {
           fmt.Printf("cmd.Run() failed with %s\n", err)
	}
        os.Setenv("MYENV",string(out))
	fmt.Printf("combined out: %s\n", string(out))
        fmt.Println("MYENV:", os.Getenv("MYENV"))

}

