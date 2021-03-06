package multimedia

//#include "multimedia.h"
import "C"
import (
	"log"
	"unsafe"
)

type QMediaServiceFeaturesInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceFeaturesInterface_ITF interface {
	QMediaServiceFeaturesInterface_PTR() *QMediaServiceFeaturesInterface
}

func (p *QMediaServiceFeaturesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceFeaturesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceFeaturesInterface(ptr QMediaServiceFeaturesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceFeaturesInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceFeaturesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceFeaturesInterface {
	var n = new(QMediaServiceFeaturesInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaServiceFeaturesInterface) QMediaServiceFeaturesInterface_PTR() *QMediaServiceFeaturesInterface {
	return ptr
}

func (ptr *QMediaServiceFeaturesInterface) DestroyQMediaServiceFeaturesInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaServiceFeaturesInterface::~QMediaServiceFeaturesInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaServiceFeaturesInterface_DestroyQMediaServiceFeaturesInterface(ptr.Pointer())
	}
}
