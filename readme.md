# DefaultBox
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/icbd/default_box/Test)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/icbd/default_box)

A tool developed for Golang.

Fill the default value of struct with struct tag.

## Install

```shell script
go get github.com/icbd/default_box@v1.0.0
```
## How to use

* Support basic types and basic type Slice and basic type Map;
* Use `[]` to mark as Slice, use `{}` to mark as Map; 
* Use `,` split items;
* Do not use `,` and `"` and `:`  within each item.

```go
package main
type User struct {
	Name    string             `default:"Bob"`
	Age     int8               `default:"10"`
	Hobbies []string           `default:"[Football, Basketball]"`
	Scores  map[string]float32 `default:"{Language: 95.55, Math: 99.50}"`
}

// Fill a existed object
u := User{}
default_box.New(&u).Fill()
fmt.Printf("%+v", u)

// Chain style
userWithDefaultValue := default_box.New(&User{}).Fill().ObjectPointer.(*User)
fmt.Printf("%+v", userWithDefaultValue)
```

### Attention

Since `default_box` is implemented by reflection, fields that use default should start with an uppercase letter.

## License

MIT, see [LICENSE](LICENSE)