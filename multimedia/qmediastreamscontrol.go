package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QMediaStreamsControl struct {
	QMediaControl
}

type QMediaStreamsControl_ITF interface {
	QMediaControl_ITF
	QMediaStreamsControl_PTR() *QMediaStreamsControl
}

func PointerFromQMediaStreamsControl(ptr QMediaStreamsControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaStreamsControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaStreamsControlFromPointer(ptr unsafe.Pointer) *QMediaStreamsControl {
	var n = new(QMediaStreamsControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaStreamsControl_") {
		n.SetObjectName("QMediaStreamsControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaStreamsControl) QMediaStreamsControl_PTR() *QMediaStreamsControl {
	return ptr
}

//QMediaStreamsControl::StreamType
type QMediaStreamsControl__StreamType int64

const (
	QMediaStreamsControl__UnknownStream    = QMediaStreamsControl__StreamType(0)
	QMediaStreamsControl__VideoStream      = QMediaStreamsControl__StreamType(1)
	QMediaStreamsControl__AudioStream      = QMediaStreamsControl__StreamType(2)
	QMediaStreamsControl__SubPictureStream = QMediaStreamsControl__StreamType(3)
	QMediaStreamsControl__DataStream       = QMediaStreamsControl__StreamType(4)
)

func (ptr *QMediaStreamsControl) ConnectActiveStreamsChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::activeStreamsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ConnectActiveStreamsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeStreamsChanged", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectActiveStreamsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::activeStreamsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DisconnectActiveStreamsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeStreamsChanged")
	}
}

//export callbackQMediaStreamsControlActiveStreamsChanged
func callbackQMediaStreamsControlActiveStreamsChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::activeStreamsChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activeStreamsChanged").(func())()
}

func (ptr *QMediaStreamsControl) IsActive(stream int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaStreamsControl_IsActive(ptr.Pointer(), C.int(stream)) != 0
	}
	return false
}

func (ptr *QMediaStreamsControl) MetaData(stream int, key string) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::metaData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QMediaStreamsControl_MetaData(ptr.Pointer(), C.int(stream), C.CString(key)))
	}
	return nil
}

func (ptr *QMediaStreamsControl) SetActive(stream int, state bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::setActive")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_SetActive(ptr.Pointer(), C.int(stream), C.int(qt.GoBoolToInt(state)))
	}
}

func (ptr *QMediaStreamsControl) StreamCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::streamCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaStreamsControl_StreamCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaStreamsControl) StreamType(stream int) QMediaStreamsControl__StreamType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::streamType")
		}
	}()

	if ptr.Pointer() != nil {
		return QMediaStreamsControl__StreamType(C.QMediaStreamsControl_StreamType(ptr.Pointer(), C.int(stream)))
	}
	return 0
}

func (ptr *QMediaStreamsControl) ConnectStreamsChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::streamsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ConnectStreamsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "streamsChanged", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectStreamsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::streamsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DisconnectStreamsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "streamsChanged")
	}
}

//export callbackQMediaStreamsControlStreamsChanged
func callbackQMediaStreamsControlStreamsChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::streamsChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "streamsChanged").(func())()
}

func (ptr *QMediaStreamsControl) DestroyQMediaStreamsControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaStreamsControl::~QMediaStreamsControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DestroyQMediaStreamsControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
