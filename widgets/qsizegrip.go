package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QSizeGrip struct {
	QWidget
}

type QSizeGrip_ITF interface {
	QWidget_ITF
	QSizeGrip_PTR() *QSizeGrip
}

func PointerFromQSizeGrip(ptr QSizeGrip_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizeGrip_PTR().Pointer()
	}
	return nil
}

func NewQSizeGripFromPointer(ptr unsafe.Pointer) *QSizeGrip {
	var n = new(QSizeGrip)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSizeGrip_") {
		n.SetObjectName("QSizeGrip_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSizeGrip) QSizeGrip_PTR() *QSizeGrip {
	return ptr
}

func NewQSizeGrip(parent QWidget_ITF) *QSizeGrip {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeGrip::QSizeGrip")
		}
	}()

	return NewQSizeGripFromPointer(C.QSizeGrip_NewQSizeGrip(PointerFromQWidget(parent)))
}

func (ptr *QSizeGrip) SetVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeGrip::setVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSizeGrip_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSizeGrip) DestroyQSizeGrip() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeGrip::~QSizeGrip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSizeGrip_DestroyQSizeGrip(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
