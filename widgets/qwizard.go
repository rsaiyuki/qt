package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QWizard struct {
	QDialog
}

type QWizard_ITF interface {
	QDialog_ITF
	QWizard_PTR() *QWizard
}

func PointerFromQWizard(ptr QWizard_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWizard_PTR().Pointer()
	}
	return nil
}

func NewQWizardFromPointer(ptr unsafe.Pointer) *QWizard {
	var n = new(QWizard)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWizard_") {
		n.SetObjectName("QWizard_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWizard) QWizard_PTR() *QWizard {
	return ptr
}

//QWizard::WizardButton
type QWizard__WizardButton int64

const (
	QWizard__BackButton       = QWizard__WizardButton(0)
	QWizard__NextButton       = QWizard__WizardButton(1)
	QWizard__CommitButton     = QWizard__WizardButton(2)
	QWizard__FinishButton     = QWizard__WizardButton(3)
	QWizard__CancelButton     = QWizard__WizardButton(4)
	QWizard__HelpButton       = QWizard__WizardButton(5)
	QWizard__CustomButton1    = QWizard__WizardButton(6)
	QWizard__CustomButton2    = QWizard__WizardButton(7)
	QWizard__CustomButton3    = QWizard__WizardButton(8)
	QWizard__Stretch          = QWizard__WizardButton(9)
	QWizard__NoButton         = QWizard__WizardButton(-1)
	QWizard__NStandardButtons = QWizard__WizardButton(6)
	QWizard__NButtons         = QWizard__WizardButton(9)
)

//QWizard::WizardOption
type QWizard__WizardOption int64

const (
	QWizard__IndependentPages             = QWizard__WizardOption(0x00000001)
	QWizard__IgnoreSubTitles              = QWizard__WizardOption(0x00000002)
	QWizard__ExtendedWatermarkPixmap      = QWizard__WizardOption(0x00000004)
	QWizard__NoDefaultButton              = QWizard__WizardOption(0x00000008)
	QWizard__NoBackButtonOnStartPage      = QWizard__WizardOption(0x00000010)
	QWizard__NoBackButtonOnLastPage       = QWizard__WizardOption(0x00000020)
	QWizard__DisabledBackButtonOnLastPage = QWizard__WizardOption(0x00000040)
	QWizard__HaveNextButtonOnLastPage     = QWizard__WizardOption(0x00000080)
	QWizard__HaveFinishButtonOnEarlyPages = QWizard__WizardOption(0x00000100)
	QWizard__NoCancelButton               = QWizard__WizardOption(0x00000200)
	QWizard__CancelButtonOnLeft           = QWizard__WizardOption(0x00000400)
	QWizard__HaveHelpButton               = QWizard__WizardOption(0x00000800)
	QWizard__HelpButtonOnRight            = QWizard__WizardOption(0x00001000)
	QWizard__HaveCustomButton1            = QWizard__WizardOption(0x00002000)
	QWizard__HaveCustomButton2            = QWizard__WizardOption(0x00004000)
	QWizard__HaveCustomButton3            = QWizard__WizardOption(0x00008000)
	QWizard__NoCancelButtonOnLastPage     = QWizard__WizardOption(0x00010000)
)

//QWizard::WizardPixmap
type QWizard__WizardPixmap int64

const (
	QWizard__WatermarkPixmap  = QWizard__WizardPixmap(0)
	QWizard__LogoPixmap       = QWizard__WizardPixmap(1)
	QWizard__BannerPixmap     = QWizard__WizardPixmap(2)
	QWizard__BackgroundPixmap = QWizard__WizardPixmap(3)
	QWizard__NPixmaps         = QWizard__WizardPixmap(4)
)

//QWizard::WizardStyle
type QWizard__WizardStyle int64

var (
	QWizard__ClassicStyle = QWizard__WizardStyle(0)
	QWizard__ModernStyle  = QWizard__WizardStyle(1)
	QWizard__MacStyle     = QWizard__WizardStyle(2)
	QWizard__AeroStyle    = QWizard__WizardStyle(3)
	QWizard__NStyles      = QWizard__WizardStyle(4)
)

func (ptr *QWizard) CurrentId() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::currentId")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QWizard_CurrentId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) HasVisitedPage(id int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::hasVisitedPage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWizard_HasVisitedPage(ptr.Pointer(), C.int(id)) != 0
	}
	return false
}

func (ptr *QWizard) Options() QWizard__WizardOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::options")
		}
	}()

	if ptr.Pointer() != nil {
		return QWizard__WizardOption(C.QWizard_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) Page(id int) *QWizardPage {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::page")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWizardPageFromPointer(C.QWizard_Page(ptr.Pointer(), C.int(id)))
	}
	return nil
}

func (ptr *QWizard) SetOptions(options QWizard__WizardOption) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setOptions")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QWizard) SetPage(id int, page QWizardPage_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setPage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetPage(ptr.Pointer(), C.int(id), PointerFromQWizardPage(page))
	}
}

func (ptr *QWizard) SetStartId(id int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setStartId")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetStartId(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) SetSubTitleFormat(format core.Qt__TextFormat) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setSubTitleFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetSubTitleFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QWizard) SetTitleFormat(format core.Qt__TextFormat) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setTitleFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetTitleFormat(ptr.Pointer(), C.int(format))
	}
}

func (ptr *QWizard) SetWizardStyle(style QWizard__WizardStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setWizardStyle")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetWizardStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QWizard) StartId() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::startId")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QWizard_StartId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) SubTitleFormat() core.Qt__TextFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::subTitleFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QWizard_SubTitleFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) TitleFormat() core.Qt__TextFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::titleFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QWizard_TitleFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) WizardStyle() QWizard__WizardStyle {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::wizardStyle")
		}
	}()

	if ptr.Pointer() != nil {
		return QWizard__WizardStyle(C.QWizard_WizardStyle(ptr.Pointer()))
	}
	return 0
}

func NewQWizard(parent QWidget_ITF, flags core.Qt__WindowType) *QWizard {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::QWizard")
		}
	}()

	return NewQWizardFromPointer(C.QWizard_NewQWizard(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QWizard) AddPage(page QWizardPage_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::addPage")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QWizard_AddPage(ptr.Pointer(), PointerFromQWizardPage(page)))
	}
	return 0
}

func (ptr *QWizard) Back() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::back")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_Back(ptr.Pointer())
	}
}

func (ptr *QWizard) Button(which QWizard__WizardButton) *QAbstractButton {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::button")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractButtonFromPointer(C.QWizard_Button(ptr.Pointer(), C.int(which)))
	}
	return nil
}

func (ptr *QWizard) ButtonText(which QWizard__WizardButton) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::buttonText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWizard_ButtonText(ptr.Pointer(), C.int(which)))
	}
	return ""
}

func (ptr *QWizard) ConnectCurrentIdChanged(f func(id int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::currentIdChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_ConnectCurrentIdChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentIdChanged", f)
	}
}

func (ptr *QWizard) DisconnectCurrentIdChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::currentIdChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectCurrentIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentIdChanged")
	}
}

//export callbackQWizardCurrentIdChanged
func callbackQWizardCurrentIdChanged(ptrName *C.char, id C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::currentIdChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentIdChanged").(func(int))(int(id))
}

func (ptr *QWizard) CurrentPage() *QWizardPage {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::currentPage")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWizardPageFromPointer(C.QWizard_CurrentPage(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWizard) ConnectCustomButtonClicked(f func(which int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::customButtonClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_ConnectCustomButtonClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "customButtonClicked", f)
	}
}

func (ptr *QWizard) DisconnectCustomButtonClicked() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::customButtonClicked")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectCustomButtonClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "customButtonClicked")
	}
}

//export callbackQWizardCustomButtonClicked
func callbackQWizardCustomButtonClicked(ptrName *C.char, which C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::customButtonClicked")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "customButtonClicked").(func(int))(int(which))
}

func (ptr *QWizard) Field(name string) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::field")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWizard_Field(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QWizard) ConnectHelpRequested(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::helpRequested")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_ConnectHelpRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "helpRequested", f)
	}
}

func (ptr *QWizard) DisconnectHelpRequested() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::helpRequested")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectHelpRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "helpRequested")
	}
}

//export callbackQWizardHelpRequested
func callbackQWizardHelpRequested(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::helpRequested")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "helpRequested").(func())()
}

func (ptr *QWizard) Next() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::next")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_Next(ptr.Pointer())
	}
}

func (ptr *QWizard) NextId() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::nextId")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QWizard_NextId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWizard) ConnectPageAdded(f func(id int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::pageAdded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_ConnectPageAdded(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pageAdded", f)
	}
}

func (ptr *QWizard) DisconnectPageAdded() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::pageAdded")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectPageAdded(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pageAdded")
	}
}

//export callbackQWizardPageAdded
func callbackQWizardPageAdded(ptrName *C.char, id C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::pageAdded")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "pageAdded").(func(int))(int(id))
}

func (ptr *QWizard) ConnectPageRemoved(f func(id int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::pageRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_ConnectPageRemoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pageRemoved", f)
	}
}

func (ptr *QWizard) DisconnectPageRemoved() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::pageRemoved")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_DisconnectPageRemoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pageRemoved")
	}
}

//export callbackQWizardPageRemoved
func callbackQWizardPageRemoved(ptrName *C.char, id C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::pageRemoved")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "pageRemoved").(func(int))(int(id))
}

func (ptr *QWizard) RemovePage(id int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::removePage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_RemovePage(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWizard) Restart() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::restart")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_Restart(ptr.Pointer())
	}
}

func (ptr *QWizard) SetButton(which QWizard__WizardButton, button QAbstractButton_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setButton")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetButton(ptr.Pointer(), C.int(which), PointerFromQAbstractButton(button))
	}
}

func (ptr *QWizard) SetButtonText(which QWizard__WizardButton, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setButtonText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetButtonText(ptr.Pointer(), C.int(which), C.CString(text))
	}
}

func (ptr *QWizard) SetDefaultProperty(className string, property string, changedSignal string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setDefaultProperty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetDefaultProperty(ptr.Pointer(), C.CString(className), C.CString(property), C.CString(changedSignal))
	}
}

func (ptr *QWizard) SetField(name string, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setField")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetField(ptr.Pointer(), C.CString(name), core.PointerFromQVariant(value))
	}
}

func (ptr *QWizard) SetOption(option QWizard__WizardOption, on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWizard) SetPixmap(which QWizard__WizardPixmap, pixmap gui.QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetPixmap(ptr.Pointer(), C.int(which), gui.PointerFromQPixmap(pixmap))
	}
}

func (ptr *QWizard) SetSideWidget(widget QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setSideWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetSideWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWizard) SetVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::setVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWizard) SideWidget() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::sideWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWizard_SideWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWizard) TestOption(option QWizard__WizardOption) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::testOption")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWizard_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QWizard) ValidateCurrentPage() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::validateCurrentPage")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWizard_ValidateCurrentPage(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWizard) DestroyQWizard() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWizard::~QWizard")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWizard_DestroyQWizard(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
