package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QPanGesture struct {
	QGesture
}

type QPanGesture_ITF interface {
	QGesture_ITF
	QPanGesture_PTR() *QPanGesture
}

func PointerFromQPanGesture(ptr QPanGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPanGesture_PTR().Pointer()
	}
	return nil
}

func NewQPanGestureFromPointer(ptr unsafe.Pointer) *QPanGesture {
	var n = new(QPanGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPanGesture_") {
		n.SetObjectName("QPanGesture_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPanGesture) QPanGesture_PTR() *QPanGesture {
	return ptr
}

func (ptr *QPanGesture) Acceleration() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPanGesture::acceleration")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QPanGesture_Acceleration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPanGesture) SetAcceleration(value float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPanGesture::setAcceleration")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPanGesture_SetAcceleration(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPanGesture) SetLastOffset(value core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPanGesture::setLastOffset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPanGesture_SetLastOffset(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPanGesture) SetOffset(value core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPanGesture::setOffset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPanGesture_SetOffset(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPanGesture) DestroyQPanGesture() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPanGesture::~QPanGesture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPanGesture_DestroyQPanGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
