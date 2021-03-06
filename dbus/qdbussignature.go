package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDBusSignature struct {
	ptr unsafe.Pointer
}

type QDBusSignature_ITF interface {
	QDBusSignature_PTR() *QDBusSignature
}

func (p *QDBusSignature) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusSignature) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusSignature(ptr QDBusSignature_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusSignature_PTR().Pointer()
	}
	return nil
}

func NewQDBusSignatureFromPointer(ptr unsafe.Pointer) *QDBusSignature {
	var n = new(QDBusSignature)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusSignature) QDBusSignature_PTR() *QDBusSignature {
	return ptr
}

func NewQDBusSignature() *QDBusSignature {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusSignature::QDBusSignature")
		}
	}()

	return NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature())
}

func NewQDBusSignature3(signature core.QLatin1String_ITF) *QDBusSignature {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusSignature::QDBusSignature")
		}
	}()

	return NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature3(core.PointerFromQLatin1String(signature)))
}

func NewQDBusSignature4(signature string) *QDBusSignature {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusSignature::QDBusSignature")
		}
	}()

	return NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature4(C.CString(signature)))
}

func NewQDBusSignature2(signature string) *QDBusSignature {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusSignature::QDBusSignature")
		}
	}()

	return NewQDBusSignatureFromPointer(C.QDBusSignature_NewQDBusSignature2(C.CString(signature)))
}

func (ptr *QDBusSignature) SetSignature(signature string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusSignature::setSignature")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusSignature_SetSignature(ptr.Pointer(), C.CString(signature))
	}
}

func (ptr *QDBusSignature) Signature() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusSignature::signature")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusSignature_Signature(ptr.Pointer()))
	}
	return ""
}
