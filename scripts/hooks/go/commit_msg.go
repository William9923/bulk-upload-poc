package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// CommitTypes shows available commit types
var CommitTypes = map[string]string{
	"feat":     "Features",
	"refactor": "Refactor",
	"improve":  "Improvements",
	"fix":      "Bug fixes",
	"perf":     "Performance",
	"revert":   "Revert",
	"docs":     "Documentations",
	"test":     "Test",
	"chore":    "Chores",
}

// CommitTitleTypeKey key to determine type in regex group
const CommitTitleTypeKey = "type"

// CommitTypeRegex build a string regex to determine the available commit type
var CommitTypeRegex = func() string {
	var commitKeys []string
	for key := range CommitTypes {
		commitKeys = append(commitKeys, key)
	}
	return fmt.Sprintf(`(?P<%s>(%s))`, CommitTitleTypeKey, strings.Join(commitKeys, "|"))
}()

const (
	preceedingRegex  = "^"
	descriptionRegex = "(.+)"
	separatorRegex   = ":"
	endingRegex      = "$"
)

func generateTypeRegex() string {
	typeRegexPrefix := "("
	typeRegexSuffix := ")"
	var commitTypeRegex string
	commitKeys := make([]string, 0, len(CommitTypes))
	for key := range CommitTypes {
		commitKeys = append(commitKeys, key)
	}
	commitTypeRegex += (typeRegexPrefix + strings.Join(commitKeys, "|") + typeRegexSuffix)
	return commitTypeRegex
}

func generateRegexCheck() string {
	var regex string
	commitTypeRegex := generateTypeRegex()

	regex += preceedingRegex

	regex += commitTypeRegex
	regex += separatorRegex
	regex += descriptionRegex

	regex += endingRegex
	return regex
}

func commitMessageRegexCheck(commitMsg string) (bool, error) {
	regex := generateRegexCheck()
	correctFormat, err := regexp.MatchString(regex, commitMsg)
	if err != nil {
		return false, err
	}
	if !correctFormat {
		return false, fmt.Errorf("wrong commit message format, regex is: %s", regex)
	}
	return correctFormat, nil
}

// Simple program to determine wether commit message is following current template
// Template: <action>: <description>
func main() {
	// Get first args from the hooks - commit message
	commitMsgFileName := os.Args[1]
	commitMsgByte, err := ioutil.ReadFile(commitMsgFileName)
	if err != nil {
		fmt.Printf("error occured when opening file: %+v\n", err)
	}
	// convert to string and trim commit message
	commitMsg := strings.Trim(string(commitMsgByte), "\n ")

	// check the regex
	passed, err := commitMessageRegexCheck(commitMsg)
	if err != nil {
		fmt.Printf("error occured when checking regex: %+v\n", err)
	}

	// determine status
	var exitStatus int
	if !passed {
		exitStatus = 1
	}
	os.Exit(exitStatus)
}
