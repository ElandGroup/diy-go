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
	result, err := XXM{}.StructToJson(book)
	fmt.Println(result)
	test.Ok(t, err)
}
