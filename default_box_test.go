package default_box

import (
	"fmt"
	"testing"
)

type User struct {
	Name    string             `default:"Bob"`
	Age     int8               `default:"10"`
	Hobbies []string           `default:"[Football, Basketball]"`
	Scores  map[string]float32 `default:"{Language: 95.55, Math: 99.50}"`
}

func TestPackDefaultBox(t *testing.T) {
	// basic use
	u := User{}
	New(&u).Fill()
	if u.Name != "Bob" {
		t.Fatal()
	}

	// chain style
	user := New(&User{}).Fill().ObjectPointer.(*User)
	if user.Name != "Bob" {
		t.Fatal()
	}
}

func TestDefaultBox_Tag(t *testing.T) {
	u := User{}
	var tag string
	var ok bool
	tag, ok = New(&u).Tag("Name")
	if ok != true || tag != "Bob" {
		t.Fatal()
	}
	tag, ok = New(&u).Tag("Scores")
	if ok != true || tag != "{Language: 95.55, Math: 99.50}" {
		t.Fatal()
	}
}

func TestDefaultBox_Fill_BasicType(t *testing.T) {
	type BasicTypeStruct struct {
		S          string  `default:"string"`
		Int8       int8    `default:"-128"`
		UInt8      uint8   `default:"255"`
		Int16      int16   `default:"-32768"`
		UInt16     uint16  `default:"65535"`
		Int32      int32   `default:"-2147483648"`
		UInt32     uint32  `default:"4294967295"`
		Int64      int64   `default:"-9223372036854775808"`
		UInt64     uint64  `default:"18446744073709551615"`
		BoolTrue1  bool    `default:"true"`
		BoolTrue2  bool    `default:"t"`
		BoolTrue3  bool    `default:"1"`
		BoolFalse1 bool    `default:"false"`
		BoolFalse2 bool    `default:"f"`
		BoolFalse3 bool    `default:"0"`
		Float32    float32 `default:"3.14159"`
		Float64    float64 `default:"3.14159265358979"`
	}
	u := BasicTypeStruct{}
	New(&u).Fill()

	if u.S != "string" {
		t.Fatal()
	}
	if u.Int8 != int8(-128) {
		t.Fatal()
	}
	if u.UInt8 != uint8(255) {
		t.Fatal()
	}
	if u.Int16 != int16(-32768) {
		t.Fatal()
	}
	if u.UInt16 != uint16(65535) {
		t.Fatal()
	}
	if u.Int32 != int32(-2147483648) {
		t.Fatal()
	}
	if u.UInt32 != uint32(4294967295) {
		t.Fatal()
	}
	if u.Int64 != int64(-9223372036854775808) {
		t.Fatal()
	}
	if u.UInt64 != uint64(18446744073709551615) {
		t.Fatal()
	}
	if u.BoolTrue1 != true {
		t.Fatal()
	}
	if u.BoolTrue2 != true {
		t.Fatal()
	}
	if u.BoolTrue3 != true {
		t.Fatal()
	}
	if u.BoolFalse1 != false {
		t.Fatal()
	}
	if u.BoolFalse2 != false {
		t.Fatal()
	}
	if u.BoolFalse3 != false {
		t.Fatal()
	}
	if u.Float32 != float32(3.14159) {
		t.Fatal()
	}
	if u.Float64 != float64(3.14159265358979) {
		t.Fatal()
	}
}

func TestDefaultBox_Fill_Slice(t *testing.T) {
	type SliceStruct struct {
		StringSlice []string  `default:"[str1, str2, str3]"`
		IntSlice    []int     `default:"[1, 1, 2, 3, 5, 8, 13]"`
		FloatSlice  []float32 `default:"[3.14159, 2.718]"`
		BoolSlice   []bool    `default:"[true, t, 1, 0]"`
	}
	u := SliceStruct{}
	New(&u).Fill()

	if len(u.StringSlice) != 3 {
		t.Fatal()
	}
	if u.StringSlice[0] != "str1" || u.StringSlice[1] != "str2" || u.StringSlice[2] != "str3" {
		t.Fatal()
	}
	if len(u.IntSlice) != 7 {
		t.Fatal()
	}
	if u.IntSlice[5] != 8 {
		t.Fatal()
	}
	if len(u.FloatSlice) != 2 {
		t.Fatal()
	}
	if u.FloatSlice[0] != float32(3.14159) {
		t.Fatal()
	}
	if len(u.BoolSlice) != 4 {
		fmt.Println(u.BoolSlice)
		t.Fatal()
	}
	if u.BoolSlice[2] != true || u.BoolSlice[3] != false {
		t.Fatal()
	}
}

func TestDefaultBox_Fill_Map(t *testing.T) {
	type MapStruct struct {
		StringStringMap map[string]string  `default:"[k1: 100, k2: 好👌]"`
		StringIntMap    map[string]int     `default:"[k1: 100, k2: 200]"`
		StringFloatMap  map[string]float32 `default:"[k1: 3.14159, k2: 200]"`
		StringBoolMap   map[string]bool    `default:"[k1: true, k2: f]"`
		IntIntMap       map[int]int        `default:"[1: 100, 2: 200]"`
		BoolIntMap      map[bool]int       `default:"[true: 100, f: 200]"`
	}
	u := MapStruct{}
	New(&u).Fill()

	if len(u.StringStringMap) != 2 {
		t.Fatal()
	}
	if u.StringStringMap["k1"] != "100" || u.StringStringMap["k2"] != "好👌" {
		t.Fatal()
	}
	if len(u.StringIntMap) != 2 {
		t.Fatal()
	}
	if u.StringIntMap["k1"] != 100 || u.StringIntMap["k2"] != 200 {
		t.Fatal()
	}
	if len(u.StringFloatMap) != 2 {
		t.Fatal()
	}
	if u.StringFloatMap["k1"] != float32(3.14159) || u.StringFloatMap["k2"] != float32(200) {
		t.Fatal()
	}
	if len(u.StringBoolMap) != 2 {
		t.Fatal()
	}
	if u.StringBoolMap["k1"] != true || u.StringBoolMap["k2"] != false {
		t.Fatal()
	}
	if len(u.IntIntMap) != 2 {
		t.Fatal()
	}
	if u.IntIntMap[1] != 100 || u.IntIntMap[2] != 200 {
		t.Fatal()
	}
	if len(u.BoolIntMap) != 2 {
		t.Fatal()
	}
	if u.BoolIntMap[true] != 100 || u.BoolIntMap[false] != 200 {
		t.Fatal()
	}
}
