package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QWinEventNotifier struct {
	QObject
}

type QWinEventNotifier_ITF interface {
	QObject_ITF
	QWinEventNotifier_PTR() *QWinEventNotifier
}

func PointerFromQWinEventNotifier(ptr QWinEventNotifier_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWinEventNotifier_PTR().Pointer()
	}
	return nil
}

func NewQWinEventNotifierFromPointer(ptr unsafe.Pointer) *QWinEventNotifier {
	var n = new(QWinEventNotifier)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWinEventNotifier_") {
		n.SetObjectName("QWinEventNotifier_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWinEventNotifier) QWinEventNotifier_PTR() *QWinEventNotifier {
	return ptr
}
