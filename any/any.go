package any

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
)

func ToInt32(val any) (int32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int32`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := findBase(valStr)
	int_, err := strconv.ParseInt(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int32(int_), nil
}

func MustToInt64(val any) int64 {
	result, err := ToInt64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToInt64(val any) (int64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int64`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := findBase(valStr)
	int_, err := strconv.ParseInt(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int_, nil
}

func MustToUint64(val any) uint64 {
	result, err := ToUint64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToUint64(val any) (uint64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to uint64`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := findBase(valStr)
	int_, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int_, nil
}

func MustToBigInt(val any) *big.Int {
	result, err := ToBigInt(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToBigInt(val any) (*big.Int, error) {
	if val == nil {
		return nil, errors.New(`nil cannot convert to *big.Int`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return big.NewInt(1), nil
	}
	if valStr == "false" {
		return big.NewInt(0), nil
	}
	bigInt, ok := new(big.Int).SetString(valStr, 10)
	if !ok {
		return nil, fmt.Errorf("cannot convert %v to *big.Int", val)
	}
	return bigInt, nil
}

func MustToUint32(val any) uint32 {
	result, err := ToUint32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToUint32(val any) (uint32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to uint32`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := findBase(valStr)
	int_, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, err
	}
	return uint32(int_), nil
}

func MustToFloat64(val any) float64 {
	result, err := ToFloat64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToFloat64(val any) (float64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to float64`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func MustToFloat32(val any) float32 {
	result, err := ToFloat32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToFloat32(val any) (float32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to float32`)
	}

	valStr := ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	result, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return 0, err
	}
	return float32(result), nil
}

func ToString(val any) string {
	value_ := reflect.ValueOf(val)
	switch value_.Kind() {
	case reflect.String:
		return value_.String()
	case reflect.Bool:
		return strconv.FormatBool(value_.Bool())
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value_.Float(), 'f', -1, 64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value_.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value_.Uint(), 10)
	case reflect.Map, reflect.Array, reflect.Struct, reflect.Slice:
		if a, ok := val.([]byte); ok {
			return string(a)
		}
		b, _ := json.Marshal(value_.Interface())
		return string(b)
	case reflect.Ptr:
		if value_.IsNil() { // IsNil 只接受 chan, func, interface, map, pointer, or slice value
			return `*nil`
		}
		return ToString(value_.Elem().Interface())
	default:
		return fmt.Sprint(val)
	}
}

func findBase(str string) (string, int) {
	base := 10
	if strings.HasPrefix(str, "0x") {
		base = 16
		str = str[2:]
	} else if strings.HasPrefix(str, "0o") {
		base = 8
		str = str[2:]
	} else if strings.HasPrefix(str, "0b") {
		base = 2
		str = str[2:]
	}
	return str, base
}

func MustToInt(val any) int {
	result, err := ToInt(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToInt(val any) (int, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := findBase(valStr)
	int_, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int(int_), nil
}

func MustToInt8(val any) int8 {
	result, err := ToInt8(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToInt8(val any) (int8, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int8`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := findBase(valStr)
	int_, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int8(int_), nil
}

func MustToBool(val any) bool {
	result, err := ToBool(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToBool(val any) (bool, error) {
	if val == nil {
		return false, errors.New(`nil cannot convert to bool`)
	}
	valStr := ToString(val)
	if valStr == "true" {
		return true, nil
	}
	if valStr == "false" {
		return false, nil
	}
	bool_, err := strconv.ParseBool(valStr)
	if err != nil {
		return false, err
	}
	return bool_, nil
}

func MustToInt32(val any) int32 {
	result, err := ToInt32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func IsStruct(v any) bool {
	if v == nil {
		return false
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() == reflect.Struct
}

// 通过 json tag 将一个结构体或者 map 转换为另一个结构体
func ToStruct(from, to any) error {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &to,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(from)
	if err != nil {
		return err
	}
	return nil
}

func StructToMap(in_ any) map[string]any {
	if in_ == nil {
		return map[string]any{}
	}
	struct_ := structs.New(in_)
	struct_.TagName = `json`

	return struct_.Map()
}
