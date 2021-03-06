package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSize struct {
	ptr unsafe.Pointer
}

type QSize_ITF interface {
	QSize_PTR() *QSize
}

func (p *QSize) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSize) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSize(ptr QSize_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSize_PTR().Pointer()
	}
	return nil
}

func NewQSizeFromPointer(ptr unsafe.Pointer) *QSize {
	var n = new(QSize)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSize) QSize_PTR() *QSize {
	return ptr
}

func NewQSize() *QSize {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::QSize")
		}
	}()

	return NewQSizeFromPointer(C.QSize_NewQSize())
}

func NewQSize2(width int, height int) *QSize {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::QSize")
		}
	}()

	return NewQSizeFromPointer(C.QSize_NewQSize2(C.int(width), C.int(height)))
}

func (ptr *QSize) Height() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::height")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSize_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSize) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSize_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSize) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSize_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSize) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSize_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSize) Rheight() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::rheight")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSize_Rheight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSize) Rwidth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::rwidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSize_Rwidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSize) Scale2(size QSize_ITF, mode Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::scale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSize_Scale2(ptr.Pointer(), PointerFromQSize(size), C.int(mode))
	}
}

func (ptr *QSize) Scale(width int, height int, mode Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::scale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSize_Scale(ptr.Pointer(), C.int(width), C.int(height), C.int(mode))
	}
}

func (ptr *QSize) SetHeight(height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::setHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSize_SetHeight(ptr.Pointer(), C.int(height))
	}
}

func (ptr *QSize) SetWidth(width int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::setWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSize_SetWidth(ptr.Pointer(), C.int(width))
	}
}

func (ptr *QSize) Transpose() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::transpose")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSize_Transpose(ptr.Pointer())
	}
}

func (ptr *QSize) Width() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSize::width")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSize_Width(ptr.Pointer()))
	}
	return 0
}
