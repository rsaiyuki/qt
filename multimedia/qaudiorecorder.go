package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QAudioRecorder struct {
	QMediaRecorder
}

type QAudioRecorder_ITF interface {
	QMediaRecorder_ITF
	QAudioRecorder_PTR() *QAudioRecorder
}

func PointerFromQAudioRecorder(ptr QAudioRecorder_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioRecorder_PTR().Pointer()
	}
	return nil
}

func NewQAudioRecorderFromPointer(ptr unsafe.Pointer) *QAudioRecorder {
	var n = new(QAudioRecorder)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioRecorder_") {
		n.SetObjectName("QAudioRecorder_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAudioRecorder) QAudioRecorder_PTR() *QAudioRecorder {
	return ptr
}

func NewQAudioRecorder(parent core.QObject_ITF) *QAudioRecorder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::QAudioRecorder")
		}
	}()

	return NewQAudioRecorderFromPointer(C.QAudioRecorder_NewQAudioRecorder(core.PointerFromQObject(parent)))
}

func (ptr *QAudioRecorder) AudioInput() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::audioInput")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_AudioInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioRecorder) ConnectAudioInputChanged(f func(name string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::audioInputChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ConnectAudioInputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "audioInputChanged", f)
	}
}

func (ptr *QAudioRecorder) DisconnectAudioInputChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::audioInputChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DisconnectAudioInputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "audioInputChanged")
	}
}

//export callbackQAudioRecorderAudioInputChanged
func callbackQAudioRecorderAudioInputChanged(ptrName *C.char, name *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::audioInputChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "audioInputChanged").(func(string))(C.GoString(name))
}

func (ptr *QAudioRecorder) AudioInputDescription(name string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::audioInputDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_AudioInputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioRecorder) AudioInputs() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::audioInputs")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAudioRecorder_AudioInputs(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QAudioRecorder) ConnectAvailableAudioInputsChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::availableAudioInputsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioRecorder_ConnectAvailableAudioInputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableAudioInputsChanged", f)
	}
}

func (ptr *QAudioRecorder) DisconnectAvailableAudioInputsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::availableAudioInputsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DisconnectAvailableAudioInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableAudioInputsChanged")
	}
}

//export callbackQAudioRecorderAvailableAudioInputsChanged
func callbackQAudioRecorderAvailableAudioInputsChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::availableAudioInputsChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "availableAudioInputsChanged").(func())()
}

func (ptr *QAudioRecorder) DefaultAudioInput() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::defaultAudioInput")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioRecorder_DefaultAudioInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioRecorder) SetAudioInput(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::setAudioInput")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioRecorder_SetAudioInput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioRecorder) DestroyQAudioRecorder() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAudioRecorder::~QAudioRecorder")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAudioRecorder_DestroyQAudioRecorder(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
