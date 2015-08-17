package configlist

import (
	"fmt"
	tk "github.com/eaciit/toolkit"
	"io/ioutil"
	"testing"
)

const (
	folder = "/users/ariefdarmawan/Temp/test"
)

type CfgTest struct {
	ConfigItemBase
	Title   string
	Command string
}

type CfgList struct {
	ConfigListBase
}

func (c *CfgList) NewItem() IConfigItem {
	d := new(CfgTest)
	return d
}

func TestWrite(t *testing.T) {
	fmt.Println("Test Write")
	cfgs := NewList(new(CfgList))
	cfgs.SetConfigFolder(folder)
	cfgs.Load()
	ct := new(CfgTest)
	ct.Id = "Test"
	ct.Title = "Test 1"
	ct.Command = "ls -al"
	e := cfgs.Set(ct)
	//e := cfgs.Write()
	if len(cfgs.Items()) == 0 {
		t.Errorf("Unable to add line")
	}
	if e != nil {
		t.Error(e.Error())
	} else {
		fmt.Println("OK")
	}
	fmt.Println("")
}

func TestLoad(t *testing.T) {
	fmt.Println("Test Load")
	cfgs := NewList(new(CfgList))
	cfgs.SetConfigFolder(folder)
	e := cfgs.Load()
	if e == nil {
		fmt.Println("OK")
		fmt.Printf("Has %d config(s) \n", len(cfgs.Items()))
		if len(cfgs.Items()) > 0 {
			fmt.Printf("Sample of 1st config: \n%v \n", tk.JsonString(cfgs.Items()[0]))
		}
	} else {
		t.Error(e.Error())
	}
	fmt.Println("")
}

func TestUnset(t *testing.T) {
	fmt.Println("Test Unset")
	cfgs := NewList(new(CfgList))
	cfgs.SetConfigFolder(folder)
	cfgs.Load()
	cfgs.Unset("Test")

	fis, e := ioutil.ReadDir(cfgs.ConfigFolder())
	if e != nil {
		t.Error(e.Error())
	} else {
		if len(fis) != len(cfgs.Items()) {
			t.Errorf("Fail. Expects %d config(s), found %d config(s)", len(fis), len(cfgs.Items()))
		} else {
			fmt.Println("OK")
		}
	}

	fmt.Println("")
}
