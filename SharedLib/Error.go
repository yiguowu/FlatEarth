package SharedLib

import "fmt"

type Severity int

const (
	INFO Severity = iota + 1
	WARNING
	ERROR
	FATAL
)

func (s Severity) String() string {
	errorLevel := [...]string{"INFO", "WARNING", "ERROR", "FATAL"}
	if s < INFO || s > FATAL {
		return "UNKNOWN"
	}
	return errorLevel[s]
}

func PanicOnError(err error, level Severity) {
	LEVEL, _ := GetenvInt("LUKKACTL_VERBOSE")
	var currentLevel = ERROR
	if LEVEL != 0 {
		currentLevel = Severity(LEVEL)
	}
	if err != nil {
		fmt.Println(err)
		if level >= currentLevel {
			panic(err)
		}
	}
}
