package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"strings"
	"unsafe"
)

type QTreeWidgetItem struct {
	ptr unsafe.Pointer
}

type QTreeWidgetItem_ITF interface {
	QTreeWidgetItem_PTR() *QTreeWidgetItem
}

func (p *QTreeWidgetItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTreeWidgetItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTreeWidgetItem(ptr QTreeWidgetItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTreeWidgetItem_PTR().Pointer()
	}
	return nil
}

func NewQTreeWidgetItemFromPointer(ptr unsafe.Pointer) *QTreeWidgetItem {
	var n = new(QTreeWidgetItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTreeWidgetItem) QTreeWidgetItem_PTR() *QTreeWidgetItem {
	return ptr
}

//QTreeWidgetItem::ChildIndicatorPolicy
type QTreeWidgetItem__ChildIndicatorPolicy int64

const (
	QTreeWidgetItem__ShowIndicator                  = QTreeWidgetItem__ChildIndicatorPolicy(0)
	QTreeWidgetItem__DontShowIndicator              = QTreeWidgetItem__ChildIndicatorPolicy(1)
	QTreeWidgetItem__DontShowIndicatorWhenChildless = QTreeWidgetItem__ChildIndicatorPolicy(2)
)

//QTreeWidgetItem::ItemType
type QTreeWidgetItem__ItemType int64

const (
	QTreeWidgetItem__Type     = QTreeWidgetItem__ItemType(0)
	QTreeWidgetItem__UserType = QTreeWidgetItem__ItemType(1000)
)

func NewQTreeWidgetItem5(parent QTreeWidget_ITF, preceding QTreeWidgetItem_ITF, ty int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem5(PointerFromQTreeWidget(parent), PointerFromQTreeWidgetItem(preceding), C.int(ty)))
}

func NewQTreeWidgetItem4(parent QTreeWidget_ITF, strin []string, ty int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem4(PointerFromQTreeWidget(parent), C.CString(strings.Join(strin, ",,,")), C.int(ty)))
}

func NewQTreeWidgetItem3(parent QTreeWidget_ITF, ty int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem3(PointerFromQTreeWidget(parent), C.int(ty)))
}

func NewQTreeWidgetItem8(parent QTreeWidgetItem_ITF, preceding QTreeWidgetItem_ITF, ty int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem8(PointerFromQTreeWidgetItem(parent), PointerFromQTreeWidgetItem(preceding), C.int(ty)))
}

func (ptr *QTreeWidgetItem) Flags() core.Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QTreeWidgetItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidgetItem) SetFlags(flags core.Qt__ItemFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setFlags")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func NewQTreeWidgetItem7(parent QTreeWidgetItem_ITF, strin []string, ty int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem7(PointerFromQTreeWidgetItem(parent), C.CString(strings.Join(strin, ",,,")), C.int(ty)))
}

func NewQTreeWidgetItem6(parent QTreeWidgetItem_ITF, ty int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem6(PointerFromQTreeWidgetItem(parent), C.int(ty)))
}

func NewQTreeWidgetItem2(strin []string, ty int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem2(C.CString(strings.Join(strin, ",,,")), C.int(ty)))
}

func NewQTreeWidgetItem9(other QTreeWidgetItem_ITF) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem9(PointerFromQTreeWidgetItem(other)))
}

func NewQTreeWidgetItem(ty int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::QTreeWidgetItem")
		}
	}()

	return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_NewQTreeWidgetItem(C.int(ty)))
}

func (ptr *QTreeWidgetItem) AddChild(child QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::addChild")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_AddChild(ptr.Pointer(), PointerFromQTreeWidgetItem(child))
	}
}

func (ptr *QTreeWidgetItem) Background(column int) *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::background")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QTreeWidgetItem_Background(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QTreeWidgetItem) CheckState(column int) core.Qt__CheckState {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::checkState")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QTreeWidgetItem_CheckState(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeWidgetItem) Child(index int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::child")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_Child(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTreeWidgetItem) ChildCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::childCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidgetItem) ChildIndicatorPolicy() QTreeWidgetItem__ChildIndicatorPolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::childIndicatorPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		return QTreeWidgetItem__ChildIndicatorPolicy(C.QTreeWidgetItem_ChildIndicatorPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidgetItem) ColumnCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_ColumnCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidgetItem) Data(column int, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTreeWidgetItem_Data(ptr.Pointer(), C.int(column), C.int(role)))
	}
	return nil
}

func (ptr *QTreeWidgetItem) Clone() *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::clone")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidgetItem) Foreground(column int) *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::foreground")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QTreeWidgetItem_Foreground(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QTreeWidgetItem) IndexOfChild(child QTreeWidgetItem_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::indexOfChild")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_IndexOfChild(ptr.Pointer(), PointerFromQTreeWidgetItem(child)))
	}
	return 0
}

func (ptr *QTreeWidgetItem) InsertChild(index int, child QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::insertChild")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_InsertChild(ptr.Pointer(), C.int(index), PointerFromQTreeWidgetItem(child))
	}
}

func (ptr *QTreeWidgetItem) IsDisabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::isDisabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsDisabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) IsExpanded() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::isExpanded")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsExpanded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) IsFirstColumnSpanned() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::isFirstColumnSpanned")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsFirstColumnSpanned(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) IsHidden() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::isHidden")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) IsSelected() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::isSelected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTreeWidgetItem_IsSelected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTreeWidgetItem) Parent() *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidgetItem) Read(in core.QDataStream_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::read")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_Read(ptr.Pointer(), core.PointerFromQDataStream(in))
	}
}

func (ptr *QTreeWidgetItem) RemoveChild(child QTreeWidgetItem_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::removeChild")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_RemoveChild(ptr.Pointer(), PointerFromQTreeWidgetItem(child))
	}
}

func (ptr *QTreeWidgetItem) SetBackground(column int, brush gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setBackground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetBackground(ptr.Pointer(), C.int(column), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QTreeWidgetItem) SetCheckState(column int, state core.Qt__CheckState) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setCheckState")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetCheckState(ptr.Pointer(), C.int(column), C.int(state))
	}
}

func (ptr *QTreeWidgetItem) SetChildIndicatorPolicy(policy QTreeWidgetItem__ChildIndicatorPolicy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setChildIndicatorPolicy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetChildIndicatorPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QTreeWidgetItem) SetData(column int, role int, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetData(ptr.Pointer(), C.int(column), C.int(role), core.PointerFromQVariant(value))
	}
}

func (ptr *QTreeWidgetItem) SetDisabled(disabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setDisabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disabled)))
	}
}

func (ptr *QTreeWidgetItem) SetExpanded(expand bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setExpanded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetExpanded(ptr.Pointer(), C.int(qt.GoBoolToInt(expand)))
	}
}

func (ptr *QTreeWidgetItem) SetFirstColumnSpanned(span bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setFirstColumnSpanned")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetFirstColumnSpanned(ptr.Pointer(), C.int(qt.GoBoolToInt(span)))
	}
}

func (ptr *QTreeWidgetItem) SetFont(column int, font gui.QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetFont(ptr.Pointer(), C.int(column), gui.PointerFromQFont(font))
	}
}

func (ptr *QTreeWidgetItem) SetForeground(column int, brush gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setForeground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetForeground(ptr.Pointer(), C.int(column), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QTreeWidgetItem) SetHidden(hide bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setHidden")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hide)))
	}
}

func (ptr *QTreeWidgetItem) SetIcon(column int, icon gui.QIcon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setIcon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetIcon(ptr.Pointer(), C.int(column), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QTreeWidgetItem) SetSelected(sele bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetSelected(ptr.Pointer(), C.int(qt.GoBoolToInt(sele)))
	}
}

func (ptr *QTreeWidgetItem) SetSizeHint(column int, size core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setSizeHint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetSizeHint(ptr.Pointer(), C.int(column), core.PointerFromQSize(size))
	}
}

func (ptr *QTreeWidgetItem) SetStatusTip(column int, statusTip string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setStatusTip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetStatusTip(ptr.Pointer(), C.int(column), C.CString(statusTip))
	}
}

func (ptr *QTreeWidgetItem) SetText(column int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetText(ptr.Pointer(), C.int(column), C.CString(text))
	}
}

func (ptr *QTreeWidgetItem) SetTextAlignment(column int, alignment int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setTextAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetTextAlignment(ptr.Pointer(), C.int(column), C.int(alignment))
	}
}

func (ptr *QTreeWidgetItem) SetToolTip(column int, toolTip string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setToolTip")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetToolTip(ptr.Pointer(), C.int(column), C.CString(toolTip))
	}
}

func (ptr *QTreeWidgetItem) SetWhatsThis(column int, whatsThis string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::setWhatsThis")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SetWhatsThis(ptr.Pointer(), C.int(column), C.CString(whatsThis))
	}
}

func (ptr *QTreeWidgetItem) SortChildren(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::sortChildren")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_SortChildren(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QTreeWidgetItem) StatusTip(column int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::statusTip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_StatusTip(ptr.Pointer(), C.int(column)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) TakeChild(index int) *QTreeWidgetItem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::takeChild")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetItemFromPointer(C.QTreeWidgetItem_TakeChild(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QTreeWidgetItem) Text(column int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_Text(ptr.Pointer(), C.int(column)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) TextAlignment(column int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::textAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_TextAlignment(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QTreeWidgetItem) ToolTip(column int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::toolTip")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_ToolTip(ptr.Pointer(), C.int(column)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) TreeWidget() *QTreeWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::treeWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTreeWidgetFromPointer(C.QTreeWidgetItem_TreeWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTreeWidgetItem) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTreeWidgetItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTreeWidgetItem) WhatsThis(column int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::whatsThis")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTreeWidgetItem_WhatsThis(ptr.Pointer(), C.int(column)))
	}
	return ""
}

func (ptr *QTreeWidgetItem) Write(out core.QDataStream_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::write")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_Write(ptr.Pointer(), core.PointerFromQDataStream(out))
	}
}

func (ptr *QTreeWidgetItem) DestroyQTreeWidgetItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTreeWidgetItem::~QTreeWidgetItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTreeWidgetItem_DestroyQTreeWidgetItem(ptr.Pointer())
	}
}
