package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QSoundEffect struct {
	core.QObject
}

type QSoundEffect_ITF interface {
	core.QObject_ITF
	QSoundEffect_PTR() *QSoundEffect
}

func PointerFromQSoundEffect(ptr QSoundEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSoundEffect_PTR().Pointer()
	}
	return nil
}

func NewQSoundEffectFromPointer(ptr unsafe.Pointer) *QSoundEffect {
	var n = new(QSoundEffect)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSoundEffect_") {
		n.SetObjectName("QSoundEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSoundEffect) QSoundEffect_PTR() *QSoundEffect {
	return ptr
}

//QSoundEffect::Loop
type QSoundEffect__Loop int64

const (
	QSoundEffect__Infinite = QSoundEffect__Loop(-2)
)

//QSoundEffect::Status
type QSoundEffect__Status int64

const (
	QSoundEffect__Null    = QSoundEffect__Status(0)
	QSoundEffect__Loading = QSoundEffect__Status(1)
	QSoundEffect__Ready   = QSoundEffect__Status(2)
	QSoundEffect__Error   = QSoundEffect__Status(3)
)

func (ptr *QSoundEffect) IsLoaded() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::isLoaded")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsLoaded(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) LoopsRemaining() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loopsRemaining")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopsRemaining(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) Play() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::play")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_Play(ptr.Pointer())
	}
}

func (ptr *QSoundEffect) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_Stop(ptr.Pointer())
	}
}

func QSoundEffect_SupportedMimeTypes() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::supportedMimeTypes")
		}
	}()

	return strings.Split(C.GoString(C.QSoundEffect_QSoundEffect_SupportedMimeTypes()), ",,,")
}

func NewQSoundEffect(parent core.QObject_ITF) *QSoundEffect {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::QSoundEffect")
		}
	}()

	return NewQSoundEffectFromPointer(C.QSoundEffect_NewQSoundEffect(core.PointerFromQObject(parent)))
}

func (ptr *QSoundEffect) Category() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::category")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSoundEffect_Category(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSoundEffect) ConnectCategoryChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::categoryChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectCategoryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "categoryChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectCategoryChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::categoryChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectCategoryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "categoryChanged")
	}
}

//export callbackQSoundEffectCategoryChanged
func callbackQSoundEffectCategoryChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::categoryChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "categoryChanged").(func())()
}

func (ptr *QSoundEffect) IsMuted() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::isMuted")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsMuted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) IsPlaying() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::isPlaying")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSoundEffect_IsPlaying(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSoundEffect) ConnectLoadedChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loadedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoadedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoadedChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loadedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoadedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadedChanged")
	}
}

//export callbackQSoundEffectLoadedChanged
func callbackQSoundEffectLoadedChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loadedChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "loadedChanged").(func())()
}

func (ptr *QSoundEffect) LoopCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loopCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSoundEffect_LoopCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectLoopCountChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loopCountChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopCountChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loopCountChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopCountChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loopCountChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopCountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loopCountChanged")
	}
}

//export callbackQSoundEffectLoopCountChanged
func callbackQSoundEffectLoopCountChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loopCountChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "loopCountChanged").(func())()
}

func (ptr *QSoundEffect) ConnectLoopsRemainingChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loopsRemainingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectLoopsRemainingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loopsRemainingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectLoopsRemainingChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loopsRemainingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectLoopsRemainingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loopsRemainingChanged")
	}
}

//export callbackQSoundEffectLoopsRemainingChanged
func callbackQSoundEffectLoopsRemainingChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::loopsRemainingChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "loopsRemainingChanged").(func())()
}

func (ptr *QSoundEffect) ConnectMutedChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectMutedChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mutedChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectMutedChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::mutedChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectMutedChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mutedChanged")
	}
}

//export callbackQSoundEffectMutedChanged
func callbackQSoundEffectMutedChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::mutedChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mutedChanged").(func())()
}

func (ptr *QSoundEffect) ConnectPlayingChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::playingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectPlayingChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playingChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectPlayingChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::playingChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectPlayingChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playingChanged")
	}
}

//export callbackQSoundEffectPlayingChanged
func callbackQSoundEffectPlayingChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::playingChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "playingChanged").(func())()
}

func (ptr *QSoundEffect) SetCategory(category string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::setCategory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetCategory(ptr.Pointer(), C.CString(category))
	}
}

func (ptr *QSoundEffect) SetLoopCount(loopCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::setLoopCount")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetLoopCount(ptr.Pointer(), C.int(loopCount))
	}
}

func (ptr *QSoundEffect) SetMuted(muted bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::setMuted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetMuted(ptr.Pointer(), C.int(qt.GoBoolToInt(muted)))
	}
}

func (ptr *QSoundEffect) SetSource(url core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::setSource")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QSoundEffect) SetVolume(volume float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::setVolume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QSoundEffect) ConnectSourceChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::sourceChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectSourceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sourceChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectSourceChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::sourceChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectSourceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sourceChanged")
	}
}

//export callbackQSoundEffectSourceChanged
func callbackQSoundEffectSourceChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::sourceChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "sourceChanged").(func())()
}

func (ptr *QSoundEffect) Status() QSoundEffect__Status {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::status")
		}
	}()

	if ptr.Pointer() != nil {
		return QSoundEffect__Status(C.QSoundEffect_Status(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectStatusChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::statusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectStatusChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::statusChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQSoundEffectStatusChanged
func callbackQSoundEffectStatusChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::statusChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "statusChanged").(func())()
}

func (ptr *QSoundEffect) Volume() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::volume")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QSoundEffect_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSoundEffect) ConnectVolumeChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::volumeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "volumeChanged", f)
	}
}

func (ptr *QSoundEffect) DisconnectVolumeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::volumeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "volumeChanged")
	}
}

//export callbackQSoundEffectVolumeChanged
func callbackQSoundEffectVolumeChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::volumeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "volumeChanged").(func())()
}

func (ptr *QSoundEffect) DestroyQSoundEffect() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSoundEffect::~QSoundEffect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSoundEffect_DestroyQSoundEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
