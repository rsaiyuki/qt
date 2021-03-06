package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QDateTimeEdit struct {
	QAbstractSpinBox
}

type QDateTimeEdit_ITF interface {
	QAbstractSpinBox_ITF
	QDateTimeEdit_PTR() *QDateTimeEdit
}

func PointerFromQDateTimeEdit(ptr QDateTimeEdit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDateTimeEdit_PTR().Pointer()
	}
	return nil
}

func NewQDateTimeEditFromPointer(ptr unsafe.Pointer) *QDateTimeEdit {
	var n = new(QDateTimeEdit)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDateTimeEdit_") {
		n.SetObjectName("QDateTimeEdit_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDateTimeEdit) QDateTimeEdit_PTR() *QDateTimeEdit {
	return ptr
}

//QDateTimeEdit::Section
type QDateTimeEdit__Section int64

const (
	QDateTimeEdit__NoSection         = QDateTimeEdit__Section(0x0000)
	QDateTimeEdit__AmPmSection       = QDateTimeEdit__Section(0x0001)
	QDateTimeEdit__MSecSection       = QDateTimeEdit__Section(0x0002)
	QDateTimeEdit__SecondSection     = QDateTimeEdit__Section(0x0004)
	QDateTimeEdit__MinuteSection     = QDateTimeEdit__Section(0x0008)
	QDateTimeEdit__HourSection       = QDateTimeEdit__Section(0x0010)
	QDateTimeEdit__DaySection        = QDateTimeEdit__Section(0x0100)
	QDateTimeEdit__MonthSection      = QDateTimeEdit__Section(0x0200)
	QDateTimeEdit__YearSection       = QDateTimeEdit__Section(0x0400)
	QDateTimeEdit__TimeSections_Mask = QDateTimeEdit__Section(QDateTimeEdit__AmPmSection | QDateTimeEdit__MSecSection | QDateTimeEdit__SecondSection | QDateTimeEdit__MinuteSection | QDateTimeEdit__HourSection)
	QDateTimeEdit__DateSections_Mask = QDateTimeEdit__Section(QDateTimeEdit__DaySection | QDateTimeEdit__MonthSection | QDateTimeEdit__YearSection)
)

func NewQDateTimeEdit3(date core.QDate_ITF, parent QWidget_ITF) *QDateTimeEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::QDateTimeEdit")
		}
	}()

	return NewQDateTimeEditFromPointer(C.QDateTimeEdit_NewQDateTimeEdit3(core.PointerFromQDate(date), PointerFromQWidget(parent)))
}

func NewQDateTimeEdit4(time core.QTime_ITF, parent QWidget_ITF) *QDateTimeEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::QDateTimeEdit")
		}
	}()

	return NewQDateTimeEditFromPointer(C.QDateTimeEdit_NewQDateTimeEdit4(core.PointerFromQTime(time), PointerFromQWidget(parent)))
}

func (ptr *QDateTimeEdit) CalendarPopup() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::calendarPopup")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDateTimeEdit_CalendarPopup(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDateTimeEdit) ClearMaximumDate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::clearMaximumDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumDate(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMaximumDateTime() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::clearMaximumDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumDateTime(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMaximumTime() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::clearMaximumTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMaximumTime(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMinimumDate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::clearMinimumDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumDate(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMinimumDateTime() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::clearMinimumDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumDateTime(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ClearMinimumTime() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::clearMinimumTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ClearMinimumTime(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) CurrentSection() QDateTimeEdit__Section {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::currentSection")
		}
	}()

	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_CurrentSection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeEdit) CurrentSectionIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::currentSectionIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDateTimeEdit_CurrentSectionIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeEdit) DateTime() *core.QDateTime {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::dateTime")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QDateTimeEdit_DateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) DisplayFormat() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::displayFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTimeEdit_DisplayFormat(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDateTimeEdit) DisplayedSections() QDateTimeEdit__Section {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::displayedSections")
		}
	}()

	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_DisplayedSections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeEdit) MaximumDateTime() *core.QDateTime {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::maximumDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QDateTimeEdit_MaximumDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) MinimumDateTime() *core.QDateTime {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::minimumDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QDateTimeEdit_MinimumDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) SectionCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::sectionCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QDateTimeEdit_SectionCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDateTimeEdit) SectionText(section QDateTimeEdit__Section) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::sectionText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QDateTimeEdit_SectionText(ptr.Pointer(), C.int(section)))
	}
	return ""
}

func (ptr *QDateTimeEdit) SetCalendarPopup(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setCalendarPopup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCalendarPopup(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QDateTimeEdit) SetCurrentSection(section QDateTimeEdit__Section) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setCurrentSection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCurrentSection(ptr.Pointer(), C.int(section))
	}
}

func (ptr *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setCurrentSectionIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCurrentSectionIndex(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QDateTimeEdit) SetDate(date core.QDate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDate(ptr.Pointer(), core.PointerFromQDate(date))
	}
}

func (ptr *QDateTimeEdit) SetDateTime(dateTime core.QDateTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateTime(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QDateTimeEdit) SetDisplayFormat(format string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setDisplayFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDisplayFormat(ptr.Pointer(), C.CString(format))
	}
}

func (ptr *QDateTimeEdit) SetMaximumDate(max core.QDate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setMaximumDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumDate(ptr.Pointer(), core.PointerFromQDate(max))
	}
}

func (ptr *QDateTimeEdit) SetMaximumDateTime(dt core.QDateTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setMaximumDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumDateTime(ptr.Pointer(), core.PointerFromQDateTime(dt))
	}
}

func (ptr *QDateTimeEdit) SetMaximumTime(max core.QTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setMaximumTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMaximumTime(ptr.Pointer(), core.PointerFromQTime(max))
	}
}

func (ptr *QDateTimeEdit) SetMinimumDate(min core.QDate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setMinimumDate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumDate(ptr.Pointer(), core.PointerFromQDate(min))
	}
}

func (ptr *QDateTimeEdit) SetMinimumDateTime(dt core.QDateTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setMinimumDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumDateTime(ptr.Pointer(), core.PointerFromQDateTime(dt))
	}
}

func (ptr *QDateTimeEdit) SetMinimumTime(min core.QTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setMinimumTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetMinimumTime(ptr.Pointer(), core.PointerFromQTime(min))
	}
}

func (ptr *QDateTimeEdit) SetTime(time core.QTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setTime")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTime(ptr.Pointer(), core.PointerFromQTime(time))
	}
}

func (ptr *QDateTimeEdit) SetTimeSpec(spec core.Qt__TimeSpec) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setTimeSpec")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTimeSpec(ptr.Pointer(), C.int(spec))
	}
}

func (ptr *QDateTimeEdit) TimeSpec() core.Qt__TimeSpec {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::timeSpec")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__TimeSpec(C.QDateTimeEdit_TimeSpec(ptr.Pointer()))
	}
	return 0
}

func NewQDateTimeEdit(parent QWidget_ITF) *QDateTimeEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::QDateTimeEdit")
		}
	}()

	return NewQDateTimeEditFromPointer(C.QDateTimeEdit_NewQDateTimeEdit(PointerFromQWidget(parent)))
}

func NewQDateTimeEdit2(datetime core.QDateTime_ITF, parent QWidget_ITF) *QDateTimeEdit {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::QDateTimeEdit")
		}
	}()

	return NewQDateTimeEditFromPointer(C.QDateTimeEdit_NewQDateTimeEdit2(core.PointerFromQDateTime(datetime), PointerFromQWidget(parent)))
}

func (ptr *QDateTimeEdit) CalendarWidget() *QCalendarWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::calendarWidget")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQCalendarWidgetFromPointer(C.QDateTimeEdit_CalendarWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDateTimeEdit) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_Clear(ptr.Pointer())
	}
}

func (ptr *QDateTimeEdit) ConnectDateTimeChanged(f func(datetime *core.QDateTime)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::dateTimeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_ConnectDateTimeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dateTimeChanged", f)
	}
}

func (ptr *QDateTimeEdit) DisconnectDateTimeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::dateTimeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_DisconnectDateTimeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dateTimeChanged")
	}
}

//export callbackQDateTimeEditDateTimeChanged
func callbackQDateTimeEditDateTimeChanged(ptrName *C.char, datetime unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::dateTimeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "dateTimeChanged").(func(*core.QDateTime))(core.NewQDateTimeFromPointer(datetime))
}

func (ptr *QDateTimeEdit) Event(event core.QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::event")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QDateTimeEdit_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QDateTimeEdit) SectionAt(index int) QDateTimeEdit__Section {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::sectionAt")
		}
	}()

	if ptr.Pointer() != nil {
		return QDateTimeEdit__Section(C.QDateTimeEdit_SectionAt(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QDateTimeEdit) SetCalendarWidget(calendarWidget QCalendarWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setCalendarWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetCalendarWidget(ptr.Pointer(), PointerFromQCalendarWidget(calendarWidget))
	}
}

func (ptr *QDateTimeEdit) SetDateRange(min core.QDate_ITF, max core.QDate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setDateRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateRange(ptr.Pointer(), core.PointerFromQDate(min), core.PointerFromQDate(max))
	}
}

func (ptr *QDateTimeEdit) SetDateTimeRange(min core.QDateTime_ITF, max core.QDateTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setDateTimeRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetDateTimeRange(ptr.Pointer(), core.PointerFromQDateTime(min), core.PointerFromQDateTime(max))
	}
}

func (ptr *QDateTimeEdit) SetSelectedSection(section QDateTimeEdit__Section) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setSelectedSection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetSelectedSection(ptr.Pointer(), C.int(section))
	}
}

func (ptr *QDateTimeEdit) SetTimeRange(min core.QTime_ITF, max core.QTime_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::setTimeRange")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_SetTimeRange(ptr.Pointer(), core.PointerFromQTime(min), core.PointerFromQTime(max))
	}
}

func (ptr *QDateTimeEdit) StepBy(steps int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::stepBy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_StepBy(ptr.Pointer(), C.int(steps))
	}
}

func (ptr *QDateTimeEdit) DestroyQDateTimeEdit() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDateTimeEdit::~QDateTimeEdit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDateTimeEdit_DestroyQDateTimeEdit(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
