# diy-go

## subject 1 date:2018-12-25

```
To achieve the following methods:
        StructToJson
        StructToMap
        MapToStruct
        XmlToJson
        XmlToJson
        XmlToMap
        StructToXml
sample
    func (XXM) StructToJson(book *diy.Book) (result string, err error) {
        b, err := json.Marshal(book)
        if err != nil {
            return
        }
        result = string(b)
        return
    }
```

## subject 2 date:2018-01-07
```
https://staging.p2shop.cn/fruit/v1/fruits
request api
        GetFruit
        PostFruit
        PutFruit
        DeleteFruit
```

