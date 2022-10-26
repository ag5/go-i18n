package i18n

import "reflect"

type empty struct{}

var PackagePath = reflect.TypeOf(empty{}).PkgPath()
