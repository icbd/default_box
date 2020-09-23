# DefaultBox

A tool developed for Golang.

Fill the default value of struct with struct tag.

## Install

```shell script
go get github.com/icbd/default_box
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
PackDefaultBox(&u).Fill()
fmt.Printf("%+v", u)

// Chain style
userWithDefaultValue := PackDefaultBox(&User{}).Fill().ObjectPointer.(*User)
fmt.Printf("%+v", userWithDefaultValue)
```