package default_box

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// defaultBox Fill by default value from tag.
type defaultBox struct {
	ObjectPointer interface{} `hint:"should be a pointer"`
	TagKey        string      `default:"default"`
}

// New Pack object into the defaultBox.
func New(object interface{}) *defaultBox {
	if reflect.TypeOf(object).Kind() != reflect.Ptr {
		panic("ObjectPointer should be a pointer to struct")
	}
	return &defaultBox{ObjectPointer: object, TagKey: "default"}
}

// Tag Fetch tag on ObjectPointer filed, return value of the key.
func (box *defaultBox) Tag(field string) (tag string, ok bool) {
	objType := reflect.TypeOf(box.ObjectPointer).Elem()
	if sf, ok := objType.FieldByName(field); ok {
		if tag, ok := sf.Tag.Lookup(box.TagKey); ok {
			return tag, true
		}
	}
	return "", false
}

// SetBasic Parse field type and set converted value, only support basic type.
func SetBasic(fieldValue reflect.Value, value string) bool {
	if !fieldValue.CanSet() {
		return false
	}
	switch fieldValue.Kind() {
	case reflect.Invalid:
		return false
	case reflect.String:
		fieldValue.SetString(value)
		return true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if i64, e := strconv.ParseInt(value, 10, 64); e == nil {
			fieldValue.SetInt(i64)
			return true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if ui64, e := strconv.ParseUint(value, 10, 64); e == nil {
			fieldValue.SetUint(ui64)
			return true
		}
	case reflect.Float32, reflect.Float64:
		if f64, e := strconv.ParseFloat(value, 64); e == nil {
			fieldValue.SetFloat(f64)
			return true
		}
	case reflect.Bool:
		if b, e := strconv.ParseBool(value); e == nil {
			fieldValue.SetBool(b)
			return true
		}
	}
	return false
}

// SetSlice Parse slice type and set converted value.
func SetSlice(fieldType reflect.StructField, fieldValue reflect.Value, value string) bool {
	if len(value) < 2 {
		return false
	}
	items := strings.Split(value[1:len(value)-1], ",")
	if len(items) < 1 {
		return false
	}
	defaultSlice := reflect.MakeSlice(fieldType.Type, len(items), len(items))
	for i := 0; i < len(items); i++ {
		SetBasic(defaultSlice.Index(i), strings.TrimSpace(items[i]))
	}
	fieldValue.Set(defaultSlice)
	return true
}

// SetMap Parse field type and set converted value.
func SetMap(fieldType reflect.StructField, fieldValue reflect.Value, value string) bool {
	if len(value) < 2 {
		return false
	}
	items := strings.Split(value[1:len(value)-1], ",")
	if len(items) < 1 {
		return false
	}
	defaultMap := reflect.MakeMap(fieldType.Type)
	for i := 0; i < len(items); i++ {
		kv := strings.Split(items[i], ":")
		if len(kv) != 2 {
			continue
		}

		mapKey := reflect.New(fieldType.Type.Key()).Elem()
		SetBasic(mapKey, strings.TrimSpace(kv[0]))
		mapValue := reflect.New(fieldType.Type.Elem()).Elem()
		SetBasic(mapValue, strings.TrimSpace(kv[1]))
		defaultMap.SetMapIndex(mapKey, mapValue)
	}
	fieldValue.Set(defaultMap)
	return true
}

// FillDefault Fill default values with tags.
func (box *defaultBox) Fill() *defaultBox {
	objValue := reflect.ValueOf(box.ObjectPointer).Elem()
	objType := reflect.TypeOf(box.ObjectPointer).Elem()
	if objValue.Kind() != reflect.Struct {
		return box
	}
	for i := 0; i < objValue.NumField(); i++ {
		fieldRunes := []rune(objType.Field(i).Name)
		if len(fieldRunes) < 1 || unicode.IsLower(fieldRunes[0]) {
			continue
		}
		defaultVal, present := objType.Field(i).Tag.Lookup(box.TagKey)
		if !present {
			continue
		}
		// basic type: string, bool, int..., uint..., float...,
		if ok := SetBasic(objValue.Field(i), defaultVal); ok {
			continue
		}
		// complex type: slice, map
		switch objValue.Field(i).Kind() {
		case reflect.Slice:
			SetSlice(objType.Field(i), objValue.Field(i), defaultVal)
		case reflect.Map:
			SetMap(objType.Field(i), objValue.Field(i), defaultVal)
		}
	}
	return box
}
