package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QObjectCleanupHandler struct {
	QObject
}

type QObjectCleanupHandler_ITF interface {
	QObject_ITF
	QObjectCleanupHandler_PTR() *QObjectCleanupHandler
}

func PointerFromQObjectCleanupHandler(ptr QObjectCleanupHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QObjectCleanupHandler_PTR().Pointer()
	}
	return nil
}

func NewQObjectCleanupHandlerFromPointer(ptr unsafe.Pointer) *QObjectCleanupHandler {
	var n = new(QObjectCleanupHandler)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QObjectCleanupHandler_") {
		n.SetObjectName("QObjectCleanupHandler_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QObjectCleanupHandler) QObjectCleanupHandler_PTR() *QObjectCleanupHandler {
	return ptr
}

func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObjectCleanupHandler::QObjectCleanupHandler")
		}
	}()

	return NewQObjectCleanupHandlerFromPointer(C.QObjectCleanupHandler_NewQObjectCleanupHandler())
}

func (ptr *QObjectCleanupHandler) Add(object QObject_ITF) *QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObjectCleanupHandler::add")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QObjectCleanupHandler_Add(ptr.Pointer(), PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QObjectCleanupHandler) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObjectCleanupHandler::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObjectCleanupHandler_Clear(ptr.Pointer())
	}
}

func (ptr *QObjectCleanupHandler) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObjectCleanupHandler::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QObjectCleanupHandler_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObjectCleanupHandler) DestroyQObjectCleanupHandler() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QObjectCleanupHandler::~QObjectCleanupHandler")
		}
	}()

	if ptr.Pointer() != nil {
		C.QObjectCleanupHandler_DestroyQObjectCleanupHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
