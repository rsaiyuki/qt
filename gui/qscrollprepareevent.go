package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QScrollPrepareEvent struct {
	core.QEvent
}

type QScrollPrepareEvent_ITF interface {
	core.QEvent_ITF
	QScrollPrepareEvent_PTR() *QScrollPrepareEvent
}

func PointerFromQScrollPrepareEvent(ptr QScrollPrepareEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollPrepareEvent_PTR().Pointer()
	}
	return nil
}

func NewQScrollPrepareEventFromPointer(ptr unsafe.Pointer) *QScrollPrepareEvent {
	var n = new(QScrollPrepareEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScrollPrepareEvent) QScrollPrepareEvent_PTR() *QScrollPrepareEvent {
	return ptr
}

func NewQScrollPrepareEvent(startPos core.QPointF_ITF) *QScrollPrepareEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollPrepareEvent::QScrollPrepareEvent")
		}
	}()

	return NewQScrollPrepareEventFromPointer(C.QScrollPrepareEvent_NewQScrollPrepareEvent(core.PointerFromQPointF(startPos)))
}

func (ptr *QScrollPrepareEvent) SetContentPos(pos core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollPrepareEvent::setContentPos")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollPrepareEvent_SetContentPos(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QScrollPrepareEvent) SetContentPosRange(rect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollPrepareEvent::setContentPosRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollPrepareEvent_SetContentPosRange(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QScrollPrepareEvent) SetViewportSize(size core.QSizeF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollPrepareEvent::setViewportSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollPrepareEvent_SetViewportSize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QScrollPrepareEvent) DestroyQScrollPrepareEvent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollPrepareEvent::~QScrollPrepareEvent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollPrepareEvent_DestroyQScrollPrepareEvent(ptr.Pointer())
	}
}
