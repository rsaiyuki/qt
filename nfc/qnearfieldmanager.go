package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QNearFieldManager struct {
	core.QObject
}

type QNearFieldManager_ITF interface {
	core.QObject_ITF
	QNearFieldManager_PTR() *QNearFieldManager
}

func PointerFromQNearFieldManager(ptr QNearFieldManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldManagerFromPointer(ptr unsafe.Pointer) *QNearFieldManager {
	var n = new(QNearFieldManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldManager_") {
		n.SetObjectName("QNearFieldManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNearFieldManager) QNearFieldManager_PTR() *QNearFieldManager {
	return ptr
}

//QNearFieldManager::TargetAccessMode
type QNearFieldManager__TargetAccessMode int64

const (
	QNearFieldManager__NoTargetAccess              = QNearFieldManager__TargetAccessMode(0x00)
	QNearFieldManager__NdefReadTargetAccess        = QNearFieldManager__TargetAccessMode(0x01)
	QNearFieldManager__NdefWriteTargetAccess       = QNearFieldManager__TargetAccessMode(0x02)
	QNearFieldManager__TagTypeSpecificTargetAccess = QNearFieldManager__TargetAccessMode(0x04)
)

func (ptr *QNearFieldManager) RegisterNdefMessageHandler(object core.QObject_ITF, method string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::registerNdefMessageHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler(ptr.Pointer(), core.PointerFromQObject(object), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) StartTargetDetection() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::startTargetDetection")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_StartTargetDetection(ptr.Pointer()) != 0
	}
	return false
}

func NewQNearFieldManager(parent core.QObject_ITF) *QNearFieldManager {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::QNearFieldManager")
		}
	}()

	return NewQNearFieldManagerFromPointer(C.QNearFieldManager_NewQNearFieldManager(core.PointerFromQObject(parent)))
}

func (ptr *QNearFieldManager) IsAvailable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::isAvailable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, object core.QObject_ITF, method string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::registerNdefMessageHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler2(ptr.Pointer(), C.int(typeNameFormat), core.PointerFromQByteArray(ty), core.PointerFromQObject(object), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler3(filter QNdefFilter_ITF, object core.QObject_ITF, method string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::registerNdefMessageHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler3(ptr.Pointer(), PointerFromQNdefFilter(filter), core.PointerFromQObject(object), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) SetTargetAccessModes(accessModes QNearFieldManager__TargetAccessMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::setTargetAccessModes")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldManager_SetTargetAccessModes(ptr.Pointer(), C.int(accessModes))
	}
}

func (ptr *QNearFieldManager) StopTargetDetection() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::stopTargetDetection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldManager_StopTargetDetection(ptr.Pointer())
	}
}

func (ptr *QNearFieldManager) TargetAccessModes() QNearFieldManager__TargetAccessMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::targetAccessModes")
		}
	}()

	if ptr.Pointer() != nil {
		return QNearFieldManager__TargetAccessMode(C.QNearFieldManager_TargetAccessModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldManager) ConnectTargetDetected(f func(target *QNearFieldTarget)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::targetDetected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetDetected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::targetDetected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldManagerTargetDetected
func callbackQNearFieldManagerTargetDetected(ptrName *C.char, target unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::targetDetected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "targetDetected").(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
}

func (ptr *QNearFieldManager) ConnectTargetLost(f func(target *QNearFieldTarget)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::targetLost")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetLost(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetLost", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetLost() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::targetLost")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetLost(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetLost")
	}
}

//export callbackQNearFieldManagerTargetLost
func callbackQNearFieldManagerTargetLost(ptrName *C.char, target unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::targetLost")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "targetLost").(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
}

func (ptr *QNearFieldManager) UnregisterNdefMessageHandler(handlerId int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::unregisterNdefMessageHandler")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_UnregisterNdefMessageHandler(ptr.Pointer(), C.int(handlerId)) != 0
	}
	return false
}

func (ptr *QNearFieldManager) DestroyQNearFieldManager() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNearFieldManager::~QNearFieldManager")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DestroyQNearFieldManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
