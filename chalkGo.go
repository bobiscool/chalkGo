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

type Chalko struct {
    codes []string
}

func (s Chalko) Add(code string) Chalko {
    s.codes = append(s.codes, code)
    return s
}

func (s Chalko) ToString(text string) string {
    return strings.Join(s.codes, "") + text + Reset
}

// style
func (s Chalko) Bold() Chalko { return s.Add(Bold) }
func (s Chalko) Dim() Chalko { return s.Add(Dim) }
func (s Chalko) Italic() Chalko { return s.Add(Italic) }
func (s Chalko) Underline() Chalko { return s.Add(Underline) }
func (s Chalko) Overline() Chalko { return s.Add(Overline) }
func (s Chalko) Inverse() Chalko { return s.Add(Inverse) }
func (s Chalko) Hidden() Chalko { return s.Add(Hidden) }
func (s Chalko) Strikethrough() Chalko { return s.Add(Strikethrough) }
func (s Chalko) Visible() Chalko { return s.Add(Visible) }

// color
func (s Chalko) Black() Chalko { return s.Add(Black) }
func (s Chalko) Red() Chalko { return s.Add(Red) }
func (s Chalko) Green() Chalko { return s.Add(Green) }
func (s Chalko) Yellow() Chalko { return s.Add(Yellow) }
func (s Chalko) Blue() Chalko { return s.Add(Blue) }
func (s Chalko) Magenta() Chalko { return s.Add(Magenta) }
func (s Chalko) Cyan() Chalko { return s.Add(Cyan) }
func (s Chalko) White() Chalko { return s.Add(White) }
func (s Chalko) BlackBright() Chalko { return s.Add(BlackBright) }
func (s Chalko) RedBright() Chalko { return s.Add(RedBright) }
func (s Chalko) GreenBright() Chalko { return s.Add(GreenBright) }
func (s Chalko) YellowBright() Chalko { return s.Add(YellowBright) }
func (s Chalko) BlueBright() Chalko { return s.Add(BlueBright) }
func (s Chalko) MagentaBright() Chalko { return s.Add(MagentaBright) }
func (s Chalko) CyanBright() Chalko { return s.Add(CyanBright) }
func (s Chalko) WhiteBright() Chalko { return s.Add(WhiteBright) }

// background color
func (s Chalko) BgBlack() Chalko { return s.Add(BgBlack) }
func (s Chalko) BgRed() Chalko { return s.Add(BgRed) }
func (s Chalko) BgGreen() Chalko { return s.Add(BgGreen) }
func (s Chalko) BgYellow() Chalko { return s.Add(BgYellow) }
func (s Chalko) BgBlue() Chalko { return s.Add(BgBlue) }
func (s Chalko) BgMagenta() Chalko { return s.Add(BgMagenta) }
func (s Chalko) BgCyan() Chalko { return s.Add(BgCyan) }
func (s Chalko) BgWhite() Chalko { return s.Add(BgWhite) }
func (s Chalko) BgBlackBright() Chalko { return s.Add(BgBlackBright) }
func (s Chalko) BgRedBright() Chalko { return s.Add(BgRedBright) }
func (s Chalko) BgGreenBright() Chalko { return s.Add(BgGreenBright) }
func (s Chalko) BgYellowBright() Chalko { return s.Add(BgYellowBright) }
func (s Chalko) BgBlueBright() Chalko { return s.Add(BgBlueBright) }
func (s Chalko) BgMagentaBright() Chalko { return s.Add(BgMagentaBright) }
func (s Chalko) BgCyanBright() Chalko { return s.Add(BgCyanBright) }
func (s Chalko) BgWhiteBright() Chalko { return s.Add(BgWhiteBright) }


// func main() {
//     fmt.Println(Chalko{}.Red().Bold().Underline().ToString("This is a red, bold, underlined text"))
// }