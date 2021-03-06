package core

//#include "core.h"
import "C"
import (
	"log"
	"strings"
	"unsafe"
)

type QUrl struct {
	ptr unsafe.Pointer
}

type QUrl_ITF interface {
	QUrl_PTR() *QUrl
}

func (p *QUrl) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUrl) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUrl(ptr QUrl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUrl_PTR().Pointer()
	}
	return nil
}

func NewQUrlFromPointer(ptr unsafe.Pointer) *QUrl {
	var n = new(QUrl)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUrl) QUrl_PTR() *QUrl {
	return ptr
}

//QUrl::ComponentFormattingOption
type QUrl__ComponentFormattingOption int64

const (
	QUrl__PrettyDecoded    = QUrl__ComponentFormattingOption(0x000000)
	QUrl__EncodeSpaces     = QUrl__ComponentFormattingOption(0x100000)
	QUrl__EncodeUnicode    = QUrl__ComponentFormattingOption(0x200000)
	QUrl__EncodeDelimiters = QUrl__ComponentFormattingOption(0x400000 | 0x800000)
	QUrl__EncodeReserved   = QUrl__ComponentFormattingOption(0x1000000)
	QUrl__DecodeReserved   = QUrl__ComponentFormattingOption(0x2000000)
	QUrl__FullyEncoded     = QUrl__ComponentFormattingOption(QUrl__EncodeSpaces | QUrl__EncodeUnicode | QUrl__EncodeDelimiters | QUrl__EncodeReserved)
	QUrl__FullyDecoded     = QUrl__ComponentFormattingOption(QUrl__FullyEncoded | QUrl__DecodeReserved | 0x4000000)
)

//QUrl::ParsingMode
type QUrl__ParsingMode int64

const (
	QUrl__TolerantMode = QUrl__ParsingMode(0)
	QUrl__StrictMode   = QUrl__ParsingMode(1)
	QUrl__DecodedMode  = QUrl__ParsingMode(2)
)

//QUrl::UrlFormattingOption
type QUrl__UrlFormattingOption int64

const (
	QUrl__None                  = QUrl__UrlFormattingOption(0x0)
	QUrl__RemoveScheme          = QUrl__UrlFormattingOption(0x1)
	QUrl__RemovePassword        = QUrl__UrlFormattingOption(0x2)
	QUrl__RemoveUserInfo        = QUrl__UrlFormattingOption(QUrl__RemovePassword | 0x4)
	QUrl__RemovePort            = QUrl__UrlFormattingOption(0x8)
	QUrl__RemoveAuthority       = QUrl__UrlFormattingOption(QUrl__RemoveUserInfo | QUrl__RemovePort | 0x10)
	QUrl__RemovePath            = QUrl__UrlFormattingOption(0x20)
	QUrl__RemoveQuery           = QUrl__UrlFormattingOption(0x40)
	QUrl__RemoveFragment        = QUrl__UrlFormattingOption(0x80)
	QUrl__PreferLocalFile       = QUrl__UrlFormattingOption(0x200)
	QUrl__StripTrailingSlash    = QUrl__UrlFormattingOption(0x400)
	QUrl__RemoveFilename        = QUrl__UrlFormattingOption(0x800)
	QUrl__NormalizePathSegments = QUrl__UrlFormattingOption(0x1000)
)

//QUrl::UserInputResolutionOption
type QUrl__UserInputResolutionOption int64

const (
	QUrl__DefaultResolution = QUrl__UserInputResolutionOption(0)
	QUrl__AssumeLocalFile   = QUrl__UserInputResolutionOption(1)
)

func NewQUrl() *QUrl {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::QUrl")
		}
	}()

	return NewQUrlFromPointer(C.QUrl_NewQUrl())
}

func NewQUrl4(other QUrl_ITF) *QUrl {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::QUrl")
		}
	}()

	return NewQUrlFromPointer(C.QUrl_NewQUrl4(PointerFromQUrl(other)))
}

func NewQUrl3(url string, parsingMode QUrl__ParsingMode) *QUrl {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::QUrl")
		}
	}()

	return NewQUrlFromPointer(C.QUrl_NewQUrl3(C.CString(url), C.int(parsingMode)))
}

func NewQUrl2(other QUrl_ITF) *QUrl {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::QUrl")
		}
	}()

	return NewQUrlFromPointer(C.QUrl_NewQUrl2(PointerFromQUrl(other)))
}

func (ptr *QUrl) Authority(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::authority")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_Authority(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_Clear(ptr.Pointer())
	}
}

func (ptr *QUrl) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUrl) FileName(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::fileName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_FileName(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) Fragment(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::fragment")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_Fragment(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func QUrl_FromAce(domain QByteArray_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::fromAce")
		}
	}()

	return C.GoString(C.QUrl_QUrl_FromAce(PointerFromQByteArray(domain)))
}

func QUrl_FromPercentEncoding(input QByteArray_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::fromPercentEncoding")
		}
	}()

	return C.GoString(C.QUrl_QUrl_FromPercentEncoding(PointerFromQByteArray(input)))
}

func (ptr *QUrl) HasFragment() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::hasFragment")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUrl_HasFragment(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUrl) HasQuery() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::hasQuery")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUrl_HasQuery(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUrl) Host(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::host")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_Host(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func QUrl_IdnWhitelist() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::idnWhitelist")
		}
	}()

	return strings.Split(C.GoString(C.QUrl_QUrl_IdnWhitelist()), ",,,")
}

func (ptr *QUrl) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUrl_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUrl) IsLocalFile() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::isLocalFile")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUrl_IsLocalFile(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUrl) IsParentOf(childUrl QUrl_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::isParentOf")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUrl_IsParentOf(ptr.Pointer(), PointerFromQUrl(childUrl)) != 0
	}
	return false
}

func (ptr *QUrl) IsRelative() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::isRelative")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUrl_IsRelative(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUrl) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUrl_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUrl) Matches(url QUrl_ITF, options QUrl__UrlFormattingOption) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::matches")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUrl_Matches(ptr.Pointer(), PointerFromQUrl(url), C.int(options)) != 0
	}
	return false
}

func (ptr *QUrl) Password(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::password")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_Password(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) Path(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::path")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_Path(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) Port(defaultPort int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::port")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QUrl_Port(ptr.Pointer(), C.int(defaultPort)))
	}
	return 0
}

func (ptr *QUrl) Query(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::query")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_Query(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) Scheme() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::scheme")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_Scheme(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUrl) SetAuthority(authority string, mode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setAuthority")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetAuthority(ptr.Pointer(), C.CString(authority), C.int(mode))
	}
}

func (ptr *QUrl) SetFragment(fragment string, mode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setFragment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetFragment(ptr.Pointer(), C.CString(fragment), C.int(mode))
	}
}

func (ptr *QUrl) SetHost(host string, mode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setHost")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetHost(ptr.Pointer(), C.CString(host), C.int(mode))
	}
}

func QUrl_SetIdnWhitelist(list []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setIdnWhitelist")
		}
	}()

	C.QUrl_QUrl_SetIdnWhitelist(C.CString(strings.Join(list, ",,,")))
}

func (ptr *QUrl) SetPassword(password string, mode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setPassword")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetPassword(ptr.Pointer(), C.CString(password), C.int(mode))
	}
}

func (ptr *QUrl) SetPath(path string, mode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetPath(ptr.Pointer(), C.CString(path), C.int(mode))
	}
}

func (ptr *QUrl) SetPort(port int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setPort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetPort(ptr.Pointer(), C.int(port))
	}
}

func (ptr *QUrl) SetQuery(query string, mode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setQuery")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetQuery(ptr.Pointer(), C.CString(query), C.int(mode))
	}
}

func (ptr *QUrl) SetQuery2(query QUrlQuery_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setQuery")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetQuery2(ptr.Pointer(), PointerFromQUrlQuery(query))
	}
}

func (ptr *QUrl) SetScheme(scheme string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setScheme")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetScheme(ptr.Pointer(), C.CString(scheme))
	}
}

func (ptr *QUrl) SetUrl(url string, parsingMode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setUrl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetUrl(ptr.Pointer(), C.CString(url), C.int(parsingMode))
	}
}

func (ptr *QUrl) SetUserInfo(userInfo string, mode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setUserInfo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetUserInfo(ptr.Pointer(), C.CString(userInfo), C.int(mode))
	}
}

func (ptr *QUrl) SetUserName(userName string, mode QUrl__ParsingMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::setUserName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_SetUserName(ptr.Pointer(), C.CString(userName), C.int(mode))
	}
}

func (ptr *QUrl) Swap(other QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_Swap(ptr.Pointer(), PointerFromQUrl(other))
	}
}

func QUrl_ToAce(domain string) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::toAce")
		}
	}()

	return NewQByteArrayFromPointer(C.QUrl_QUrl_ToAce(C.CString(domain)))
}

func (ptr *QUrl) ToDisplayString(options QUrl__UrlFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::toDisplayString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_ToDisplayString(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) ToEncoded(options QUrl__UrlFormattingOption) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::toEncoded")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QUrl_ToEncoded(ptr.Pointer(), C.int(options)))
	}
	return nil
}

func (ptr *QUrl) ToLocalFile() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::toLocalFile")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_ToLocalFile(ptr.Pointer()))
	}
	return ""
}

func QUrl_ToPercentEncoding(input string, exclude QByteArray_ITF, include QByteArray_ITF) *QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::toPercentEncoding")
		}
	}()

	return NewQByteArrayFromPointer(C.QUrl_QUrl_ToPercentEncoding(C.CString(input), PointerFromQByteArray(exclude), PointerFromQByteArray(include)))
}

func (ptr *QUrl) ToString(options QUrl__UrlFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_ToString(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) TopLevelDomain(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::topLevelDomain")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_TopLevelDomain(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) Url(options QUrl__UrlFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::url")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_Url(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) UserInfo(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::userInfo")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_UserInfo(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) UserName(options QUrl__ComponentFormattingOption) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::userName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUrl_UserName(ptr.Pointer(), C.int(options)))
	}
	return ""
}

func (ptr *QUrl) DestroyQUrl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUrl::~QUrl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUrl_DestroyQUrl(ptr.Pointer())
	}
}
