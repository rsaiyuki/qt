package quick

//#include "quick.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSGOpacityNode struct {
	QSGNode
}

type QSGOpacityNode_ITF interface {
	QSGNode_ITF
	QSGOpacityNode_PTR() *QSGOpacityNode
}

func PointerFromQSGOpacityNode(ptr QSGOpacityNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpacityNode_PTR().Pointer()
	}
	return nil
}

func NewQSGOpacityNodeFromPointer(ptr unsafe.Pointer) *QSGOpacityNode {
	var n = new(QSGOpacityNode)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGOpacityNode) QSGOpacityNode_PTR() *QSGOpacityNode {
	return ptr
}

func NewQSGOpacityNode() *QSGOpacityNode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpacityNode::QSGOpacityNode")
		}
	}()

	return NewQSGOpacityNodeFromPointer(C.QSGOpacityNode_NewQSGOpacityNode())
}

func (ptr *QSGOpacityNode) Opacity() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpacityNode::opacity")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QSGOpacityNode_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpacityNode) SetOpacity(opacity float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpacityNode::setOpacity")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGOpacityNode_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNode() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGOpacityNode::~QSGOpacityNode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGOpacityNode_DestroyQSGOpacityNode(ptr.Pointer())
	}
}
