package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QFontComboBox struct {
	QComboBox
}

type QFontComboBox_ITF interface {
	QComboBox_ITF
	QFontComboBox_PTR() *QFontComboBox
}

func PointerFromQFontComboBox(ptr QFontComboBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFontComboBox_PTR().Pointer()
	}
	return nil
}

func NewQFontComboBoxFromPointer(ptr unsafe.Pointer) *QFontComboBox {
	var n = new(QFontComboBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFontComboBox_") {
		n.SetObjectName("QFontComboBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFontComboBox) QFontComboBox_PTR() *QFontComboBox {
	return ptr
}

//QFontComboBox::FontFilter
type QFontComboBox__FontFilter int64

const (
	QFontComboBox__AllFonts          = QFontComboBox__FontFilter(0)
	QFontComboBox__ScalableFonts     = QFontComboBox__FontFilter(0x1)
	QFontComboBox__NonScalableFonts  = QFontComboBox__FontFilter(0x2)
	QFontComboBox__MonospacedFonts   = QFontComboBox__FontFilter(0x4)
	QFontComboBox__ProportionalFonts = QFontComboBox__FontFilter(0x8)
)

func (ptr *QFontComboBox) FontFilters() QFontComboBox__FontFilter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFontComboBox::fontFilters")
		}
	}()

	if ptr.Pointer() != nil {
		return QFontComboBox__FontFilter(C.QFontComboBox_FontFilters(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFontComboBox) SetCurrentFont(font gui.QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFontComboBox::setCurrentFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFontComboBox_SetCurrentFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QFontComboBox) SetFontFilters(filters QFontComboBox__FontFilter) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFontComboBox::setFontFilters")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFontComboBox_SetFontFilters(ptr.Pointer(), C.int(filters))
	}
}

func (ptr *QFontComboBox) SetWritingSystem(script gui.QFontDatabase__WritingSystem) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFontComboBox::setWritingSystem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFontComboBox_SetWritingSystem(ptr.Pointer(), C.int(script))
	}
}

func (ptr *QFontComboBox) WritingSystem() gui.QFontDatabase__WritingSystem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFontComboBox::writingSystem")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.QFontDatabase__WritingSystem(C.QFontComboBox_WritingSystem(ptr.Pointer()))
	}
	return 0
}

func NewQFontComboBox(parent QWidget_ITF) *QFontComboBox {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFontComboBox::QFontComboBox")
		}
	}()

	return NewQFontComboBoxFromPointer(C.QFontComboBox_NewQFontComboBox(PointerFromQWidget(parent)))
}

func (ptr *QFontComboBox) DestroyQFontComboBox() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFontComboBox::~QFontComboBox")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFontComboBox_DestroyQFontComboBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
