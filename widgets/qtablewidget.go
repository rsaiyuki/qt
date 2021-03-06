package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QTableWidget struct {
	QTableView
}

type QTableWidget_ITF interface {
	QTableView_ITF
	QTableWidget_PTR() *QTableWidget
}

func PointerFromQTableWidget(ptr QTableWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTableWidget_PTR().Pointer()
	}
	return nil
}

func NewQTableWidgetFromPointer(ptr unsafe.Pointer) *QTableWidget {
	var n = new(QTableWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTableWidget_") {
		n.SetObjectName("QTableWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTableWidget) QTableWidget_PTR() *QTableWidget {
	return ptr
}

func (ptr *QTableWidget) ItemAt(point core.QPoint_ITF) *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_ItemAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QTableWidget) ConnectCellActivated(f func(row int, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellActivated", f)
	}
}

func (ptr *QTableWidget) DisconnectCellActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellActivated")
	}
}

//export callbackQTableWidgetCellActivated
func callbackQTableWidgetCellActivated(ptrName *C.char, row C.int, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellActivated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cellActivated").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellChanged(f func(row int, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCellChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellChanged")
	}
}

//export callbackQTableWidgetCellChanged
func callbackQTableWidgetCellChanged(ptrName *C.char, row C.int, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cellChanged").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellClicked(f func(row int, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectCellClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellClicked")
	}
}

//export callbackQTableWidgetCellClicked
func callbackQTableWidgetCellClicked(ptrName *C.char, row C.int, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cellClicked").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellDoubleClicked(f func(row int, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellDoubleClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectCellDoubleClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellDoubleClicked")
	}
}

//export callbackQTableWidgetCellDoubleClicked
func callbackQTableWidgetCellDoubleClicked(ptrName *C.char, row C.int, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellDoubleClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cellDoubleClicked").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellEntered(f func(row int, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellEntered", f)
	}
}

func (ptr *QTableWidget) DisconnectCellEntered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellEntered")
	}
}

//export callbackQTableWidgetCellEntered
func callbackQTableWidgetCellEntered(ptrName *C.char, row C.int, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellEntered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cellEntered").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) ConnectCellPressed(f func(row int, column int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCellPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cellPressed", f)
	}
}

func (ptr *QTableWidget) DisconnectCellPressed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCellPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cellPressed")
	}
}

//export callbackQTableWidgetCellPressed
func callbackQTableWidgetCellPressed(ptrName *C.char, row C.int, column C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellPressed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cellPressed").(func(int, int))(int(row), int(column))
}

func (ptr *QTableWidget) CellWidget(row int, column int) *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::cellWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QTableWidget_CellWidget(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_Clear(ptr.Pointer())
	}
}

func (ptr *QTableWidget) ClearContents() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::clearContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ClearContents(ptr.Pointer())
	}
}

func (ptr *QTableWidget) ClosePersistentEditor(item QTableWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::closePersistentEditor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ClosePersistentEditor(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) Column(item QTableWidgetItem_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::column")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_Column(ptr.Pointer(), PointerFromQTableWidgetItem(item)))
	}
	return 0
}

func (ptr *QTableWidget) ColumnCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidget) ConnectCurrentCellChanged(f func(currentRow int, currentColumn int, previousRow int, previousColumn int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentCellChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCurrentCellChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentCellChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCurrentCellChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentCellChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCurrentCellChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentCellChanged")
	}
}

//export callbackQTableWidgetCurrentCellChanged
func callbackQTableWidgetCurrentCellChanged(ptrName *C.char, currentRow C.int, currentColumn C.int, previousRow C.int, previousColumn C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentCellChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentCellChanged").(func(int, int, int, int))(int(currentRow), int(currentColumn), int(previousRow), int(previousColumn))
}

func (ptr *QTableWidget) CurrentColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_CurrentColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidget) CurrentItem() *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_CurrentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidget) ConnectCurrentItemChanged(f func(current *QTableWidgetItem, previous *QTableWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentItemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectCurrentItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentItemChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectCurrentItemChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentItemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectCurrentItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentItemChanged")
	}
}

//export callbackQTableWidgetCurrentItemChanged
func callbackQTableWidgetCurrentItemChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentItemChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentItemChanged").(func(*QTableWidgetItem, *QTableWidgetItem))(NewQTableWidgetItemFromPointer(current), NewQTableWidgetItemFromPointer(previous))
}

func (ptr *QTableWidget) CurrentRow() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::currentRow")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_CurrentRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidget) EditItem(item QTableWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::editItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_EditItem(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::horizontalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_HorizontalHeaderItem(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) InsertColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::insertColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_InsertColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableWidget) InsertRow(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::insertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_InsertRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableWidget) Item(row int, column int) *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::item")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_Item(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemActivated(f func(item *QTableWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemActivated", f)
	}
}

func (ptr *QTableWidget) DisconnectItemActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemActivated")
	}
}

//export callbackQTableWidgetItemActivated
func callbackQTableWidgetItemActivated(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemActivated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemActivated").(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ItemAt2(ax int, ay int) *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_ItemAt2(ptr.Pointer(), C.int(ax), C.int(ay)))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemChanged(f func(item *QTableWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectItemChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemChanged")
	}
}

//export callbackQTableWidgetItemChanged
func callbackQTableWidgetItemChanged(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemChanged").(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ConnectItemClicked(f func(item *QTableWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectItemClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemClicked")
	}
}

//export callbackQTableWidgetItemClicked
func callbackQTableWidgetItemClicked(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemClicked").(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ConnectItemDoubleClicked(f func(item *QTableWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemDoubleClicked", f)
	}
}

func (ptr *QTableWidget) DisconnectItemDoubleClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemDoubleClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemDoubleClicked")
	}
}

//export callbackQTableWidgetItemDoubleClicked
func callbackQTableWidgetItemDoubleClicked(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemDoubleClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemDoubleClicked").(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ConnectItemEntered(f func(item *QTableWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemEntered", f)
	}
}

func (ptr *QTableWidget) DisconnectItemEntered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemEntered")
	}
}

//export callbackQTableWidgetItemEntered
func callbackQTableWidgetItemEntered(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemEntered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemEntered").(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ConnectItemPressed(f func(item *QTableWidgetItem)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemPressed", f)
	}
}

func (ptr *QTableWidget) DisconnectItemPressed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemPressed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemPressed")
	}
}

//export callbackQTableWidgetItemPressed
func callbackQTableWidgetItemPressed(ptrName *C.char, item unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemPressed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemPressed").(func(*QTableWidgetItem))(NewQTableWidgetItemFromPointer(item))
}

func (ptr *QTableWidget) ItemPrototype() *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemPrototype")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_ItemPrototype(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTableWidget) ConnectItemSelectionChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemSelectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ConnectItemSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "itemSelectionChanged", f)
	}
}

func (ptr *QTableWidget) DisconnectItemSelectionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemSelectionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DisconnectItemSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "itemSelectionChanged")
	}
}

//export callbackQTableWidgetItemSelectionChanged
func callbackQTableWidgetItemSelectionChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::itemSelectionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "itemSelectionChanged").(func())()
}

func (ptr *QTableWidget) OpenPersistentEditor(item QTableWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::openPersistentEditor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_OpenPersistentEditor(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) RemoveCellWidget(row int, column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::removeCellWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveCellWidget(ptr.Pointer(), C.int(row), C.int(column))
	}
}

func (ptr *QTableWidget) RemoveColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::removeColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QTableWidget) RemoveRow(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::removeRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_RemoveRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QTableWidget) Row(item QTableWidgetItem_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::row")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_Row(ptr.Pointer(), PointerFromQTableWidgetItem(item)))
	}
	return 0
}

func (ptr *QTableWidget) RowCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_RowCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTableWidget) ScrollToItem(item QTableWidgetItem_ITF, hint QAbstractItemView__ScrollHint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::scrollToItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_ScrollToItem(ptr.Pointer(), PointerFromQTableWidgetItem(item), C.int(hint))
	}
}

func (ptr *QTableWidget) SetCellWidget(row int, column int, widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setCellWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCellWidget(ptr.Pointer(), C.int(row), C.int(column), PointerFromQWidget(widget))
	}
}

func (ptr *QTableWidget) SetColumnCount(columns int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setColumnCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetColumnCount(ptr.Pointer(), C.int(columns))
	}
}

func (ptr *QTableWidget) SetCurrentCell(row int, column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setCurrentCell")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentCell(ptr.Pointer(), C.int(row), C.int(column))
	}
}

func (ptr *QTableWidget) SetCurrentCell2(row int, column int, command core.QItemSelectionModel__SelectionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setCurrentCell")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentCell2(ptr.Pointer(), C.int(row), C.int(column), C.int(command))
	}
}

func (ptr *QTableWidget) SetCurrentItem(item QTableWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setCurrentItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentItem(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetCurrentItem2(item QTableWidgetItem_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setCurrentItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetCurrentItem2(ptr.Pointer(), PointerFromQTableWidgetItem(item), C.int(command))
	}
}

func (ptr *QTableWidget) SetHorizontalHeaderItem(column int, item QTableWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setHorizontalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetHorizontalHeaderItem(ptr.Pointer(), C.int(column), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetHorizontalHeaderLabels(labels []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setHorizontalHeaderLabels")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetHorizontalHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QTableWidget) SetItem(row int, column int, item QTableWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetItem(ptr.Pointer(), C.int(row), C.int(column), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetItemPrototype(item QTableWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setItemPrototype")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetItemPrototype(ptr.Pointer(), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetRangeSelected(ran QTableWidgetSelectionRange_ITF, sele bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setRangeSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetRangeSelected(ptr.Pointer(), PointerFromQTableWidgetSelectionRange(ran), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QTableWidget) SetRowCount(rows int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setRowCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetRowCount(ptr.Pointer(), C.int(rows))
	}
}

func (ptr *QTableWidget) SetVerticalHeaderItem(row int, item QTableWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setVerticalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetVerticalHeaderItem(ptr.Pointer(), C.int(row), PointerFromQTableWidgetItem(item))
	}
}

func (ptr *QTableWidget) SetVerticalHeaderLabels(labels []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::setVerticalHeaderLabels")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SetVerticalHeaderLabels(ptr.Pointer(), C.CString(strings.Join(labels, ",,,")))
	}
}

func (ptr *QTableWidget) SortItems(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::sortItems")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_SortItems(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::takeHorizontalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_TakeHorizontalHeaderItem(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) TakeItem(row int, column int) *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::takeItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_TakeItem(ptr.Pointer(), C.int(row), C.int(column)))
	}
	return nil
}

func (ptr *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::takeVerticalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_TakeVerticalHeaderItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::verticalHeaderItem")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTableWidgetItemFromPointer(C.QTableWidget_VerticalHeaderItem(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QTableWidget) VisualColumn(logicalColumn int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::visualColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_VisualColumn(ptr.Pointer(), C.int(logicalColumn)))
	}
	return 0
}

func (ptr *QTableWidget) VisualRow(logicalRow int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::visualRow")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTableWidget_VisualRow(ptr.Pointer(), C.int(logicalRow)))
	}
	return 0
}

func (ptr *QTableWidget) DestroyQTableWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTableWidget::~QTableWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTableWidget_DestroyQTableWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
