package SharedLib

import (
	"io/ioutil"
	"os"
	"serviceGroup/SharedLib"
	"strconv"
)

func ReadFile(fileName string, optional ...string) []byte {
	path := ""
	if len(optional) > 0 {
		path = optional[0]
		fileName = path + "/" + fileName
	}
	dat, err := ioutil.ReadFile(fileName)
	PanicOnError(err, WARNING)
	return dat
}

func WriteFile(dat []byte, fileName string, optional ...string) {
	path := ""
	mode := os.FileMode.Perm(0644)
	if len(optional) > 0 {
		fileMode, err := strconv.Atoi(optional[0])
		SharedLib.PanicOnError(err, SharedLib.WARNING)
		if err != nil {
			return
		}
		mode = os.FileMode.Perm(fileMode)
	}
	if len(optional) > 1 {
		path = optional[1]
		fileName = path + "/" + fileName
	}
	err := ioutil.WriteFile(fileName, dat, mode)
	SharedLib.PanicOnError(err, SharedLib.WARNING)
}
