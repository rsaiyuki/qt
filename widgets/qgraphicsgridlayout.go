package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QGraphicsGridLayout struct {
	QGraphicsLayout
}

type QGraphicsGridLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsGridLayout_PTR() *QGraphicsGridLayout
}

func PointerFromQGraphicsGridLayout(ptr QGraphicsGridLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsGridLayout_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsGridLayoutFromPointer(ptr unsafe.Pointer) *QGraphicsGridLayout {
	var n = new(QGraphicsGridLayout)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsGridLayout) QGraphicsGridLayout_PTR() *QGraphicsGridLayout {
	return ptr
}

func NewQGraphicsGridLayout(parent QGraphicsLayoutItem_ITF) *QGraphicsGridLayout {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::QGraphicsGridLayout")
		}
	}()

	return NewQGraphicsGridLayoutFromPointer(C.QGraphicsGridLayout_NewQGraphicsGridLayout(PointerFromQGraphicsLayoutItem(parent)))
}

func (ptr *QGraphicsGridLayout) AddItem2(item QGraphicsLayoutItem_ITF, row int, column int, alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_AddItem2(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(row), C.int(column), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) AddItem(item QGraphicsLayoutItem_ITF, row int, column int, rowSpan int, columnSpan int, alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::addItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_AddItem(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(row), C.int(column), C.int(rowSpan), C.int(columnSpan), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) Alignment(item QGraphicsLayoutItem_ITF) core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_Alignment(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnAlignment(column int) core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::columnAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_ColumnAlignment(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnMaximumWidth(column int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::columnMaximumWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_ColumnMaximumWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnMinimumWidth(column int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::columnMinimumWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_ColumnMinimumWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnPreferredWidth(column int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::columnPreferredWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_ColumnPreferredWidth(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnSpacing(column int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::columnSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_ColumnSpacing(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::columnStretchFactor")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_ColumnStretchFactor(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) HorizontalSpacing() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::horizontalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_HorizontalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) Invalidate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::invalidate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_Invalidate(ptr.Pointer())
	}
}

func (ptr *QGraphicsGridLayout) ItemAt2(index int) *QGraphicsLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsGridLayout_ItemAt2(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QGraphicsGridLayout) ItemAt(row int, column int) *QGraphicsLayoutItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutItemFromPointer(C.QGraphicsGridLayout_ItemAt(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QGraphicsGridLayout) RemoveAt(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::removeAt")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_RemoveAt(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QGraphicsGridLayout) RemoveItem(item QGraphicsLayoutItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::removeItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_RemoveItem(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item))
	}
}

func (ptr *QGraphicsGridLayout) RowAlignment(row int) core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::rowAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsGridLayout_RowAlignment(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowMaximumHeight(row int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::rowMaximumHeight")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_RowMaximumHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowMinimumHeight(row int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::rowMinimumHeight")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_RowMinimumHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowPreferredHeight(row int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::rowPreferredHeight")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_RowPreferredHeight(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowSpacing(row int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::rowSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_RowSpacing(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) RowStretchFactor(row int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::rowStretchFactor")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QGraphicsGridLayout_RowStretchFactor(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) SetAlignment(item QGraphicsLayoutItem_ITF, alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetAlignment(ptr.Pointer(), PointerFromQGraphicsLayoutItem(item), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnAlignment(column int, alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setColumnAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnAlignment(ptr.Pointer(), C.int(column), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnFixedWidth(column int, width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setColumnFixedWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnFixedWidth(ptr.Pointer(), C.int(column), C.double(width))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnMaximumWidth(column int, width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setColumnMaximumWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnMaximumWidth(ptr.Pointer(), C.int(column), C.double(width))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnMinimumWidth(column int, width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setColumnMinimumWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnMinimumWidth(ptr.Pointer(), C.int(column), C.double(width))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnPreferredWidth(column int, width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setColumnPreferredWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnPreferredWidth(ptr.Pointer(), C.int(column), C.double(width))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnSpacing(column int, spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setColumnSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnSpacing(ptr.Pointer(), C.int(column), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setColumnStretchFactor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetColumnStretchFactor(ptr.Pointer(), C.int(column), C.int(stretch))
	}
}

func (ptr *QGraphicsGridLayout) SetGeometry(rect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setGeometry")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetGeometry(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setHorizontalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetHorizontalSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) SetRowAlignment(row int, alignment core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setRowAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowAlignment(ptr.Pointer(), C.int(row), C.int(alignment))
	}
}

func (ptr *QGraphicsGridLayout) SetRowFixedHeight(row int, height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setRowFixedHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowFixedHeight(ptr.Pointer(), C.int(row), C.double(height))
	}
}

func (ptr *QGraphicsGridLayout) SetRowMaximumHeight(row int, height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setRowMaximumHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowMaximumHeight(ptr.Pointer(), C.int(row), C.double(height))
	}
}

func (ptr *QGraphicsGridLayout) SetRowMinimumHeight(row int, height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setRowMinimumHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowMinimumHeight(ptr.Pointer(), C.int(row), C.double(height))
	}
}

func (ptr *QGraphicsGridLayout) SetRowPreferredHeight(row int, height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setRowPreferredHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowPreferredHeight(ptr.Pointer(), C.int(row), C.double(height))
	}
}

func (ptr *QGraphicsGridLayout) SetRowSpacing(row int, spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setRowSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowSpacing(ptr.Pointer(), C.int(row), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setRowStretchFactor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetRowStretchFactor(ptr.Pointer(), C.int(row), C.int(stretch))
	}
}

func (ptr *QGraphicsGridLayout) SetSpacing(spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) SetVerticalSpacing(spacing float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::setVerticalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_SetVerticalSpacing(ptr.Pointer(), C.double(spacing))
	}
}

func (ptr *QGraphicsGridLayout) VerticalSpacing() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::verticalSpacing")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsGridLayout_VerticalSpacing(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsGridLayout) DestroyQGraphicsGridLayout() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsGridLayout::~QGraphicsGridLayout")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsGridLayout_DestroyQGraphicsGridLayout(ptr.Pointer())
	}
}
