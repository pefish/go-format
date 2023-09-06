package go_format

import (
	"fmt"
	"github.com/pefish/go-test-assert"
	"testing"
)

type Test struct {
	UserId       uint64  `json:"user_id"`
	Type         uint64  `json:"type"`
	OrderNumber  string  `json:"order_number"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	TransferMemo string  `json:"tranfer_memo"`
	Status       uint64  `json:"status"`

	BaseModel
}

type BaseModel struct {
	Id        uint64 `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func TestFormatClass_StructToMap(t *testing.T) {
	type Test struct {
		A uint64 `json:"a"`
		B string `json:"haha"`
	}
	testObj := Test{
		A: 100,
		B: `1111`,
	}
	testMap := FormatInstance.StructToMap(testObj)
	test.Equal(t, true, testMap["a"].(uint64) == 100)
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
	test.Equal(t, nil, err)
	test.Equal(t, true, testObj.A == 100)
	test.Equal(t, "1111", testObj.B)
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
	test.Equal(t, nil, err)
	test.Equal(t, true, testObj[0].A == 100)
	test.Equal(t, "1111", testObj[0].B)
}

func TestFormatType_MustToInt64(t *testing.T) {
	a := `1222222`
	test.Equal(t, int64(1222222), FormatInstance.MustToInt64(a))

	a1 := `0x16`
	test.Equal(t, int64(22), FormatInstance.MustToInt64(a1))
	a2 := `0o17`
	test.Equal(t, int64(15), FormatInstance.MustToInt64(a2))
	a3 := `0b101`
	test.Equal(t, int64(5), FormatInstance.MustToInt64(a3))

	var b int = 12
	test.Equal(t, int64(12), FormatInstance.MustToInt64(b))

	var c int8 = 12
	test.Equal(t, int64(12), FormatInstance.MustToInt64(c))

	var d int16 = 12
	test.Equal(t, int64(12), FormatInstance.MustToInt64(d))

	var f int32 = 12
	test.Equal(t, int64(12), FormatInstance.MustToInt64(f))

	var g uint8 = 12
	test.Equal(t, int64(12), FormatInstance.MustToInt64(g))

	var h uint16 = 12
	test.Equal(t, int64(12), FormatInstance.MustToInt64(h))

	var i uint32 = 12
	test.Equal(t, int64(12), FormatInstance.MustToInt64(i))

	var m bool = true
	test.Equal(t, int64(1), FormatInstance.MustToInt64(m))

	m_, err := FormatInstance.ToInt64(m)
	test.Equal(t, nil, err)
	test.Equal(t, int64(1), m_)

}

func TestFormatType_MustToUint64(t *testing.T) {
	a := `1222222`
	test.Equal(t, uint64(1222222), FormatInstance.MustToUint64(a))

	var b int = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(b))

	var c int8 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(c))

	var d int16 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(d))

	var f int32 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(f))

	var g uint8 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(g))

	var h uint16 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(h))

	var i uint32 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(i))

	var j uint64 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(j))

	var k float32 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(k))

	var l float64 = 12
	test.Equal(t, uint64(12), FormatInstance.MustToUint64(l))

	var m bool = true
	test.Equal(t, uint64(1), FormatInstance.MustToUint64(m))

	var n string = "0xc00007a000"
	test.Equal(t, uint64(824634220544), FormatInstance.MustToUint64(n))

	m_, err := FormatInstance.ToUint64(n)
	test.Equal(t, nil, err)
	test.Equal(t, uint64(824634220544), m_)
}

func TestFormatType_ToString(t *testing.T) {
	var a *float64
	b := 0.34
	a = &b
	test.Equal(t, "0.34", FormatInstance.ToString(a))

	type BType struct {
		B1 int
	}
	a1 := struct {
		A string
		B BType
	}{`1`, BType{2}}
	test.Equal(t, "{1 {2}}", FormatInstance.ToString(a1))

	type Test struct {
		Test1 *string `json:"test1"`
	}
	test_ := Test{}
	test.Equal(t, "*nil", FormatInstance.ToString(test_.Test1))

	a2 := 625462456
	test.Equal(t, "625462456", FormatInstance.ToString(a2))

	a3 := 0xf43f2
	test.Equal(t, "1000434", FormatInstance.ToString(a3))
}

func TestFormatType_GetValuesInTagFromStruct(t *testing.T) {
	// []*Test{}
	test_ := []*Test{}
	fields := FormatInstance.GetValuesInTagFromStruct(test_, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))

	// Test{}
	test1 := Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(test1, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))

	// *Test{}
	test2 := Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(&test2, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))

	// []Test{}
	test3 := []Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(test3, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))

	// *[]Test{}
	test4 := []Test{}
	fields = FormatInstance.GetValuesInTagFromStruct(&test4, `json`)
	test.Equal(t, "[user_id type order_number price amount tranfer_memo status id created_at updated_at]", fmt.Sprint(fields))
}

func TestFormatType_MustToBool(t *testing.T) {
	a := `true`
	test.Equal(t, true, FormatInstance.MustToBool(a))
}

func TestFormatType_MustToInt(t *testing.T) {
	a := "4546"
	test.Equal(t, 4546, FormatInstance.MustToInt(a))
}

func TestFormatType_MustToInt8(t *testing.T) {
	a := "12"
	test.Equal(t, int8(12), FormatInstance.MustToInt8(a))
}

func TestFormatType_MustToInt32(t *testing.T) {
	a := "4546"
	test.Equal(t, int32(4546), FormatInstance.MustToInt32(a))
}

func TestFormatType_MustToUint32(t *testing.T) {
	a := "4546"
	test.Equal(t, uint32(4546), FormatInstance.MustToUint32(a))
}

func TestFormatType_MustToFloat64(t *testing.T) {
	a := "4546.3526"
	test.Equal(t, 4546.3526, FormatInstance.MustToFloat64(a))
}

func TestFormatType_MustToFloat32(t *testing.T) {
	a := "4546.3526"
	test.Equal(t, float32(4546.3526), FormatInstance.MustToFloat32(a))
}

func TestFormatType_GroupStringSlice(t *testing.T) {
	strSlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	group := FormatInstance.GroupStringSlice(strSlice, 3)
	test.Equal(t, 3, len(group))
	test.Equal(t, 3, len(group[0]))
	test.Equal(t, "a", group[0][0])
	test.Equal(t, "c", group[0][2])
	test.Equal(t, 3, len(group[1]))
	test.Equal(t, "d", group[1][0])
	test.Equal(t, "e", group[1][1])
	test.Equal(t, 1, len(group[2]))
	test.Equal(t, "g", group[2][0])
}
