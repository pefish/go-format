package p_format

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"gitee.com/pefish/p-go-error"
	"gitee.com/pefish/p-go-reflect"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"time"
)

type FormatClass struct {
}

var Format = FormatClass{}

func (this *FormatClass) Int64ToString(int_ int64) string {
	return strconv.FormatInt(int_, 10)
}

func (this *FormatClass) StringToInt64(str string) int64 {
	int_, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return int_
}

func (this *FormatClass) CharToInt64(char uint8) int64 {
	int_, err := strconv.ParseInt(string(char), 10, 64)
	if err != nil {
		panic(err)
	}
	return int_
}

func (this *FormatClass) StringToUint64(str string) uint64 {
	uint_, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint_
}

func (this *FormatClass) Uint64ToString(uint_ uint64) string {
	return strconv.FormatUint(uint_, 10)
}

func (this *FormatClass) Float64ToString(f_ float64) string {
	return strconv.FormatFloat(f_, 'f', -1, 64)
}

func (this *FormatClass) BoolToString(bool_ bool) string {
	return strconv.FormatBool(bool_)
}

func (this *FormatClass) StringToBool(str string) bool {
	bool_, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return bool_
}

func (this *FormatClass) StringToFloat64(str string) float64 {
	float_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return float_
}

func (this *FormatClass) EncodeBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (this *FormatClass) DecodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

/**
struct中export的字段才能被转换出来, 通过`json:"abc"`可以控制出来的字段名
*/
func (this *FormatClass) StructToMap(in_ interface{}) map[string]interface{} {
	var result map[string]interface{}
	inrec, err := json.Marshal(in_)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(inrec, &result)
	return result
}

func (this *FormatClass) StructToSlice(in_ interface{}) []interface{} {
	var result []interface{}
	inrec, err := json.Marshal(in_)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(inrec, &result)
	return result
}

func (this *FormatClass) StructToMapString(in_ interface{}) (out map[string]string) {
	var result map[string]interface{}
	inrec, err := json.Marshal(in_)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(inrec, &result)

	out = map[string]string{}
	for key, val := range result {
		out[key] = p_reflect.Reflect.ToString(val)
	}
	return
}

func (this *FormatClass) MapStringToMapInterface(mapInterface map[string]interface{}, mapString map[string]string) {
	for key, val := range mapString {
		mapInterface[key] = val
	}
}

func (this *FormatClass) MapStringToStruct(struct_ interface{}, map_ map[string]string) {
	inrec, err := json.Marshal(map_)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(inrec, struct_)
}

func (this *FormatClass) MapStringToStructByKey(struct_ interface{}, map_ map[string]string) {
	t := reflect.ValueOf(struct_).Elem()
	for k, v := range map_ {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}

func (this *FormatClass) MapToStruct(struct_ interface{}, map_ interface{}) {
	if reflect.TypeOf(map_).Kind() != reflect.Map {
		p_error.ThrowInternal(`map_ not a map`)
	}
	if reflect.TypeOf(struct_).Kind() != reflect.Ptr {
		p_error.ThrowInternal(`struct_ not a ptr`)
	}
	inrec, err := json.Marshal(map_)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(inrec, struct_)
}

func (this *FormatClass) SliceToStruct(struct_ interface{}, slice_ interface{}) {
	if reflect.TypeOf(slice_).Kind() != reflect.Slice {
		p_error.ThrowInternal(`map_ not a slice`)
	}
	if reflect.TypeOf(struct_).Kind() != reflect.Ptr {
		p_error.ThrowInternal(`struct_ not a ptr`)
	}
	inrec, err := json.Marshal(slice_)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(inrec, struct_)
}

func (this *FormatClass) StringToTime(str string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	return t
}

func (this *FormatClass) NullfloatToFloat(nf sql.NullFloat64) float64 {
	if !nf.Valid {
		panic(errors.New(`Nullfloat64 is not valid`))
	}
	return nf.Float64
}

func (this *FormatClass) FloatToNullfloat(f float64) sql.NullFloat64 {
	return sql.NullFloat64{
		f,
		true,
	}
}

func (this *FormatClass) MapToSortedQueryString(map_ map[string]string) string {
	var buf bytes.Buffer
	keys := make([]string, 0, len(map_))
	for k := range map_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := map_[k]
		prefix := url.QueryEscape(k) + "="
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(prefix)
		buf.WriteString(url.QueryEscape(vs))
	}
	return buf.String()
}

func (this *FormatClass) MapInterfaceToMapString(map_ map[string]interface{}) (out map[string]string) {
	out = map[string]string{}
	for key, val := range map_ {
		out[key] = p_reflect.Reflect.ToString(val)
	}
	return
}

func (this *FormatClass) SliceInterfaceToSliceString(slice_ []interface{}) (out []string) {
	out = []string{}
	for _, val := range slice_ {
		out = append(out, p_reflect.Reflect.ToString(val))
	}
	return
}

func (this *FormatClass) SliceInterfaceToSliceMapInterface(slice_ []interface{}) (out []map[string]interface{}) {
	out = []map[string]interface{}{}
	for _, val := range slice_ {
		out = append(out, val.(map[string]interface{}))
	}
	return
}

func (this *FormatClass) MapGetString(map_ map[string]interface{}, key string) (exist bool, out string) {
	exist = false
	if map_ == nil || len(map_) == 0 {
		return
	}
	if v, ok := map_[key]; ok {
		out = v.(string)
		exist = true
	}
	return
}

func (this *FormatClass) MapGetInt(map_ map[string]interface{}, key string) (exist bool, out int) {
	exist = false
	if map_ == nil || len(map_) == 0 {
		return
	}
	if v, ok := map_[key]; ok {
		out = int(v.(float64))
		exist = true
	}
	return
}

func (this *FormatClass) MapGetMap(map_ map[string]interface{}, key string) (exist bool, out map[string]interface{}) {
	exist = false
	if map_ == nil || len(map_) == 0 {
		return
	}
	if v, ok := map_[key]; ok {
		out = v.(map[string]interface{})
		exist = true
	}
	return
}
