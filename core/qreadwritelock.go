package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QReadWriteLock struct {
	ptr unsafe.Pointer
}

type QReadWriteLock_ITF interface {
	QReadWriteLock_PTR() *QReadWriteLock
}

func (p *QReadWriteLock) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QReadWriteLock) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQReadWriteLock(ptr QReadWriteLock_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QReadWriteLock_PTR().Pointer()
	}
	return nil
}

func NewQReadWriteLockFromPointer(ptr unsafe.Pointer) *QReadWriteLock {
	var n = new(QReadWriteLock)
	n.SetPointer(ptr)
	return n
}

func (ptr *QReadWriteLock) QReadWriteLock_PTR() *QReadWriteLock {
	return ptr
}

//QReadWriteLock::RecursionMode
type QReadWriteLock__RecursionMode int64

const (
	QReadWriteLock__NonRecursive = QReadWriteLock__RecursionMode(0)
	QReadWriteLock__Recursive    = QReadWriteLock__RecursionMode(1)
)

func NewQReadWriteLock(recursionMode QReadWriteLock__RecursionMode) *QReadWriteLock {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::QReadWriteLock")
		}
	}()

	return NewQReadWriteLockFromPointer(C.QReadWriteLock_NewQReadWriteLock(C.int(recursionMode)))
}

func (ptr *QReadWriteLock) LockForRead() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::lockForRead")
		}
	}()

	if ptr.Pointer() != nil {
		C.QReadWriteLock_LockForRead(ptr.Pointer())
	}
}

func (ptr *QReadWriteLock) LockForWrite() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::lockForWrite")
		}
	}()

	if ptr.Pointer() != nil {
		C.QReadWriteLock_LockForWrite(ptr.Pointer())
	}
}

func (ptr *QReadWriteLock) TryLockForRead() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::tryLockForRead")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QReadWriteLock_TryLockForRead(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QReadWriteLock) TryLockForRead2(timeout int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::tryLockForRead")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QReadWriteLock_TryLockForRead2(ptr.Pointer(), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QReadWriteLock) TryLockForWrite() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::tryLockForWrite")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QReadWriteLock_TryLockForWrite(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QReadWriteLock) TryLockForWrite2(timeout int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::tryLockForWrite")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QReadWriteLock_TryLockForWrite2(ptr.Pointer(), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QReadWriteLock) Unlock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::unlock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QReadWriteLock_Unlock(ptr.Pointer())
	}
}

func (ptr *QReadWriteLock) DestroyQReadWriteLock() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QReadWriteLock::~QReadWriteLock")
		}
	}()

	if ptr.Pointer() != nil {
		C.QReadWriteLock_DestroyQReadWriteLock(ptr.Pointer())
	}
}
