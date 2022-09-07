package lang

// First check
//go:generate $GOPATH/bin/gdev-i18n -type ErrorCode -tomlpath i18n -check

// Second generation
//go:generate $GOPATH/bin/gdev-i18n -type ErrorCode -ctxkey i18n -tomlpath i18n -defaultlocale zh_cn -output i118n_stringer.go

// or write like this, attention COMMENT prefix #
//#go:generate $GOPATH/bin/gdev-i18n -type=ErrorCode -ctxkey=i18n -tomlpath=i18n -defaultlocale=zh_cn -output=i118n_stringer.go

// ErrorCode Define a custom shaping type that supports i18n as an error code
type ErrorCode uint

// I18nErrorCodeDetail Define a custom shaping type that supports i18n as a struct error code
type I18nErrorCodeDetail struct {
	ErrString string // wrap another error
	ErrCode   uint   // custom shaping type Val
	Locale    string // i18n locale set
}

// Specify the error code of a part of the range as the replacement value content of the formatted output
// NOTE: The range 10000 to 20000 is used as a formatted output replacement component
const (
	ComUserName ErrorCode = 10000 + iota
	ComUserPwd
)

// Assign error code value range
const (
	Ok ErrorCode = 20000 + iota
	UserLoginInvalid
	MerchantLoginInvalid
	SubDirectoryTest
)
