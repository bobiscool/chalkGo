package github.com/bobiscool/chalkGo

import (
    "strings"
    "fmt"
)

const (
    Reset          = "\033[0m"
    Bold           = "\033[1m"
    Dim            = "\033[2m"
    Italic         = "\033[3m"
    Underline      = "\033[4m"
    Overline       = "\033[53m"
    Inverse        = "\033[7m"
    Hidden         = "\033[8m"
    Strikethrough  = "\033[9m"
    Visible        = "\033[28m"

    Black          = "\033[30m"
    Red            = "\033[31m"
    Green          = "\033[32m"
    Yellow         = "\033[33m"
    Blue           = "\033[34m"
    Magenta        = "\033[35m"
    Cyan           = "\033[36m"
    White          = "\033[37m"

    BlackBright    = "\033[90m"
    RedBright      = "\033[91m"
    GreenBright    = "\033[92m"
    YellowBright   = "\033[93m"
    BlueBright     = "\033[94m"
    MagentaBright  = "\033[95m"
    CyanBright     = "\033[96m"
    WhiteBright    = "\033[97m"

    BgBlack        = "\033[40m"
    BgRed          = "\033[41m"
    BgGreen        = "\033[42m"
    BgYellow       = "\033[43m"
    BgBlue         = "\033[44m"
    BgMagenta      = "\033[45m"
    BgCyan         = "\033[46m"
    BgWhite        = "\033[47m"

    BgBlackBright  = "\033[100m"
    BgRedBright    = "\033[101m"
    BgGreenBright  = "\033[102m"
    BgYellowBright = "\033[103m"
    BgBlueBright   = "\033[104m"
    BgMagentaBright= "\033[105m"
    BgCyanBright   = "\033[106m"
    BgWhiteBright  = "\033[107m"
)


const version = "v0.0.1"

func Version() string {
	return version
}

type ChalkGo struct {
    codes []string
}

func (s ChalkGo) Add(code string) ChalkGo {
    s.codes = append(s.codes, code)
    return s
}

func (s ChalkGo) ToString(text string) string {
    return strings.Join(s.codes, "") + text + Reset
}

// style
func (s ChalkGo) Bold() ChalkGo { return s.Add(Bold) }
func (s ChalkGo) Dim() ChalkGo { return s.Add(Dim) }
func (s ChalkGo) Italic() ChalkGo { return s.Add(Italic) }
func (s ChalkGo) Underline() ChalkGo { return s.Add(Underline) }
func (s ChalkGo) Overline() ChalkGo { return s.Add(Overline) }
func (s ChalkGo) Inverse() ChalkGo { return s.Add(Inverse) }
func (s ChalkGo) Hidden() ChalkGo { return s.Add(Hidden) }
func (s ChalkGo) Strikethrough() ChalkGo { return s.Add(Strikethrough) }
func (s ChalkGo) Visible() ChalkGo { return s.Add(Visible) }

// color
func (s ChalkGo) Black() ChalkGo { return s.Add(Black) }
func (s ChalkGo) Red() ChalkGo { return s.Add(Red) }
func (s ChalkGo) Green() ChalkGo { return s.Add(Green) }
func (s ChalkGo) Yellow() ChalkGo { return s.Add(Yellow) }
func (s ChalkGo) Blue() ChalkGo { return s.Add(Blue) }
func (s ChalkGo) Magenta() ChalkGo { return s.Add(Magenta) }
func (s ChalkGo) Cyan() ChalkGo { return s.Add(Cyan) }
func (s ChalkGo) White() ChalkGo { return s.Add(White) }
func (s ChalkGo) BlackBright() ChalkGo { return s.Add(BlackBright) }
func (s ChalkGo) RedBright() ChalkGo { return s.Add(RedBright) }
func (s ChalkGo) GreenBright() ChalkGo { return s.Add(GreenBright) }
func (s ChalkGo) YellowBright() ChalkGo { return s.Add(YellowBright) }
func (s ChalkGo) BlueBright() ChalkGo { return s.Add(BlueBright) }
func (s ChalkGo) MagentaBright() ChalkGo { return s.Add(MagentaBright) }
func (s ChalkGo) CyanBright() ChalkGo { return s.Add(CyanBright) }
func (s ChalkGo) WhiteBright() ChalkGo { return s.Add(WhiteBright) }

// background color
func (s ChalkGo) BgBlack() ChalkGo { return s.Add(BgBlack) }
func (s ChalkGo) BgRed() ChalkGo { return s.Add(BgRed) }
func (s ChalkGo) BgGreen() ChalkGo { return s.Add(BgGreen) }
func (s ChalkGo) BgYellow() ChalkGo { return s.Add(BgYellow) }
func (s ChalkGo) BgBlue() ChalkGo { return s.Add(BgBlue) }
func (s ChalkGo) BgMagenta() ChalkGo { return s.Add(BgMagenta) }
func (s ChalkGo) BgCyan() ChalkGo { return s.Add(BgCyan) }
func (s ChalkGo) BgWhite() ChalkGo { return s.Add(BgWhite) }
func (s ChalkGo) BgBlackBright() ChalkGo { return s.Add(BgBlackBright) }
func (s ChalkGo) BgRedBright() ChalkGo { return s.Add(BgRedBright) }
func (s ChalkGo) BgGreenBright() ChalkGo { return s.Add(BgGreenBright) }
func (s ChalkGo) BgYellowBright() ChalkGo { return s.Add(BgYellowBright) }
func (s ChalkGo) BgBlueBright() ChalkGo { return s.Add(BgBlueBright) }
func (s ChalkGo) BgMagentaBright() ChalkGo { return s.Add(BgMagentaBright) }
func (s ChalkGo) BgCyanBright() ChalkGo { return s.Add(BgCyanBright) }
func (s ChalkGo) BgWhiteBright() ChalkGo { return s.Add(BgWhiteBright) }


// func main() {
//     fmt.Println(ChalkGo{}.Red().Bold().Underline().ToString("This is a red, bold, underlined text"))
// }