package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QCompass struct {
	QSensor
}

type QCompass_ITF interface {
	QSensor_ITF
	QCompass_PTR() *QCompass
}

func PointerFromQCompass(ptr QCompass_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompass_PTR().Pointer()
	}
	return nil
}

func NewQCompassFromPointer(ptr unsafe.Pointer) *QCompass {
	var n = new(QCompass)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCompass_") {
		n.SetObjectName("QCompass_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCompass) QCompass_PTR() *QCompass {
	return ptr
}

func (ptr *QCompass) Reading() *QCompassReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompass::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCompassReadingFromPointer(C.QCompass_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQCompass(parent core.QObject_ITF) *QCompass {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompass::QCompass")
		}
	}()

	return NewQCompassFromPointer(C.QCompass_NewQCompass(core.PointerFromQObject(parent)))
}

func (ptr *QCompass) DestroyQCompass() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompass::~QCompass")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompass_DestroyQCompass(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
