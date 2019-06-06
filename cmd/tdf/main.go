// This is a simple utility that returns an absolute filename to a
// test resource within the test directory independent of where the
// current working directory is within a go project.  If the file does
// not exist an empty name is returned.
package main

import (
	"fmt"
	"os"

	"github.com/borud/tdf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\nUsage: %s <filename>\n\n", os.Args[0])
		return
	}
	fmt.Printf("%s\n", tdf.Find(os.Args[1]))
}
