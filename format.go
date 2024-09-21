package go_format

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// 下划线单词转为驼峰单词
func UnderscoreToCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	caser := cases.Title(language.BrazilianPortuguese)
	s = caser.String(s)
	return strings.Replace(s, " ", "", -1)
}

// 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}

			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}

func EncodeBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func DecodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func IsZeroValue(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return val.IsNil()
	default:
		return val.IsZero()
	}
}

func StructToMap(in_ interface{}) map[string]interface{} {
	if in_ == nil {
		return map[string]interface{}{}
	}
	struct_ := structs.New(in_)
	struct_.TagName = `json`

	return struct_.Map()
}

func MapToStruct(dest interface{}, map_ map[string]interface{}) error {
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

func SliceToStruct(dest interface{}, slice_ []interface{}) error {
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
func GetValuesInTagFromStruct(interf interface{}, tag string) []string {
	result := make([]string, 0)
	return getValuesInTagFromStruct(result, reflect.TypeOf(interf), tag)
}

func getValuesInTagFromStruct(result []string, type_ reflect.Type, tagName string) []string {
	realValKind := type_.Kind()
	switch realValKind {
	case reflect.Ptr, reflect.Slice:
		result = getValuesInTagFromStruct(result, type_.Elem(), tagName)
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
			result = getValuesInTagFromStruct(result, fieldType, tagName)
		}
	default:
		return result
	}

	return result
}

func MustToInt(val interface{}) int {
	result, err := ToInt(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToInt(val interface{}) (int, error) {
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

func MustToInt8(val interface{}) int8 {
	result, err := ToInt8(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToInt8(val interface{}) (int8, error) {
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

func MustToBool(val interface{}) bool {
	result, err := ToBool(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToBool(val interface{}) (bool, error) {
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

func MustToInt32(val interface{}) int32 {
	result, err := ToInt32(val)
	if err != nil {
		panic(err)
	}
	return result
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

func ToInt32(val interface{}) (int32, error) {
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

func MustToInt64(val interface{}) int64 {
	result, err := ToInt64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToInt64(val interface{}) (int64, error) {
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

func MustToUint64(val interface{}) uint64 {
	result, err := ToUint64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToUint64(val interface{}) (uint64, error) {
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

func MustToUint32(val interface{}) uint32 {
	result, err := ToUint32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToUint32(val interface{}) (uint32, error) {
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

func MustToFloat64(val interface{}) float64 {
	result, err := ToFloat64(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToFloat64(val interface{}) (float64, error) {
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

func MustToFloat32(val interface{}) float32 {
	result, err := ToFloat32(val)
	if err != nil {
		panic(err)
	}
	return result
}

func ToFloat32(val interface{}) (float32, error) {
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

func ToString(val interface{}) string {
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

func GroupSlice[T any](slice []T, countPerGroup int) [][]T {
	resultGroup := make([][]T, 0)

	intValues := GroupInt(len(slice), countPerGroup)

	for i, intValue := range intValues {
		start := 0
		if i > 0 {
			start = i * intValues[i-1]
		}
		resultGroup = append(resultGroup, slice[start:start+intValue])
	}
	return resultGroup
}

// 对数值进行分组。例如 35 使用 10 分组结果是 [10,10,10,5]
func GroupInt[T int | uint | int64 | uint64](number T, sliceBy T) []T {
	results := make([]T, 0)
	var start, end T = 0, 0
	for {
		start = end
		end += sliceBy
		if end > number {
			end = number
		}
		results = append(results, end-start)
		if end >= number {
			break
		}
	}
	return results
}
