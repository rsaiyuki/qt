package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QHttpMultiPart struct {
	core.QObject
}

type QHttpMultiPart_ITF interface {
	core.QObject_ITF
	QHttpMultiPart_PTR() *QHttpMultiPart
}

func PointerFromQHttpMultiPart(ptr QHttpMultiPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpMultiPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpMultiPartFromPointer(ptr unsafe.Pointer) *QHttpMultiPart {
	var n = new(QHttpMultiPart)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHttpMultiPart_") {
		n.SetObjectName("QHttpMultiPart_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHttpMultiPart) QHttpMultiPart_PTR() *QHttpMultiPart {
	return ptr
}

//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int64

const (
	QHttpMultiPart__MixedType       = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType = QHttpMultiPart__ContentType(3)
)

func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObject_ITF) *QHttpMultiPart {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpMultiPart::QHttpMultiPart")
		}
	}()

	return NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart2(C.int(contentType), core.PointerFromQObject(parent)))
}

func NewQHttpMultiPart(parent core.QObject_ITF) *QHttpMultiPart {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpMultiPart::QHttpMultiPart")
		}
	}()

	return NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart(core.PointerFromQObject(parent)))
}

func (ptr *QHttpMultiPart) Append(httpPart QHttpPart_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpMultiPart::append")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_Append(ptr.Pointer(), PointerFromQHttpPart(httpPart))
	}
}

func (ptr *QHttpMultiPart) Boundary() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpMultiPart::boundary")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QHttpMultiPart_Boundary(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHttpMultiPart) SetBoundary(boundary core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpMultiPart::setBoundary")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetBoundary(ptr.Pointer(), core.PointerFromQByteArray(boundary))
	}
}

func (ptr *QHttpMultiPart) SetContentType(contentType QHttpMultiPart__ContentType) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpMultiPart::setContentType")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetContentType(ptr.Pointer(), C.int(contentType))
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpMultiPart::~QHttpMultiPart")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DestroyQHttpMultiPart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
