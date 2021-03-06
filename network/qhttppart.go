package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QHttpPart struct {
	ptr unsafe.Pointer
}

type QHttpPart_ITF interface {
	QHttpPart_PTR() *QHttpPart
}

func (p *QHttpPart) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHttpPart) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHttpPart(ptr QHttpPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpPartFromPointer(ptr unsafe.Pointer) *QHttpPart {
	var n = new(QHttpPart)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHttpPart) QHttpPart_PTR() *QHttpPart {
	return ptr
}

func NewQHttpPart() *QHttpPart {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpPart::QHttpPart")
		}
	}()

	return NewQHttpPartFromPointer(C.QHttpPart_NewQHttpPart())
}

func NewQHttpPart2(other QHttpPart_ITF) *QHttpPart {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpPart::QHttpPart")
		}
	}()

	return NewQHttpPartFromPointer(C.QHttpPart_NewQHttpPart2(PointerFromQHttpPart(other)))
}

func (ptr *QHttpPart) SetBody(body core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpPart::setBody")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpPart_SetBody(ptr.Pointer(), core.PointerFromQByteArray(body))
	}
}

func (ptr *QHttpPart) SetBodyDevice(device core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpPart::setBodyDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpPart_SetBodyDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QHttpPart) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpPart::setHeader")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpPart_SetHeader(ptr.Pointer(), C.int(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QHttpPart) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpPart::setRawHeader")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpPart_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QHttpPart) Swap(other QHttpPart_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpPart::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpPart_Swap(ptr.Pointer(), PointerFromQHttpPart(other))
	}
}

func (ptr *QHttpPart) DestroyQHttpPart() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHttpPart::~QHttpPart")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHttpPart_DestroyQHttpPart(ptr.Pointer())
	}
}
