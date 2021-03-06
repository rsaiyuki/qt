package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTabletEvent struct {
	QInputEvent
}

type QTabletEvent_ITF interface {
	QInputEvent_ITF
	QTabletEvent_PTR() *QTabletEvent
}

func PointerFromQTabletEvent(ptr QTabletEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTabletEvent_PTR().Pointer()
	}
	return nil
}

func NewQTabletEventFromPointer(ptr unsafe.Pointer) *QTabletEvent {
	var n = new(QTabletEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTabletEvent) QTabletEvent_PTR() *QTabletEvent {
	return ptr
}

//QTabletEvent::PointerType
type QTabletEvent__PointerType int64

const (
	QTabletEvent__UnknownPointer = QTabletEvent__PointerType(0)
	QTabletEvent__Pen            = QTabletEvent__PointerType(1)
	QTabletEvent__Cursor         = QTabletEvent__PointerType(2)
	QTabletEvent__Eraser         = QTabletEvent__PointerType(3)
)

//QTabletEvent::TabletDevice
type QTabletEvent__TabletDevice int64

const (
	QTabletEvent__NoDevice       = QTabletEvent__TabletDevice(0)
	QTabletEvent__Puck           = QTabletEvent__TabletDevice(1)
	QTabletEvent__Stylus         = QTabletEvent__TabletDevice(2)
	QTabletEvent__Airbrush       = QTabletEvent__TabletDevice(3)
	QTabletEvent__FourDMouse     = QTabletEvent__TabletDevice(4)
	QTabletEvent__XFreeEraser    = QTabletEvent__TabletDevice(5)
	QTabletEvent__RotationStylus = QTabletEvent__TabletDevice(6)
)

func (ptr *QTabletEvent) Button() core.Qt__MouseButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::button")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QTabletEvent_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) Buttons() core.Qt__MouseButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::buttons")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QTabletEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) Device() QTabletEvent__TabletDevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::device")
		}
	}()

	if ptr.Pointer() != nil {
		return QTabletEvent__TabletDevice(C.QTabletEvent_Device(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) GlobalX() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::globalX")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_GlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) GlobalY() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::globalY")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_GlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) HiResGlobalX() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::hiResGlobalX")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTabletEvent_HiResGlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) HiResGlobalY() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::hiResGlobalY")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTabletEvent_HiResGlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) PointerType() QTabletEvent__PointerType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::pointerType")
		}
	}()

	if ptr.Pointer() != nil {
		return QTabletEvent__PointerType(C.QTabletEvent_PointerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) Pressure() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::pressure")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTabletEvent_Pressure(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) Rotation() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::rotation")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTabletEvent_Rotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) TangentialPressure() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::tangentialPressure")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTabletEvent_TangentialPressure(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) X() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::x")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) XTilt() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::xTilt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_XTilt(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) Y() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::y")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) YTilt() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::yTilt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_YTilt(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTabletEvent) Z() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTabletEvent::z")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_Z(ptr.Pointer()))
	}
	return 0
}
