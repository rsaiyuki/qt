package xml

//#include "xml.h"
import "C"
import (
	"log"
	"unsafe"
)

type QXmlLocator struct {
	ptr unsafe.Pointer
}

type QXmlLocator_ITF interface {
	QXmlLocator_PTR() *QXmlLocator
}

func (p *QXmlLocator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlLocator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlLocator(ptr QXmlLocator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLocator_PTR().Pointer()
	}
	return nil
}

func NewQXmlLocatorFromPointer(ptr unsafe.Pointer) *QXmlLocator {
	var n = new(QXmlLocator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlLocator) QXmlLocator_PTR() *QXmlLocator {
	return ptr
}

func (ptr *QXmlLocator) ColumnNumber() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLocator::columnNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_ColumnNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlLocator) LineNumber() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLocator::lineNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_LineNumber(ptr.Pointer()))
	}
	return 0
}

func (ptr *QXmlLocator) DestroyQXmlLocator() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlLocator::~QXmlLocator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlLocator_DestroyQXmlLocator(ptr.Pointer())
	}
}
