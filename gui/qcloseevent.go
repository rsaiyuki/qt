package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QCloseEvent struct {
	core.QEvent
}

type QCloseEvent_ITF interface {
	core.QEvent_ITF
	QCloseEvent_PTR() *QCloseEvent
}

func PointerFromQCloseEvent(ptr QCloseEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCloseEvent_PTR().Pointer()
	}
	return nil
}

func NewQCloseEventFromPointer(ptr unsafe.Pointer) *QCloseEvent {
	var n = new(QCloseEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCloseEvent) QCloseEvent_PTR() *QCloseEvent {
	return ptr
}

func NewQCloseEvent() *QCloseEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCloseEvent::QCloseEvent")
		}
	}()

	return NewQCloseEventFromPointer(C.QCloseEvent_NewQCloseEvent())
}
