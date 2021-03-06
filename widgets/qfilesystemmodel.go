package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QFileSystemModel struct {
	core.QAbstractItemModel
}

type QFileSystemModel_ITF interface {
	core.QAbstractItemModel_ITF
	QFileSystemModel_PTR() *QFileSystemModel
}

func PointerFromQFileSystemModel(ptr QFileSystemModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSystemModel_PTR().Pointer()
	}
	return nil
}

func NewQFileSystemModelFromPointer(ptr unsafe.Pointer) *QFileSystemModel {
	var n = new(QFileSystemModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFileSystemModel_") {
		n.SetObjectName("QFileSystemModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileSystemModel) QFileSystemModel_PTR() *QFileSystemModel {
	return ptr
}

//QFileSystemModel::Roles
type QFileSystemModel__Roles int64

var (
	QFileSystemModel__FileIconRole    = QFileSystemModel__Roles(core.Qt__DecorationRole)
	QFileSystemModel__FilePathRole    = QFileSystemModel__Roles(C.QFileSystemModel_FilePathRole_Type())
	QFileSystemModel__FileNameRole    = QFileSystemModel__Roles(C.QFileSystemModel_FileNameRole_Type())
	QFileSystemModel__FilePermissions = QFileSystemModel__Roles(C.QFileSystemModel_FilePermissions_Type())
)

func (ptr *QFileSystemModel) IsReadOnly() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::isReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileSystemModel) NameFilterDisables() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::nameFilterDisables")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_NameFilterDisables(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileSystemModel) ResolveSymlinks() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::resolveSymlinks")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_ResolveSymlinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileSystemModel) Rmdir(index core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::rmdir")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_Rmdir(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) SetNameFilterDisables(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::setNameFilterDisables")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetNameFilterDisables(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileSystemModel) SetReadOnly(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::setReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileSystemModel) SetResolveSymlinks(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::setResolveSymlinks")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetResolveSymlinks(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func NewQFileSystemModel(parent core.QObject_ITF) *QFileSystemModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::QFileSystemModel")
		}
	}()

	return NewQFileSystemModelFromPointer(C.QFileSystemModel_NewQFileSystemModel(core.PointerFromQObject(parent)))
}

func (ptr *QFileSystemModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::canFetchMore")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) ColumnCount(parent core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFileSystemModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QFileSystemModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QFileSystemModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QFileSystemModel) ConnectDirectoryLoaded(f func(path string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::directoryLoaded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectDirectoryLoaded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directoryLoaded", f)
	}
}

func (ptr *QFileSystemModel) DisconnectDirectoryLoaded() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::directoryLoaded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectDirectoryLoaded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directoryLoaded")
	}
}

//export callbackQFileSystemModelDirectoryLoaded
func callbackQFileSystemModelDirectoryLoaded(ptrName *C.char, path *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::directoryLoaded")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "directoryLoaded").(func(string))(C.GoString(path))
}

func (ptr *QFileSystemModel) DropMimeData(data core.QMimeData_ITF, action core.Qt__DropAction, row int, column int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::dropMimeData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_DropMimeData(ptr.Pointer(), core.PointerFromQMimeData(data), C.int(action), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) FetchMore(parent core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::fetchMore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QFileSystemModel) FileName(index core.QModelIndex_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::fileName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_FileName(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return ""
}

func (ptr *QFileSystemModel) FilePath(index core.QModelIndex_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::filePath")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_FilePath(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return ""
}

func (ptr *QFileSystemModel) ConnectFileRenamed(f func(path string, oldName string, newName string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::fileRenamed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectFileRenamed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fileRenamed", f)
	}
}

func (ptr *QFileSystemModel) DisconnectFileRenamed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::fileRenamed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectFileRenamed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fileRenamed")
	}
}

//export callbackQFileSystemModelFileRenamed
func callbackQFileSystemModelFileRenamed(ptrName *C.char, path *C.char, oldName *C.char, newName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::fileRenamed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "fileRenamed").(func(string, string, string))(C.GoString(path), C.GoString(oldName), C.GoString(newName))
}

func (ptr *QFileSystemModel) Filter() core.QDir__Filter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return core.QDir__Filter(C.QFileSystemModel_Filter(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileSystemModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QFileSystemModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QFileSystemModel) HasChildren(parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::hasChildren")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_HasChildren(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::headerData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QFileSystemModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QFileSystemModel) IconProvider() *QFileIconProvider {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::iconProvider")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQFileIconProviderFromPointer(C.QFileSystemModel_IconProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileSystemModel) Index2(path string, column int) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::index")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_Index2(ptr.Pointer(), C.CString(path), C.int(column)))
	}
	return nil
}

func (ptr *QFileSystemModel) Index(row int, column int, parent core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::index")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_Index(ptr.Pointer(), C.int(row), C.int(column), core.PointerFromQModelIndex(parent)))
	}
	return nil
}

func (ptr *QFileSystemModel) IsDir(index core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::isDir")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_IsDir(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) LastModified(index core.QModelIndex_ITF) *core.QDateTime {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::lastModified")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QFileSystemModel_LastModified(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QFileSystemModel) MimeTypes() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::mimeTypes")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemModel_MimeTypes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemModel) Mkdir(parent core.QModelIndex_ITF, name string) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::mkdir")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_Mkdir(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.CString(name)))
	}
	return nil
}

func (ptr *QFileSystemModel) MyComputer(role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::myComputer")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QFileSystemModel_MyComputer(ptr.Pointer(), C.int(role)))
	}
	return nil
}

func (ptr *QFileSystemModel) NameFilters() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::nameFilters")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemModel_NameFilters(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemModel) Parent(index core.QModelIndex_ITF) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::parent")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_Parent(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QFileSystemModel) RootDirectory() *core.QDir {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::rootDirectory")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQDirFromPointer(C.QFileSystemModel_RootDirectory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileSystemModel) RootPath() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::rootPath")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_RootPath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileSystemModel) ConnectRootPathChanged(f func(newPath string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::rootPathChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_ConnectRootPathChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rootPathChanged", f)
	}
}

func (ptr *QFileSystemModel) DisconnectRootPathChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::rootPathChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DisconnectRootPathChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rootPathChanged")
	}
}

//export callbackQFileSystemModelRootPathChanged
func callbackQFileSystemModelRootPathChanged(ptrName *C.char, newPath *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::rootPathChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "rootPathChanged").(func(string))(C.GoString(newPath))
}

func (ptr *QFileSystemModel) RowCount(parent core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QFileSystemModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QFileSystemModel) SetData(idx core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::setData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileSystemModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(idx), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QFileSystemModel) SetFilter(filters core.QDir__Filter) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::setFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetFilter(ptr.Pointer(), C.int(filters))
	}
}

func (ptr *QFileSystemModel) SetIconProvider(provider QFileIconProvider_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::setIconProvider")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetIconProvider(ptr.Pointer(), PointerFromQFileIconProvider(provider))
	}
}

func (ptr *QFileSystemModel) SetNameFilters(filters []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::setNameFilters")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_SetNameFilters(ptr.Pointer(), C.CString(strings.Join(filters, ",,,")))
	}
}

func (ptr *QFileSystemModel) SetRootPath(newPath string) *core.QModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::setRootPath")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QFileSystemModel_SetRootPath(ptr.Pointer(), C.CString(newPath)))
	}
	return nil
}

func (ptr *QFileSystemModel) Sort(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::sort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QFileSystemModel) SupportedDropActions() core.Qt__DropAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::supportedDropActions")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QFileSystemModel_SupportedDropActions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileSystemModel) Type(index core.QModelIndex_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::type")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileSystemModel_Type(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return ""
}

func (ptr *QFileSystemModel) DestroyQFileSystemModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileSystemModel::~QFileSystemModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileSystemModel_DestroyQFileSystemModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
