package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QCameraViewfinderSettingsControl struct {
	QMediaControl
}

type QCameraViewfinderSettingsControl_ITF interface {
	QMediaControl_ITF
	QCameraViewfinderSettingsControl_PTR() *QCameraViewfinderSettingsControl
}

func PointerFromQCameraViewfinderSettingsControl(ptr QCameraViewfinderSettingsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraViewfinderSettingsControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraViewfinderSettingsControlFromPointer(ptr unsafe.Pointer) *QCameraViewfinderSettingsControl {
	var n = new(QCameraViewfinderSettingsControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraViewfinderSettingsControl_") {
		n.SetObjectName("QCameraViewfinderSettingsControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraViewfinderSettingsControl) QCameraViewfinderSettingsControl_PTR() *QCameraViewfinderSettingsControl {
	return ptr
}

//QCameraViewfinderSettingsControl::ViewfinderParameter
type QCameraViewfinderSettingsControl__ViewfinderParameter int64

const (
	QCameraViewfinderSettingsControl__Resolution       = QCameraViewfinderSettingsControl__ViewfinderParameter(0)
	QCameraViewfinderSettingsControl__PixelAspectRatio = QCameraViewfinderSettingsControl__ViewfinderParameter(1)
	QCameraViewfinderSettingsControl__MinimumFrameRate = QCameraViewfinderSettingsControl__ViewfinderParameter(2)
	QCameraViewfinderSettingsControl__MaximumFrameRate = QCameraViewfinderSettingsControl__ViewfinderParameter(3)
	QCameraViewfinderSettingsControl__PixelFormat      = QCameraViewfinderSettingsControl__ViewfinderParameter(4)
	QCameraViewfinderSettingsControl__UserParameter    = QCameraViewfinderSettingsControl__ViewfinderParameter(1000)
)

func (ptr *QCameraViewfinderSettingsControl) IsViewfinderParameterSupported(parameter QCameraViewfinderSettingsControl__ViewfinderParameter) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettingsControl::isViewfinderParameterSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(ptr.Pointer(), C.int(parameter)) != 0
	}
	return false
}

func (ptr *QCameraViewfinderSettingsControl) SetViewfinderParameter(parameter QCameraViewfinderSettingsControl__ViewfinderParameter, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettingsControl::setViewfinderParameter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_SetViewfinderParameter(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value))
	}
}

func (ptr *QCameraViewfinderSettingsControl) ViewfinderParameter(parameter QCameraViewfinderSettingsControl__ViewfinderParameter) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettingsControl::viewfinderParameter")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QCameraViewfinderSettingsControl_ViewfinderParameter(ptr.Pointer(), C.int(parameter)))
	}
	return nil
}

func (ptr *QCameraViewfinderSettingsControl) DestroyQCameraViewfinderSettingsControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraViewfinderSettingsControl::~QCameraViewfinderSettingsControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
