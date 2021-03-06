package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSound struct {
	core.QObject
}

type QSound_ITF interface {
	core.QObject_ITF
	QSound_PTR() *QSound
}

func PointerFromQSound(ptr QSound_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSound_PTR().Pointer()
	}
	return nil
}

func NewQSoundFromPointer(ptr unsafe.Pointer) *QSound {
	var n = new(QSound)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSound_") {
		n.SetObjectName("QSound_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSound) QSound_PTR() *QSound {
	return ptr
}

//QSound::Loop
type QSound__Loop int64

const (
	QSound__Infinite = QSound__Loop(-1)
)

func (ptr *QSound) SetLoops(number int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::setLoops")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSound_SetLoops(ptr.Pointer(), C.int(number))
	}
}

func NewQSound(filename string, parent core.QObject_ITF) *QSound {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::QSound")
		}
	}()

	return NewQSoundFromPointer(C.QSound_NewQSound(C.CString(filename), core.PointerFromQObject(parent)))
}

func (ptr *QSound) FileName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::fileName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSound_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSound) IsFinished() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::isFinished")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSound_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSound) Loops() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::loops")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSound_Loops(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSound) LoopsRemaining() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::loopsRemaining")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSound_LoopsRemaining(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSound) Play2() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::play")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSound_Play2(ptr.Pointer())
	}
}

func QSound_Play(filename string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::play")
		}
	}()

	C.QSound_QSound_Play(C.CString(filename))
}

func (ptr *QSound) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSound_Stop(ptr.Pointer())
	}
}

func (ptr *QSound) DestroyQSound() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSound::~QSound")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSound_DestroyQSound(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
