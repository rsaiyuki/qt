package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QWriteLocker struct {
	ptr unsafe.Pointer
}

type QWriteLocker_ITF interface {
	QWriteLocker_PTR() *QWriteLocker
}

func (p *QWriteLocker) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWriteLocker) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWriteLocker(ptr QWriteLocker_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWriteLocker_PTR().Pointer()
	}
	return nil
}

func NewQWriteLockerFromPointer(ptr unsafe.Pointer) *QWriteLocker {
	var n = new(QWriteLocker)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWriteLocker) QWriteLocker_PTR() *QWriteLocker {
	return ptr
}

func NewQWriteLocker(lock QReadWriteLock_ITF) *QWriteLocker {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWriteLocker::QWriteLocker")
		}
	}()

	return NewQWriteLockerFromPointer(C.QWriteLocker_NewQWriteLocker(PointerFromQReadWriteLock(lock)))
}

func (ptr *QWriteLocker) ReadWriteLock() *QReadWriteLock {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWriteLocker::readWriteLock")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQReadWriteLockFromPointer(C.QWriteLocker_ReadWriteLock(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWriteLocker) Relock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWriteLocker::relock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWriteLocker_Relock(ptr.Pointer())
	}
}

func (ptr *QWriteLocker) Unlock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWriteLocker::unlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWriteLocker_Unlock(ptr.Pointer())
	}
}

func (ptr *QWriteLocker) DestroyQWriteLocker() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWriteLocker::~QWriteLocker")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWriteLocker_DestroyQWriteLocker(ptr.Pointer())
	}
}
