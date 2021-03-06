package multimediawidgets

//#include "multimediawidgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"log"
	"unsafe"
)

type QVideoWidgetControl struct {
	multimedia.QMediaControl
}

type QVideoWidgetControl_ITF interface {
	multimedia.QMediaControl_ITF
	QVideoWidgetControl_PTR() *QVideoWidgetControl
}

func PointerFromQVideoWidgetControl(ptr QVideoWidgetControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWidgetControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoWidgetControlFromPointer(ptr unsafe.Pointer) *QVideoWidgetControl {
	var n = new(QVideoWidgetControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWidgetControl_") {
		n.SetObjectName("QVideoWidgetControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoWidgetControl) QVideoWidgetControl_PTR() *QVideoWidgetControl {
	return ptr
}

func (ptr *QVideoWidgetControl) AspectRatioMode() core.Qt__AspectRatioMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::aspectRatioMode")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWidgetControl_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) Brightness() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::brightness")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Brightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectBrightnessChanged(f func(brightness int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::brightnessChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectBrightnessChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::brightnessChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWidgetControlBrightnessChanged
func callbackQVideoWidgetControlBrightnessChanged(ptrName *C.char, brightness C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::brightnessChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "brightnessChanged").(func(int))(int(brightness))
}

func (ptr *QVideoWidgetControl) Contrast() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::contrast")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectContrastChanged(f func(contrast int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::contrastChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectContrastChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::contrastChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWidgetControlContrastChanged
func callbackQVideoWidgetControlContrastChanged(ptrName *C.char, contrast C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::contrastChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "contrastChanged").(func(int))(int(contrast))
}

func (ptr *QVideoWidgetControl) ConnectFullScreenChanged(f func(fullScreen bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::fullScreenChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectFullScreenChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::fullScreenChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWidgetControlFullScreenChanged
func callbackQVideoWidgetControlFullScreenChanged(ptrName *C.char, fullScreen C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::fullScreenChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "fullScreenChanged").(func(bool))(int(fullScreen) != 0)
}

func (ptr *QVideoWidgetControl) Hue() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::hue")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectHueChanged(f func(hue int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::hueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectHueChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::hueChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWidgetControlHueChanged
func callbackQVideoWidgetControlHueChanged(ptrName *C.char, hue C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::hueChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "hueChanged").(func(int))(int(hue))
}

func (ptr *QVideoWidgetControl) IsFullScreen() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::isFullScreen")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWidgetControl) Saturation() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::saturation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectSaturationChanged(f func(saturation int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::saturationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSaturationChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::saturationChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWidgetControlSaturationChanged
func callbackQVideoWidgetControlSaturationChanged(ptrName *C.char, saturation C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::saturationChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "saturationChanged").(func(int))(int(saturation))
}

func (ptr *QVideoWidgetControl) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::setAspectRatioMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoWidgetControl) SetBrightness(brightness int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::setBrightness")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidgetControl) SetContrast(contrast int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::setContrast")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidgetControl) SetFullScreen(fullScreen bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::setFullScreen")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidgetControl) SetHue(hue int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::setHue")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidgetControl) SetSaturation(saturation int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::setSaturation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWidgetControl) VideoWidget() *widgets.QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::videoWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QVideoWidgetControl_VideoWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidgetControl) DestroyQVideoWidgetControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoWidgetControl::~QVideoWidgetControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DestroyQVideoWidgetControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
