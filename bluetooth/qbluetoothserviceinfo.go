package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"log"
	"unsafe"
)

type QBluetoothServiceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothServiceInfo_ITF interface {
	QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo
}

func (p *QBluetoothServiceInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothServiceInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothServiceInfo(ptr QBluetoothServiceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServiceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothServiceInfo {
	var n = new(QBluetoothServiceInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetoothServiceInfo) QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo {
	return ptr
}

//QBluetoothServiceInfo::AttributeId
type QBluetoothServiceInfo__AttributeId int64

var (
	QBluetoothServiceInfo__ServiceRecordHandle              = QBluetoothServiceInfo__AttributeId(0x0000)
	QBluetoothServiceInfo__ServiceClassIds                  = QBluetoothServiceInfo__AttributeId(0x0001)
	QBluetoothServiceInfo__ServiceRecordState               = QBluetoothServiceInfo__AttributeId(0x0002)
	QBluetoothServiceInfo__ServiceId                        = QBluetoothServiceInfo__AttributeId(0x0003)
	QBluetoothServiceInfo__ProtocolDescriptorList           = QBluetoothServiceInfo__AttributeId(0x0004)
	QBluetoothServiceInfo__BrowseGroupList                  = QBluetoothServiceInfo__AttributeId(0x0005)
	QBluetoothServiceInfo__LanguageBaseAttributeIdList      = QBluetoothServiceInfo__AttributeId(0x0006)
	QBluetoothServiceInfo__ServiceInfoTimeToLive            = QBluetoothServiceInfo__AttributeId(0x0007)
	QBluetoothServiceInfo__ServiceAvailability              = QBluetoothServiceInfo__AttributeId(0x0008)
	QBluetoothServiceInfo__BluetoothProfileDescriptorList   = QBluetoothServiceInfo__AttributeId(0x0009)
	QBluetoothServiceInfo__DocumentationUrl                 = QBluetoothServiceInfo__AttributeId(0x000A)
	QBluetoothServiceInfo__ClientExecutableUrl              = QBluetoothServiceInfo__AttributeId(0x000B)
	QBluetoothServiceInfo__IconUrl                          = QBluetoothServiceInfo__AttributeId(0x000C)
	QBluetoothServiceInfo__AdditionalProtocolDescriptorList = QBluetoothServiceInfo__AttributeId(0x000D)
	QBluetoothServiceInfo__PrimaryLanguageBase              = QBluetoothServiceInfo__AttributeId(0x0100)
	QBluetoothServiceInfo__ServiceName                      = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceName_Type())
	QBluetoothServiceInfo__ServiceDescription               = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceDescription_Type())
	QBluetoothServiceInfo__ServiceProvider                  = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceProvider_Type())
)

//QBluetoothServiceInfo::Protocol
type QBluetoothServiceInfo__Protocol int64

const (
	QBluetoothServiceInfo__UnknownProtocol = QBluetoothServiceInfo__Protocol(0)
	QBluetoothServiceInfo__L2capProtocol   = QBluetoothServiceInfo__Protocol(1)
	QBluetoothServiceInfo__RfcommProtocol  = QBluetoothServiceInfo__Protocol(2)
)

func (ptr *QBluetoothServiceInfo) ServiceDescription() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::serviceDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::serviceName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceProvider() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::serviceProvider")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceProvider(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) SetServiceDescription(description string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::setServiceDescription")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::setServiceName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceProvider(provider string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::setServiceProvider")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceProvider(ptr.Pointer(), C.CString(provider))
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceUuid(uuid QBluetoothUuid_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::setServiceUuid")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func NewQBluetoothServiceInfo() *QBluetoothServiceInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::QBluetoothServiceInfo")
		}
	}()

	return NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo())
}

func NewQBluetoothServiceInfo2(other QBluetoothServiceInfo_ITF) *QBluetoothServiceInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::QBluetoothServiceInfo")
		}
	}()

	return NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo2(PointerFromQBluetoothServiceInfo(other)))
}

func (ptr *QBluetoothServiceInfo) IsComplete() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::isComplete")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsRegistered() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::isRegistered")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsRegistered(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ProtocolServiceMultiplexer() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::protocolServiceMultiplexer")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServiceInfo_ProtocolServiceMultiplexer(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) RegisterService(localAdapter QBluetoothAddress_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::registerService")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_RegisterService(ptr.Pointer(), PointerFromQBluetoothAddress(localAdapter)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ServerChannel() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::serverChannel")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServiceInfo_ServerChannel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) SetDevice(device QBluetoothDeviceInfo_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::setDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetDevice(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(device))
	}
}

func (ptr *QBluetoothServiceInfo) SocketProtocol() QBluetoothServiceInfo__Protocol {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::socketProtocol")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServiceInfo_SocketProtocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) UnregisterService() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::unregisterService")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_UnregisterService(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) DestroyQBluetoothServiceInfo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceInfo::~QBluetoothServiceInfo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(ptr.Pointer())
	}
}
