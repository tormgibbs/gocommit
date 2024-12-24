package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"os/exec"
)

func createFile(commitNumber string) {
	file, err := os.OpenFile("gocommit.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer file.Close()

	now := time.Now().Format("02/01/06 15:04:05")
	content := fmt.Sprintf("Commit %s occurred at %s", commitNumber, now)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("failed to write to file: %s\n", err)
		return
	}
}

func commitChanges(commitMsg string) {
	cmd := exec.Command("git", "add", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("git add failed with error: %s\nOutput: %s\n", err, output)
		return
	}

	cmd = exec.Command("git", "commit", "-m", commitMsg)
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("git commit failed with error: %s\nOutput: %s\n", err, output)
		return
	}
}

func pushChanges() {
	cmd := exec.Command("git", "push", "-u", "origin", "main", "-f")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("git push failed with error: %s\nOutput: %s\n", err, output)
		return
	}
}

func main() {
	now := time.Now().Format("02/01/06 15:04:05")
	commitMsg := fmt.Sprintf("random message %s", now)

	for n := range 3 {
		createFile(strconv.Itoa(n))
		commitChanges(commitMsg)
	}

	pushChanges()
}
