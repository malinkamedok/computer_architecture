package comp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Util struct{}

func NewUtil() *Util {
	return &Util{}
}

func (u *Util) ReadFile(fileName string) ([]string, error) {
	readFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var result []string
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}
	return result, nil
}

func (u *Util) WriteFile(data []IsaFormat) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("cannot create json file mode: %w", err)
	}

	err = ioutil.WriteFile("result.json", file, 0644)
	return err
}
