package golib

import (
	"fmt"
	"github.com/maprost/testbox/must"
	"strconv"
	"testing"
)

func Equal(t testing.TB, msg string, actual int, expected string) {
	t.Helper()
	fmt.Println(msg, actual)
	must.BeEqual(t, strconv.Itoa(actual), expected)
}
