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
	testMap := FormatInstance.StructToMap(testObj)
	go_test_.Equal(t, true, testMap["a"].(uint64) == 100)
	go_test_.Equal(t, "q", testMap["c"].(string))
}

func TestFormatClass_MapToStruct(t *testing.T) {
	type Test struct {
		A uint64 `json:"a"`
		B string `json:"haha"`
	}
	testObj := Test{}
	err := FormatInstance.MapToStruct(
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
	err := FormatInstance.SliceToStruct(
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
	go_test_.Equal(t, int64(1222222), FormatInstance.MustToInt64(a))

	a1 := `0x16`
	go_test_.Equal(t, int64(22), FormatInstance.MustToInt64(a1))
	a2 := `0o17`
	go_test_.Equal(t, int64(15), FormatInstance.MustToInt64(a2))
	a3 := `0b101`
	go_test_.Equal(t, int64(5), FormatInstance.MustToInt64(a3))

	var b int = 12
	go_test_.Equal(t, int64(12), FormatInstance.MustToInt64(b))

	var c int8 = 12
	go_test_.Equal(t, int64(12), FormatInstance.MustToInt64(c))

	var d int16 = 12
	go_test_.Equal(t, int64(12), FormatInstance.MustToInt64(d))

	var f int32 = 12
	go_test_.Equal(t, int64(12), FormatInstance.MustToInt64(f))

	var g uint8 = 12
	go_test_.Equal(t, int64(12), FormatInstance.MustToInt64(g))

	var h uint16 = 12
	go_test_.Equal(t, int64(12), FormatInstance.MustToInt64(h))

	var i uint32 = 12
	go_test_.Equal(t, int64(12), FormatInstance.MustToInt64(i))

	go_test_.Equal(t, int64(1), FormatInstance.MustToInt64(true))

	m_, err := FormatInstance.ToInt64(true)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, int64(1), m_)

}

func TestFormatType_MustToUint64(t *testing.T) {
	a := `1222222`
	go_test_.Equal(t, uint64(1222222), FormatInstance.MustToUint64(a))

	var b int = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(b))

	var c int8 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(c))

	var d int16 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(d))

	var f int32 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(f))

	var g uint8 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(g))

	var h uint16 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(h))

	var i uint32 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(i))

	var j uint64 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(j))

	var k float32 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(k))

	var l float64 = 12
	go_test_.Equal(t, uint64(12), FormatInstance.MustToUint64(l))

	var m bool = true
	go_test_.Equal(t, uint64(1), FormatInstance.MustToUint64(m))

	var n string = "0xc00007a000"
	go_test_.Equal(t, uint64(824634220544), FormatInstance.MustToUint64(n))

	m_, err := FormatInstance.ToUint64(n)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, uint64(824634220544), m_)
}

func TestFormatType_ToString(t *testing.T) {
	var a *float64
	b := 0.34
	a = &b
	go_test_.Equal(t, "0.34", FormatInstance.ToString(a))

	type BType struct {
		B1 int
	}
	a1 := struct {
		A string
		B BType
	}{`1`, BType{2}}
	go_test_.Equal(t, "{\"A\":\"1\",\"B\":{\"B1\":2}}", FormatInstance.ToString(a1))
	a11 := []BType{
		{
			B1: 1,
		},
		{
			B1: 2,
		},
	}
	go_test_.Equal(t, "[{\"B1\":1},{\"B1\":2}]", FormatInstance.ToString(a11))

	type Test struct {
		Test1 *string `json:"test1"`
	}
	test_ := Test{}
	go_test_.Equal(t, "*nil", FormatInstance.ToString(test_.Test1))

	a2 := 625462456
	go_test_.Equal(t, "625462456", FormatInstance.ToString(a2))

	a3 := 0xf43f2
	go_test_.Equal(t, "1000434", FormatInstance.ToString(a3))

	a4 := map[string]interface{}{
		"go_test_": "go_test_",
	}
	go_test_.Equal(t, `{"go_test_":"go_test_"}`, FormatInstance.ToString(a4))
}

func TestFormatType_GetValuesInTagFromStruct(t *testing.T) {
	// []*Test{}
	test_ := []*Test{}
	fields := FormatInstance.GetValuesInTagFromStruct(test_, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// Test{}
	test1 := Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(test1, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *Test{}
	test2 := Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(&test2, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// []Test{}
	test3 := []Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(test3, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *[]Test{}
	test4 := []Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(&test4, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))

	// *[]*Test{}
	test5 := []*Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(&test5, `json`)
	go_test_.Equal(t, "[user_id type order_number price amount tranfer_memo status time id created_at updated_at]", fmt.Sprint(fields))
}

func TestFormatType_MustToBool(t *testing.T) {
	go_test_.Equal(t, true, FormatInstance.MustToBool(`true`))
	go_test_.Equal(t, false, FormatInstance.MustToBool(`false`))
}

func TestFormatType_MustToInt(t *testing.T) {
	a := "4546"
	go_test_.Equal(t, 4546, FormatInstance.MustToInt(a))
}

func TestFormatType_MustToInt8(t *testing.T) {
	a := "12"
	go_test_.Equal(t, int8(12), FormatInstance.MustToInt8(a))
}

func TestFormatType_MustToInt32(t *testing.T) {
	a := "4546"
	go_test_.Equal(t, int32(4546), FormatInstance.MustToInt32(a))
}

func TestFormatType_MustToUint32(t *testing.T) {
	a := "4546"
	go_test_.Equal(t, uint32(4546), FormatInstance.MustToUint32(a))
}

func TestFormatType_MustToFloat64(t *testing.T) {
	go_test_.Equal(t, 4546.3526, FormatInstance.MustToFloat64("4546.3526"))
	go_test_.Equal(t, float64(4546), FormatInstance.MustToFloat64("4546.0000"))
	go_test_.Equal(t, float64(0.0042), FormatInstance.MustToFloat64("0.0042"))
}

func TestFormatType_MustToFloat32(t *testing.T) {
	a := "4546.3526"
	go_test_.Equal(t, float32(4546.3526), FormatInstance.MustToFloat32(a))
}

func TestFormatType_GroupStringSlice(t *testing.T) {
	formatInstance := NewFormatInstance[string]()
	strSlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	group := formatInstance.GroupSlice(strSlice, 3)
	go_test_.Equal(t, 3, len(group))
	go_test_.Equal(t, 3, len(group[0]))
	go_test_.Equal(t, "a", group[0][0])
	go_test_.Equal(t, "c", group[0][2])
	go_test_.Equal(t, 3, len(group[1]))
	go_test_.Equal(t, "d", group[1][0])
	go_test_.Equal(t, "e", group[1][1])
	go_test_.Equal(t, 1, len(group[2]))
	go_test_.Equal(t, "g", group[2][0])

	group1 := formatInstance.GroupSlice(strSlice, 10)
	go_test_.Equal(t, 1, len(group1))
	go_test_.Equal(t, 7, len(group1[0]))
	go_test_.Equal(t, "a", group1[0][0])
	go_test_.Equal(t, "g", group1[0][6])
}
