package go_format

import (
	"crypto/rc4"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
	"unicode"

	"github.com/google/uuid"
	go_format_string "github.com/pefish/go-format/string"
	go_format_type "github.com/pefish/go-format/type"
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

func CamelCaseToWords(s string) []string {
	results := make([]string, 0)
	startIndex := 0
	for i, r := range s {
		if i == 0 {
			continue
		}
		if unicode.IsUpper(r) {
			results = append(results, s[startIndex:i])
			startIndex = i
		}
		if i == len(s)-1 {
			results = append(results, s[startIndex:])
		}
	}
	return results
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

// FetchTags
//
//	@Description: 通过反射获取所有 tag 值，只搜寻一层。 支持 []*Test{}、*[]*Test{}、Test{}、*Test{}、[]Test{}、*[]Test{}
//	@param interf
//	@param tag
//	@return []string
func FetchTags(interf any, tagName string) []string {
	type_ := reflect.TypeOf(interf)
	// 剥离 slice 和 指针
strip:
	for {
		switch type_.Kind() {
		case reflect.Ptr, reflect.Slice:
			type_ = type_.Elem()
		default:
			break strip
		}
	}

	if type_.Kind() != reflect.Struct {
		return make([]string, 0)
	}

	results := make([]string, 0)

	// struct
	for i := 0; i < type_.NumField(); i++ {
		fieldType := type_.Field(i).Type
		tagValue := type_.Field(i).Tag.Get(tagName)
		if tagValue == "" {
			if fieldType.Kind() != reflect.Struct {
				continue
			}
			// 元素是 struct，再搜寻一层
			for i := 0; i < fieldType.NumField(); i++ {
				tagValue := fieldType.Field(i).Tag.Get(tagName)
				if tagValue == "" {
					continue
				}
				tagValues := strings.Split(tagValue, ",")
				results = append(results, tagValues[0])
			}
			continue
		}

		tagValues := strings.Split(tagValue, ",")
		results = append(results, tagValues[0])
	}

	return results
}

// 对数据进行编码，每次编码结果不一样，但是都可以自解码出原来的明文
func EncodePefish(data string) string {
	pass := uuid.New().String()

	c, _ := rc4.NewCipher([]byte(pass))
	src := []byte(data)
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	rc4Result := hex.EncodeToString(dst)

	passGroups := go_format_string.Group(pass, &go_format_type.GroupOpts{
		GroupCount: 3,
	})
	for i, passGroup := range passGroups {
		passGroups[i] = fmt.Sprintf("_%s=", passGroup)
	}

	insertIndexs := randomCountInt(0, len(rc4Result), 3)
	sort.Ints(insertIndexs)

	rc4Result = go_format_string.MustInsert(passGroups[0], rc4Result, insertIndexs[0])
	rc4Result = go_format_string.MustInsert(passGroups[1], rc4Result, insertIndexs[1]+len(passGroups[0]))
	rc4Result = go_format_string.MustInsert(passGroups[2], rc4Result, insertIndexs[2]+len(passGroups[0])+len(passGroups[1]))

	return EncodeBase64(rc4Result)
}

func DecodePefish(data string) (string, error) {
	b, err := DecodeBase64(data)
	if err != nil {
		return "", err
	}
	d := string(b)

	passGroup := go_format_string.BetweenAnd(d, "_", "=")
	pass := strings.Join(passGroup, "")

	for _, passEle := range passGroup {
		d = strings.ReplaceAll(d, fmt.Sprintf("_%s=", passEle), "")
	}

	rc4Data := d

	c, err := rc4.NewCipher([]byte(pass))
	if err != nil {
		return "", err
	}
	inputBytes, err := hex.DecodeString(rc4Data)
	if err != nil {
		return "", err
	}
	dst := make([]byte, len(inputBytes))
	c.XORKeyStream(dst, inputBytes)

	return string(dst), nil
}

func randomCountInt(start int, end int, count int) []int {
	map_ := make(map[int]int, 0)
	for i := start; i < end; i++ {
		map_[i] = i
	}

	results := make([]int, 0)
	for _, v := range map_ {
		if len(results) == count {
			break
		}
		results = append(results, v)
	}

	return results
}

func SyncMapToMap(m *sync.Map) map[string]any {
	result := make(map[string]any, 0)
	m.Range(func(key, value any) bool {
		result[key.(string)] = value
		return true
	})

	return result
}
