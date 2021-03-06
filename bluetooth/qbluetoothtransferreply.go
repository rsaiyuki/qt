package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QBluetoothTransferReply struct {
	core.QObject
}

type QBluetoothTransferReply_ITF interface {
	core.QObject_ITF
	QBluetoothTransferReply_PTR() *QBluetoothTransferReply
}

func PointerFromQBluetoothTransferReply(ptr QBluetoothTransferReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferReply_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) *QBluetoothTransferReply {
	var n = new(QBluetoothTransferReply)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothTransferReply_") {
		n.SetObjectName("QBluetoothTransferReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothTransferReply) QBluetoothTransferReply_PTR() *QBluetoothTransferReply {
	return ptr
}

//QBluetoothTransferReply::TransferError
type QBluetoothTransferReply__TransferError int64

const (
	QBluetoothTransferReply__NoError                   = QBluetoothTransferReply__TransferError(0)
	QBluetoothTransferReply__UnknownError              = QBluetoothTransferReply__TransferError(1)
	QBluetoothTransferReply__FileNotFoundError         = QBluetoothTransferReply__TransferError(2)
	QBluetoothTransferReply__HostNotFoundError         = QBluetoothTransferReply__TransferError(3)
	QBluetoothTransferReply__UserCanceledTransferError = QBluetoothTransferReply__TransferError(4)
	QBluetoothTransferReply__IODeviceNotReadableError  = QBluetoothTransferReply__TransferError(5)
	QBluetoothTransferReply__ResourceBusyError         = QBluetoothTransferReply__TransferError(6)
	QBluetoothTransferReply__SessionError              = QBluetoothTransferReply__TransferError(7)
)

func (ptr *QBluetoothTransferReply) Abort() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::abort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothTransferReply) Error() QBluetoothTransferReply__TransferError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothTransferReply__TransferError(C.QBluetoothTransferReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothTransferReply) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothTransferReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothTransferReplyFinished
func callbackQBluetoothTransferReplyFinished(ptrName *C.char, reply unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
}

func (ptr *QBluetoothTransferReply) IsFinished() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::isFinished")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) IsRunning() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::isRunning")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::manager")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferReply_Manager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferReply::~QBluetoothTransferReply")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
