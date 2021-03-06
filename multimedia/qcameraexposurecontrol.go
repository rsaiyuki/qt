package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QCameraExposureControl struct {
	QMediaControl
}

type QCameraExposureControl_ITF interface {
	QMediaControl_ITF
	QCameraExposureControl_PTR() *QCameraExposureControl
}

func PointerFromQCameraExposureControl(ptr QCameraExposureControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraExposureControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraExposureControlFromPointer(ptr unsafe.Pointer) *QCameraExposureControl {
	var n = new(QCameraExposureControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraExposureControl_") {
		n.SetObjectName("QCameraExposureControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraExposureControl) QCameraExposureControl_PTR() *QCameraExposureControl {
	return ptr
}

//QCameraExposureControl::ExposureParameter
type QCameraExposureControl__ExposureParameter int64

const (
	QCameraExposureControl__ISO                       = QCameraExposureControl__ExposureParameter(0)
	QCameraExposureControl__Aperture                  = QCameraExposureControl__ExposureParameter(1)
	QCameraExposureControl__ShutterSpeed              = QCameraExposureControl__ExposureParameter(2)
	QCameraExposureControl__ExposureCompensation      = QCameraExposureControl__ExposureParameter(3)
	QCameraExposureControl__FlashPower                = QCameraExposureControl__ExposureParameter(4)
	QCameraExposureControl__FlashCompensation         = QCameraExposureControl__ExposureParameter(5)
	QCameraExposureControl__TorchPower                = QCameraExposureControl__ExposureParameter(6)
	QCameraExposureControl__SpotMeteringPoint         = QCameraExposureControl__ExposureParameter(7)
	QCameraExposureControl__ExposureMode              = QCameraExposureControl__ExposureParameter(8)
	QCameraExposureControl__MeteringMode              = QCameraExposureControl__ExposureParameter(9)
	QCameraExposureControl__ExtendedExposureParameter = QCameraExposureControl__ExposureParameter(1000)
)

func (ptr *QCameraExposureControl) ActualValue(parameter QCameraExposureControl__ExposureParameter) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::actualValue")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraExposureControl_ActualValue(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraExposureControl) ConnectActualValueChanged(f func(parameter int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::actualValueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectActualValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actualValueChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectActualValueChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::actualValueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectActualValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actualValueChanged")
	}
}

//export callbackQCameraExposureControlActualValueChanged
func callbackQCameraExposureControlActualValueChanged(ptrName *C.char, parameter C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::actualValueChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "actualValueChanged").(func(int))(int(parameter))
}

func (ptr *QCameraExposureControl) IsParameterSupported(parameter QCameraExposureControl__ExposureParameter) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::isParameterSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraExposureControl_IsParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraExposureControl) ConnectParameterRangeChanged(f func(parameter int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::parameterRangeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectParameterRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "parameterRangeChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectParameterRangeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::parameterRangeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectParameterRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "parameterRangeChanged")
	}
}

//export callbackQCameraExposureControlParameterRangeChanged
func callbackQCameraExposureControlParameterRangeChanged(ptrName *C.char, parameter C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::parameterRangeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "parameterRangeChanged").(func(int))(int(parameter))
}

func (ptr *QCameraExposureControl) RequestedValue(parameter QCameraExposureControl__ExposureParameter) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::requestedValue")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraExposureControl_RequestedValue(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraExposureControl) ConnectRequestedValueChanged(f func(parameter int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::requestedValueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_ConnectRequestedValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "requestedValueChanged", f)
	}
}

func (ptr *QCameraExposureControl) DisconnectRequestedValueChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::requestedValueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DisconnectRequestedValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "requestedValueChanged")
	}
}

//export callbackQCameraExposureControlRequestedValueChanged
func callbackQCameraExposureControlRequestedValueChanged(ptrName *C.char, parameter C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::requestedValueChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "requestedValueChanged").(func(int))(int(parameter))
}

func (ptr *QCameraExposureControl) SetValue(parameter QCameraExposureControl__ExposureParameter, value core.QVariant_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::setValue")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraExposureControl_SetValue(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QCameraExposureControl) DestroyQCameraExposureControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraExposureControl::~QCameraExposureControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraExposureControl_DestroyQCameraExposureControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
