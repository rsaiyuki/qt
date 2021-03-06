package script

//#include "script.h"
import "C"
import (
	"log"
	"unsafe"
)

type QScriptSyntaxCheckResult struct {
	ptr unsafe.Pointer
}

type QScriptSyntaxCheckResult_ITF interface {
	QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult
}

func (p *QScriptSyntaxCheckResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptSyntaxCheckResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptSyntaxCheckResult(ptr QScriptSyntaxCheckResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptSyntaxCheckResult_PTR().Pointer()
	}
	return nil
}

func NewQScriptSyntaxCheckResultFromPointer(ptr unsafe.Pointer) *QScriptSyntaxCheckResult {
	var n = new(QScriptSyntaxCheckResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptSyntaxCheckResult) QScriptSyntaxCheckResult_PTR() *QScriptSyntaxCheckResult {
	return ptr
}

//QScriptSyntaxCheckResult::State
type QScriptSyntaxCheckResult__State int64

const (
	QScriptSyntaxCheckResult__Error        = QScriptSyntaxCheckResult__State(0)
	QScriptSyntaxCheckResult__Intermediate = QScriptSyntaxCheckResult__State(1)
	QScriptSyntaxCheckResult__Valid        = QScriptSyntaxCheckResult__State(2)
)

func NewQScriptSyntaxCheckResult(other QScriptSyntaxCheckResult_ITF) *QScriptSyntaxCheckResult {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptSyntaxCheckResult::QScriptSyntaxCheckResult")
		}
	}()

	return NewQScriptSyntaxCheckResultFromPointer(C.QScriptSyntaxCheckResult_NewQScriptSyntaxCheckResult(PointerFromQScriptSyntaxCheckResult(other)))
}

func (ptr *QScriptSyntaxCheckResult) ErrorColumnNumber() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptSyntaxCheckResult::errorColumnNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QScriptSyntaxCheckResult_ErrorColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) ErrorLineNumber() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptSyntaxCheckResult::errorLineNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QScriptSyntaxCheckResult_ErrorLineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScriptSyntaxCheckResult) ErrorMessage() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptSyntaxCheckResult::errorMessage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QScriptSyntaxCheckResult_ErrorMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QScriptSyntaxCheckResult) DestroyQScriptSyntaxCheckResult() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptSyntaxCheckResult::~QScriptSyntaxCheckResult")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScriptSyntaxCheckResult_DestroyQScriptSyntaxCheckResult(ptr.Pointer())
	}
}
