package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QMutex struct {
	ptr unsafe.Pointer
}

type QMutex_ITF interface {
	QMutex_PTR() *QMutex
}

func (p *QMutex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutex(ptr QMutex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutex_PTR().Pointer()
	}
	return nil
}

func NewQMutexFromPointer(ptr unsafe.Pointer) *QMutex {
	var n = new(QMutex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutex) QMutex_PTR() *QMutex {
	return ptr
}

//QMutex::RecursionMode
type QMutex__RecursionMode int64

const (
	QMutex__NonRecursive = QMutex__RecursionMode(0)
	QMutex__Recursive    = QMutex__RecursionMode(1)
)

func (ptr *QMutex) Lock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutex::lock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMutex_Lock(ptr.Pointer())
	}
}

func (ptr *QMutex) TryLock(timeout int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutex::tryLock")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMutex_TryLock(ptr.Pointer(), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QMutex) Unlock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutex::unlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMutex_Unlock(ptr.Pointer())
	}
}

func NewQMutex(mode QMutex__RecursionMode) *QMutex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutex::QMutex")
		}
	}()

	return NewQMutexFromPointer(C.QMutex_NewQMutex(C.int(mode)))
}

func (ptr *QMutex) IsRecursive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMutex::isRecursive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMutex_IsRecursive(ptr.Pointer()) != 0
	}
	return false
}
