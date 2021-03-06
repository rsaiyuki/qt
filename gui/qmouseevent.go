package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QMouseEvent struct {
	QInputEvent
}

type QMouseEvent_ITF interface {
	QInputEvent_ITF
	QMouseEvent_PTR() *QMouseEvent
}

func PointerFromQMouseEvent(ptr QMouseEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMouseEvent_PTR().Pointer()
	}
	return nil
}

func NewQMouseEventFromPointer(ptr unsafe.Pointer) *QMouseEvent {
	var n = new(QMouseEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMouseEvent) QMouseEvent_PTR() *QMouseEvent {
	return ptr
}

func NewQMouseEvent(ty core.QEvent__Type, localPos core.QPointF_ITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::QMouseEvent")
		}
	}()

	return NewQMouseEventFromPointer(C.QMouseEvent_NewQMouseEvent(C.int(ty), core.PointerFromQPointF(localPos), C.int(button), C.int(buttons), C.int(modifiers)))
}

func NewQMouseEvent2(ty core.QEvent__Type, localPos core.QPointF_ITF, screenPos core.QPointF_ITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::QMouseEvent")
		}
	}()

	return NewQMouseEventFromPointer(C.QMouseEvent_NewQMouseEvent2(C.int(ty), core.PointerFromQPointF(localPos), core.PointerFromQPointF(screenPos), C.int(button), C.int(buttons), C.int(modifiers)))
}

func NewQMouseEvent3(ty core.QEvent__Type, localPos core.QPointF_ITF, windowPos core.QPointF_ITF, screenPos core.QPointF_ITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::QMouseEvent")
		}
	}()

	return NewQMouseEventFromPointer(C.QMouseEvent_NewQMouseEvent3(C.int(ty), core.PointerFromQPointF(localPos), core.PointerFromQPointF(windowPos), core.PointerFromQPointF(screenPos), C.int(button), C.int(buttons), C.int(modifiers)))
}

func (ptr *QMouseEvent) Button() core.Qt__MouseButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::button")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEvent_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) Buttons() core.Qt__MouseButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::buttons")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) Flags() core.Qt__MouseEventFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MouseEventFlag(C.QMouseEvent_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) GlobalX() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::globalX")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMouseEvent_GlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) GlobalY() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::globalY")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMouseEvent_GlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) Source() core.Qt__MouseEventSource {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::source")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MouseEventSource(C.QMouseEvent_Source(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) X() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::x")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMouseEvent_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEvent) Y() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMouseEvent::y")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMouseEvent_Y(ptr.Pointer()))
	}
	return 0
}
