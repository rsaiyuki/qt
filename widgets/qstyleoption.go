package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QStyleOption struct {
	ptr unsafe.Pointer
}

type QStyleOption_ITF interface {
	QStyleOption_PTR() *QStyleOption
}

func (p *QStyleOption) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStyleOption) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStyleOption(ptr QStyleOption_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOption_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionFromPointer(ptr unsafe.Pointer) *QStyleOption {
	var n = new(QStyleOption)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOption) QStyleOption_PTR() *QStyleOption {
	return ptr
}

//QStyleOption::OptionType
type QStyleOption__OptionType int64

var (
	QStyleOption__SO_Default           = QStyleOption__OptionType(0)
	QStyleOption__SO_FocusRect         = QStyleOption__OptionType(1)
	QStyleOption__SO_Button            = QStyleOption__OptionType(2)
	QStyleOption__SO_Tab               = QStyleOption__OptionType(3)
	QStyleOption__SO_MenuItem          = QStyleOption__OptionType(4)
	QStyleOption__SO_Frame             = QStyleOption__OptionType(5)
	QStyleOption__SO_ProgressBar       = QStyleOption__OptionType(6)
	QStyleOption__SO_ToolBox           = QStyleOption__OptionType(7)
	QStyleOption__SO_Header            = QStyleOption__OptionType(8)
	QStyleOption__SO_DockWidget        = QStyleOption__OptionType(9)
	QStyleOption__SO_ViewItem          = QStyleOption__OptionType(10)
	QStyleOption__SO_TabWidgetFrame    = QStyleOption__OptionType(11)
	QStyleOption__SO_TabBarBase        = QStyleOption__OptionType(12)
	QStyleOption__SO_RubberBand        = QStyleOption__OptionType(13)
	QStyleOption__SO_ToolBar           = QStyleOption__OptionType(14)
	QStyleOption__SO_GraphicsItem      = QStyleOption__OptionType(15)
	QStyleOption__SO_Complex           = QStyleOption__OptionType(0xf0000)
	QStyleOption__SO_Slider            = QStyleOption__OptionType(C.QStyleOption_SO_Slider_Type())
	QStyleOption__SO_SpinBox           = QStyleOption__OptionType(C.QStyleOption_SO_SpinBox_Type())
	QStyleOption__SO_ToolButton        = QStyleOption__OptionType(C.QStyleOption_SO_ToolButton_Type())
	QStyleOption__SO_ComboBox          = QStyleOption__OptionType(C.QStyleOption_SO_ComboBox_Type())
	QStyleOption__SO_TitleBar          = QStyleOption__OptionType(C.QStyleOption_SO_TitleBar_Type())
	QStyleOption__SO_GroupBox          = QStyleOption__OptionType(C.QStyleOption_SO_GroupBox_Type())
	QStyleOption__SO_SizeGrip          = QStyleOption__OptionType(C.QStyleOption_SO_SizeGrip_Type())
	QStyleOption__SO_CustomBase        = QStyleOption__OptionType(0xf00)
	QStyleOption__SO_ComplexCustomBase = QStyleOption__OptionType(0xf000000)
)

//QStyleOption::StyleOptionType
type QStyleOption__StyleOptionType int64

var (
	QStyleOption__Type = QStyleOption__StyleOptionType(QStyleOption__SO_Default)
)

//QStyleOption::StyleOptionVersion
type QStyleOption__StyleOptionVersion int64

var (
	QStyleOption__Version = QStyleOption__StyleOptionVersion(1)
)

func NewQStyleOption2(other QStyleOption_ITF) *QStyleOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOption::QStyleOption")
		}
	}()

	return NewQStyleOptionFromPointer(C.QStyleOption_NewQStyleOption2(PointerFromQStyleOption(other)))
}

func NewQStyleOption(version int, ty int) *QStyleOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOption::QStyleOption")
		}
	}()

	return NewQStyleOptionFromPointer(C.QStyleOption_NewQStyleOption(C.int(version), C.int(ty)))
}

func (ptr *QStyleOption) InitFrom(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOption::initFrom")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleOption_InitFrom(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStyleOption) DestroyQStyleOption() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOption::~QStyleOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QStyleOption_DestroyQStyleOption(ptr.Pointer())
	}
}
