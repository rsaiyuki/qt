package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QGraphicsProxyWidget struct {
	QGraphicsWidget
}

type QGraphicsProxyWidget_ITF interface {
	QGraphicsWidget_ITF
	QGraphicsProxyWidget_PTR() *QGraphicsProxyWidget
}

func PointerFromQGraphicsProxyWidget(ptr QGraphicsProxyWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsProxyWidget_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsProxyWidgetFromPointer(ptr unsafe.Pointer) *QGraphicsProxyWidget {
	var n = new(QGraphicsProxyWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsProxyWidget_") {
		n.SetObjectName("QGraphicsProxyWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsProxyWidget) QGraphicsProxyWidget_PTR() *QGraphicsProxyWidget {
	return ptr
}

func NewQGraphicsProxyWidget(parent QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QGraphicsProxyWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsProxyWidget::QGraphicsProxyWidget")
		}
	}()

	return NewQGraphicsProxyWidgetFromPointer(C.QGraphicsProxyWidget_NewQGraphicsProxyWidget(PointerFromQGraphicsItem(parent), C.int(wFlags)))
}

func (ptr *QGraphicsProxyWidget) CreateProxyForChildWidget(child QWidget_ITF) *QGraphicsProxyWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsProxyWidget::createProxyForChildWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsProxyWidgetFromPointer(C.QGraphicsProxyWidget_CreateProxyForChildWidget(ptr.Pointer(), PointerFromQWidget(child)))
	}
	return nil
}

func (ptr *QGraphicsProxyWidget) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsProxyWidget::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsProxyWidget) SetGeometry(rect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsProxyWidget::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_SetGeometry(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsProxyWidget) SetWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsProxyWidget::setWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsProxyWidget) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsProxyWidget::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsProxyWidget_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsProxyWidget) Widget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsProxyWidget::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QGraphicsProxyWidget_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsProxyWidget) DestroyQGraphicsProxyWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsProxyWidget::~QGraphicsProxyWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DestroyQGraphicsProxyWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
