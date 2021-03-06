package network

//#include "network.h"
import "C"
import (
	"log"
	"unsafe"
)

type QDnsTextRecord struct {
	ptr unsafe.Pointer
}

type QDnsTextRecord_ITF interface {
	QDnsTextRecord_PTR() *QDnsTextRecord
}

func (p *QDnsTextRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsTextRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsTextRecord(ptr QDnsTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsTextRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsTextRecordFromPointer(ptr unsafe.Pointer) *QDnsTextRecord {
	var n = new(QDnsTextRecord)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDnsTextRecord) QDnsTextRecord_PTR() *QDnsTextRecord {
	return ptr
}

func NewQDnsTextRecord() *QDnsTextRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsTextRecord::QDnsTextRecord")
		}
	}()

	return NewQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord())
}

func NewQDnsTextRecord2(other QDnsTextRecord_ITF) *QDnsTextRecord {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsTextRecord::QDnsTextRecord")
		}
	}()

	return NewQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord2(PointerFromQDnsTextRecord(other)))
}

func (ptr *QDnsTextRecord) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsTextRecord::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsTextRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsTextRecord) Swap(other QDnsTextRecord_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsTextRecord::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsTextRecord_Swap(ptr.Pointer(), PointerFromQDnsTextRecord(other))
	}
}

func (ptr *QDnsTextRecord) DestroyQDnsTextRecord() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDnsTextRecord::~QDnsTextRecord")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDnsTextRecord_DestroyQDnsTextRecord(ptr.Pointer())
	}
}
