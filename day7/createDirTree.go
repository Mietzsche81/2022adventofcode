package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type FileRecord struct {
	name string
	size int
}

type Directory struct {
	name     string
	size     int
	parent   *Directory
	files    []FileRecord
	children map[string]*Directory
}

type Command struct {
	directory *Directory
	command   string
	output    []string
}

func main() {

	//
	// Read input data
	//

	fileName := strings.TrimSpace(os.Args[1])

	//
	// Parse commands, build file tree, and sort by size
	//

	commandSequence := ParseData(fileName)
	fileTree := TraverseCommands(commandSequence)
	CalculateDirectorySize(&fileTree)

	//
	// Report sum of smaller than 100KB
	//
	smallFolders := FilterBySize(&fileTree, 100000)
	total := 0
	for _, size := range smallFolders {
		total += size
	}
	fmt.Println(total)
}

func ParseData(fileName string) (commands []Command) {
	// Read from file
	fmt.Printf("Processing input from: '%s'\n", fileName)
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	scanner := bufio.NewScanner(fin)

	// Read each command
	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '$' {
			// Parse command
			commands = append(commands,
				Command{
					command: strings.TrimSpace(line[1:]),
				},
			)
		} else {
			// append output
			commands[len(commands)-1].output = append(
				commands[len(commands)-1].output,
				strings.TrimSpace(line),
			)
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func TraverseCommands(commands []Command) (root Directory) {
	// Create root
	root.name = "/"
	root.children = make(map[string]*Directory)
	var pwd *Directory = &root
	for i := range commands {
		step := &commands[i]
		step.directory = pwd
		if strings.HasPrefix(step.command, "cd ") {
			// change directory
			dest := strings.TrimSpace(step.command[3:])
			if dest == "/" {
				// Change root
				pwd = &root
			} else if dest == ".." {
				// Change up
				pwd = pwd.parent
			} else {
				// Change down
				if child, exists := pwd.children[dest]; exists {
					pwd = child
				} else {
					// Add directory to tree if not indexed yet
					pwd.children[dest] = &Directory{
						name:     dest,
						parent:   pwd,
						children: make(map[string]*Directory),
					}
					pwd = pwd.children[dest]
				}
			}
		} else if strings.HasPrefix(step.command, "ls") {
			for _, entry := range step.output {
				words := strings.Fields(entry)
				spec := words[0]
				dest := strings.TrimSpace(entry[len(spec)+1:])
				if spec == "dir" {
					if _, exists := pwd.children[dest]; !exists {
						// Add directory to tree if not indexed yet
						pwd.children[dest] = &Directory{
							name:     dest,
							parent:   pwd,
							children: make(map[string]*Directory),
						}
					}
				} else {
					// Add files
					size, err := strconv.Atoi(spec)
					if err != nil {
						log.Fatal(err)
					}
					pwd.files = append(pwd.files,
						FileRecord{
							name: strings.TrimSpace(entry[len(spec)+1:]),
							size: size,
						},
					)
				}
			}
		} else {
			err := fmt.Errorf("TraverseCommands: Unknown command '%s'", step.command)
			log.Fatal(err)
		}
	}

	return
}

func CalculateDirectorySize(dir *Directory) int {
	total_size := 0
	// Add local files
	for i := range dir.files {
		total_size += dir.files[i].size
	}
	// Recurse child folders
	for i := range dir.children {
		total_size += CalculateDirectorySize(dir.children[i])
	}
	dir.size = total_size
	return total_size
}

func FilterBySize(root *Directory, maxSize int) map[*Directory]int {
	results := make(map[*Directory]int)
	var recurse func(pwd *Directory, cumulate *map[*Directory]int)
	recurse = func(pwd *Directory, cumulate *map[*Directory]int) {
		if pwd.size < maxSize {
			(*cumulate)[pwd] = pwd.size
		}
		for i := range pwd.children {
			recurse(pwd.children[i], cumulate)
		}
	}
	recurse(root, &results)
	return results
}
