package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSGTextureProvider struct {
	core.QObject
}

type QSGTextureProvider_ITF interface {
	core.QObject_ITF
	QSGTextureProvider_PTR() *QSGTextureProvider
}

func PointerFromQSGTextureProvider(ptr QSGTextureProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureProvider_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureProviderFromPointer(ptr unsafe.Pointer) *QSGTextureProvider {
	var n = new(QSGTextureProvider)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSGTextureProvider_") {
		n.SetObjectName("QSGTextureProvider_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSGTextureProvider) QSGTextureProvider_PTR() *QSGTextureProvider {
	return ptr
}

func (ptr *QSGTextureProvider) Texture() *QSGTexture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGTextureProvider::texture")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGTextureProvider_Texture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureProvider) ConnectTextureChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGTextureProvider::textureChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectTextureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textureChanged", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTextureChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGTextureProvider::textureChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectTextureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textureChanged")
	}
}

//export callbackQSGTextureProviderTextureChanged
func callbackQSGTextureProviderTextureChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSGTextureProvider::textureChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "textureChanged").(func())()
}
