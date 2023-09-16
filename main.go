package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func main() {

	prompt := promptui.Select{
		Label: "Select file type",
		Items: []string{"Images", "Videos", "Documents"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	files, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if result == "Images" {
		imageExt := []string{"jpg", "jpeg", "png", "gif", "webp"}
		searchFileByType(files, imageExt)
	}

	if result == "Videos" {
		videoExt := []string{"mp4", "mkv", "avi", "webm", "mov"}
		searchFileByType(files, videoExt)
	}

	if result == "Documents" {
		documentExt := []string{"pdf", "docx", "doc", "ppt", "pptx", "xlsx", "xls", "txt"}
		searchFileByType(files, documentExt)
	}

}

func searchFileByType(files []os.DirEntry, validExtension []string) {
	for _, file := range files {
		if !file.IsDir() {
			fileExt := filepath.Ext(file.Name())
			for _, ext := range validExtension {
				if fileExt == "."+ext {
					color.Green(file.Name())
				}
			}
		}
	}
}
