package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QGraphicsBlurEffect struct {
	QGraphicsEffect
}

type QGraphicsBlurEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsBlurEffect_PTR() *QGraphicsBlurEffect
}

func PointerFromQGraphicsBlurEffect(ptr QGraphicsBlurEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsBlurEffect_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsBlurEffectFromPointer(ptr unsafe.Pointer) *QGraphicsBlurEffect {
	var n = new(QGraphicsBlurEffect)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsBlurEffect_") {
		n.SetObjectName("QGraphicsBlurEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsBlurEffect) QGraphicsBlurEffect_PTR() *QGraphicsBlurEffect {
	return ptr
}

//QGraphicsBlurEffect::BlurHint
type QGraphicsBlurEffect__BlurHint int64

const (
	QGraphicsBlurEffect__PerformanceHint = QGraphicsBlurEffect__BlurHint(0x00)
	QGraphicsBlurEffect__QualityHint     = QGraphicsBlurEffect__BlurHint(0x01)
	QGraphicsBlurEffect__AnimationHint   = QGraphicsBlurEffect__BlurHint(0x02)
)

func (ptr *QGraphicsBlurEffect) BlurHints() QGraphicsBlurEffect__BlurHint {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::blurHints")
		}
	}()

	if ptr.Pointer() != nil {
		return QGraphicsBlurEffect__BlurHint(C.QGraphicsBlurEffect_BlurHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsBlurEffect) BlurRadius() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::blurRadius")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsBlurEffect_BlurRadius(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsBlurEffect) SetBlurHints(hints QGraphicsBlurEffect__BlurHint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::setBlurHints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_SetBlurHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QGraphicsBlurEffect) SetBlurRadius(blurRadius float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::setBlurRadius")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_SetBlurRadius(ptr.Pointer(), C.double(blurRadius))
	}
}

func NewQGraphicsBlurEffect(parent core.QObject_ITF) *QGraphicsBlurEffect {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::QGraphicsBlurEffect")
		}
	}()

	return NewQGraphicsBlurEffectFromPointer(C.QGraphicsBlurEffect_NewQGraphicsBlurEffect(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsBlurEffect) ConnectBlurHintsChanged(f func(hints QGraphicsBlurEffect__BlurHint)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::blurHintsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_ConnectBlurHintsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "blurHintsChanged", f)
	}
}

func (ptr *QGraphicsBlurEffect) DisconnectBlurHintsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::blurHintsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_DisconnectBlurHintsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "blurHintsChanged")
	}
}

//export callbackQGraphicsBlurEffectBlurHintsChanged
func callbackQGraphicsBlurEffectBlurHintsChanged(ptrName *C.char, hints C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::blurHintsChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "blurHintsChanged").(func(QGraphicsBlurEffect__BlurHint))(QGraphicsBlurEffect__BlurHint(hints))
}

func (ptr *QGraphicsBlurEffect) DestroyQGraphicsBlurEffect() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsBlurEffect::~QGraphicsBlurEffect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_DestroyQGraphicsBlurEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
