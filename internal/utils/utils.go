package utils

import (
	"FileTimeChanger/internal/validator"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// HandleGenericErr is used to improve the readability of other functions by
// reducing duplication of basic error handling code.
// When err is nil, HandleGenericErr will print the success string (to let
// the user know the method above was successful) unless the success string
// is set to an empty string.
// When err is not nil, HandleGenericErr prints the error context to std out
// and if the "FTC_DEV" flag is set the error stack trace is printed for
// debugging purposes.
// TODO: Print error stack trace
func HandleGenericErr(successStr, errContext string, err error) {
	if err == nil {
		if successStr != "" {
			fmt.Printf("%s\n", successStr)
		}

		return
	}

	// User-friendly error message about what went wrong
	fmt.Printf("Error: %s: %s\n", errContext, err)

	// Print extra debug information for devs
	_, envPresent := os.LookupEnv("FTC_DEV")
	if envPresent {
		// TODO: Print error stack trace
	}
}

// HandleFatalErr calls HandleGenericErr and then terminates the program
// if an error has occurred.
// This method should not be used unless the program is unable to continue
// in the current state.
func HandleFatalErr(successStr, errContext string, err error) {
	HandleGenericErr(successStr, errContext, err)

	if err != nil {
		fmt.Println("Exiting Program...")
		os.Exit(1)
	}
}

// ReadString is a wrapper around the bufio readString function that strips the
// delimiter and trims the string.
// ReadString returns a string representing the input from the reader and the
// first error that occurred.
func ReadString(reader *bufio.Reader, delim byte) (string, error) {
	str, err := reader.ReadString(delim)
	if err != nil {
		return "", err
	}

	if len(str) > 0 && str[len(str)-1] == delim {
		str = str[:len(str)-1]
	}

	return strings.TrimSpace(str), nil
}

// Prompt is used to handle requesting input from Stdin.
// Prompt will request input from the user (message) and then validate the
// input. If the input is valid, a message will be shown to the user informing
// them the data was validated, but if the input is invalid the method will
// inform the user the data is invalid and prompt the user again until the data
// from the user is valid.
// TODO: Implement quit and max tries.
// TODO: Is it possible to pass the return type wanted as a param?
func Prompt(message, successStr, errMsg string, validate validator.InputValidator) string {
	// Initialize the buffered reader to read from standard in
	stdInReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", message)
		input, err := ReadString(stdInReader, '\n')
		HandleGenericErr("", "Error reading input", err)

		if validate(input) {
			fmt.Printf("%s\n", successStr)
			return input
		}

		fmt.Println()
		fmt.Printf("Error: %s, please try again...\n", errMsg)
	}
}
