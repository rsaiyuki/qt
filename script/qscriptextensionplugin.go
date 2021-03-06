package script

//#include "script.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QScriptExtensionPlugin struct {
	core.QObject
}

type QScriptExtensionPlugin_ITF interface {
	core.QObject_ITF
	QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin
}

func PointerFromQScriptExtensionPlugin(ptr QScriptExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQScriptExtensionPluginFromPointer(ptr unsafe.Pointer) *QScriptExtensionPlugin {
	var n = new(QScriptExtensionPlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScriptExtensionPlugin_") {
		n.SetObjectName("QScriptExtensionPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QScriptExtensionPlugin) QScriptExtensionPlugin_PTR() *QScriptExtensionPlugin {
	return ptr
}

func (ptr *QScriptExtensionPlugin) Initialize(key string, engine QScriptEngine_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptExtensionPlugin::initialize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_Initialize(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine))
	}
}

func (ptr *QScriptExtensionPlugin) Keys() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptExtensionPlugin::keys")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QScriptExtensionPlugin_Keys(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QScriptExtensionPlugin) SetupPackage(key string, engine QScriptEngine_ITF) *QScriptValue {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptExtensionPlugin::setupPackage")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQScriptValueFromPointer(C.QScriptExtensionPlugin_SetupPackage(ptr.Pointer(), C.CString(key), PointerFromQScriptEngine(engine)))
	}
	return nil
}

func (ptr *QScriptExtensionPlugin) DestroyQScriptExtensionPlugin() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QScriptExtensionPlugin::~QScriptExtensionPlugin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QScriptExtensionPlugin_DestroyQScriptExtensionPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
