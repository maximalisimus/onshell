package main
	
import (
    "fmt"
    "os"
    "os/exec"
    "strings"
    "path/filepath"
    "log"
)

func main() {
    shell := os.Getenv("SHELL")
    replacer := strings.NewReplacer("/", "", "usr", "","bin","")
    isshell := replacer.Replace(shell)
    if len(os.Args) >= 2 {
		abs, err := filepath.Abs(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		if _, err := os.Stat(abs); err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Error! File %s does not exists!\n",os.Args[1])
				fmt.Printf("Please, input the file!\n")
				// file does not exist
				os.Exit(2)
			} else {
				fmt.Printf("Unknow error!")
				// other error
				os.Exit(255)
			}
		}
		
		script := os.Args[1]
		cmd := exec.Command(isshell, script)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else
	{
		fmt.Printf("Error! Please, input the argument!\n")
		fmt.Printf("Shell = %s",isshell)
	}
}

