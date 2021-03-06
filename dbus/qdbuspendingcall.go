package dbus

//#include "dbus.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDBusPendingCall struct {
	ptr unsafe.Pointer
}

type QDBusPendingCall_ITF interface {
	QDBusPendingCall_PTR() *QDBusPendingCall
}

func (p *QDBusPendingCall) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusPendingCall) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusPendingCall(ptr QDBusPendingCall_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusPendingCall_PTR().Pointer()
	}
	return nil
}

func NewQDBusPendingCallFromPointer(ptr unsafe.Pointer) *QDBusPendingCall {
	var n = new(QDBusPendingCall)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusPendingCall) QDBusPendingCall_PTR() *QDBusPendingCall {
	return ptr
}

func NewQDBusPendingCall(other QDBusPendingCall_ITF) *QDBusPendingCall {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCall::QDBusPendingCall")
		}
	}()

	return NewQDBusPendingCallFromPointer(C.QDBusPendingCall_NewQDBusPendingCall(PointerFromQDBusPendingCall(other)))
}

func (ptr *QDBusPendingCall) Swap(other QDBusPendingCall_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCall::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusPendingCall_Swap(ptr.Pointer(), PointerFromQDBusPendingCall(other))
	}
}

func (ptr *QDBusPendingCall) DestroyQDBusPendingCall() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusPendingCall::~QDBusPendingCall")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusPendingCall_DestroyQDBusPendingCall(ptr.Pointer())
	}
}
