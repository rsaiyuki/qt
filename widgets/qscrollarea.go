package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QScrollArea struct {
	QAbstractScrollArea
}

type QScrollArea_ITF interface {
	QAbstractScrollArea_ITF
	QScrollArea_PTR() *QScrollArea
}

func PointerFromQScrollArea(ptr QScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQScrollAreaFromPointer(ptr unsafe.Pointer) *QScrollArea {
	var n = new(QScrollArea)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScrollArea_") {
		n.SetObjectName("QScrollArea_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScrollArea) QScrollArea_PTR() *QScrollArea {
	return ptr
}

func (ptr *QScrollArea) Alignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QScrollArea_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScrollArea) SetAlignment(v core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollArea_SetAlignment(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QScrollArea) SetWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::setWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QScrollArea) SetWidgetResizable(resizable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::setWidgetResizable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidgetResizable(ptr.Pointer(), C.int(qt.GoBoolToInt(resizable)))
	}
}

func (ptr *QScrollArea) WidgetResizable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::widgetResizable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QScrollArea_WidgetResizable(ptr.Pointer()) != 0
	}
	return false
}

func NewQScrollArea(parent QWidget_ITF) *QScrollArea {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::QScrollArea")
		}
	}()

	return NewQScrollAreaFromPointer(C.QScrollArea_NewQScrollArea(PointerFromQWidget(parent)))
}

func (ptr *QScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::ensureVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollArea_EnsureVisible(ptr.Pointer(), C.int(x), C.int(y), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QScrollArea) EnsureWidgetVisible(childWidget QWidget_ITF, xmargin int, ymargin int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::ensureWidgetVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollArea_EnsureWidgetVisible(ptr.Pointer(), PointerFromQWidget(childWidget), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QScrollArea) FocusNextPrevChild(next bool) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::focusNextPrevChild")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QScrollArea_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QScrollArea) TakeWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::takeWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QScrollArea_TakeWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) Widget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QScrollArea_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) DestroyQScrollArea() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScrollArea::~QScrollArea")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScrollArea_DestroyQScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
