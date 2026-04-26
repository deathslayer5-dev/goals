package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Goals struct {
	time    string
	content []string
}

func main() {
	directory := "/home/deathslayer/goals/goal_list/"
	scanner := bufio.NewScanner(os.Stdin)
	var Time string
	if len(os.Args) <= 1 {
		Time = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Local).Format("02-01-2006")
	} else {
		Time = os.Args[1]
	}
	output := make([]string, 0)
	goals := Goals{
		Time,
		output,
	}
	fileName := fmt.Sprintf("%sgoals_%s.txt", directory, goals.time)
	head := fmt.Sprintf("Goals for: %v\n", Time)
	fmt.Println(head)
	fmt.Println("Input ? to exit")
	output = append(output, head)
	goalCount := 0
	for {
		fmt.Print("Enter goal: ")
		if scanner.Scan() {
			input := scanner.Text()
			switch input {
			case "?":
				fmt.Println("? pressed, exiting ...")
				goals.content = output
				content := strings.Join(goals.content, "\n")
				fmt.Println(content)
				err := os.WriteFile(fileName, []byte(content), 0644)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Goals saved to ", fileName)
				}
				return
			case "?g":
				fmt.Println("Current Goals:", goalCount)
				content := strings.Join(output, "\n")
				fmt.Println(content)
			case "?d":
				fmt.Println("Deleting last goal")
				goalCount--
				if goalCount <= 0 {
					goalCount = 0
					output = []string{head}
				}
				output = output[:goalCount]
				content := strings.Join(output, "\n")
				fmt.Println(content)
			default:
				goalCount++
				goal := fmt.Sprintf("%d: %s", goalCount, input)
				output = append(output, goal)
			}

		}
		if err := scanner.Err(); err != nil {
			_, err := fmt.Fprintln(os.Stderr, "reading standard input:", err)
			if err != nil {
				return
			}
		}
	}

}
