package vm

import (
	"ca/internal/comp"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Util struct{}

func (u *Util) ReadISAFile(fileName string) ([]comp.IsaFormat, error) {
	jsonFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot read file with commands")
	}

	var isa []comp.IsaFormat
	err = json.Unmarshal(jsonFile, &isa)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal json")
	}
	return isa, nil
}
