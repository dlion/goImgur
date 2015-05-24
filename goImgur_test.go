package goImgur_test

import (
	"github.com/dlion/goImgur"
	"reflect"
	"testing"
)

func TestUpload(t *testing.T) {
	str, err := goImgur.Upload("image_test.png", "ed6e565e7bd497e")
	if err != nil {
		t.Error(err)
	}

	typeStr := reflect.TypeOf(*str).String()

	if typeStr != "string" {
		t.Error("str is not a string")
	}
}
