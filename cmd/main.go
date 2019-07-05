package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"FileTimeChanger/internal/file"
	"FileTimeChanger/internal/utils"
)

func main() {
	fileData := gatherData()

	// TODO: Try to keep aTime the same if possible
	aTime := time.Now()

	err := os.Chtimes(fileData.GetFilepath(), aTime, fileData.GetModifiedTime())
	utils.HandleGenericErr("Sorry, there was an error changing the modified time of the file", err)

	fmt.Printf("The time for %s has been changed to %s\n", fileData.GetFilepath(), fileData.GetModifiedTime().Format(time.RFC1123))
}

func gatherData() *file.Data {
	// Initialize the buffered reader to read from standard in
	stdInReader := bufio.NewReader(os.Stdin)

	// Prompt for the filepath
	fmt.Printf("Please enter a the path to a file: ")
	filepath, err := utils.ReadString(stdInReader, '\n')
	utils.HandleGenericErr("Sorry, there was an error reading the path to the file", err)

	fi, err := os.Stat(filepath)
	utils.HandleGenericErr("Sorry, there was an error getting important information on the file you selected", err)

	fmt.Printf("Please enter the new time modified for the file (e.g. %s): ", time.RFC1123)
	timeStr, err := utils.ReadString(stdInReader, '\n')
	mTime, err := time.ParseInLocation(time.RFC1123, timeStr, fi.ModTime().Location())
	utils.HandleGenericErr("Sorry, there was an error reading the time you entered", err)

	return file.NewData(filepath, mTime)
}
