package printprogramname

import "os"

func PrintProgramname() string {
	return os.Args[0]
}