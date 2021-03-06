package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QCameraImageProcessing struct {
	core.QObject
}

type QCameraImageProcessing_ITF interface {
	core.QObject_ITF
	QCameraImageProcessing_PTR() *QCameraImageProcessing
}

func PointerFromQCameraImageProcessing(ptr QCameraImageProcessing_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraImageProcessing_PTR().Pointer()
	}
	return nil
}

func NewQCameraImageProcessingFromPointer(ptr unsafe.Pointer) *QCameraImageProcessing {
	var n = new(QCameraImageProcessing)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraImageProcessing_") {
		n.SetObjectName("QCameraImageProcessing_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCameraImageProcessing) QCameraImageProcessing_PTR() *QCameraImageProcessing {
	return ptr
}

//QCameraImageProcessing::ColorFilter
type QCameraImageProcessing__ColorFilter int64

const (
	QCameraImageProcessing__ColorFilterNone       = QCameraImageProcessing__ColorFilter(0)
	QCameraImageProcessing__ColorFilterGrayscale  = QCameraImageProcessing__ColorFilter(1)
	QCameraImageProcessing__ColorFilterNegative   = QCameraImageProcessing__ColorFilter(2)
	QCameraImageProcessing__ColorFilterSolarize   = QCameraImageProcessing__ColorFilter(3)
	QCameraImageProcessing__ColorFilterSepia      = QCameraImageProcessing__ColorFilter(4)
	QCameraImageProcessing__ColorFilterPosterize  = QCameraImageProcessing__ColorFilter(5)
	QCameraImageProcessing__ColorFilterWhiteboard = QCameraImageProcessing__ColorFilter(6)
	QCameraImageProcessing__ColorFilterBlackboard = QCameraImageProcessing__ColorFilter(7)
	QCameraImageProcessing__ColorFilterAqua       = QCameraImageProcessing__ColorFilter(8)
	QCameraImageProcessing__ColorFilterVendor     = QCameraImageProcessing__ColorFilter(1000)
)

//QCameraImageProcessing::WhiteBalanceMode
type QCameraImageProcessing__WhiteBalanceMode int64

const (
	QCameraImageProcessing__WhiteBalanceAuto        = QCameraImageProcessing__WhiteBalanceMode(0)
	QCameraImageProcessing__WhiteBalanceManual      = QCameraImageProcessing__WhiteBalanceMode(1)
	QCameraImageProcessing__WhiteBalanceSunlight    = QCameraImageProcessing__WhiteBalanceMode(2)
	QCameraImageProcessing__WhiteBalanceCloudy      = QCameraImageProcessing__WhiteBalanceMode(3)
	QCameraImageProcessing__WhiteBalanceShade       = QCameraImageProcessing__WhiteBalanceMode(4)
	QCameraImageProcessing__WhiteBalanceTungsten    = QCameraImageProcessing__WhiteBalanceMode(5)
	QCameraImageProcessing__WhiteBalanceFluorescent = QCameraImageProcessing__WhiteBalanceMode(6)
	QCameraImageProcessing__WhiteBalanceFlash       = QCameraImageProcessing__WhiteBalanceMode(7)
	QCameraImageProcessing__WhiteBalanceSunset      = QCameraImageProcessing__WhiteBalanceMode(8)
	QCameraImageProcessing__WhiteBalanceVendor      = QCameraImageProcessing__WhiteBalanceMode(1000)
)

func (ptr *QCameraImageProcessing) ColorFilter() QCameraImageProcessing__ColorFilter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::colorFilter")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraImageProcessing__ColorFilter(C.QCameraImageProcessing_ColorFilter(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) Contrast() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::contrast")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) DenoisingLevel() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::denoisingLevel")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_DenoisingLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) IsAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::isAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) IsColorFilterSupported(filter QCameraImageProcessing__ColorFilter) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::isColorFilterSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsColorFilterSupported(ptr.Pointer(), C.int(filter)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) IsWhiteBalanceModeSupported(mode QCameraImageProcessing__WhiteBalanceMode) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::isWhiteBalanceModeSupported")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCameraImageProcessing_IsWhiteBalanceModeSupported(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QCameraImageProcessing) ManualWhiteBalance() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::manualWhiteBalance")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_ManualWhiteBalance(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) Saturation() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::saturation")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) SetColorFilter(filter QCameraImageProcessing__ColorFilter) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::setColorFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetColorFilter(ptr.Pointer(), C.int(filter))
	}
}

func (ptr *QCameraImageProcessing) SetContrast(value float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::setContrast")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetContrast(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QCameraImageProcessing) SetDenoisingLevel(level float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::setDenoisingLevel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetDenoisingLevel(ptr.Pointer(), C.double(level))
	}
}

func (ptr *QCameraImageProcessing) SetManualWhiteBalance(colorTemperature float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::setManualWhiteBalance")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetManualWhiteBalance(ptr.Pointer(), C.double(colorTemperature))
	}
}

func (ptr *QCameraImageProcessing) SetSaturation(value float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::setSaturation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetSaturation(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QCameraImageProcessing) SetSharpeningLevel(level float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::setSharpeningLevel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetSharpeningLevel(ptr.Pointer(), C.double(level))
	}
}

func (ptr *QCameraImageProcessing) SetWhiteBalanceMode(mode QCameraImageProcessing__WhiteBalanceMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::setWhiteBalanceMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCameraImageProcessing_SetWhiteBalanceMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCameraImageProcessing) SharpeningLevel() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::sharpeningLevel")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QCameraImageProcessing_SharpeningLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraImageProcessing) WhiteBalanceMode() QCameraImageProcessing__WhiteBalanceMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCameraImageProcessing::whiteBalanceMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QCameraImageProcessing__WhiteBalanceMode(C.QCameraImageProcessing_WhiteBalanceMode(ptr.Pointer()))
	}
	return 0
}
