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

type QCompleter struct {
	core.QObject
}

type QCompleter_ITF interface {
	core.QObject_ITF
	QCompleter_PTR() *QCompleter
}

func PointerFromQCompleter(ptr QCompleter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCompleter_PTR().Pointer()
	}
	return nil
}

func NewQCompleterFromPointer(ptr unsafe.Pointer) *QCompleter {
	var n = new(QCompleter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCompleter_") {
		n.SetObjectName("QCompleter_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCompleter) QCompleter_PTR() *QCompleter {
	return ptr
}

//QCompleter::CompletionMode
type QCompleter__CompletionMode int64

const (
	QCompleter__PopupCompletion           = QCompleter__CompletionMode(0)
	QCompleter__UnfilteredPopupCompletion = QCompleter__CompletionMode(1)
	QCompleter__InlineCompletion          = QCompleter__CompletionMode(2)
)

//QCompleter::ModelSorting
type QCompleter__ModelSorting int64

const (
	QCompleter__UnsortedModel                = QCompleter__ModelSorting(0)
	QCompleter__CaseSensitivelySortedModel   = QCompleter__ModelSorting(1)
	QCompleter__CaseInsensitivelySortedModel = QCompleter__ModelSorting(2)
)

func (ptr *QCompleter) CaseSensitivity() core.Qt__CaseSensitivity {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::caseSensitivity")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__CaseSensitivity(C.QCompleter_CaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) CompletionColumn() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::completionColumn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionColumn(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) CompletionMode() QCompleter__CompletionMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::completionMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QCompleter__CompletionMode(C.QCompleter_CompletionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) CompletionPrefix() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::completionPrefix")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_CompletionPrefix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCompleter) CompletionRole() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::completionRole")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) FilterMode() core.Qt__MatchFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::filterMode")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__MatchFlag(C.QCompleter_FilterMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) MaxVisibleItems() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::maxVisibleItems")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCompleter_MaxVisibleItems(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) ModelSorting() QCompleter__ModelSorting {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::modelSorting")
		}
	}()

	if ptr.Pointer() != nil {
		return QCompleter__ModelSorting(C.QCompleter_ModelSorting(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) SetCaseSensitivity(caseSensitivity core.Qt__CaseSensitivity) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setCaseSensitivity")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetCaseSensitivity(ptr.Pointer(), C.int(caseSensitivity))
	}
}

func (ptr *QCompleter) SetCompletionColumn(column int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setCompletionColumn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionColumn(ptr.Pointer(), C.int(column))
	}
}

func (ptr *QCompleter) SetCompletionMode(mode QCompleter__CompletionMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setCompletionMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QCompleter) SetCompletionPrefix(prefix string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setCompletionPrefix")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionPrefix(ptr.Pointer(), C.CString(prefix))
	}
}

func (ptr *QCompleter) SetCompletionRole(role int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setCompletionRole")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetCompletionRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QCompleter) SetFilterMode(filterMode core.Qt__MatchFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setFilterMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetFilterMode(ptr.Pointer(), C.int(filterMode))
	}
}

func (ptr *QCompleter) SetMaxVisibleItems(maxItems int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setMaxVisibleItems")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetMaxVisibleItems(ptr.Pointer(), C.int(maxItems))
	}
}

func (ptr *QCompleter) SetModelSorting(sorting QCompleter__ModelSorting) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setModelSorting")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetModelSorting(ptr.Pointer(), C.int(sorting))
	}
}

func (ptr *QCompleter) SetWrapAround(wrap bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setWrapAround")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetWrapAround(ptr.Pointer(), C.int(qt.GoBoolToInt(wrap)))
	}
}

func (ptr *QCompleter) WrapAround() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::wrapAround")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCompleter_WrapAround(ptr.Pointer()) != 0
	}
	return false
}

func NewQCompleter2(model core.QAbstractItemModel_ITF, parent core.QObject_ITF) *QCompleter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::QCompleter")
		}
	}()

	return NewQCompleterFromPointer(C.QCompleter_NewQCompleter2(core.PointerFromQAbstractItemModel(model), core.PointerFromQObject(parent)))
}

func NewQCompleter(parent core.QObject_ITF) *QCompleter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::QCompleter")
		}
	}()

	return NewQCompleterFromPointer(C.QCompleter_NewQCompleter(core.PointerFromQObject(parent)))
}

func NewQCompleter3(list []string, parent core.QObject_ITF) *QCompleter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::QCompleter")
		}
	}()

	return NewQCompleterFromPointer(C.QCompleter_NewQCompleter3(C.CString(strings.Join(list, ",,,")), core.PointerFromQObject(parent)))
}

func (ptr *QCompleter) ConnectActivated(f func(text string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::activated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QCompleter) DisconnectActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::activated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQCompleterActivated
func callbackQCompleterActivated(ptrName *C.char, text *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::activated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activated").(func(string))(C.GoString(text))
}

func (ptr *QCompleter) Complete(rect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::complete")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_Complete(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QCompleter) CompletionCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::completionCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCompleter_CompletionCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) CompletionModel() *core.QAbstractItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::completionModel")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QCompleter_CompletionModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) CurrentCompletion() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::currentCompletion")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_CurrentCompletion(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCompleter) CurrentIndex() *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QCompleter_CurrentIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) CurrentRow() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::currentRow")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCompleter_CurrentRow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCompleter) ConnectHighlighted(f func(text string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::highlighted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_ConnectHighlighted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "highlighted", f)
	}
}

func (ptr *QCompleter) DisconnectHighlighted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::highlighted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_DisconnectHighlighted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "highlighted")
	}
}

//export callbackQCompleterHighlighted
func callbackQCompleterHighlighted(ptrName *C.char, text *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::highlighted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "highlighted").(func(string))(C.GoString(text))
}

func (ptr *QCompleter) Model() *core.QAbstractItemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::model")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QCompleter_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) PathFromIndex(index core.QModelIndex_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::pathFromIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QCompleter_PathFromIndex(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return ""
}

func (ptr *QCompleter) Popup() *QAbstractItemView {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::popup")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemViewFromPointer(C.QCompleter_Popup(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) SetCurrentRow(row int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setCurrentRow")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCompleter_SetCurrentRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QCompleter) SetModel(model core.QAbstractItemModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QCompleter) SetPopup(popup QAbstractItemView_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setPopup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetPopup(ptr.Pointer(), PointerFromQAbstractItemView(popup))
	}
}

func (ptr *QCompleter) SetWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::setWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QCompleter) SplitPath(path string) []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::splitPath")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCompleter_SplitPath(ptr.Pointer(), C.CString(path))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCompleter) Widget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::widget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QCompleter_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCompleter) DestroyQCompleter() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCompleter::~QCompleter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCompleter_DestroyQCompleter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
