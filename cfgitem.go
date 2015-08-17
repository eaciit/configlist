package configlist

import (
	"encoding/json"
	"fmt"
	er "github.com/eaciit/errorlib"
	"io/ioutil"
)

type IConfigItem interface {
	GetId() string
}

type ConfigItemBase struct {
	Id string
}

func (c *ConfigItemBase) GetId() string {
	return c.Id
}

func NewFromFile(filename string, d IConfigItem) error {
	_ = "breakpoint"
	bs, e := ioutil.ReadFile(filename)
	if e != nil {
		return er.Error(packageName, "", "NewFromFile.(ReadFile)", e.Error())
	}

	e = json.Unmarshal(bs, d)
	if e != nil {
		return er.Error(packageName, "", "NewFromFile.(Unmarshal)", e.Error())
	}

	return nil
}

func WriteItem(d IConfigItem, filename string) error {
	//_ = "breakpoint"
	bs, e := json.MarshalIndent(d, "", "\t")
	if e != nil {
		return er.Error(packageName, objConfig, "Write",
			fmt.Sprintf("Unable to write to %s : %s", filename, e.Error()))
	}
	e = ioutil.WriteFile(filename, bs, 0644)
	if e != nil {
		return er.Error(packageName, objConfig, "Write",
			fmt.Sprintf("Unable to write to %s : %s", filename, e.Error()))
	}
	return nil
}
