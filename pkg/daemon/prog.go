package daemon

import (
	"fmt"
	"os"
	"runtime"
)

func Exit(code int, msg string) {
	var callerDetails string
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		file, line := details.FileLine(pc)
		callerDetails = fmt.Sprintf("%s:%d", file, line)
	}
	if code != -1 {
		if msg != "" {
			fmt.Fprintf(os.Stderr, "%s\t %s\n", callerDetails, msg)
		}
		os.Exit(code)
	}

	if msg != "" {
		fmt.Fprintf(os.Stdout, "%s\t%v\n", callerDetails, msg)
	}
	os.Exit(-1)
}
