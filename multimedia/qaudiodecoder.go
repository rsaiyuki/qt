package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAudioDecoder struct {
	QMediaObject
}

type QAudioDecoder_ITF interface {
	QMediaObject_ITF
	QAudioDecoder_PTR() *QAudioDecoder
}

func PointerFromQAudioDecoder(ptr QAudioDecoder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioDecoder_PTR().Pointer()
	}
	return nil
}

func NewQAudioDecoderFromPointer(ptr unsafe.Pointer) *QAudioDecoder {
	var n = new(QAudioDecoder)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioDecoder_") {
		n.SetObjectName("QAudioDecoder_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioDecoder) QAudioDecoder_PTR() *QAudioDecoder {
	return ptr
}

//QAudioDecoder::Error
type QAudioDecoder__Error int64

const (
	QAudioDecoder__NoError             = QAudioDecoder__Error(0)
	QAudioDecoder__ResourceError       = QAudioDecoder__Error(1)
	QAudioDecoder__FormatError         = QAudioDecoder__Error(2)
	QAudioDecoder__AccessDeniedError   = QAudioDecoder__Error(3)
	QAudioDecoder__ServiceMissingError = QAudioDecoder__Error(4)
)

//QAudioDecoder::State
type QAudioDecoder__State int64

const (
	QAudioDecoder__StoppedState  = QAudioDecoder__State(0)
	QAudioDecoder__DecodingState = QAudioDecoder__State(1)
)

func (ptr *QAudioDecoder) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoder_ErrorString(ptr.Pointer()))
	}
	return ""
}

func NewQAudioDecoder(parent core.QObject_ITF) *QAudioDecoder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::QAudioDecoder")
		}
	}()

	return NewQAudioDecoderFromPointer(C.QAudioDecoder_NewQAudioDecoder(core.PointerFromQObject(parent)))
}

func (ptr *QAudioDecoder) BufferAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::bufferAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAudioDecoder_BufferAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAudioDecoder) ConnectBufferAvailableChanged(f func(available bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::bufferAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectBufferAvailableChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferAvailableChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectBufferAvailableChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::bufferAvailableChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectBufferAvailableChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferAvailableChanged")
	}
}

//export callbackQAudioDecoderBufferAvailableChanged
func callbackQAudioDecoderBufferAvailableChanged(ptrName *C.char, available C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::bufferAvailableChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "bufferAvailableChanged").(func(bool))(int(available) != 0)
}

func (ptr *QAudioDecoder) ConnectBufferReady(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::bufferReady")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectBufferReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bufferReady", f)
	}
}

func (ptr *QAudioDecoder) DisconnectBufferReady() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::bufferReady")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectBufferReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bufferReady")
	}
}

//export callbackQAudioDecoderBufferReady
func callbackQAudioDecoderBufferReady(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::bufferReady")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "bufferReady").(func())()
}

func (ptr *QAudioDecoder) Error() QAudioDecoder__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QAudioDecoder__Error(C.QAudioDecoder_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioDecoder) ConnectFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QAudioDecoder) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQAudioDecoderFinished
func callbackQAudioDecoderFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QAudioDecoder) SetAudioFormat(format QAudioFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::setAudioFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetAudioFormat(ptr.Pointer(), PointerFromQAudioFormat(format))
	}
}

func (ptr *QAudioDecoder) SetSourceDevice(device core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::setSourceDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetSourceDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioDecoder) SetSourceFilename(fileName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::setSourceFilename")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_SetSourceFilename(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QAudioDecoder) ConnectSourceChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::sourceChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectSourceChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::sourceChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQAudioDecoderSourceChanged
func callbackQAudioDecoderSourceChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::sourceChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sourceChanged").(func())()
}

func (ptr *QAudioDecoder) SourceDevice() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::sourceDevice")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioDecoder_SourceDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioDecoder) SourceFilename() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::sourceFilename")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioDecoder_SourceFilename(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioDecoder) Start() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_Start(ptr.Pointer())
	}
}

func (ptr *QAudioDecoder) ConnectStateChanged(f func(state QAudioDecoder__State)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAudioDecoder) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAudioDecoderStateChanged
func callbackQAudioDecoderStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QAudioDecoder__State))(QAudioDecoder__State(state))
}

func (ptr *QAudioDecoder) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioDecoder) DestroyQAudioDecoder() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioDecoder::~QAudioDecoder")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioDecoder_DestroyQAudioDecoder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
