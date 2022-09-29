package api

import (
	"github.com/apache/arrow/go/v9/arrow"
)

// Other side can dynamically/runtime construct an Table
// PRO Smidigt
// CON small speed penality constructing when Target side create Table struct pointing into Arrow records ,
type CallWithTable interface {
	Call(schema arrow.Schema) error
	//	Process(Arrow.Sche reader io.Reader, customParams interface{}) bool
}
