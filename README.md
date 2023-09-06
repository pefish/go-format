# go-format

### Description

类型或格式转换工具

### Start

```golang
var a *float64
b := 0.34
a = &b
test.Equal(t, "0.34", Reflect.ToString(a))


type BType struct {
    B1 int
}
a1 := struct {
    A string
    B BType
}{`1`, BType{2}}
test.Equal(t, "{1 {2}}", Reflect.ToString(a1))



type Test struct {
    Test1 *string `json:"test1"`
}
test_ := Test{}
test.Equal(t, "*nil", Reflect.ToString(test_.Test1))



a2 := 625462456
test.Equal(t, "625462456", Reflect.ToString(a2))



a3 := 0xf43f2
test.Equal(t, "1000434", Reflect.ToString(a3))
```
