package sensors

//#include "sensors.h"
import "C"
import (
	"log"
	"unsafe"
)

type QRotationFilter struct {
	QSensorFilter
}

type QRotationFilter_ITF interface {
	QSensorFilter_ITF
	QRotationFilter_PTR() *QRotationFilter
}

func PointerFromQRotationFilter(ptr QRotationFilter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationFilter_PTR().Pointer()
	}
	return nil
}

func NewQRotationFilterFromPointer(ptr unsafe.Pointer) *QRotationFilter {
	var n = new(QRotationFilter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRotationFilter) QRotationFilter_PTR() *QRotationFilter {
	return ptr
}

func (ptr *QRotationFilter) Filter(reading QRotationReading_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationFilter::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRotationFilter_Filter(ptr.Pointer(), PointerFromQRotationReading(reading)) != 0
	}
	return false
}
