package files

import (
	"fmt"
	"io/ioutil"
)

func read(fileName string) ([]byte, error) {
	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %v", err)
	}

	return b, nil
}
