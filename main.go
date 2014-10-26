package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Usage: renamer file1 [file2] [file3] [...]")
		os.Exit(1)
	}

	// check the files exist

	for _, path := range files {
		_, err := os.Lstat(path)
		if err != nil {
			fmt.Printf("file %s does not exist\n", path)
			os.Exit(1)
		}
	}

	// select an editor

	editor := "vi"

	if value := os.Getenv("EDITOR"); value != "" {
		editor = value
	}

	if value := os.Getenv("VISUAL"); value != "" {
		editor = value
	}

	// write list of files to tmp file

	f, err := ioutil.TempFile(os.TempDir(), "renamer")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(strings.Join(files, "\n") + "\n")
	if err != nil {
		panic(err)
	}

	defer os.Remove(f.Name())

	// launch the editor

	cmd := exec.Command(editor, f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// read the file list

	b, err := ioutil.ReadFile(f.Name())
	if err != nil {
		panic(err)
	}
	newFiles := strings.Split(string(b), "\n")

	// ignore the extra newline
	if newFiles[len(newFiles)-1] == "" {
		newFiles = newFiles[0 : len(newFiles)-1]
	}

	if len(files) != len(newFiles) {
		fmt.Printf("number of files in input and output do not match")
		os.Exit(1)
	}

	// rename files

	for i := 0; i < len(files); i++ {
		err := os.Rename(files[i], newFiles[i])
		if err != nil {
			panic(err)
		}
	}
}
