// Package tdf (test data find) is a utility to locate the test files
// regardless of where you are executing the tests from.  This code
// assumes that you store your test data under "test" as per the
// standard project layout proposed here:
// https://github.com/golang-standards/project-layout
package tdf

import (
	"log"
	"os"
	"os/user"
	"path"
	"strings"
)

// Find searches from the current working directory to the user's home
// directory to see if we can find the file within the test directory
// given by filename.  If we can't find the file it will return an
// empty string.  If we are unable to get the current working
// directory or the current home directory we terminate with a
// log.Fatalf().
func Find(filename string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to get current working directory: %v", err)
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get home directory for %s: %v", usr.Username, err)
	}

	levelsToProbe := len(strings.Split(usr.HomeDir, string(os.PathSeparator)))
	elements := strings.Split(currentDir, string(os.PathSeparator))
	for {
		if len(elements) < levelsToProbe {
			return ""
		}

		probe := path.Join(strings.Join(elements, string(os.PathSeparator)), "test", filename)
		_, err := os.Stat(probe)
		if err == nil {
			return probe
		}

		elements = elements[:len(elements)-1]
	}
}
