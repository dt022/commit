// a program for automatically commiting and pushing to git by doing 'git add -A'.
// there will be two flags:
// -m will do 'git commit -m [argv]'.
// -p will do the above including 'git push'.

package commit

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func commit(message string) bool {
	// Run the command 'git add -A'
	cmd := exec.Command("git", "add", "-A")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Run the command 'git commit -m [commitMsg]'
	cmd = exec.Command("git", "commit", "-m", message)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return true
}

func Commit() {
	// Define flags
	commitMsg := flag.String("m", "", "Commit with message")
	pushFlag := flag.String("p", "", "Commit and Push with message")

	// print usage
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	// Parse command-line arguments
	flag.Parse()

	// if commitMsg, then commit
	if *commitMsg != "" {
		commit(*commitMsg)
		fmt.Printf("Successfully committed with the message {%v}.", *commitMsg)
	}

	// Push to remote repository if -p flag is provided
	if *pushFlag != "" {
		commit(*pushFlag)
		// Run the command 'git push'
		cmd := exec.Command("git", "push")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully committed & pushed with the message {%v}.", *pushFlag)
	}
}
