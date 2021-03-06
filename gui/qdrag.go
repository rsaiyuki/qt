package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDrag struct {
	core.QObject
}

type QDrag_ITF interface {
	core.QObject_ITF
	QDrag_PTR() *QDrag
}

func PointerFromQDrag(ptr QDrag_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDrag_PTR().Pointer()
	}
	return nil
}

func NewQDragFromPointer(ptr unsafe.Pointer) *QDrag {
	var n = new(QDrag)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDrag_") {
		n.SetObjectName("QDrag_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDrag) QDrag_PTR() *QDrag {
	return ptr
}

func NewQDrag(dragSource core.QObject_ITF) *QDrag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::QDrag")
		}
	}()

	return NewQDragFromPointer(C.QDrag_NewQDrag(core.PointerFromQObject(dragSource)))
}

func (ptr *QDrag) ConnectActionChanged(f func(action core.Qt__DropAction)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::actionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_ConnectActionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actionChanged", f)
	}
}

func (ptr *QDrag) DisconnectActionChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::actionChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_DisconnectActionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actionChanged")
	}
}

//export callbackQDragActionChanged
func callbackQDragActionChanged(ptrName *C.char, action C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::actionChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "actionChanged").(func(core.Qt__DropAction))(core.Qt__DropAction(action))
}

func (ptr *QDrag) Exec(supportedActions core.Qt__DropAction) core.Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::exec")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_Exec(ptr.Pointer(), C.int(supportedActions)))
	}
	return 0
}

func (ptr *QDrag) Exec2(supportedActions core.Qt__DropAction, defaultDropAction core.Qt__DropAction) core.Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::exec")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_Exec2(ptr.Pointer(), C.int(supportedActions), C.int(defaultDropAction)))
	}
	return 0
}

func (ptr *QDrag) MimeData() *core.QMimeData {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::mimeData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQMimeDataFromPointer(C.QDrag_MimeData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDrag) SetDragCursor(cursor QPixmap_ITF, action core.Qt__DropAction) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::setDragCursor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_SetDragCursor(ptr.Pointer(), PointerFromQPixmap(cursor), C.int(action))
	}
}

func (ptr *QDrag) SetHotSpot(hotspot core.QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::setHotSpot")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_SetHotSpot(ptr.Pointer(), core.PointerFromQPoint(hotspot))
	}
}

func (ptr *QDrag) SetMimeData(data core.QMimeData_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::setMimeData")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_SetMimeData(ptr.Pointer(), core.PointerFromQMimeData(data))
	}
}

func (ptr *QDrag) SetPixmap(pixmap QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::setPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_SetPixmap(ptr.Pointer(), PointerFromQPixmap(pixmap))
	}
}

func (ptr *QDrag) Source() *core.QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::source")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QDrag_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDrag) SupportedActions() core.Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::supportedActions")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QDrag_SupportedActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDrag) Target() *core.QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::target")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QDrag_Target(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDrag) ConnectTargetChanged(f func(newTarget *core.QObject)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::targetChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_ConnectTargetChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetChanged", f)
	}
}

func (ptr *QDrag) DisconnectTargetChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::targetChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_DisconnectTargetChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetChanged")
	}
}

//export callbackQDragTargetChanged
func callbackQDragTargetChanged(ptrName *C.char, newTarget unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::targetChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "targetChanged").(func(*core.QObject))(core.NewQObjectFromPointer(newTarget))
}

func (ptr *QDrag) DestroyQDrag() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDrag::~QDrag")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDrag_DestroyQDrag(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
