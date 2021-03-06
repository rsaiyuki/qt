package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QBluetoothLocalDevice struct {
	core.QObject
}

type QBluetoothLocalDevice_ITF interface {
	core.QObject_ITF
	QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice
}

func PointerFromQBluetoothLocalDevice(ptr QBluetoothLocalDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothLocalDevice_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) *QBluetoothLocalDevice {
	var n = new(QBluetoothLocalDevice)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothLocalDevice_") {
		n.SetObjectName("QBluetoothLocalDevice_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothLocalDevice) QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice {
	return ptr
}

//QBluetoothLocalDevice::Error
type QBluetoothLocalDevice__Error int64

const (
	QBluetoothLocalDevice__NoError      = QBluetoothLocalDevice__Error(0)
	QBluetoothLocalDevice__PairingError = QBluetoothLocalDevice__Error(1)
	QBluetoothLocalDevice__UnknownError = QBluetoothLocalDevice__Error(100)
)

//QBluetoothLocalDevice::HostMode
type QBluetoothLocalDevice__HostMode int64

const (
	QBluetoothLocalDevice__HostPoweredOff                 = QBluetoothLocalDevice__HostMode(0)
	QBluetoothLocalDevice__HostConnectable                = QBluetoothLocalDevice__HostMode(1)
	QBluetoothLocalDevice__HostDiscoverable               = QBluetoothLocalDevice__HostMode(2)
	QBluetoothLocalDevice__HostDiscoverableLimitedInquiry = QBluetoothLocalDevice__HostMode(3)
)

//QBluetoothLocalDevice::Pairing
type QBluetoothLocalDevice__Pairing int64

const (
	QBluetoothLocalDevice__Unpaired         = QBluetoothLocalDevice__Pairing(0)
	QBluetoothLocalDevice__Paired           = QBluetoothLocalDevice__Pairing(1)
	QBluetoothLocalDevice__AuthorizedPaired = QBluetoothLocalDevice__Pairing(2)
)

func (ptr *QBluetoothLocalDevice) ConnectError(f func(error QBluetoothLocalDevice__Error)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectError() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::error")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQBluetoothLocalDeviceError
func callbackQBluetoothLocalDeviceError(ptrName *C.char, error C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::error")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "error").(func(QBluetoothLocalDevice__Error))(QBluetoothLocalDevice__Error(error))
}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::hostModeStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectHostModeStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hostModeStateChanged", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectHostModeStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::hostModeStateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectHostModeStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hostModeStateChanged")
	}
}

//export callbackQBluetoothLocalDeviceHostModeStateChanged
func callbackQBluetoothLocalDeviceHostModeStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::hostModeStateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "hostModeStateChanged").(func(QBluetoothLocalDevice__HostMode))(QBluetoothLocalDevice__HostMode(state))
}

func (ptr *QBluetoothLocalDevice) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDevice() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::~QBluetoothLocalDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::QBluetoothLocalDevice")
		}
	}()

	return NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(core.PointerFromQObject(parent)))
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::QBluetoothLocalDevice")
		}
	}()

	return NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(PointerFromQBluetoothAddress(address), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::hostMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__HostMode(C.QBluetoothLocalDevice_HostMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothLocalDevice_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(accept bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::pairingConfirmation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmation(ptr.Pointer(), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddress_ITF) QBluetoothLocalDevice__Pairing {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::pairingStatus")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__Pairing(C.QBluetoothLocalDevice_PairingStatus(ptr.Pointer(), PointerFromQBluetoothAddress(address)))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) PowerOn() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::powerOn")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PowerOn(ptr.Pointer())
	}
}

func (ptr *QBluetoothLocalDevice) RequestPairing(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::requestPairing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_RequestPairing(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.int(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) SetHostMode(mode QBluetoothLocalDevice__HostMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothLocalDevice::setHostMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_SetHostMode(ptr.Pointer(), C.int(mode))
	}
}
