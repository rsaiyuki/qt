package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QFileDevice struct {
	QIODevice
}

type QFileDevice_ITF interface {
	QIODevice_ITF
	QFileDevice_PTR() *QFileDevice
}

func PointerFromQFileDevice(ptr QFileDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileDevice_PTR().Pointer()
	}
	return nil
}

func NewQFileDeviceFromPointer(ptr unsafe.Pointer) *QFileDevice {
	var n = new(QFileDevice)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFileDevice_") {
		n.SetObjectName("QFileDevice_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileDevice) QFileDevice_PTR() *QFileDevice {
	return ptr
}

//QFileDevice::FileError
type QFileDevice__FileError int64

const (
	QFileDevice__NoError          = QFileDevice__FileError(0)
	QFileDevice__ReadError        = QFileDevice__FileError(1)
	QFileDevice__WriteError       = QFileDevice__FileError(2)
	QFileDevice__FatalError       = QFileDevice__FileError(3)
	QFileDevice__ResourceError    = QFileDevice__FileError(4)
	QFileDevice__OpenError        = QFileDevice__FileError(5)
	QFileDevice__AbortError       = QFileDevice__FileError(6)
	QFileDevice__TimeOutError     = QFileDevice__FileError(7)
	QFileDevice__UnspecifiedError = QFileDevice__FileError(8)
	QFileDevice__RemoveError      = QFileDevice__FileError(9)
	QFileDevice__RenameError      = QFileDevice__FileError(10)
	QFileDevice__PositionError    = QFileDevice__FileError(11)
	QFileDevice__ResizeError      = QFileDevice__FileError(12)
	QFileDevice__PermissionsError = QFileDevice__FileError(13)
	QFileDevice__CopyError        = QFileDevice__FileError(14)
)

//QFileDevice::FileHandleFlag
type QFileDevice__FileHandleFlag int64

const (
	QFileDevice__AutoCloseHandle = QFileDevice__FileHandleFlag(0x0001)
	QFileDevice__DontCloseHandle = QFileDevice__FileHandleFlag(0)
)

//QFileDevice::MemoryMapFlags
type QFileDevice__MemoryMapFlags int64

const (
	QFileDevice__NoOptions        = QFileDevice__MemoryMapFlags(0)
	QFileDevice__MapPrivateOption = QFileDevice__MemoryMapFlags(0x0001)
)

//QFileDevice::Permission
type QFileDevice__Permission int64

const (
	QFileDevice__ReadOwner  = QFileDevice__Permission(0x4000)
	QFileDevice__WriteOwner = QFileDevice__Permission(0x2000)
	QFileDevice__ExeOwner   = QFileDevice__Permission(0x1000)
	QFileDevice__ReadUser   = QFileDevice__Permission(0x0400)
	QFileDevice__WriteUser  = QFileDevice__Permission(0x0200)
	QFileDevice__ExeUser    = QFileDevice__Permission(0x0100)
	QFileDevice__ReadGroup  = QFileDevice__Permission(0x0040)
	QFileDevice__WriteGroup = QFileDevice__Permission(0x0020)
	QFileDevice__ExeGroup   = QFileDevice__Permission(0x0010)
	QFileDevice__ReadOther  = QFileDevice__Permission(0x0004)
	QFileDevice__WriteOther = QFileDevice__Permission(0x0002)
	QFileDevice__ExeOther   = QFileDevice__Permission(0x0001)
)

func (ptr *QFileDevice) AtEnd() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::atEnd")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDevice_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDevice) Close() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::close")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDevice_Close(ptr.Pointer())
	}
}

func (ptr *QFileDevice) Error() QFileDevice__FileError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QFileDevice__FileError(C.QFileDevice_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDevice) FileName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::fileName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDevice_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDevice) Flush() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::flush")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDevice_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDevice) Handle() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::handle")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFileDevice_Handle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDevice) IsSequential() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::isSequential")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDevice_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDevice) Permissions() QFileDevice__Permission {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::permissions")
		}
	}()

	if ptr.Pointer() != nil {
		return QFileDevice__Permission(C.QFileDevice_Permissions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDevice) SetPermissions(permissions QFileDevice__Permission) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::setPermissions")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDevice_SetPermissions(ptr.Pointer(), C.int(permissions)) != 0
	}
	return false
}

func (ptr *QFileDevice) UnsetError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::unsetError")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDevice_UnsetError(ptr.Pointer())
	}
}

func (ptr *QFileDevice) DestroyQFileDevice() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDevice::~QFileDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDevice_DestroyQFileDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
