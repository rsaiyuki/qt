package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QTextList struct {
	QTextBlockGroup
}

type QTextList_ITF interface {
	QTextBlockGroup_ITF
	QTextList_PTR() *QTextList
}

func PointerFromQTextList(ptr QTextList_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextList_PTR().Pointer()
	}
	return nil
}

func NewQTextListFromPointer(ptr unsafe.Pointer) *QTextList {
	var n = new(QTextList)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTextList_") {
		n.SetObjectName("QTextList_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTextList) QTextList_PTR() *QTextList {
	return ptr
}

func (ptr *QTextList) ItemNumber(block QTextBlock_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextList::itemNumber")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextList_ItemNumber(ptr.Pointer(), PointerFromQTextBlock(block)))
	}
	return 0
}

func (ptr *QTextList) ItemText(block QTextBlock_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextList::itemText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextList_ItemText(ptr.Pointer(), PointerFromQTextBlock(block)))
	}
	return ""
}

func (ptr *QTextList) Add(block QTextBlock_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextList::add")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextList_Add(ptr.Pointer(), PointerFromQTextBlock(block))
	}
}

func (ptr *QTextList) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextList::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextList_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextList) RemoveItem(i int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextList::removeItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextList_RemoveItem(ptr.Pointer(), C.int(i))
	}
}

func (ptr *QTextList) SetFormat(format QTextListFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextList::setFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextList_SetFormat(ptr.Pointer(), PointerFromQTextListFormat(format))
	}
}
