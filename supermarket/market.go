package supermarket

import (
	"fmt"
	s "strings"
)

var item = map[string]interface{}{
	"flour":  90.10,
	"cheese": 35.60,
	"meat":   349.99,
	"onions": 30.50,
}

func Get(name string) {
	s.ToLower(name)
	//fmt.Println(item)
	if item[name] != nil {
		fmt.Println("Price of", name, "is", item[name])
	} else {
		fmt.Println("Item not found or deleted")
	}
}

func Post(name string, value interface{}) {
	s.ToLower(name)
	if item[name] != nil {
		fmt.Println("Item already exists")
	} else {
		if item == nil {
			item = make(map[string]interface{})
			item[name] = value
		} else {
			item[name] = value
		}
	}

}

func Put(name string, value interface{}) {
	s.ToLower(name)
	if item[name] == nil {
		fmt.Println("Item doesnot exist")
	} else {
		item[name] = value
	}
}

func Delete(name string) {
	s.ToLower(name)
	if item[name] != nil {
		delete(item, name)
	} else {
		fmt.Println("Item doesn't exist or already deleted")
	}
}
func Print() {
	fmt.Println("Items in supermarket:")
	fmt.Println(item)
	//return item
}
