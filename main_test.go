package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	var i *int
	x(&i)

	t.Fatal("SUCCEEDED")
}

func x(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println(reflect.ValueOf(i).IsNil())
	fmt.Println(t)
}
