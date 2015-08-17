# ConfigList

ConfigList is a tool to create list of json file, normally for config, written in Go

## 0. include
```
import "github.com/eaciit/configlist"
```

## 1. setup new list & item
```
type AutoNumber struct {
  configlist.ConfigItemBase
  Title string
  CurrentNo int
}

type AutoNumberList struct{
  configlist.ConfigListBase
}
```

## 2. extends configitem type
```
type (a *AutoNumberList) NewItem() configlist.IConfigItem {
  i := new(AutoNumber)
  return i
}
```

## 3. setup list for further usage and load
```
var l := configlist.NewList(new(AutonumberList))
l.Load()
```

### 4. add list item
```
var a := new(AutoNumber)
a.Id = "NewNumber"
a.Title = "New number system"
a.CurrentNo = 1
l.Set(a)
```

### 5. update item
```
a, i, f := l.Get("NewNumber")

//--- found
if f {
  a.CurretNo++
} else {
  a = new(AutoNumber)
  a.Id = "NewNumber"
  a.Title = "New number system"
  a.CurrentNo = 1
}
l.Set(a)
```

### 6. delete an item
```
e := l.Unset("NewNumber")
```
