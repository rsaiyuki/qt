package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAudioOutput struct {
	core.QObject
}

type QAudioOutput_ITF interface {
	core.QObject_ITF
	QAudioOutput_PTR() *QAudioOutput
}

func PointerFromQAudioOutput(ptr QAudioOutput_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioOutput_PTR().Pointer()
	}
	return nil
}

func NewQAudioOutputFromPointer(ptr unsafe.Pointer) *QAudioOutput {
	var n = new(QAudioOutput)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioOutput_") {
		n.SetObjectName("QAudioOutput_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioOutput) QAudioOutput_PTR() *QAudioOutput {
	return ptr
}

func NewQAudioOutput2(audioDevice QAudioDeviceInfo_ITF, format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioOutput {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::QAudioOutput")
		}
	}()

	return NewQAudioOutputFromPointer(C.QAudioOutput_NewQAudioOutput2(PointerFromQAudioDeviceInfo(audioDevice), PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func NewQAudioOutput(format QAudioFormat_ITF, parent core.QObject_ITF) *QAudioOutput {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::QAudioOutput")
		}
	}()

	return NewQAudioOutputFromPointer(C.QAudioOutput_NewQAudioOutput(PointerFromQAudioFormat(format), core.PointerFromQObject(parent)))
}

func (ptr *QAudioOutput) BufferSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::bufferSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) BytesFree() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::bytesFree")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_BytesFree(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Category() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::category")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioOutput_Category(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioOutput) ConnectNotify(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::notify")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_ConnectNotify(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "notify", f)
	}
}

func (ptr *QAudioOutput) DisconnectNotify() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::notify")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_DisconnectNotify(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "notify")
	}
}

//export callbackQAudioOutputNotify
func callbackQAudioOutputNotify(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::notify")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "notify").(func())()
}

func (ptr *QAudioOutput) NotifyInterval() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::notifyInterval")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_NotifyInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) PeriodSize() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::periodSize")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAudioOutput_PeriodSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) Reset() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::reset")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_Reset(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Resume() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::resume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_Resume(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) SetBufferSize(value int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::setBufferSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetBufferSize(ptr.Pointer(), C.int(value))
	}
}

func (ptr *QAudioOutput) SetCategory(category string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::setCategory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetCategory(ptr.Pointer(), C.CString(category))
	}
}

func (ptr *QAudioOutput) SetNotifyInterval(ms int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::setNotifyInterval")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetNotifyInterval(ptr.Pointer(), C.int(ms))
	}
}

func (ptr *QAudioOutput) SetVolume(volume float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::setVolume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QAudioOutput) Start2() *core.QIODevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::start")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAudioOutput_Start2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAudioOutput) Start(device core.QIODevice_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_Start(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QAudioOutput) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_Stop(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Suspend() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::suspend")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_Suspend(ptr.Pointer())
	}
}

func (ptr *QAudioOutput) Volume() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::volume")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QAudioOutput_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAudioOutput) DestroyQAudioOutput() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioOutput::~QAudioOutput")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioOutput_DestroyQAudioOutput(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
