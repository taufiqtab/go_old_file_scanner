package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are you Ready ? (Y/n)")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	fmt.Println("Enter file path ? (Ex : /Users/name/folder1/subfolder/)")
	directory, _ := reader.ReadString('\n')
	directory = strings.TrimSuffix(directory, "\n")

	if text == "y" {
		fmt.Println("Scanning files")
	} else {
		return
	}

	//fetch all file inside given directory :
	entries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	//loop all the file result
	for _, e := range entries {
		//make sure it's note directory
		if e.IsDir() != true {
			//fetch file info using os.Stat every file iteration
			file_info, err2 := os.Stat(e.Name())
			if err2 != nil {
				log.Fatal(err2)
			}
			//get last modification time using ModTime
			lastEdit := file_info.ModTime()

			now := time.Now()
			// now = now.Add(25 * (time.Hour * 24)) //modified date by +days
			diff := now.Sub(lastEdit)
			days := math.Round(diff.Hours() / 24)
			msg := ""

			if days >= 365 {
				msg = "(DELETED!)"
			} else {
				msg = "(KEEP.)"
			}

			fmt.Println("[", file_info.Name(), "]", "- [ Age : ", days, " Days ] -", msg)

		}
	}
}
