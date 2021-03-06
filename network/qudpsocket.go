package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QUdpSocket struct {
	QAbstractSocket
}

type QUdpSocket_ITF interface {
	QAbstractSocket_ITF
	QUdpSocket_PTR() *QUdpSocket
}

func PointerFromQUdpSocket(ptr QUdpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUdpSocket_PTR().Pointer()
	}
	return nil
}

func NewQUdpSocketFromPointer(ptr unsafe.Pointer) *QUdpSocket {
	var n = new(QUdpSocket)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QUdpSocket_") {
		n.SetObjectName("QUdpSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QUdpSocket) QUdpSocket_PTR() *QUdpSocket {
	return ptr
}

func NewQUdpSocket(parent core.QObject_ITF) *QUdpSocket {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUdpSocket::QUdpSocket")
		}
	}()

	return NewQUdpSocketFromPointer(C.QUdpSocket_NewQUdpSocket(core.PointerFromQObject(parent)))
}

func (ptr *QUdpSocket) HasPendingDatagrams() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUdpSocket::hasPendingDatagrams")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUdpSocket_HasPendingDatagrams(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUdpSocket::joinMulticastGroup")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUdpSocket::joinMulticastGroup")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUdpSocket::leaveMulticastGroup")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUdpSocket::leaveMulticastGroup")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUdpSocket::setMulticastInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetMulticastInterface(ptr.Pointer(), PointerFromQNetworkInterface(iface))
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocket() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUdpSocket::~QUdpSocket")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUdpSocket_DestroyQUdpSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
