package sensors

//#include "sensors.h"
import "C"
import (
	"log"
	"unsafe"
)

type QCompassFilter struct {
	QSensorFilter
}

type QCompassFilter_ITF interface {
	QSensorFilter_ITF
	QCompassFilter_PTR() *QCompassFilter
}

func PointerFromQCompassFilter(ptr QCompassFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompassFilter_PTR().Pointer()
	}
	return nil
}

func NewQCompassFilterFromPointer(ptr unsafe.Pointer) *QCompassFilter {
	var n = new(QCompassFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCompassFilter) QCompassFilter_PTR() *QCompassFilter {
	return ptr
}

func (ptr *QCompassFilter) Filter(reading QCompassReading_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompassFilter::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCompassFilter_Filter(ptr.Pointer(), PointerFromQCompassReading(reading)) != 0
	}
	return false
}
