package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QItemEditorCreatorBase struct {
	ptr unsafe.Pointer
}

type QItemEditorCreatorBase_ITF interface {
	QItemEditorCreatorBase_PTR() *QItemEditorCreatorBase
}

func (p *QItemEditorCreatorBase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QItemEditorCreatorBase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQItemEditorCreatorBase(ptr QItemEditorCreatorBase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QItemEditorCreatorBase_PTR().Pointer()
	}
	return nil
}

func NewQItemEditorCreatorBaseFromPointer(ptr unsafe.Pointer) *QItemEditorCreatorBase {
	var n = new(QItemEditorCreatorBase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QItemEditorCreatorBase) QItemEditorCreatorBase_PTR() *QItemEditorCreatorBase {
	return ptr
}

func (ptr *QItemEditorCreatorBase) DestroyQItemEditorCreatorBase() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorCreatorBase::~QItemEditorCreatorBase")
		}
	}()

	if ptr.Pointer() != nil {
		C.QItemEditorCreatorBase_DestroyQItemEditorCreatorBase(ptr.Pointer())
	}
}

func (ptr *QItemEditorCreatorBase) CreateWidget(parent QWidget_ITF) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorCreatorBase::createWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QItemEditorCreatorBase_CreateWidget(ptr.Pointer(), PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QItemEditorCreatorBase) ValuePropertyName() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QItemEditorCreatorBase::valuePropertyName")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QItemEditorCreatorBase_ValuePropertyName(ptr.Pointer()))
	}
	return nil
}
