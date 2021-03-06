package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QLayout struct {
	QLayoutItem
	core.QObject
}

type QLayout_ITF interface {
	QLayoutItem_ITF
	core.QObject_ITF
	QLayout_PTR() *QLayout
}

func (p *QLayout) Pointer() unsafe.Pointer {
	return p.QLayoutItem_PTR().Pointer()
}

func (p *QLayout) SetPointer(ptr unsafe.Pointer) {
	p.QLayoutItem_PTR().SetPointer(ptr)
	p.QObject_PTR().SetPointer(ptr)
}

func PointerFromQLayout(ptr QLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLayout_PTR().Pointer()
	}
	return nil
}

func NewQLayoutFromPointer(ptr unsafe.Pointer) *QLayout {
	var n = new(QLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLayout_") {
		n.SetObjectName("QLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLayout) QLayout_PTR() *QLayout {
	return ptr
}

//QLayout::SizeConstraint
type QLayout__SizeConstraint int64

const (
	QLayout__SetDefaultConstraint = QLayout__SizeConstraint(0)
	QLayout__SetNoConstraint      = QLayout__SizeConstraint(1)
	QLayout__SetMinimumSize       = QLayout__SizeConstraint(2)
	QLayout__SetFixedSize         = QLayout__SizeConstraint(3)
	QLayout__SetMaximumSize       = QLayout__SizeConstraint(4)
	QLayout__SetMinAndMaxSize     = QLayout__SizeConstraint(5)
)

func (ptr *QLayout) SetSizeConstraint(v QLayout__SizeConstraint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setSizeConstraint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_SetSizeConstraint(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLayout) SetSpacing(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_SetSpacing(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLayout) SizeConstraint() QLayout__SizeConstraint {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::sizeConstraint")
		}
	}()

	if ptr.Pointer() != nil {
		return QLayout__SizeConstraint(C.QLayout_SizeConstraint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Spacing() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::spacing")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Activate() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::activate")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLayout_Activate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) AddItem(item QLayoutItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QLayout) AddWidget(w QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::addWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_AddWidget(ptr.Pointer(), PointerFromQWidget(w))
	}
}

func (ptr *QLayout) ControlTypes() QSizePolicy__ControlType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::controlTypes")
		}
	}()

	if ptr.Pointer() != nil {
		return QSizePolicy__ControlType(C.QLayout_ControlTypes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) ExpandingDirections() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::expandingDirections")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLayout) GetContentsMargins(left int, top int, right int, bottom int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::getContentsMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_GetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLayout) IndexOf(widget QWidget_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::indexOf")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLayout_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QLayout) Invalidate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::invalidate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QLayout) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLayout_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) IsEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::isEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLayout_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLayout) ItemAt(index int) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_ItemAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QLayout) Layout() *QLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::layout")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutFromPointer(C.QLayout_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) MenuBar() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::menuBar")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayout_MenuBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) ParentWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::parentWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLayout_ParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLayout) RemoveItem(item QLayoutItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::removeItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_RemoveItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QLayout) RemoveWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::removeWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QLayout) ReplaceWidget(from QWidget_ITF, to QWidget_ITF, options core.Qt__FindChildOption) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::replaceWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_ReplaceWidget(ptr.Pointer(), PointerFromQWidget(from), PointerFromQWidget(to), C.int(options)))
	}
	return nil
}

func (ptr *QLayout) SetAlignment2(l QLayout_ITF, alignment core.Qt__AlignmentFlag) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLayout_SetAlignment2(ptr.Pointer(), PointerFromQLayout(l), C.int(alignment)) != 0
	}
	return false
}

func (ptr *QLayout) SetAlignment(w QWidget_ITF, alignment core.Qt__AlignmentFlag) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLayout_SetAlignment(ptr.Pointer(), PointerFromQWidget(w), C.int(alignment)) != 0
	}
	return false
}

func (ptr *QLayout) SetContentsMargins2(margins core.QMargins_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setContentsMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setContentsMargins")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_SetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QLayout) SetEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QLayout) SetGeometry(r core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(r))
	}
}

func (ptr *QLayout) SetMenuBar(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::setMenuBar")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_SetMenuBar(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QLayout) TakeAt(index int) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::takeAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QLayout) Update() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLayout::update")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLayout_Update(ptr.Pointer())
	}
}
