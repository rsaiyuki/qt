package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAccessibleTextRemoveEvent struct {
	QAccessibleTextCursorEvent
}

type QAccessibleTextRemoveEvent_ITF interface {
	QAccessibleTextCursorEvent_ITF
	QAccessibleTextRemoveEvent_PTR() *QAccessibleTextRemoveEvent
}

func PointerFromQAccessibleTextRemoveEvent(ptr QAccessibleTextRemoveEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextRemoveEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleTextRemoveEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextRemoveEvent {
	var n = new(QAccessibleTextRemoveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextRemoveEvent) QAccessibleTextRemoveEvent_PTR() *QAccessibleTextRemoveEvent {
	return ptr
}

func NewQAccessibleTextRemoveEvent2(iface QAccessibleInterface_ITF, position int, text string) *QAccessibleTextRemoveEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent")
		}
	}()

	return NewQAccessibleTextRemoveEventFromPointer(C.QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent2(PointerFromQAccessibleInterface(iface), C.int(position), C.CString(text)))
}

func NewQAccessibleTextRemoveEvent(object core.QObject_ITF, position int, text string) *QAccessibleTextRemoveEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent")
		}
	}()

	return NewQAccessibleTextRemoveEventFromPointer(C.QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent(core.PointerFromQObject(object), C.int(position), C.CString(text)))
}

func (ptr *QAccessibleTextRemoveEvent) ChangePosition() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTextRemoveEvent::changePosition")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextRemoveEvent_ChangePosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTextRemoveEvent) TextRemoved() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleTextRemoveEvent::textRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTextRemoveEvent_TextRemoved(ptr.Pointer()))
	}
	return ""
}
