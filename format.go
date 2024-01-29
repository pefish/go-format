package go_format

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
	"strings"
)

type FormatType[T any] struct {
}

var FormatInstance = NewFormatInstance[interface{}]()

func NewFormatInstance[T any]() *FormatType[T] {
	return &FormatType[T]{}
}

func (ft *FormatType[T]) EncodeBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (ft *FormatType[T]) DecodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func (ft *FormatType[T]) IsZeroValue(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return val.IsNil()
	default:
		return val.IsZero()
	}
}

func (ft *FormatType[T]) StructToMap(in_ interface{}) map[string]interface{} {
	if in_ == nil {
		return map[string]interface{}{}
	}
	struct_ := structs.New(in_)
	struct_.TagName = `json`

	return struct_.Map()
}

func (ft *FormatType[T]) MapToStruct(dest interface{}, map_ map[string]interface{}) error {
	if map_ == nil {
		return fmt.Errorf("map is nil")
	}
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &dest,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(map_)
	if err != nil {
		return err
	}
	return nil
}

func (ft *FormatType[T]) SliceToStruct(dest interface{}, slice_ []interface{}) error {
	if slice_ == nil {
		return fmt.Errorf("slice is nil")
	}
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &dest,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(slice_)
	if err != nil {
		return err
	}
	return nil
}

// GetValuesInTagFromStruct
//
//	@Description: 通过反射获取指针中struct的所有tag值. 支持 []*Test{}、*[]*Test{}、Test{}、*Test{}、[]Test{}、*[]Test{}
//	@receiver ft
//	@param interf
//	@param tag
//	@return []string
func (ft *FormatType[T]) GetValuesInTagFromStruct(interf interface{}, tag string) []string {
	result := make([]string, 0)
	return ft.getValuesInTagFromStruct(result, reflect.TypeOf(interf), tag)
}

func (ft *FormatType[T]) getValuesInTagFromStruct(result []string, type_ reflect.Type, tagName string) []string {
	realValKind := type_.Kind()
	switch realValKind {
	case reflect.Ptr, reflect.Slice:
		result = ft.getValuesInTagFromStruct(result, type_.Elem(), tagName)
	case reflect.Struct:
		if type_.String() == "time.Time" {
			return result
		}
		for i := 0; i < type_.NumField(); i++ {
			fieldType := type_.Field(i).Type
			tagValue := type_.Field(i).Tag.Get(tagName)
			if tagValue != `` && (fieldType.Kind() != reflect.Struct || fieldType.String() == "time.Time") {
				tagValues := strings.Split(tagValue, ",")
				result = append(result, tagValues[0])
			}
			result = ft.getValuesInTagFromStruct(result, fieldType, tagName)
		}
	default:
		return result
	}

	return result
}

func (ft *FormatType[T]) MustToInt(val interface{}) int {
	result, err := ft.ToInt(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) ToInt(val interface{}) (int, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int`)
	}
	valStr := ft.ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := ft.findBase(valStr)
	int_, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int(int_), nil
}

func (ft *FormatType[T]) MustToInt8(val interface{}) int8 {
	result, err := ft.ToInt8(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) ToInt8(val interface{}) (int8, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int8`)
	}
	valStr := ft.ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := ft.findBase(valStr)
	int_, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int8(int_), nil
}

func (ft *FormatType[T]) MustToBool(val interface{}) bool {
	result, err := ft.ToBool(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) ToBool(val interface{}) (bool, error) {
	if val == nil {
		return false, errors.New(`nil cannot convert to bool`)
	}
	valStr := ft.ToString(val)
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

func (ft *FormatType[T]) MustToInt32(val interface{}) int32 {
	result, err := ft.ToInt32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) findBase(str string) (string, int) {
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

func (ft *FormatType[T]) ToInt32(val interface{}) (int32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int32`)
	}
	valStr := ft.ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := ft.findBase(valStr)
	int_, err := strconv.ParseInt(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int32(int_), nil
}

func (ft *FormatType[T]) MustToInt64(val interface{}) int64 {
	result, err := ft.ToInt64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) ToInt64(val interface{}) (int64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to int64`)
	}
	valStr := ft.ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := ft.findBase(valStr)
	int_, err := strconv.ParseInt(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int_, nil
}

func (ft *FormatType[T]) MustToUint64(val interface{}) uint64 {
	result, err := ft.ToUint64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) ToUint64(val interface{}) (uint64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to uint64`)
	}
	valStr := ft.ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := ft.findBase(valStr)
	int_, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, err
	}
	return int_, nil
}

func (ft *FormatType[T]) MustToUint32(val interface{}) uint32 {
	result, err := ft.ToUint32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) ToUint32(val interface{}) (uint32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to uint32`)
	}
	valStr := ft.ToString(val)
	if valStr == "true" {
		return 1, nil
	}
	if valStr == "false" {
		return 0, nil
	}
	str, base := ft.findBase(valStr)
	int_, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, err
	}
	return uint32(int_), nil
}

func (ft *FormatType[T]) MustToFloat64(val interface{}) float64 {
	result, err := ft.ToFloat64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) ToFloat64(val interface{}) (float64, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to float64`)
	}
	valStr := ft.ToString(val)
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

func (ft *FormatType[T]) MustToFloat32(val interface{}) float32 {
	result, err := ft.ToFloat32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func (ft *FormatType[T]) ToFloat32(val interface{}) (float32, error) {
	if val == nil {
		return 0, errors.New(`nil cannot convert to float32`)
	}

	valStr := ft.ToString(val)
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

func (ft *FormatType[T]) ToString(val interface{}) string {
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
	case reflect.Ptr:
		if value_.IsNil() { // IsNil 只接受 chan, func, interface, map, pointer, or slice value
			return `*nil`
		}
		return ft.ToString(value_.Elem().Interface())
	default:
		return fmt.Sprint(val)
	}
}

func (ft *FormatType[T]) GroupSlice(slice []T, countPerGroup uint64) [][]T {
	resultGroup := make([][]T, 0)
	start, end := 0, 0
	for {
		start = end
		end += int(countPerGroup)
		if end > len(slice) {
			end = len(slice)
		}

		thisGroup := slice[start:end]
		resultGroup = append(resultGroup, thisGroup)
		if end >= len(slice) {
			break
		}
	}
	return resultGroup
}
