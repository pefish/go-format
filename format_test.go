package go_format

import (
	"fmt"
	"testing"
	"time"

	go_test_ "github.com/pefish/go-test"
)

type Test struct {
	UserId       uint64    `json:"user_id,omitempty"`
	Type         uint64    `json:"type"`
	OrderNumber  string    `json:"order_number"`
	Price        float64   `json:"price"`
	Amount       float64   `json:"amount"`
	TransferMemo string    `json:"tranfer_memo"`
	Status       uint64    `json:"status"`
	Time         time.Time `json:"time"`

	BaseModel
}

type BaseModel struct {
	Id        uint64 `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func TestFormatClass_StructToMap(t *testing.T) {
	type Nest struct {
		C string `json:"c"`
	}
	type Test struct {
		A    uint64 `json:"a"`
		B    string `json:"haha"`
		Nest `json:"nest,flatten"`
	}

	testObj := Test{
		A:    100,
		B:    `1111`,
		Nest: Nest{C: "q"},
	}
	testMap := StructToMap(testObj)
	go_test_.Equal(t, true, testMap["a"].(uint64) == 100)
	go_test_.Equal(t, "q", testMap["c"].(string))
}

func TestFormatClass_MapToStruct(t *testing.T) {
	type Test struct {
		A uint64 `json:"a"`
		B string `json:"haha"`
	}
	testObj := Test{}
	err := MapToStruct(
		&testObj,
		map[string]any{
			`a`:    100,
			`haha`: `1111`,
		},
	)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, true, testObj.A == 100)
	go_test_.Equal(t, "1111", testObj.B)
}

func TestFormatClass_SliceToStruct(t *testing.T) {
	type Test struct {
		A uint64 `json:"a"`
		B string `json:"haha"`
	}
	testObj := []Test{}
	err := SliceToStruct(
		&testObj,
		[]any{
			map[string]any{
				`a`:    100,
				`haha`: `1111`,
			},
			map[string]any{
				`a`:    100,
				`haha`: `1111`,
			},
		},
	)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, true, testObj[0].A == 100)
	go_test_.Equal(t, "1111", testObj[0].B)
}

func TestFormatType_FetchTags(t *testing.T) {
	// []*Test{}
	test_ := []*Test{}
	fields := FetchTags(test_, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// Test{}
	test1 := Test{}
	fields = FetchTags(test1, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *Test{}
	test2 := Test{}
	fields = FetchTags(&test2, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// []Test{}
	test3 := []Test{}
	fields = FetchTags(test3, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *[]Test{}
	test4 := []Test{}
	fields = FetchTags(&test4, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *[]*Test{}
	test5 := []*Test{}
	fields = FetchTags(&test5, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	test6 := struct {
		A string `json:"a"`
		B []struct {
			C string `json:"c"`
		} `json:"b"`
	}{}
	fields = FetchTags(&test6, `json`)
	go_test_.Equal(t, "[a b]", fmt.Sprint(fields))

}

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	result := UnderscoreToCamelCase("gsfghs_bfgbsg_sgg")
	fmt.Println(result)
	go_test_.Equal(t, "GsfghsBfgbsgSgg", result)
}

func TestCamelCaseToUnderscore(t *testing.T) {
	result := CamelCaseToUnderscore("FuckYou")
	go_test_.Equal(t, "fuck_you", result)
}

func TestCamelCaseToWords(t *testing.T) {
	results := CamelCaseToWords("fuckYou988MyGod")
	go_test_.Equal(t, "fuck", results[0])
	go_test_.Equal(t, "You988", results[1])
	go_test_.Equal(t, "My", results[2])
	go_test_.Equal(t, "God", results[3])
}

func TestEncodePefish(t *testing.T) {
	result := EncodePefish("fuckYou988MyGod")
	fmt.Println(result)
}

func TestDecodePefish(t *testing.T) {
	result, err := DecodePefish("YzFlY185YzFjMjVmOC1jY2I9MTlfZS00NDJmLTk1OWYtPTg4MmY4YzA3YV8xZWEwOTUxOWUwZDk9MGYzZjFlODMyNjkxYmIy")
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "fuckYou988MyGod", result)
}
