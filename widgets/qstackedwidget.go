package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QStackedWidget struct {
	QFrame
}

type QStackedWidget_ITF interface {
	QFrame_ITF
	QStackedWidget_PTR() *QStackedWidget
}

func PointerFromQStackedWidget(ptr QStackedWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackedWidget_PTR().Pointer()
	}
	return nil
}

func NewQStackedWidgetFromPointer(ptr unsafe.Pointer) *QStackedWidget {
	var n = new(QStackedWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStackedWidget_") {
		n.SetObjectName("QStackedWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStackedWidget) QStackedWidget_PTR() *QStackedWidget {
	return ptr
}

func (ptr *QStackedWidget) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedWidget) CurrentIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QStackedWidget) SetCurrentIndex(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::setCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedWidget_SetCurrentIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QStackedWidget) SetCurrentWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::setCurrentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedWidget_SetCurrentWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func NewQStackedWidget(parent QWidget_ITF) *QStackedWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::QStackedWidget")
		}
	}()

	return NewQStackedWidgetFromPointer(C.QStackedWidget_NewQStackedWidget(PointerFromQWidget(parent)))
}

func (ptr *QStackedWidget) AddWidget(widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::addWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_AddWidget(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedWidget) ConnectCurrentChanged(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedWidget_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QStackedWidget) DisconnectCurrentChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedWidget_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQStackedWidgetCurrentChanged
func callbackQStackedWidgetCurrentChanged(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::currentChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(int))(int(index))
}

func (ptr *QStackedWidget) CurrentWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::currentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedWidget_CurrentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QStackedWidget) IndexOf(widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::indexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedWidget) InsertWidget(index int, widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::insertWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QStackedWidget_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStackedWidget) RemoveWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::removeWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedWidget_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStackedWidget) Widget(index int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QStackedWidget_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QStackedWidget) ConnectWidgetRemoved(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::widgetRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedWidget_ConnectWidgetRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "widgetRemoved", f)
	}
}

func (ptr *QStackedWidget) DisconnectWidgetRemoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::widgetRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedWidget_DisconnectWidgetRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "widgetRemoved")
	}
}

//export callbackQStackedWidgetWidgetRemoved
func callbackQStackedWidgetWidgetRemoved(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::widgetRemoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "widgetRemoved").(func(int))(int(index))
}

func (ptr *QStackedWidget) DestroyQStackedWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStackedWidget::~QStackedWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStackedWidget_DestroyQStackedWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
