package SharedLib

import (
	"errors"
	"os"
	"strconv"
)

func getenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, errors.New("Environment variable is empty")
	}
	return v, nil
}

func GetenvInt(key string) (int, error) {
	s, err := getenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GetenvBool(key string) (bool, error) {
	s, err := getenvStr(key)
	if err != nil {
		return false, nil
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}
