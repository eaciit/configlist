package configlist

import (
	"fmt"
	err "github.com/eaciit/errorlib"
	//tk "github.com/eaciit/toolkit"
	"io/ioutil"
	"os"
	"path/filepath"
	//"strings"
)

const (
	objConfigList = "ConfigList"
	objConfig     = "ConfigItem"
)

type IConfigList interface {
	Validate() error
	NewItem() IConfigItem
	Self() IConfigList
	SetSelf(IConfigList)

	ConfigFolder() string
	SetConfigFolder(string)

	Items() []IConfigItem

	Get(string) (IConfigItem, int, bool)
	Set(IConfigItem) error
	Unset(string) error
	Write() error
	Load() error
}

type ConfigListBase struct {
	configFolder string
	items        []IConfigItem
	self         IConfigList
}

func NewList(d IConfigList) IConfigList {
	d.SetSelf(d)
	return d
}

func (c *ConfigListBase) SetSelf(d IConfigList) {
	c.self = d
}

func (c *ConfigListBase) Self() IConfigList {
	return c.self
}

func (l *ConfigListBase) Validate() error {
	if _, e := os.Stat(l.configFolder); e != nil {
		if os.IsNotExist(e) {
			return err.Error(packageName, objConfigList, "Validate",
				fmt.Sprintf("Directory %s is not exitst", l.configFolder))
		} else {
			return err.Error(packageName, objConfigList, "Validate", e.Error())
		}
	}

	return nil
}

func (c *ConfigListBase) ConfigFolder() string {
	return c.configFolder
}

func (c *ConfigListBase) SetConfigFolder(folder string) {
	c.configFolder = folder
}

func (c *ConfigListBase) Items() []IConfigItem {
	c.Inititems()
	return c.items
}

func (l *ConfigListBase) Get(id string) (IConfigItem, int, bool) {
	exist := false
	loop := true
	i := 0
	for loop {
		if i >= len(l.items) {
			loop = false
		} else {
			if l.items[i].GetId() == id {
				loop = false
				exist = true
			} else {
				i++
			}
		}
	}

	if exist {
		return l.items[i], i, true
	} else {
		return nil, -1, exist
	}
}

func (c *ConfigListBase) Inititems() {
	if c.items == nil {
		c.items = make([]IConfigItem, 0)
	}
}

func (c *ConfigListBase) NewItem() IConfigItem {
	return nil
}

func (l *ConfigListBase) Set(item IConfigItem) error {
	var e error
	l.Inititems()

	_, i, exist := l.Get(item.GetId())
	if exist {
		l.items[i] = item
	} else {
		l.items = append(l.items, item)
	}
	e = l.Write()
	if e != nil {
		return e
	}
	return nil
}

func (l *ConfigListBase) Unset(id string) error {
	l.Inititems()

	_, i, exist := l.Get(id)
	if exist {
		filename := filepath.Join(l.ConfigFolder(), fmt.Sprintf("%s.json", id))
		e := os.Remove(filename)
		if e != nil {
			return err.Error(packageName, objConfigList, "Unset",
				fmt.Sprintf("Unable to delete file %s : %s", filename, e.Error()))
		}

		if i == 0 {
			l.items = l.items[1:]
		} else {
			l.items = append(l.items[:i], l.items[i+1:]...)
		}
	}

	return nil
}

func (l *ConfigListBase) Load() error {
	e := l.Self().Validate()
	if e != nil {
		return e
	}

	if l.Self().NewItem() == nil {
		return err.Error(packageName, objConfigList, "Load", "No implementation of NewItem()")
	}

	fis, e := ioutil.ReadDir(l.configFolder)
	if e != nil {
		return err.Error(packageName, objConfigList, "Load", e.Error())
	}

	l.Inititems()
	for _, fi := range fis {
		item := l.Self().NewItem()
		e = NewFromFile(filepath.Join(l.configFolder, fi.Name()), item)
		if e == nil {
			l.Set(item)
		}
	}

	return nil
}

func (l *ConfigListBase) Write() error {
	//_ = "breakpoint"
	e := l.Self().Validate()
	if e != nil {
		return e
	}

	if l.items == nil {
		l.items = make([]IConfigItem, 0)
	}

	for _, item := range l.items {
		fn := filepath.Join(l.configFolder, fmt.Sprintf("%s.json", item.GetId()))
		e = WriteItem(item, fn)
		if e != nil {
			return e
		}
	}

	return nil
}
