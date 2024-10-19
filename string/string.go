package go_format_string

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func Desensitize(str string) string {
	index := strings.Index(str, `@`)
	if index == -1 {
		return DesensitizeMobile(str)
	} else {
		return MustDesensitizeEmail(str)
	}
}

/*
>7        前3后4中间4个*
<=7 && >4 前2后2中间4个*
<=4 && >2 前1后1中间2个*
*/
func DesensitizeMobile(str string) string {
	result := ``
	length := len(str)
	if length > 7 {
		result = str[:3] + `****` + str[length-4:]
	} else if length <= 7 && length > 4 {
		result = str[:2] + `****` + str[length-2:]
	} else if length <= 4 && length > 2 {
		result = str[:1] + `**` + str[length-1:]
	} else {
		result = "*"
	}
	return result
}

func MustDesensitizeEmail(str string) string {
	result, err := DesensitizeEmail(str)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@前字符串长度>3   前4 中4个* 后@后面所有
@前字符串长度<=3  前@前面所有 中4个* 后@后面所有
*/
func DesensitizeEmail(str string) (string, error) {
	result := ``
	index := strings.Index(str, `@`)
	if index == -1 {
		return "", errors.New(`Not email.`)
	}
	preAt := str[:index]
	if len(preAt) > 3 {
		result = str[:4] + `****` + str[index:]
	} else {
		result = preAt + `****` + str[index:]
	}
	return result, nil
}

func RemoveLast(str string, num int) string {
	return str[:len(str)-num]
}

func RemoveFirst(str string, num int) string {
	return str[num:]
}

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func ReplaceAll(str string, oldStr string, newStr string) (result string) {
	return strings.Replace(str, oldStr, newStr, -1)
}

func MustSpanLeft(str string, length int, fillChar string) string {
	result, err := SpanLeft(str, length, fillChar)
	if err != nil {
		panic(err)
	}
	return result
}

func SpanLeft(str string, length int, fillChar string) (string, error) {
	if len(str) > length {
		return "", errors.New(`Length is too small.`)
	}
	if len(fillChar) != 1 {
		return "", errors.New(`Length of fillChar must be 1.`)
	}
	result := ``
	for i := 0; i < length-len(str); i++ {
		result += fillChar
	}
	return result + str, nil
}

func MustSpanRight(str string, length int, fillChar string) string {
	result, err := SpanRight(str, length, fillChar)
	if err != nil {
		panic(err)
	}
	return result
}

func SpanRight(str string, length int, fillChar string) (string, error) {
	if len(str) > length {
		return "", errors.New(`Length is too small.`)
	}
	if len(fillChar) != 1 {
		return "", errors.New(`Length of fillChar must be 1.`)
	}
	result := str
	for i := 0; i < length-len(str); i++ {
		result += fillChar
	}
	return result, nil
}

func StartWith(str string, substr string) bool {
	return strings.HasPrefix(str, substr)
}

func Indexes(str string, substr string) []int {
	results := make([]int, 0)
	for {
		i := strings.Index(str, substr)
		if i == -1 {
			break
		}
		index := i
		if len(results) > 0 {
			index += results[len(results)-1] + len(substr)
		}
		results = append(results, index)
		str = str[i+len(substr):]
	}
	return results
}

// 倒着找，substrs 中谁先找到就返回谁的 index
func LastIndex(str string, substrs []string) int {
	result := -1
	for _, substr := range substrs {
		index := strings.LastIndex(str, substr)
		if index != -1 && index > result {
			result = index
		}
	}
	return result
}

func Index(str string, substrs []string) int {
	result := -1
	for _, substr := range substrs {
		index := strings.Index(str, substr)
		if index != -1 && index < result {
			result = index
		}
	}
	return result
}

func BetweenAnd(str string, startStr string, endStr string) []string {
	arr := strings.Split(str, startStr)
	if len(arr) <= 1 {
		return nil
	}
	results := make([]string, 0)
	for i := 1; i < len(arr); i++ {
		index := strings.Index(arr[i], endStr)
		if index == -1 {
			continue
		}
		results = append(results, arr[i][:index])
	}
	return results
}

func EndWith(str string, substr string) bool {
	return strings.HasSuffix(str, substr)
}

func UserIdToInviteCode(userId uint64, length int) (string, error) {
	userIdStr := strconv.FormatUint(userId, 10)
	if len(userIdStr) > length {
		return "", errors.New("Length is too small.")
	}
	result := ""
	for _, a := range userIdStr {
		result += string(a + 20)
	}
	if len(userIdStr) != length {
		r, err := randomStringFromDic("ABCDEFGHIJKLMNOPQRSTUVWXYZ", int32(length-len(userIdStr)))
		if err != nil {
			return "", err
		}
		result += r
	}
	return result, nil
}

func randomStringFromDic(dictionary string, count int32) (string, error) {
	b := make([]byte, count)
	l := len(dictionary)

	_, err := rand.New(rand.NewSource(time.Now().UnixNano())).Read(b)

	if err != nil {
		return "", err
	}
	for i, v := range b {
		b[i] = dictionary[v%byte(l)]
	}

	return string(b), nil
}

func Insert(source string, target string, index int) (string, error) {
	if index > len(target) {
		return "", errors.Errorf("Index too large, target string length is <%d>", len(target))
	}
	return target[:index] + source + target[index:], nil
}

func MustInsert(source string, target string, index int) string {
	result, err := Insert(source, target, index)
	if err != nil {
		panic(err)
	}
	return result
}
