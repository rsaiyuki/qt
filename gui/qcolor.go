package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QColor struct {
	ptr unsafe.Pointer
}

type QColor_ITF interface {
	QColor_PTR() *QColor
}

func (p *QColor) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QColor) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQColor(ptr QColor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColor_PTR().Pointer()
	}
	return nil
}

func NewQColorFromPointer(ptr unsafe.Pointer) *QColor {
	var n = new(QColor)
	n.SetPointer(ptr)
	return n
}

func (ptr *QColor) QColor_PTR() *QColor {
	return ptr
}

//QColor::NameFormat
type QColor__NameFormat int64

const (
	QColor__HexRgb  = QColor__NameFormat(0)
	QColor__HexArgb = QColor__NameFormat(1)
)

//QColor::Spec
type QColor__Spec int64

const (
	QColor__Invalid = QColor__Spec(0)
	QColor__Rgb     = QColor__Spec(1)
	QColor__Hsv     = QColor__Spec(2)
	QColor__Cmyk    = QColor__Spec(3)
	QColor__Hsl     = QColor__Spec(4)
)

func (ptr *QColor) ConvertTo(colorSpec QColor__Spec) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::convertTo")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QColor_ConvertTo(ptr.Pointer(), C.int(colorSpec)))
	}
	return nil
}

func (ptr *QColor) SetRgbF(r float64, g float64, b float64, a float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setRgbF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetRgbF(ptr.Pointer(), C.double(r), C.double(g), C.double(b), C.double(a))
	}
}

func NewQColor() *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::QColor")
		}
	}()

	return NewQColorFromPointer(C.QColor_NewQColor())
}

func NewQColor8(color core.Qt__GlobalColor) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::QColor")
		}
	}()

	return NewQColorFromPointer(C.QColor_NewQColor8(C.int(color)))
}

func NewQColor6(color QColor_ITF) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::QColor")
		}
	}()

	return NewQColorFromPointer(C.QColor_NewQColor6(PointerFromQColor(color)))
}

func NewQColor4(name string) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::QColor")
		}
	}()

	return NewQColorFromPointer(C.QColor_NewQColor4(C.CString(name)))
}

func NewQColor5(name string) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::QColor")
		}
	}()

	return NewQColorFromPointer(C.QColor_NewQColor5(C.CString(name)))
}

func NewQColor2(r int, g int, b int, a int) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::QColor")
		}
	}()

	return NewQColorFromPointer(C.QColor_NewQColor2(C.int(r), C.int(g), C.int(b), C.int(a)))
}

func (ptr *QColor) Alpha() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::alpha")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Alpha(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) AlphaF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::alphaF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_AlphaF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) Black() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::black")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Black(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) BlackF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::blackF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_BlackF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) Blue() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::blue")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Blue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) BlueF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::blueF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_BlueF(ptr.Pointer()))
	}
	return 0
}

func QColor_ColorNames() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::colorNames")
		}
	}()

	return strings.Split(C.GoString(C.QColor_QColor_ColorNames()), ",,,")
}

func (ptr *QColor) Cyan() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::cyan")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Cyan(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) CyanF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::cyanF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_CyanF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) Darker(factor int) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::darker")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QColor_Darker(ptr.Pointer(), C.int(factor)))
	}
	return nil
}

func QColor_FromCmyk(c int, m int, y int, k int, a int) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::fromCmyk")
		}
	}()

	return NewQColorFromPointer(C.QColor_QColor_FromCmyk(C.int(c), C.int(m), C.int(y), C.int(k), C.int(a)))
}

func QColor_FromCmykF(c float64, m float64, y float64, k float64, a float64) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::fromCmykF")
		}
	}()

	return NewQColorFromPointer(C.QColor_QColor_FromCmykF(C.double(c), C.double(m), C.double(y), C.double(k), C.double(a)))
}

func QColor_FromHsl(h int, s int, l int, a int) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::fromHsl")
		}
	}()

	return NewQColorFromPointer(C.QColor_QColor_FromHsl(C.int(h), C.int(s), C.int(l), C.int(a)))
}

func QColor_FromHslF(h float64, s float64, l float64, a float64) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::fromHslF")
		}
	}()

	return NewQColorFromPointer(C.QColor_QColor_FromHslF(C.double(h), C.double(s), C.double(l), C.double(a)))
}

func QColor_FromHsv(h int, s int, v int, a int) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::fromHsv")
		}
	}()

	return NewQColorFromPointer(C.QColor_QColor_FromHsv(C.int(h), C.int(s), C.int(v), C.int(a)))
}

func QColor_FromHsvF(h float64, s float64, v float64, a float64) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::fromHsvF")
		}
	}()

	return NewQColorFromPointer(C.QColor_QColor_FromHsvF(C.double(h), C.double(s), C.double(v), C.double(a)))
}

func QColor_FromRgb2(r int, g int, b int, a int) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::fromRgb")
		}
	}()

	return NewQColorFromPointer(C.QColor_QColor_FromRgb2(C.int(r), C.int(g), C.int(b), C.int(a)))
}

func QColor_FromRgbF(r float64, g float64, b float64, a float64) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::fromRgbF")
		}
	}()

	return NewQColorFromPointer(C.QColor_QColor_FromRgbF(C.double(r), C.double(g), C.double(b), C.double(a)))
}

func (ptr *QColor) GetCmyk(c int, m int, y int, k int, a int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::getCmyk")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_GetCmyk(ptr.Pointer(), C.int(c), C.int(m), C.int(y), C.int(k), C.int(a))
	}
}

func (ptr *QColor) GetHsl(h int, s int, l int, a int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::getHsl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_GetHsl(ptr.Pointer(), C.int(h), C.int(s), C.int(l), C.int(a))
	}
}

func (ptr *QColor) GetHsv(h int, s int, v int, a int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::getHsv")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_GetHsv(ptr.Pointer(), C.int(h), C.int(s), C.int(v), C.int(a))
	}
}

func (ptr *QColor) GetRgb(r int, g int, b int, a int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::getRgb")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_GetRgb(ptr.Pointer(), C.int(r), C.int(g), C.int(b), C.int(a))
	}
}

func (ptr *QColor) Green() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::green")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Green(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) GreenF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::greenF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_GreenF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HslHue() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hslHue")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_HslHue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HslHueF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hslHueF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_HslHueF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HslSaturation() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hslSaturation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_HslSaturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HslSaturationF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hslSaturationF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_HslSaturationF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HsvHue() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hsvHue")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_HsvHue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HsvHueF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hsvHueF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_HsvHueF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HsvSaturation() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hsvSaturation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_HsvSaturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HsvSaturationF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hsvSaturationF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_HsvSaturationF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) Hue() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hue")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) HueF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::hueF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_HueF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QColor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func QColor_IsValidColor(name string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::isValidColor")
		}
	}()

	return C.QColor_QColor_IsValidColor(C.CString(name)) != 0
}

func (ptr *QColor) Lighter(factor int) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::lighter")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QColor_Lighter(ptr.Pointer(), C.int(factor)))
	}
	return nil
}

func (ptr *QColor) Lightness() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::lightness")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Lightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) LightnessF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::lightnessF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_LightnessF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) Magenta() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::magenta")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Magenta(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) MagentaF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::magentaF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_MagentaF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QColor_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QColor) Name2(format QColor__NameFormat) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QColor_Name2(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QColor) Red() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::red")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Red(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) RedF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::redF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_RedF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) Saturation() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::saturation")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) SaturationF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::saturationF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_SaturationF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) SetAlpha(alpha int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setAlpha")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetAlpha(ptr.Pointer(), C.int(alpha))
	}
}

func (ptr *QColor) SetAlphaF(alpha float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setAlphaF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetAlphaF(ptr.Pointer(), C.double(alpha))
	}
}

func (ptr *QColor) SetBlue(blue int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setBlue")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetBlue(ptr.Pointer(), C.int(blue))
	}
}

func (ptr *QColor) SetBlueF(blue float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setBlueF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetBlueF(ptr.Pointer(), C.double(blue))
	}
}

func (ptr *QColor) SetCmyk(c int, m int, y int, k int, a int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setCmyk")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetCmyk(ptr.Pointer(), C.int(c), C.int(m), C.int(y), C.int(k), C.int(a))
	}
}

func (ptr *QColor) SetCmykF(c float64, m float64, y float64, k float64, a float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setCmykF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetCmykF(ptr.Pointer(), C.double(c), C.double(m), C.double(y), C.double(k), C.double(a))
	}
}

func (ptr *QColor) SetGreen(green int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setGreen")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetGreen(ptr.Pointer(), C.int(green))
	}
}

func (ptr *QColor) SetGreenF(green float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setGreenF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetGreenF(ptr.Pointer(), C.double(green))
	}
}

func (ptr *QColor) SetHsl(h int, s int, l int, a int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setHsl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetHsl(ptr.Pointer(), C.int(h), C.int(s), C.int(l), C.int(a))
	}
}

func (ptr *QColor) SetHslF(h float64, s float64, l float64, a float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setHslF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetHslF(ptr.Pointer(), C.double(h), C.double(s), C.double(l), C.double(a))
	}
}

func (ptr *QColor) SetHsv(h int, s int, v int, a int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setHsv")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetHsv(ptr.Pointer(), C.int(h), C.int(s), C.int(v), C.int(a))
	}
}

func (ptr *QColor) SetHsvF(h float64, s float64, v float64, a float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setHsvF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetHsvF(ptr.Pointer(), C.double(h), C.double(s), C.double(v), C.double(a))
	}
}

func (ptr *QColor) SetNamedColor(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setNamedColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetNamedColor(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QColor) SetRed(red int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setRed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetRed(ptr.Pointer(), C.int(red))
	}
}

func (ptr *QColor) SetRedF(red float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setRedF")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetRedF(ptr.Pointer(), C.double(red))
	}
}

func (ptr *QColor) SetRgb(r int, g int, b int, a int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::setRgb")
		}
	}()

	if ptr.Pointer() != nil {
		C.QColor_SetRgb(ptr.Pointer(), C.int(r), C.int(g), C.int(b), C.int(a))
	}
}

func (ptr *QColor) Spec() QColor__Spec {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::spec")
		}
	}()

	if ptr.Pointer() != nil {
		return QColor__Spec(C.QColor_Spec(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) ToCmyk() *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::toCmyk")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QColor_ToCmyk(ptr.Pointer()))
	}
	return nil
}

func (ptr *QColor) ToHsl() *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::toHsl")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QColor_ToHsl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QColor) ToHsv() *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::toHsv")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QColor_ToHsv(ptr.Pointer()))
	}
	return nil
}

func (ptr *QColor) ToRgb() *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::toRgb")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QColor_ToRgb(ptr.Pointer()))
	}
	return nil
}

func (ptr *QColor) Value() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::value")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Value(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) ValueF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::valueF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_ValueF(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) Yellow() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::yellow")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QColor_Yellow(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColor) YellowF() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QColor::yellowF")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QColor_YellowF(ptr.Pointer()))
	}
	return 0
}
