package api

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
)

// Other side can dynamically/runtime construct an Table
// PRO Smidigt
// CON small speed penality constructing when Target side create Table struct pointing into Arrow records ,
type CallWithTable interface {
	Call(schema arrow.Schema, date32 array.Date32) error
	//	Process(Arrow.Sche reader io.Reader, customParams interface{}) bool
}
