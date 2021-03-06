package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QFormLayout struct {
	QLayout
}

type QFormLayout_ITF interface {
	QLayout_ITF
	QFormLayout_PTR() *QFormLayout
}

func PointerFromQFormLayout(ptr QFormLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFormLayout_PTR().Pointer()
	}
	return nil
}

func NewQFormLayoutFromPointer(ptr unsafe.Pointer) *QFormLayout {
	var n = new(QFormLayout)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFormLayout_") {
		n.SetObjectName("QFormLayout_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFormLayout) QFormLayout_PTR() *QFormLayout {
	return ptr
}

//QFormLayout::FieldGrowthPolicy
type QFormLayout__FieldGrowthPolicy int64

const (
	QFormLayout__FieldsStayAtSizeHint  = QFormLayout__FieldGrowthPolicy(0)
	QFormLayout__ExpandingFieldsGrow   = QFormLayout__FieldGrowthPolicy(1)
	QFormLayout__AllNonFixedFieldsGrow = QFormLayout__FieldGrowthPolicy(2)
)

//QFormLayout::ItemRole
type QFormLayout__ItemRole int64

const (
	QFormLayout__LabelRole    = QFormLayout__ItemRole(0)
	QFormLayout__FieldRole    = QFormLayout__ItemRole(1)
	QFormLayout__SpanningRole = QFormLayout__ItemRole(2)
)

//QFormLayout::RowWrapPolicy
type QFormLayout__RowWrapPolicy int64

const (
	QFormLayout__DontWrapRows = QFormLayout__RowWrapPolicy(0)
	QFormLayout__WrapLongRows = QFormLayout__RowWrapPolicy(1)
	QFormLayout__WrapAllRows  = QFormLayout__RowWrapPolicy(2)
)

func (ptr *QFormLayout) FieldGrowthPolicy() QFormLayout__FieldGrowthPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::fieldGrowthPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QFormLayout__FieldGrowthPolicy(C.QFormLayout_FieldGrowthPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) FormAlignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::formAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QFormLayout_FormAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) HorizontalSpacing() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::horizontalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFormLayout_HorizontalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) LabelAlignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::labelAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QFormLayout_LabelAlignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) RowWrapPolicy() QFormLayout__RowWrapPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::rowWrapPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QFormLayout__RowWrapPolicy(C.QFormLayout_RowWrapPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) SetFieldGrowthPolicy(policy QFormLayout__FieldGrowthPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setFieldGrowthPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetFieldGrowthPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QFormLayout) SetFormAlignment(alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setFormAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetFormAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QFormLayout) SetHorizontalSpacing(spacing int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setHorizontalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetHorizontalSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QFormLayout) SetLabelAlignment(alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setLabelAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetLabelAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QFormLayout) SetRowWrapPolicy(policy QFormLayout__RowWrapPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setRowWrapPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetRowWrapPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QFormLayout) SetVerticalSpacing(spacing int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setVerticalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetVerticalSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QFormLayout) VerticalSpacing() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::verticalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFormLayout_VerticalSpacing(ptr.Pointer()))
	}
	return 0
}

func NewQFormLayout(parent QWidget_ITF) *QFormLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::QFormLayout")
		}
	}()

	return NewQFormLayoutFromPointer(C.QFormLayout_NewQFormLayout(PointerFromQWidget(parent)))
}

func (ptr *QFormLayout) AddItem(item QLayoutItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_AddItem(ptr.Pointer(), PointerFromQLayoutItem(item))
	}
}

func (ptr *QFormLayout) AddRow6(layout QLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::addRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow6(ptr.Pointer(), PointerFromQLayout(layout))
	}
}

func (ptr *QFormLayout) AddRow2(label QWidget_ITF, field QLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::addRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow2(ptr.Pointer(), PointerFromQWidget(label), PointerFromQLayout(field))
	}
}

func (ptr *QFormLayout) AddRow(label QWidget_ITF, field QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::addRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow(ptr.Pointer(), PointerFromQWidget(label), PointerFromQWidget(field))
	}
}

func (ptr *QFormLayout) AddRow5(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::addRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow5(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QFormLayout) AddRow4(labelText string, field QLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::addRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow4(ptr.Pointer(), C.CString(labelText), PointerFromQLayout(field))
	}
}

func (ptr *QFormLayout) AddRow3(labelText string, field QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::addRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_AddRow3(ptr.Pointer(), C.CString(labelText), PointerFromQWidget(field))
	}
}

func (ptr *QFormLayout) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFormLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) ExpandingDirections() core.Qt__Orientation {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::expandingDirections")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QFormLayout_ExpandingDirections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) HasHeightForWidth() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::hasHeightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFormLayout_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFormLayout) HeightForWidth(width int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::heightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFormLayout_HeightForWidth(ptr.Pointer(), C.int(width)))
	}
	return 0
}

func (ptr *QFormLayout) InsertRow6(row int, layout QLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow6(ptr.Pointer(), C.int(row), PointerFromQLayout(layout))
	}
}

func (ptr *QFormLayout) InsertRow2(row int, label QWidget_ITF, field QLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow2(ptr.Pointer(), C.int(row), PointerFromQWidget(label), PointerFromQLayout(field))
	}
}

func (ptr *QFormLayout) InsertRow(row int, label QWidget_ITF, field QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow(ptr.Pointer(), C.int(row), PointerFromQWidget(label), PointerFromQWidget(field))
	}
}

func (ptr *QFormLayout) InsertRow5(row int, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow5(ptr.Pointer(), C.int(row), PointerFromQWidget(widget))
	}
}

func (ptr *QFormLayout) InsertRow4(row int, labelText string, field QLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow4(ptr.Pointer(), C.int(row), C.CString(labelText), PointerFromQLayout(field))
	}
}

func (ptr *QFormLayout) InsertRow3(row int, labelText string, field QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_InsertRow3(ptr.Pointer(), C.int(row), C.CString(labelText), PointerFromQWidget(field))
	}
}

func (ptr *QFormLayout) Invalidate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::invalidate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QFormLayout) ItemAt2(index int) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QFormLayout_ItemAt2(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QFormLayout) ItemAt(row int, role QFormLayout__ItemRole) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QFormLayout_ItemAt(ptr.Pointer(), C.int(row), C.int(role)))
	}
	return nil
}

func (ptr *QFormLayout) LabelForField2(field QLayout_ITF) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::labelForField")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QFormLayout_LabelForField2(ptr.Pointer(), PointerFromQLayout(field)))
	}
	return nil
}

func (ptr *QFormLayout) LabelForField(field QWidget_ITF) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::labelForField")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QFormLayout_LabelForField(ptr.Pointer(), PointerFromQWidget(field)))
	}
	return nil
}

func (ptr *QFormLayout) RowCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFormLayout_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) SetGeometry(rect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QFormLayout) SetItem(row int, role QFormLayout__ItemRole, item QLayoutItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetItem(ptr.Pointer(), C.int(row), C.int(role), PointerFromQLayoutItem(item))
	}
}

func (ptr *QFormLayout) SetLayout(row int, role QFormLayout__ItemRole, layout QLayout_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetLayout(ptr.Pointer(), C.int(row), C.int(role), PointerFromQLayout(layout))
	}
}

func (ptr *QFormLayout) SetSpacing(spacing int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetSpacing(ptr.Pointer(), C.int(spacing))
	}
}

func (ptr *QFormLayout) SetWidget(row int, role QFormLayout__ItemRole, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::setWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_SetWidget(ptr.Pointer(), C.int(row), C.int(role), PointerFromQWidget(widget))
	}
}

func (ptr *QFormLayout) Spacing() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::spacing")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFormLayout_Spacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFormLayout) TakeAt(index int) *QLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::takeAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQLayoutItemFromPointer(C.QFormLayout_TakeAt(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QFormLayout) DestroyQFormLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFormLayout::~QFormLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFormLayout_DestroyQFormLayout(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
