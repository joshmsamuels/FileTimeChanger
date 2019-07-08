package main

import (
	"fmt"
	"time"

	"FileTimeChanger/internal/file"
	"FileTimeChanger/internal/utils"
	"FileTimeChanger/internal/validator"
)

func main() {
	fileData := file.NewDataBlank()

	fileData.SetFilepath(
		utils.Prompt("Please enter a the path to a file",
			"Filepath was stored successfully",
			"Unable to read the filepath",
			validator.ValidateFilepath),
	)

	timeStr := utils.Prompt(
		fmt.Sprintf("Please enter the new time I should set for %s (format: %s)", fileData.GetFilepath(), file.TimeFormat),
		"Time modified was entered successfully",
		"Unable to store the time entered",
		validator.ValidateDate,
	)

	mTime, err := time.Parse(file.TimeFormat, timeStr)

	fileData.SetModifiedTime(mTime)

	err = fileData.ChangeFileTime()
	utils.HandleGenericErr(fmt.Sprintf("The time for %s was successfully changed to %s", fileData.GetFilepath(), fileData.GetModifiedTime()), "Sorry, there was an error changing the file time", err)
}
