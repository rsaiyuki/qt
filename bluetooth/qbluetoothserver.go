package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QBluetoothServer struct {
	core.QObject
}

type QBluetoothServer_ITF interface {
	core.QObject_ITF
	QBluetoothServer_PTR() *QBluetoothServer
}

func PointerFromQBluetoothServer(ptr QBluetoothServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServer_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServerFromPointer(ptr unsafe.Pointer) *QBluetoothServer {
	var n = new(QBluetoothServer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothServer_") {
		n.SetObjectName("QBluetoothServer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothServer) QBluetoothServer_PTR() *QBluetoothServer {
	return ptr
}

//QBluetoothServer::Error
type QBluetoothServer__Error int64

const (
	QBluetoothServer__NoError                       = QBluetoothServer__Error(0)
	QBluetoothServer__UnknownError                  = QBluetoothServer__Error(1)
	QBluetoothServer__PoweredOffError               = QBluetoothServer__Error(2)
	QBluetoothServer__InputOutputError              = QBluetoothServer__Error(3)
	QBluetoothServer__ServiceAlreadyRegisteredError = QBluetoothServer__Error(4)
	QBluetoothServer__UnsupportedProtocolError      = QBluetoothServer__Error(5)
)

func NewQBluetoothServer(serverType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothServer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::QBluetoothServer")
		}
	}()

	return NewQBluetoothServerFromPointer(C.QBluetoothServer_NewQBluetoothServer(C.int(serverType), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::newConnection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QBluetoothServer) DisconnectNewConnection() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::newConnection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQBluetoothServerNewConnection
func callbackQBluetoothServerNewConnection(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::newConnection")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "newConnection").(func())()
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothServer__Error(C.QBluetoothServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) IsListening() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::isListening")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::maxPendingConnections")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::serverType")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServer_ServerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::~QBluetoothServer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DestroyQBluetoothServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) Close() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::close")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServer_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) HasPendingConnections() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::hasPendingConnections")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::nextPendingConnection")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBluetoothSocketFromPointer(C.QBluetoothServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServer) SetMaxPendingConnections(numConnections int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServer::setMaxPendingConnections")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}
