package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QTapReading struct {
	QSensorReading
}

type QTapReading_ITF interface {
	QSensorReading_ITF
	QTapReading_PTR() *QTapReading
}

func PointerFromQTapReading(ptr QTapReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapReading_PTR().Pointer()
	}
	return nil
}

func NewQTapReadingFromPointer(ptr unsafe.Pointer) *QTapReading {
	var n = new(QTapReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTapReading_") {
		n.SetObjectName("QTapReading_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTapReading) QTapReading_PTR() *QTapReading {
	return ptr
}

//QTapReading::TapDirection
type QTapReading__TapDirection int64

const (
	QTapReading__Undefined = QTapReading__TapDirection(0)
	QTapReading__X         = QTapReading__TapDirection(0x0001)
	QTapReading__Y         = QTapReading__TapDirection(0x0002)
	QTapReading__Z         = QTapReading__TapDirection(0x0004)
	QTapReading__X_Pos     = QTapReading__TapDirection(0x0011)
	QTapReading__Y_Pos     = QTapReading__TapDirection(0x0022)
	QTapReading__Z_Pos     = QTapReading__TapDirection(0x0044)
	QTapReading__X_Neg     = QTapReading__TapDirection(0x0101)
	QTapReading__Y_Neg     = QTapReading__TapDirection(0x0202)
	QTapReading__Z_Neg     = QTapReading__TapDirection(0x0404)
	QTapReading__X_Both    = QTapReading__TapDirection(0x0111)
	QTapReading__Y_Both    = QTapReading__TapDirection(0x0222)
	QTapReading__Z_Both    = QTapReading__TapDirection(0x0444)
)

func (ptr *QTapReading) IsDoubleTap() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapReading::isDoubleTap")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTapReading_IsDoubleTap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapReading) TapDirection() QTapReading__TapDirection {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapReading::tapDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return QTapReading__TapDirection(C.QTapReading_TapDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTapReading) SetDoubleTap(doubleTap bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapReading::setDoubleTap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTapReading_SetDoubleTap(ptr.Pointer(), C.int(qt.GoBoolToInt(doubleTap)))
	}
}

func (ptr *QTapReading) SetTapDirection(tapDirection QTapReading__TapDirection) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapReading::setTapDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTapReading_SetTapDirection(ptr.Pointer(), C.int(tapDirection))
	}
}
