package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"log"
	"unsafe"
)

type QMediaPlaylist struct {
	core.QObject
	QMediaBindableInterface
}

type QMediaPlaylist_ITF interface {
	core.QObject_ITF
	QMediaBindableInterface_ITF
	QMediaPlaylist_PTR() *QMediaPlaylist
}

func (p *QMediaPlaylist) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QMediaPlaylist) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
}

func PointerFromQMediaPlaylist(ptr QMediaPlaylist_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaPlaylist_PTR().Pointer()
	}
	return nil
}

func NewQMediaPlaylistFromPointer(ptr unsafe.Pointer) *QMediaPlaylist {
	var n = new(QMediaPlaylist)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaPlaylist_") {
		n.SetObjectName("QMediaPlaylist_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaPlaylist) QMediaPlaylist_PTR() *QMediaPlaylist {
	return ptr
}

//QMediaPlaylist::Error
type QMediaPlaylist__Error int64

const (
	QMediaPlaylist__NoError                 = QMediaPlaylist__Error(0)
	QMediaPlaylist__FormatError             = QMediaPlaylist__Error(1)
	QMediaPlaylist__FormatNotSupportedError = QMediaPlaylist__Error(2)
	QMediaPlaylist__NetworkError            = QMediaPlaylist__Error(3)
	QMediaPlaylist__AccessDeniedError       = QMediaPlaylist__Error(4)
)

//QMediaPlaylist::PlaybackMode
type QMediaPlaylist__PlaybackMode int64

const (
	QMediaPlaylist__CurrentItemOnce   = QMediaPlaylist__PlaybackMode(0)
	QMediaPlaylist__CurrentItemInLoop = QMediaPlaylist__PlaybackMode(1)
	QMediaPlaylist__Sequential        = QMediaPlaylist__PlaybackMode(2)
	QMediaPlaylist__Loop              = QMediaPlaylist__PlaybackMode(3)
	QMediaPlaylist__Random            = QMediaPlaylist__PlaybackMode(4)
)

func (ptr *QMediaPlaylist) PlaybackMode() QMediaPlaylist__PlaybackMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::playbackMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QMediaPlaylist__PlaybackMode(C.QMediaPlaylist_PlaybackMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) SetPlaybackMode(mode QMediaPlaylist__PlaybackMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::setPlaybackMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_SetPlaybackMode(ptr.Pointer(), C.int(mode))
	}
}

func NewQMediaPlaylist(parent core.QObject_ITF) *QMediaPlaylist {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::QMediaPlaylist")
		}
	}()

	return NewQMediaPlaylistFromPointer(C.QMediaPlaylist_NewQMediaPlaylist(core.PointerFromQObject(parent)))
}

func (ptr *QMediaPlaylist) AddMedia(content QMediaContent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::addMedia")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_AddMedia(ptr.Pointer(), PointerFromQMediaContent(content)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Clear() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::clear")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Clear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) CurrentIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::currentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_CurrentIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectCurrentIndexChanged(f func(position int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::currentIndexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectCurrentIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIndexChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectCurrentIndexChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::currentIndexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectCurrentIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIndexChanged")
	}
}

//export callbackQMediaPlaylistCurrentIndexChanged
func callbackQMediaPlaylistCurrentIndexChanged(ptrName *C.char, position C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::currentIndexChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentIndexChanged").(func(int))(int(position))
}

func (ptr *QMediaPlaylist) Error() QMediaPlaylist__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QMediaPlaylist__Error(C.QMediaPlaylist_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaPlaylist_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaPlaylist) InsertMedia(pos int, content QMediaContent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::insertMedia")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_InsertMedia(ptr.Pointer(), C.int(pos), PointerFromQMediaContent(content)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) IsReadOnly() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::isReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Load3(device core.QIODevice_ITF, format string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::load")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load3(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) Load(request network.QNetworkRequest_ITF, format string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::load")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load(ptr.Pointer(), network.PointerFromQNetworkRequest(request), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) Load2(location core.QUrl_ITF, format string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::load")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Load2(ptr.Pointer(), core.PointerFromQUrl(location), C.CString(format))
	}
}

func (ptr *QMediaPlaylist) ConnectLoadFailed(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::loadFailed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectLoadFailed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loadFailed", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectLoadFailed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::loadFailed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectLoadFailed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loadFailed")
	}
}

//export callbackQMediaPlaylistLoadFailed
func callbackQMediaPlaylistLoadFailed(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::loadFailed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "loadFailed").(func())()
}

func (ptr *QMediaPlaylist) ConnectLoaded(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::loaded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectLoaded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "loaded", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectLoaded() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::loaded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectLoaded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "loaded")
	}
}

//export callbackQMediaPlaylistLoaded
func callbackQMediaPlaylistLoaded(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::loaded")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "loaded").(func())()
}

func (ptr *QMediaPlaylist) ConnectMediaAboutToBeInserted(f func(start int, end int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaAboutToBeInserted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaAboutToBeInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaAboutToBeInserted", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaAboutToBeInserted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaAboutToBeInserted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaAboutToBeInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaAboutToBeInserted")
	}
}

//export callbackQMediaPlaylistMediaAboutToBeInserted
func callbackQMediaPlaylistMediaAboutToBeInserted(ptrName *C.char, start C.int, end C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaAboutToBeInserted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mediaAboutToBeInserted").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) ConnectMediaAboutToBeRemoved(f func(start int, end int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaAboutToBeRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaAboutToBeRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaAboutToBeRemoved", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaAboutToBeRemoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaAboutToBeRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaAboutToBeRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaAboutToBeRemoved")
	}
}

//export callbackQMediaPlaylistMediaAboutToBeRemoved
func callbackQMediaPlaylistMediaAboutToBeRemoved(ptrName *C.char, start C.int, end C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaAboutToBeRemoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mediaAboutToBeRemoved").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) ConnectMediaChanged(f func(start int, end int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaChanged")
	}
}

//export callbackQMediaPlaylistMediaChanged
func callbackQMediaPlaylistMediaChanged(ptrName *C.char, start C.int, end C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mediaChanged").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) MediaCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_MediaCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectMediaInserted(f func(start int, end int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaInserted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaInserted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaInserted", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaInserted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaInserted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaInserted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaInserted")
	}
}

//export callbackQMediaPlaylistMediaInserted
func callbackQMediaPlaylistMediaInserted(ptrName *C.char, start C.int, end C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaInserted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mediaInserted").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) MediaObject() *QMediaObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaObject")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QMediaPlaylist_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMediaPlaylist) ConnectMediaRemoved(f func(start int, end int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectMediaRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mediaRemoved", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectMediaRemoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectMediaRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mediaRemoved")
	}
}

//export callbackQMediaPlaylistMediaRemoved
func callbackQMediaPlaylistMediaRemoved(ptrName *C.char, start C.int, end C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::mediaRemoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "mediaRemoved").(func(int, int))(int(start), int(end))
}

func (ptr *QMediaPlaylist) Next() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::next")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Next(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) NextIndex(steps int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::nextIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_NextIndex(ptr.Pointer(), C.int(steps)))
	}
	return 0
}

func (ptr *QMediaPlaylist) ConnectPlaybackModeChanged(f func(mode QMediaPlaylist__PlaybackMode)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::playbackModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_ConnectPlaybackModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "playbackModeChanged", f)
	}
}

func (ptr *QMediaPlaylist) DisconnectPlaybackModeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::playbackModeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DisconnectPlaybackModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "playbackModeChanged")
	}
}

//export callbackQMediaPlaylistPlaybackModeChanged
func callbackQMediaPlaylistPlaybackModeChanged(ptrName *C.char, mode C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::playbackModeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "playbackModeChanged").(func(QMediaPlaylist__PlaybackMode))(QMediaPlaylist__PlaybackMode(mode))
}

func (ptr *QMediaPlaylist) Previous() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::previous")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Previous(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) PreviousIndex(steps int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::previousIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QMediaPlaylist_PreviousIndex(ptr.Pointer(), C.int(steps)))
	}
	return 0
}

func (ptr *QMediaPlaylist) RemoveMedia(pos int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::removeMedia")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_RemoveMedia(ptr.Pointer(), C.int(pos)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) RemoveMedia2(start int, end int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::removeMedia")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_RemoveMedia2(ptr.Pointer(), C.int(start), C.int(end)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Save2(device core.QIODevice_ITF, format string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::save")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Save2(ptr.Pointer(), core.PointerFromQIODevice(device), C.CString(format)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) Save(location core.QUrl_ITF, format string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::save")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMediaPlaylist_Save(ptr.Pointer(), core.PointerFromQUrl(location), C.CString(format)) != 0
	}
	return false
}

func (ptr *QMediaPlaylist) SetCurrentIndex(playlistPosition int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::setCurrentIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_SetCurrentIndex(ptr.Pointer(), C.int(playlistPosition))
	}
}

func (ptr *QMediaPlaylist) Shuffle() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::shuffle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_Shuffle(ptr.Pointer())
	}
}

func (ptr *QMediaPlaylist) DestroyQMediaPlaylist() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaPlaylist::~QMediaPlaylist")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaPlaylist_DestroyQMediaPlaylist(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
