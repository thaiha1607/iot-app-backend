package utils

import (
	"reflect"
	"runtime"
)

// GetFuncName returns the name of the function passed as argument
func GetFuncName(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
