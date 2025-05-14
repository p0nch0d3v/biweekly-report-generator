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
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) >= 1 {
		file_path := argsWithoutProg[0]

		file, err := os.Open(file_path)
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
				triple_tab := strings.Index(line, "\t\t\t")
				single_tab_space := strings.Index(line, "    ")
				double_tab_space := strings.Index(line, "        ")
				triple_tab_space := strings.Index(line, "            ")

				if triple_tab == -1 && triple_tab_space == -1 {
					if (single_tab == 0 || single_tab_space == 0) && (double_tab == -1 && double_tab_space == -1) {
						clean_project := strings.ReplaceAll(line, "\t", "")
						clean_project = strings.ReplaceAll(clean_project, "    ", "")
						if slices.Index(projects, clean_project) == -1 {
							projects = append(projects, clean_project)
						}
						current_project = clean_project
					}

					if (single_tab == 0 || single_tab_space == 0) && (double_tab == 0 || double_tab_space == 0) {
						clean_task := strings.ReplaceAll(line, "\t\t", "")
						clean_task = strings.ReplaceAll(clean_task, "        ", "")
						project_task := fmt.Sprintf("%s %s", current_project, clean_task)
						if slices.Index(tasks, project_task) == -1 {
							tasks = append(tasks, project_task)
						}
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
	} else {
    fmt.Println("No input file specified")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
  }
}
