package androidextras

//#include "androidextras_android.h"
import "C"
import (
	"log"
	"unsafe"
)

type QAndroidJniEnvironment struct {
	ptr unsafe.Pointer
}

type QAndroidJniEnvironment_ITF interface {
	QAndroidJniEnvironment_PTR() *QAndroidJniEnvironment
}

func (p *QAndroidJniEnvironment) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAndroidJniEnvironment) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAndroidJniEnvironment(ptr QAndroidJniEnvironment_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidJniEnvironment_PTR().Pointer()
	}
	return nil
}

func NewQAndroidJniEnvironmentFromPointer(ptr unsafe.Pointer) *QAndroidJniEnvironment {
	var n = new(QAndroidJniEnvironment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAndroidJniEnvironment) QAndroidJniEnvironment_PTR() *QAndroidJniEnvironment {
	return ptr
}

func NewQAndroidJniEnvironment() *QAndroidJniEnvironment {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniEnvironment::QAndroidJniEnvironment")
		}
	}()

	return NewQAndroidJniEnvironmentFromPointer(C.QAndroidJniEnvironment_NewQAndroidJniEnvironment())
}

func QAndroidJniEnvironment_JavaVM() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniEnvironment::javaVM")
		}
	}()

	return unsafe.Pointer(C.QAndroidJniEnvironment_QAndroidJniEnvironment_JavaVM())
}

func (ptr *QAndroidJniEnvironment) DestroyQAndroidJniEnvironment() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAndroidJniEnvironment::~QAndroidJniEnvironment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAndroidJniEnvironment_DestroyQAndroidJniEnvironment(ptr.Pointer())
	}
}
