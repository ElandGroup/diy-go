package xiaoxinmiao

import (
	diy "diy-go"
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"

	"github.com/relax-space/go-kit/base"

	"github.com/relax-space/go-kitt/mapstruct"

	xj "github.com/basgys/goxml2json"
	"github.com/fatih/structs"
)

func StructToJson(book *diy.Book) (result string, err error) {
	b, err := json.Marshal(book)
	if err != nil {
		return
	}
	result = string(b)
	return
}

func StructToMap(book *diy.Book) (result map[string]interface{}) {
	s := structs.New(book)
	s.TagName = "json"
	result = s.Map()
	return
}

func StructToXml(book *diy.Book) (result string, err error) {
	b, err := xml.Marshal(book)
	result = string(b)
	return
}

func MapToJson(maps map[string]interface{}) (result string, err error) {
	b, err := json.Marshal(maps)
	result = string(b)
	return
}

func MapToStruct(maps map[string]interface{}) (book *diy.Book, err error) {
	book = &diy.Book{}
	err = mapstruct.Decode(maps, book)
	return
}

func MapToXml(maps map[string]interface{}) (result string, err error) {
	xmlMap := base.XmlMap{}
	for k, v := range maps {
		xmlMap[k] = v
	}
	b, err := xml.MarshalIndent(xmlMap, "", " ")
	result = string(b)
	return
}

func JsonToStruct(jsonStr string) (book *diy.Book, err error) {
	book = &diy.Book{}
	err = json.Unmarshal([]byte(jsonStr), book)
	return
}

func JsonToMap(jsonStr string) (result *map[string]interface{}, err error) {
	result = &map[string]interface{}{}
	err = json.Unmarshal([]byte(jsonStr), result)
	return
}

func JsonToXml(jsonStr string) (result string, err error) {
	var v map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &v)
	if err != nil {
		return
	}
	result, err = MapToXml(v)
	return
}

func XmlToJson(xmlStr string) (result string, err error) {
	b, err := xj.Convert(strings.NewReader(xmlStr))
	result = b.String()
	return
}

func XmlToStruct(xmlStr string) (book *diy.Book, err error) {
	book = &diy.Book{}
	err = xml.Unmarshal([]byte(xmlStr), book)
	return
}

func XmlToMap(xmlStr string) (result map[string]interface{}, err error) {
	data, err := base.XmlToMap(xmlStr)
	if err != nil {
		return
	}
	v, has := data["xml"]
	if !has {
		err = errors.New("xml tag is not found.")
		return
	}
	result = v.(map[string]interface{})
	return
}
