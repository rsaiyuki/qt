package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QStatusTipEvent struct {
	core.QEvent
}

type QStatusTipEvent_ITF interface {
	core.QEvent_ITF
	QStatusTipEvent_PTR() *QStatusTipEvent
}

func PointerFromQStatusTipEvent(ptr QStatusTipEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatusTipEvent_PTR().Pointer()
	}
	return nil
}

func NewQStatusTipEventFromPointer(ptr unsafe.Pointer) *QStatusTipEvent {
	var n = new(QStatusTipEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStatusTipEvent) QStatusTipEvent_PTR() *QStatusTipEvent {
	return ptr
}

func NewQStatusTipEvent(tip string) *QStatusTipEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusTipEvent::QStatusTipEvent")
		}
	}()

	return NewQStatusTipEventFromPointer(C.QStatusTipEvent_NewQStatusTipEvent(C.CString(tip)))
}

func (ptr *QStatusTipEvent) Tip() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStatusTipEvent::tip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QStatusTipEvent_Tip(ptr.Pointer()))
	}
	return ""
}
