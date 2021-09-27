package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	path,err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	printDirectory(path,0)
}

func printDirectory(path string, depth int) {
	stat,err := os.Lstat(path)
	if err != nil {
		panic(err)
	}
	printNameWithDepth(stat.Name(), depth)
	if stat.IsDir() {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		dirs,err := file.Readdirnames(0)
		if err != nil {
			panic(err)
		}
		sort.Strings(dirs)
		for _, dir := range dirs {
			printDirectory(path + "/" + dir, depth + 1)
		}
	}
}

func printNameWithDepth(name string, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("---")
	}
	fmt.Println(name)
}