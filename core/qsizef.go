package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSizeF struct {
	ptr unsafe.Pointer
}

type QSizeF_ITF interface {
	QSizeF_PTR() *QSizeF
}

func (p *QSizeF) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSizeF) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSizeF(ptr QSizeF_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSizeF_PTR().Pointer()
	}
	return nil
}

func NewQSizeFFromPointer(ptr unsafe.Pointer) *QSizeF {
	var n = new(QSizeF)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSizeF) QSizeF_PTR() *QSizeF {
	return ptr
}

func NewQSizeF() *QSizeF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::QSizeF")
		}
	}()

	return NewQSizeFFromPointer(C.QSizeF_NewQSizeF())
}

func NewQSizeF2(size QSize_ITF) *QSizeF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::QSizeF")
		}
	}()

	return NewQSizeFFromPointer(C.QSizeF_NewQSizeF2(PointerFromQSize(size)))
}

func NewQSizeF3(width float64, height float64) *QSizeF {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::QSizeF")
		}
	}()

	return NewQSizeFFromPointer(C.QSizeF_NewQSizeF3(C.double(width), C.double(height)))
}

func (ptr *QSizeF) Height() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::height")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QSizeF_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizeF) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSizeF_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSizeF) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSizeF_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSizeF) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSizeF_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSizeF) Rheight() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::rheight")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QSizeF_Rheight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizeF) Rwidth() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::rwidth")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QSizeF_Rwidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSizeF) Scale2(size QSizeF_ITF, mode Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::scale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSizeF_Scale2(ptr.Pointer(), PointerFromQSizeF(size), C.int(mode))
	}
}

func (ptr *QSizeF) Scale(width float64, height float64, mode Qt__AspectRatioMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::scale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSizeF_Scale(ptr.Pointer(), C.double(width), C.double(height), C.int(mode))
	}
}

func (ptr *QSizeF) SetHeight(height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::setHeight")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSizeF_SetHeight(ptr.Pointer(), C.double(height))
	}
}

func (ptr *QSizeF) SetWidth(width float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::setWidth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSizeF_SetWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QSizeF) Transpose() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::transpose")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSizeF_Transpose(ptr.Pointer())
	}
}

func (ptr *QSizeF) Width() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSizeF::width")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QSizeF_Width(ptr.Pointer()))
	}
	return 0
}
