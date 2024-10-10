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

	BaseModel `json:"baseModel"`
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
		map[string]interface{}{
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
		[]interface{}{
			map[string]interface{}{
				`a`:    100,
				`haha`: `1111`,
			},
			map[string]interface{}{
				`a`:    100,
				`haha`: `1111`,
			},
		},
	)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, true, testObj[0].A == 100)
	go_test_.Equal(t, "1111", testObj[0].B)
}

func TestFormatType_MustToInt64(t *testing.T) {
	a := `1222222`
	go_test_.Equal(t, int64(1222222), MustToInt64(a))

	a1 := `0x16`
	go_test_.Equal(t, int64(22), MustToInt64(a1))
	a2 := `0o17`
	go_test_.Equal(t, int64(15), MustToInt64(a2))
	a3 := `0b101`
	go_test_.Equal(t, int64(5), MustToInt64(a3))

	var b int = 12
	go_test_.Equal(t, int64(12), MustToInt64(b))

	var c int8 = 12
	go_test_.Equal(t, int64(12), MustToInt64(c))

	var d int16 = 12
	go_test_.Equal(t, int64(12), MustToInt64(d))

	var f int32 = 12
	go_test_.Equal(t, int64(12), MustToInt64(f))

	var g uint8 = 12
	go_test_.Equal(t, int64(12), MustToInt64(g))

	var h uint16 = 12
	go_test_.Equal(t, int64(12), MustToInt64(h))

	var i uint32 = 12
	go_test_.Equal(t, int64(12), MustToInt64(i))

	go_test_.Equal(t, int64(1), MustToInt64(true))

	m_, err := ToInt64(true)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, int64(1), m_)

}

func TestFormatType_MustToUint64(t *testing.T) {
	a := `1222222`
	go_test_.Equal(t, uint64(1222222), MustToUint64(a))

	var b int = 12
	go_test_.Equal(t, uint64(12), MustToUint64(b))

	var c int8 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(c))

	var d int16 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(d))

	var f int32 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(f))

	var g uint8 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(g))

	var h uint16 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(h))

	var i uint32 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(i))

	var j uint64 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(j))

	var k float32 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(k))

	var l float64 = 12
	go_test_.Equal(t, uint64(12), MustToUint64(l))

	var m bool = true
	go_test_.Equal(t, uint64(1), MustToUint64(m))

	var n string = "0xc00007a000"
	go_test_.Equal(t, uint64(824634220544), MustToUint64(n))

	m_, err := ToUint64(n)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, uint64(824634220544), m_)
}

func TestFormatType_ToString(t *testing.T) {
	var a *float64
	b := 0.34
	a = &b
	go_test_.Equal(t, "0.34", ToString(a))

	type BType struct {
		B1 int
	}
	a1 := struct {
		A string
		B BType
	}{`1`, BType{2}}
	go_test_.Equal(t, "{\"A\":\"1\",\"B\":{\"B1\":2}}", ToString(a1))
	a11 := []BType{
		{
			B1: 1,
		},
		{
			B1: 2,
		},
	}
	go_test_.Equal(t, "[{\"B1\":1},{\"B1\":2}]", ToString(a11))

	type Test struct {
		Test1 *string `json:"test1"`
	}
	test_ := Test{}
	go_test_.Equal(t, "*nil", ToString(test_.Test1))

	a2 := 625462456
	go_test_.Equal(t, "625462456", ToString(a2))

	a3 := 0xf43f2
	go_test_.Equal(t, "1000434", ToString(a3))

	a4 := map[string]interface{}{
		"go_test_": "go_test_",
	}
	go_test_.Equal(t, `{"go_test_":"go_test_"}`, ToString(a4))

	go_test_.Equal(t, `["1","2"]`, ToString([]string{"1", "2"}))

	go_test_.Equal(t, `Aa`, ToString([]byte{65, 97}))
}

func TestFormatType_GetValuesInTagFromStruct(t *testing.T) {
	// []*Test{}
	test_ := []*Test{}
	fields := GetValuesInTagFromStruct(test_, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// Test{}
	test1 := Test{}
	fields = GetValuesInTagFromStruct(test1, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *Test{}
	test2 := Test{}
	fields = GetValuesInTagFromStruct(&test2, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// []Test{}
	test3 := []Test{}
	fields = GetValuesInTagFromStruct(test3, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *[]Test{}
	test4 := []Test{}
	fields = GetValuesInTagFromStruct(&test4, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *[]*Test{}
	test5 := []*Test{}
	fields = GetValuesInTagFromStruct(&test5, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))
}

func TestFormatType_MustToBool(t *testing.T) {
	go_test_.Equal(t, true, MustToBool(`true`))
	go_test_.Equal(t, false, MustToBool(`false`))
}

func TestFormatType_MustToInt(t *testing.T) {
	a := "4546"
	go_test_.Equal(t, 4546, MustToInt(a))
}

func TestFormatType_MustToInt8(t *testing.T) {
	a := "12"
	go_test_.Equal(t, int8(12), MustToInt8(a))
}

func TestFormatType_MustToInt32(t *testing.T) {
	a := "4546"
	go_test_.Equal(t, int32(4546), MustToInt32(a))
}

func TestFormatType_MustToUint32(t *testing.T) {
	a := "4546"
	go_test_.Equal(t, uint32(4546), MustToUint32(a))
}

func TestFormatType_MustToFloat64(t *testing.T) {
	go_test_.Equal(t, 4546.3526, MustToFloat64("4546.3526"))
	go_test_.Equal(t, float64(4546), MustToFloat64("4546.0000"))
	go_test_.Equal(t, float64(0.0042), MustToFloat64("0.0042"))
}

func TestFormatType_MustToFloat32(t *testing.T) {
	a := "4546.3526"
	go_test_.Equal(t, float32(4546.3526), MustToFloat32(a))
}

func TestFormatType_GroupSlice(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	group := GroupSlice(strSlice, &GroupOpts{
		CountPerGroup: 3,
	})
	go_test_.Equal(t, 3, len(group))
	go_test_.Equal(t, 3, len(group[0]))
	go_test_.Equal(t, "a", group[0][0])
	go_test_.Equal(t, "c", group[0][2])
	go_test_.Equal(t, 3, len(group[1]))
	go_test_.Equal(t, "d", group[1][0])
	go_test_.Equal(t, "e", group[1][1])
	go_test_.Equal(t, 1, len(group[2]))
	go_test_.Equal(t, "g", group[2][0])

	group1 := GroupSlice(strSlice, &GroupOpts{
		CountPerGroup: 10,
	})
	go_test_.Equal(t, 1, len(group1))
	go_test_.Equal(t, 7, len(group1[0]))
	go_test_.Equal(t, "a", group1[0][0])
	go_test_.Equal(t, "g", group1[0][6])
}

func TestGroupInt(t *testing.T) {
	results := GroupInt(34, 10)
	fmt.Println(results)
}

func TestGroupString(t *testing.T) {
	str := "gsd449998d88gsgsrt"
	results := GroupString(str, &GroupOpts{
		GroupCount: 3,
	})
	fmt.Println(results)
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
