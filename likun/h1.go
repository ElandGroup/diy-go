package likun

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MyBook struct {
	Name  string
	Price float64
}

func getBook() MyBook {
	book := MyBook{Name: "go", Price: 99}
	return book
}

type What struct {
}

/**struct => json**/
func (What) Struct2Josn(book *MyBook) (result string, err error) {
	mybook, err := json.Marshal(book)
	if err != nil {
		return
	}
	result = string(mybook)
	return
}

/**struct => map[string]interface{}**/
func Struct2MapInt(book MyBook) map[string]interface{} {
	mapbook := map[string]interface{}{
		"Name":  book.Name,
		"Price": book.Price,
	}
	return mapbook
}
func Struct2MapInt2(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

/**map[string]interface{} => struct**/
func MapIner2Struct(mapbook map[string]interface{}) MyBook {
	b, _ := json.Marshal(mapbook)
	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))
	strData := <-ch
	mb := &MyBook{}
	err := json.Unmarshal([]byte(strData), &mb)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(err)
	}
	return *mb
}

/**xml=> json**/
// func Xml2Json((xmlData) byte[]) (result string, err error) {
// 	data := &MyBook{}
// 	err := xml.Unmarshal(xmlData, data)
// 	if nil != err {
// 		fmt.Println("Error unmarshalling from XML", err)
// 		return
// 	}
// 	result, err := json.Marshal(data)
// 	if nil != err {
// 		fmt.Println("Error marshalling to JSON", err)
// 		return
// 	}
// 	fmt.Printf("%v+\n", result)
// }
