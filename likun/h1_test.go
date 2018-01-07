package likun

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_Struct2Josn(t *testing.T) {
	book := getBook()
	result, err := What{}.Struct2Josn(&book)
	fmt.Printf("%+v\n", result)
	test.Ok(t, err)
}

func Test_Struct2MapInt(t *testing.T) {
	book := getBook()
	mapbook := Struct2MapInt(book)
	fmt.Printf("%+v\n", mapbook)
	test.Ok(t, nil)
}

func Test_Struct2MapInt2(t *testing.T) {
	book := getBook()
	mapbook := Struct2MapInt2(book)
	fmt.Printf("%+v\n", mapbook)
}

func Test_MapIner2Struct(t *testing.T) {
	mapbook := map[string]interface{}{
		"Name":  "go",
		"Price": 99,
	}
	book := MapIner2Struct(mapbook)
	fmt.Printf("%+v\n", book)
}

/**xml=> json**/
func Test_Xml2Json(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
    <MyBook>
        <Name>go</Name>
        <Price>99</Price>
    </MyBook>`)
	data := &MyBook{}
	err := xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", data)
}

/**xml => struct**/
func Test_Xml2Struct(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
		<MyBook>
			<Name>go</Name>
			<Price>99</Price>
		</MyBook>`)
	data := &MyBook{}
	err := xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println(err)
		return
	}
	b, _ := json.Marshal(data)
	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))
	strData := <-ch
	mb := &MyBook{}
	err1 := json.Unmarshal([]byte(strData), &mb)
	if nil != err1 {
		fmt.Println(err1)
		return
	}
	fmt.Printf("%+v\n", *mb)
}

/**xml=>map[string]interface{}**/
func Test_Xml2Map(t *testing.T) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
		<MyBook>
			<Name>go</Name>
			<Price>99</Price>
		</MyBook>`)
	data := &MyBook{}
	err := xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println(err)
		return
	}
	mapbook := Struct2MapInt(*data)
	fmt.Printf("%+v\n", mapbook)
}

/**struct => xml**/
func Test_Strut2Xml(t *testing.T) {
	book := getBook()
	output, err := xml.MarshalIndent(book, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	str_output := string(output)
	fmt.Printf("%+v\n", str_output)
}
