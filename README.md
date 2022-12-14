# gdev-i18n

gdev-i18n is a golang international language pack tool, refer to the official [stinger](https://github.com/golang/tools/tree/master/cmd/stringer) tool

golang国际化语言包工具，参考了官方stringer工具

# 一、Usage/使用

## 1.1、Build/下载编译

因未使用`ioutil`包，构建二进制可执行文件要求go最低版本1.16，因生成的代码使用了`context.Context`，生成的代码支持1.7及其以上版本

Because `ioutil` is not used,
to build a binary executable file the minimum version of Go 1.16 is required,
because of the use of `context.Context` the generated code can be used for version 1.7 and above

go Version 1.16 and above for install/使用go1.16及其以上版本编译安装
````
go install github.com/gqgo/gdev-i18n@latest
````

或者也可以直接从release下载对应平台编译好的二进制
可执行程序放置于环境变量目录下
即可使用：[https://github.com/gqgo/gdev-i18n/releases](https://github.com/gqgo/gdev-i18n/releases)

---

The `example` directory gives an example of usage witch used `gomod`

`example`目录给出一种示例，该示例使用了`gomod`，请参考
````
cd example
# Note that the files generated in `example` directory are ignored, 
# you need to execute the following command to generate
go generate ./...
go run main.go
````

开发调试可以使用`make`指令，可以调试`test`目录下的多个示例

You can use the `make` command for development and debugging,
can debug multiple examples in the `test` directory
````
make debug
````

## 1.2、Define Numerical shaping Constant/定义整数数值型常量

````
type Code int

const (
    ERROROFYOU Code = iota + 1   
    ERROROFMINE Code = 10000
)
````

> Numerical shaping: such as `int`, `uint`, `uint32`, etc.

## 1.3、Define Language Package/定义语言包

> Use only TOML format files

* 定义语言包目录：语言包目录位于定义常量源码文件的同级目录下的子目录，默认语言包目录名称为`i18n`
* The language package directory is located in a subdirectory of the same level directory that defines the constant source code file. The default language package directory name is `i18n`
* 语言包目录下使用TOML文件定义i18n翻译的键值对；
* Use TOML files in the language package directory to define key-value pairs for i18n translation
* 语言包目录下若使用子目录，则子目录名将被作为语言类型标记，子目录下TOML文件名和文件数目不做限制，例如：`en`；
* If a subdirectory is used in the language package directory, the name of the subdirectory will be marked as the language type, TOML file name and number of files in subdirectories are not limited. For example: `en`;
* 语言包目录下不使用子目录直接定义TOML文件的，则TOML文件的文件名将被作为语言类型标记，例如：`en.toml`；
* If the TOML file is directly defined in language package directory, the file name of the TOML file will be marked as the language type. For example: `en.toml`;
* 语言包键值对的键名使用常量字面量，上述例子中`ERROROFYOU`就将作为键名；
* Use constant literals for the key names of language pack key-value pairs. In the above example, `ERROROFYOU` will be used as the key name

## 1.4、代码生成/Generate Code

Go to the directory that defines the constant under the terminal and execute it

终端下进入到定义常量的目录直接执行：
````
$GOPATH/bin/gdev-i18n -type Code -tomlpath i18n
````

You can also use the `go generate` command directly in the constant source code

你也可以配合`go generate`指令直接在定于常量的源码里使用
````
// write this code in your golang source code, then use `go generate` command
//go:generate $GOPATH/bin/gdev-i18n -type Code -tomlpath i18n
````

## 1.5、开发调试/Dev and debugging

可以使用`-check`指令参数核对语言包缺失的键值对

You can use the `-check` command
to check the missing key-value pairs of the language pack

````
$GOPATH/bin/gdev-i18n -type Code -tomlpath i18n -check
````

带`-check`的命令执行后如果有缺失键值对或无用键值对，则可能会输出如下提示以协助开发

After the command with `-check` is executed,
if there are missing key-value pairs or useless key-value pairs,
the following prompt may be output to assist development

````
gdev-i18n: Check Fail
gdev-i18n: The missing key-value pair information as follows
gdev-i18n: You can copy and fill it to the corresponding TOML file
************TYPE `Single` locale `zh-hk` missing key-value pair************
Sig01=""
Sig02=""
Sig03=""
gdev-i18n: Check Warning
gdev-i18n: key-value pairs that will not be used because there is no corresponding constant
gdev-i18n: You can delete the key-value pairs in the corresponding TOML file
************Can be deleted TOML keys of locale `en`************
HELLO
WORLD
````

## 1.6、指令详情/command details

Get more help information about commands

获取更多命令使用帮助信息：
````
$GOPATH/bin/gdev-i18n --help
````

````
Usage of gdev-i18n:
        gdev-i18n [flags] -type T [directory]
        gdev-i18n [flags] -type T -tomlpath DIR -check # just for check
        gdev-i18n [flags] -type T -defaultlocale LOCALE -tomlpath DIR files... # Must be a single package
For more information, see:
        https://github.com/gqgo/gdev-i18n
Flags:
  -check
        Check missing or useless key-value pairs in TOML
  -ctxkey string
        key used by context.Value for get locale; default i18nLocale
  -defaultlocale string
        set default locale name; default naturally sorted first
  -output string
        output file name; default srcdir/<type>_i18n_string.go
  -tags string
        comma-separated list of build tags to apply
  -tomlpath string
        set toml i18n file path; default srcdir/i18n
  -type string
        comma-separated list of type names; must be set
````

> If your GOBIN directory has been added to the environment variable, the above `$GOPATH/bin/` can also be omitted

> 如果你的GOBIN目录已加入环境变量，上述`$GOPATH/bin/`也是可以省略的

## 1.7、调用/Code call

> For example

Directory tree
````
.
├── i18n
│     └── en.toml
│     ├── zh_cn.toml
│     └── zh_hk
│     │     ├── user.toml
│     │     └── merchant.toml
└── code.go
````

Given the name of a (signed or unsigned) integer type T that has constants defined at file `code.go`
````
package foo

type Pill int

const (
    Placebo Pill = iota
    Aspirin
    Ibuprofen
    Paracetamol
    Acetaminophen = Paracetamol // NOTE: with the same value will be ignored, do not use same value
)
````

Define TOML key-value pairs in all locale TOML file,
example for `i18n/en.toml`
````
Placebo="en locale Placebo"
Aspirin="en locale Aspirin"
Ibuprofen="en locale Ibuprofen"
Acetaminophen="en locale Acetaminophen"
````

running this command
````
gdev-i18n -type=Pill
````

in the same directory will create the file `pill_i18n_string.go`, in package `foo`,
containing a definition of, and a struct `I18nPillErrorWrap` will also be created

type `Pill` Added method
````
func (Pill) String() string
func (Pill) Error() string
func (Pill) Wrap(err error, locale string, args ...interface{}) I18nPillErrorWrap
func (Pill) WrapWithContext(ctx context.Context, err error, args ...interface{}) I18nPillErrorWrap
func (Pill) IsLocaleSupport(locale string) bool
func (Pill) Lang(ctx context.Context, args ...interface{}) string
func (Pill) Trans(locale string, args ...interface{}) string
````

Now you can use type `Pill`'s methods with the locale identifier to get the translation value

因部分翻译文本中可能会使用诸如`%s`类型的替换佔位符在代码中实时更改，建议规划好整形数值区间，
某些区间的值专门用于替换`%s`的。

Because some translation texts may use replacement placeholders such as `%s` to change in the code in real time, 
it is recommended to plan the integer value range, and this range values are specifically used to replace `%s`.

# 二、TOML规范支持/TOML Specification Support

TOML Link : [https://toml.io/en/](https://toml.io/en/)

* TOML文件仅支持`Basic strings`形式的字符串键值对，形如：`Key="value"`，也就意味着一对键值对只能位于一行
* TOML file only support `Basic strings` key/value, shaped like `Key="value"`, do not support `Multi-line basic strings`, pair K/V only be located on one line
* TOML键名仅支持`裸键`，键名只能包含ASCII字母，ASCII数字，下划线和短横线（`A-Za-z0-9_-`）
* TOML file only support `Bare keys`,only contain ASCII letters, ASCII digits, underscores, and dashes(`A-Za-z0-9_-`)
* 区块也就是TOML官方的`Table`将被忽略
* The block section, which is the TOML official `Table`, will be ignored
* 支持`#`开头的注释，注释将被忽略
* Support comments starting with `#`, comments will be ignored
