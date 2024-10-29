package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("./test.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var projects []string
	var tasks []string
	var current_project string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			single_tab := strings.Index(line, "\t")
			double_tab := strings.Index(line, "\t\t")

			if single_tab > -1 && double_tab == -1 {
				clean_project := strings.ReplaceAll(line, "\t", "")
				if slices.Index(projects, clean_project) == -1 {
					projects = append(projects, clean_project)
				}
				current_project = clean_project
			}

			if single_tab > -1 && double_tab > -1 {
				clean_task := strings.ReplaceAll(line, "\t\t", "")
				project_task := fmt.Sprintf("%s %s", current_project, clean_task)
				if slices.Index(tasks, project_task) == -1 {
					tasks = append(tasks, project_task)
				}
			}
		}
	}

	slices.Sort(projects)

	fmt.Println(" ")
	for _, value := range projects {
		fmt.Println(value)
	}

	slices.Sort(tasks)

	fmt.Println(" ")
	for _, value := range tasks {
		fmt.Println(value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
