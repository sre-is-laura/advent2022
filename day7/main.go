package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MAXDIRSIZE = 100000
	DISKSIZE   = 70000000
	NEEDED     = 30000000
)

type FileInfo struct {
	path string
	name string
	size int
}

// Main function
func main() {
	//fi := readInput("./testinput.txt")
	fi := readInput("./input.txt")
	//fmt.Printf("FileInfo: %+v\n", fi)

	sizesByDir := getSizesByDir((fi))
	//sum := part1(sizesByDir)
	//fmt.Printf("Sum %d\n", sum)

	del := part2(sizesByDir)
	fmt.Printf("Size to delete %d\n", del)
}

func part2(sizesByDir map[string]int) int {
	spaceFree := DISKSIZE - sizesByDir["/"]
	spaceNeeded := NEEDED - spaceFree
	fmt.Printf("Space needed %d\n", spaceNeeded)

	smallestSufficient := sizesByDir["/"]
	for _, size := range sizesByDir {
		if size > spaceNeeded && size < smallestSufficient {
			smallestSufficient = size
		}
	}

	return smallestSufficient
}

func part1(sizesByDir map[string]int) int {
	sum := 0
	for _, size := range sizesByDir {
		if size <= MAXDIRSIZE {
			sum += size
		}
	}
	return sum
}

func getSizesByDir(fi []FileInfo) map[string]int {
	result := make(map[string]int)

	for _, f := range fi {
		path := f.path
		for ; path != "/"; path = cd(path, "..") {
			result[path] += f.size
		}
		result[path] += f.size
	}

	return result
}

func readInput(path string) []FileInfo {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	val := scanner.Text()

	if val != "$ cd /" {
		log.Fatal("Don't know initial directory")
	}
	curdir := "/"
	lastcmd := "cd"

	result := make([]FileInfo, 0)

	for scanner.Scan() {
		val := scanner.Text()

		fields := strings.Split(val, " ")
		if fields[0] == "$" { // command
			lastcmd = fields[1]
			if fields[1] == "cd" {
				curdir = cd(curdir, fields[2])
			} else if fields[1] == "ls" {
				continue
			} else {
				log.Fatal("Don't know command %s", fields[1])
			}
		} else { // command output
			if lastcmd == "ls" {
				if fields[1] == "dir" {
					continue
				} else {
					size, _ := strconv.Atoi(fields[0])
					fi := FileInfo{
						size: size,
						path: curdir,
						name: fields[1],
					}
					result = append(result, fi)
				}

			} else {
				log.Fatal("Don't know last command %s", lastcmd)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func cd(cur string, param string) string {

	result := ""
	if param == ".." {
		lastslash := strings.LastIndex(cur, "/")
		result = cur[0:lastslash]
	} else if strings.HasPrefix(param, "/") {
		result = param
	} else {
		if cur == "/" {
			result = cur + param
		} else {
			result = cur + "/" + param
		}
	}

	if len(result) == 0 {
		result = "/"
	}

	//fmt.Printf("cd %s from %s gives %s\n", param, cur, result)
	return result
}
