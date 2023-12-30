package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"getcharzp.cn/helper"
)

func TestRandGenerate(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	println(s)
}

func TestCheckGoCodeValid(t *testing.T) {
	valid, err := helper.CheckGoCodeValid("../code/code-user/main.go")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(valid)
}
