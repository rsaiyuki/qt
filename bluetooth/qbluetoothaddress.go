package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"log"
	"unsafe"
)

type QBluetoothAddress struct {
	ptr unsafe.Pointer
}

type QBluetoothAddress_ITF interface {
	QBluetoothAddress_PTR() *QBluetoothAddress
}

func (p *QBluetoothAddress) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothAddress) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothAddress(ptr QBluetoothAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothAddress_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothAddressFromPointer(ptr unsafe.Pointer) *QBluetoothAddress {
	var n = new(QBluetoothAddress)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetoothAddress) QBluetoothAddress_PTR() *QBluetoothAddress {
	return ptr
}

func NewQBluetoothAddress() *QBluetoothAddress {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothAddress::QBluetoothAddress")
		}
	}()

	return NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress())
}

func NewQBluetoothAddress4(other QBluetoothAddress_ITF) *QBluetoothAddress {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothAddress::QBluetoothAddress")
		}
	}()

	return NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress4(PointerFromQBluetoothAddress(other)))
}

func NewQBluetoothAddress3(address string) *QBluetoothAddress {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothAddress::QBluetoothAddress")
		}
	}()

	return NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress3(C.CString(address)))
}

func (ptr *QBluetoothAddress) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothAddress::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothAddress) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothAddress::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothAddress) ToString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothAddress::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothAddress) DestroyQBluetoothAddress() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothAddress::~QBluetoothAddress")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothAddress_DestroyQBluetoothAddress(ptr.Pointer())
	}
}
