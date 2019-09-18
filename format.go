package go_format

import (
	"encoding/base64"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"time"
)

type FormatClass struct {
}

var Format = FormatClass{}

func (this *FormatClass) EncodeBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (this *FormatClass) DecodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func (this *FormatClass) StructToMap(in_ interface{}) map[string]interface{} {
	struct_ := structs.New(in_)
	struct_.TagName = `json`
	return struct_.Map()
}

func (this *FormatClass) MapToStruct(map_ map[string]interface{}, dest interface{}) {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &dest,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(map_)
	if err != nil {
		panic(err)
	}
}

func (this *FormatClass) StringToTime(str string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	return t
}