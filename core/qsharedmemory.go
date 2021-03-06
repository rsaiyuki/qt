package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QSharedMemory struct {
	QObject
}

type QSharedMemory_ITF interface {
	QObject_ITF
	QSharedMemory_PTR() *QSharedMemory
}

func PointerFromQSharedMemory(ptr QSharedMemory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedMemory_PTR().Pointer()
	}
	return nil
}

func NewQSharedMemoryFromPointer(ptr unsafe.Pointer) *QSharedMemory {
	var n = new(QSharedMemory)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSharedMemory_") {
		n.SetObjectName("QSharedMemory_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSharedMemory) QSharedMemory_PTR() *QSharedMemory {
	return ptr
}

//QSharedMemory::AccessMode
type QSharedMemory__AccessMode int64

const (
	QSharedMemory__ReadOnly  = QSharedMemory__AccessMode(0)
	QSharedMemory__ReadWrite = QSharedMemory__AccessMode(1)
)

//QSharedMemory::SharedMemoryError
type QSharedMemory__SharedMemoryError int64

const (
	QSharedMemory__NoError          = QSharedMemory__SharedMemoryError(0)
	QSharedMemory__PermissionDenied = QSharedMemory__SharedMemoryError(1)
	QSharedMemory__InvalidSize      = QSharedMemory__SharedMemoryError(2)
	QSharedMemory__KeyError         = QSharedMemory__SharedMemoryError(3)
	QSharedMemory__AlreadyExists    = QSharedMemory__SharedMemoryError(4)
	QSharedMemory__NotFound         = QSharedMemory__SharedMemoryError(5)
	QSharedMemory__LockError        = QSharedMemory__SharedMemoryError(6)
	QSharedMemory__OutOfResources   = QSharedMemory__SharedMemoryError(7)
	QSharedMemory__UnknownError     = QSharedMemory__SharedMemoryError(8)
)

func NewQSharedMemory2(parent QObject_ITF) *QSharedMemory {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::QSharedMemory")
		}
	}()

	return NewQSharedMemoryFromPointer(C.QSharedMemory_NewQSharedMemory2(PointerFromQObject(parent)))
}

func NewQSharedMemory(key string, parent QObject_ITF) *QSharedMemory {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::QSharedMemory")
		}
	}()

	return NewQSharedMemoryFromPointer(C.QSharedMemory_NewQSharedMemory(C.CString(key), PointerFromQObject(parent)))
}

func (ptr *QSharedMemory) Attach(mode QSharedMemory__AccessMode) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::attach")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Attach(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSharedMemory) ConstData() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::constData")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSharedMemory_ConstData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSharedMemory) Create(size int, mode QSharedMemory__AccessMode) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::create")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Create(ptr.Pointer(), C.int(size), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSharedMemory) Data() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::data")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSharedMemory_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSharedMemory) Data2() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::data")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSharedMemory_Data2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSharedMemory) Detach() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::detach")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Detach(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSharedMemory) Error() QSharedMemory__SharedMemoryError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QSharedMemory__SharedMemoryError(C.QSharedMemory_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSharedMemory) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSharedMemory) IsAttached() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::isAttached")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSharedMemory_IsAttached(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSharedMemory) Key() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::key")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_Key(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSharedMemory) Lock() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::lock")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Lock(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSharedMemory) NativeKey() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::nativeKey")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_NativeKey(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSharedMemory) SetKey(key string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::setKey")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSharedMemory_SetKey(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QSharedMemory) SetNativeKey(key string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::setNativeKey")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSharedMemory_SetNativeKey(ptr.Pointer(), C.CString(key))
	}
}

func (ptr *QSharedMemory) Size() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::size")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSharedMemory_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSharedMemory) Unlock() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::unlock")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSharedMemory_Unlock(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSharedMemory) DestroyQSharedMemory() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSharedMemory::~QSharedMemory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSharedMemory_DestroyQSharedMemory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
