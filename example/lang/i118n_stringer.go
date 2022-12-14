// Code generated by "gdev-i18n -type ErrorCode -ctxkey i18n -tomlpath i18n -defaultlocale zh_cn -output i118n_stringer.go"; DO NOT EDIT.

package lang

import (
	"context"
	"fmt"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the gdev-i18n command to generate them again.
	var x [1]struct{}
	_ = x[ComUserName-10000]
	_ = x[ComUserPwd-10001]
	_ = x[Ok-20000]
	_ = x[UserLoginInvalid-20001]
	_ = x[MerchantLoginInvalid-20002]
	_ = x[SubDirectoryTest-20003]
}

const (
	_ErrorCode_En_name_0   = "user nameuser password"
	_ErrorCode_ZhCn_name_0 = "用户名登录密码"
	_ErrorCode_En_name_1   = "okYour %s format is wrong, please try again laterYour merchant's %s format is incorrectsub directory test"
	_ErrorCode_ZhCn_name_1 = "操作成功您的%s格式错误，请稍后再试商户端%s格式有误子目录测试"
)

var (
	_ErrorCode_En_index_0   = [...]uint8{0, 9, 22}
	_ErrorCode_ZhCn_index_0 = [...]uint8{0, 9, 21}
	_ErrorCode_En_index_1   = [...]uint8{0, 2, 49, 87, 105}
	_ErrorCode_ZhCn_index_1 = [...]uint8{0, 12, 50, 73, 88}
)

// _transOne translate one CONST
func (i ErrorCode) _transOne(locale string) string {
	switch locale {
	case "en":
		switch {
		case 10000 <= i && i <= 10001:
			i -= 10000
			return _ErrorCode_En_name_0[_ErrorCode_En_index_0[i]:_ErrorCode_En_index_0[i+1]]
		case 20000 <= i && i <= 20003:
			i -= 20000
			return _ErrorCode_En_name_1[_ErrorCode_En_index_1[i]:_ErrorCode_En_index_1[i+1]]
		default:
			return "ErrorCode[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
		}
	case "zh_cn":
		switch {
		case 10000 <= i && i <= 10001:
			i -= 10000
			return _ErrorCode_ZhCn_name_0[_ErrorCode_ZhCn_index_0[i]:_ErrorCode_ZhCn_index_0[i+1]]
		case 20000 <= i && i <= 20003:
			i -= 20000
			return _ErrorCode_ZhCn_name_1[_ErrorCode_ZhCn_index_1[i]:_ErrorCode_ZhCn_index_1[i+1]]
		default:
			return "ErrorCode[" + locale + "](" + strconv.FormatInt(int64(i), 10) + ")"
		}
	default:
		// Normally unreachable, should not happen but be cautious
		return ""
	}
}

// _ErrorCode_supported All supported locales record
var _ErrorCode_supported = map[string]int{"en": 0, "zh_cn": 1}

// _ErrorCode_defaultLocale default locale
// generated pass by gdev-i18n flag -defaultlocale, Don't assign directly
var _ErrorCode_defaultLocale = "zh_cn"

// _ErrorCode_ctxKey Key from context.Context Value get locale
// generated pass by gdev-i18n flag -ctxkey, Don't assign directly
var _ErrorCode_ctxKey = "i18n"

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method Error.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the fmt.Stringer interface, so that you can output it directly by package fmt,
//  - If you understand the above mechanism then you can use this method with confidence
func (i ErrorCode) String() string {
	return i._trans(_ErrorCode_defaultLocale)
}

// WARNING: You should use Trans, Lang, Wrap, WrapWithContext method instead
//  - You should not use this method in an internationalized language environment, as well as method String.
//  - Because this method always returns the translation value of the default language.
//  - This method implements the error interface, so that you can return the value as an error,
//  - If you understand the above mechanism then you can use this method with confidence
func (i ErrorCode) Error() string {
	return i._trans(_ErrorCode_defaultLocale)
}

// Code get original type uint value
func (i ErrorCode) Code() uint {
	return uint(i)
}

// Wrap another error with locale set for i18n TYPE Const
//  - err another error
//  - locale i18n locale name
//  - args optional formatting component
func (i ErrorCode) Wrap(err error, locale string, args ...interface{}) *I18nErrorCodeErrorWrap {
	return &I18nErrorCodeErrorWrap{err: err, origin: i, locale: locale, args: args}
}

// 新增
type I18nErrorCodeErrorDetail struct {
	ErrString string // wrap another error
	ErrCode   uint   // custom shaping type Val
	Locale    string // i18n locale set
}

func (i ErrorCode) TransErrorDetail() *I18nErrorCodeDetail {
	return &I18nErrorCodeDetail{ErrString: i._trans(_ErrorCode_defaultLocale), ErrCode: uint(i), Locale: _ErrorCode_defaultLocale}
}

// WrapWithContext wrap another error with context.Context set for i18n TYPE Const
//  - ctx context with Value use Key from _ErrorCode_ctxKey, which pass by gdev-i18n flag -ctxkey
//  - err another error
//  - args optional formatting component
func (i ErrorCode) WrapWithContext(ctx context.Context, err error, args ...interface{}) *I18nErrorCodeErrorWrap {
	return &I18nErrorCodeErrorWrap{err: err, origin: i, locale: _ErrorCode_localeFromCtxWithFallback(ctx), args: args}
}

// I18nErrorCodeErrorWrap type i18n error wrapper
//   WARNING
//   This struct ONLY used to wrap the CONST generated by the gdev-i18n tool,
//   Pass easily obtain internationalized translations through Error, String, Translate
//   WARNING
type I18nErrorCodeErrorWrap struct {
	err    error         // wrap another error
	origin ErrorCode     // custom shaping type Val
	locale string        // i18n locale set
	args   []interface{} // formatted output replacement component
}

// Translate get translated string
func (e *I18nErrorCodeErrorWrap) Translate() string {
	return e.origin.Trans(e.locale, e.args...)
}

// String implement fmt.Stringer, get translated string use Translate
func (e *I18nErrorCodeErrorWrap) String() string {
	return e.Translate()
}

// Error struct as error, get typed message wrap with inside error message
//  - this method will be formatted wrap error if exist.
//  - Only for development and debugging, or logging full error message
//  - if you want to get typed message, please use method String or Translate
func (e *I18nErrorCodeErrorWrap) Error() string {
	if e.err == nil {
		return e.Translate()
	}
	return fmt.Sprintf("%s (%s)", e.Translate(), e.err.Error())
}

// Format alias for method Error
//  - this method will be formatted wrap error if exist.
//  - Only for development and debugging, or logging full error message
//  - if you want to get typed message, please use method String or Translate
func (e *I18nErrorCodeErrorWrap) Format() string {
	return e.Error()
}

// Value get original type value
func (e *I18nErrorCodeErrorWrap) Value() ErrorCode {
	return e.origin
}

// Unwrap an error. Get the error inside
func (e *I18nErrorCodeErrorWrap) Unwrap() error {
	return e.err
}

// IsLocaleSupport Check if the specified locale is supported
func (i ErrorCode) IsLocaleSupport(locale string) bool {
	return _ErrorCode_isLocaleSupport(locale)
}

// Lang get target translate text use context.Context
//  - ctx  context with Value use Key from _ErrorCode_ctxKey, which pass by gdev-i18n flag -ctxkey
//  - args Optional placeholder replacement value, value type of ErrorCode, or type of string
func (i ErrorCode) Lang(ctx context.Context, args ...interface{}) string {
	return i._trans(_ErrorCode_localeFromCtxWithFallback(ctx), args...)
}

// Trans get target translate text use specified language locale identifier
//  - locale specified language locale identifier, need pass by IsLocaleSupport
//  - args Optional placeholder replacement value, value type of ErrorCode, or type of string
func (i ErrorCode) Trans(locale string, args ...interface{}) string {
	if !_ErrorCode_isLocaleSupport(locale) {
		locale = _ErrorCode_defaultLocale
	}
	return i._trans(locale, args...)
}

func _ErrorCode_isLocaleSupport(locale string) bool {
	_, ok := _ErrorCode_supported[locale]
	return ok
}

// _ErrorCode_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _ErrorCode_isLocaleSupport is false
func _ErrorCode_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _ErrorCode_defaultLocale
	}
	v := ctx.Value(_ErrorCode_ctxKey)
	if v == nil {
		return _ErrorCode_defaultLocale
	}
	if vv, ok := v.(string); ok && _ErrorCode_isLocaleSupport(vv) {
		return vv
	}
	return _ErrorCode_defaultLocale
}

// _trans trustworthy parameters inside method
//   - locale i18n local
//   - args   value type of ErrorCode, or type of string
func (i ErrorCode) _trans(locale string, args ...interface{}) string {
	msg := i._transOne(locale)
	if len(args) > 0 {
		var com []interface{}
		for _, arg := range args {
			if typ, ok := arg.(ErrorCode); ok {
				com = append(com, typ._transOne(locale))
			} else {
				com = append(com, arg) // arg as string scalar
			}
		}
		return fmt.Sprintf(msg, com...)
	}
	return msg
}
