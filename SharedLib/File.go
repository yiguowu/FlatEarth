package SharedLib

import "io/ioutil"

func ReadFile(fileName string, optional ...string) []byte {
	path := ""
	if len(optional) > 0 {
		path = optional[0]
		fileName = path + "/" + fileName
	}
	dat, err := ioutil.ReadFile(fileName)
	PanicOnError(err, WARNING)
	return []byte(dat)
}
