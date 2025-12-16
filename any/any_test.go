package any

import (
	"testing"

	go_test_ "github.com/pefish/go-test"
)

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

	a4 := map[string]any{
		"go_test_": "go_test_",
	}
	go_test_.Equal(t, `{"go_test_":"go_test_"}`, ToString(a4))

	go_test_.Equal(t, `["1","2"]`, ToString([]string{"1", "2"}))

	go_test_.Equal(t, `Aa`, ToString([]byte{65, 97}))
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

func TestFormatType_ToStruct(t *testing.T) {
	type A struct {
		A1 string `json:"a1"`
		A2 int    `json:"a2"`
	}
	type B struct {
		A3 string `json:"a1"`
		A4 int    `json:"a2"`
	}
	var a any = A{
		A1: "test",
		A2: 123,
	}
	var b B
	err := ToStruct(a, &b)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "test", b.A3)
	go_test_.Equal(t, 123, b.A4)

	aMap := map[string]any{
		`a1`: "map_test",
		`a2`: 456,
	}
	var b1 B
	err = ToStruct(aMap, &b1)
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "map_test", b1.A3)
	go_test_.Equal(t, 456, b1.A4)

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
