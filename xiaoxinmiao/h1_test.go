package xiaoxinmiao

import (
	diy "diy-go"
	"fmt"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_StructToJson(t *testing.T) {
	book := &diy.Book{
		Name:  "go",
		Price: 12.5,
	}
	result, err := StructToJson(book)
	fmt.Println(result)
	test.Ok(t, err)
}

func Test_StructToMap(t *testing.T) {
	book := &diy.Book{
		Name:  "go",
		Price: 12.5,
	}
	result := StructToMap(book)
	fmt.Println(result)
}

func Test_StructToXml(t *testing.T) {
	book := &diy.Book{
		Name:  "go",
		Price: 12.5,
	}
	result, err := StructToXml(book)
	fmt.Println(result)
	test.Ok(t, err)
}

func Test_MapToJson(t *testing.T) {
	maps := map[string]interface{}{
		"name":  "go",
		"price": 12.6,
	}
	result, err := MapToJson(maps)
	fmt.Println(result)
	test.Ok(t, err)
}

func Test_MapToStruct(t *testing.T) {
	maps := map[string]interface{}{
		"name":  "go",
		"price": 12.6,
	}
	book, err := MapToStruct(maps)
	fmt.Println(book)
	test.Ok(t, err)
}

func Test_MapToXml(t *testing.T) {
	maps := map[string]interface{}{
		"name":  "go",
		"price": 12.6,
	}
	result, err := MapToXml(maps)
	fmt.Println(result)
	test.Ok(t, err)
}

func Test_JsonToStruct(t *testing.T) {
	jsonStr := `{
		"name":  "go",
		"price": 12.6
	}`
	book, err := JsonToStruct(jsonStr)
	fmt.Println(book)
	test.Ok(t, err)
}

func Test_JsonToMap(t *testing.T) {
	jsonStr := `{
		"name":  "go",
		"price": 12.6
	}`
	result, err := JsonToMap(jsonStr)
	fmt.Println(result)
	test.Ok(t, err)

}

func Test_JsonToXml(t *testing.T) {
	jsonStr := `{
		"name":  "go",
		"price": 12.6
	}`
	result, err := JsonToXml(jsonStr)
	fmt.Println(result)
	test.Ok(t, err)

}

func Test_XmlToJson(t *testing.T) {
	xmlStr := `
	<name>go</name>
	<price>12.7</price>
	`
	result, err := XmlToJson(xmlStr)
	fmt.Println(result)
	test.Ok(t, err)
}

func Test_XmlToStruct(t *testing.T) {
	xmlStr := `<xml>
	<name>go</name>
	<price>12.7</price>
	</xml>`
	book, err := XmlToStruct(xmlStr)
	fmt.Println(book)
	test.Ok(t, err)
}

func Test_XmlToMap(t *testing.T) {
	xmlStr := `<xml>
	<name>go</name>
	<price>12.7</price>
	</xml>`
	result, err := XmlToMap(xmlStr)
	fmt.Println(result)
	test.Ok(t, err)
}
