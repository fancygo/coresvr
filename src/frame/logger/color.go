package logger

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

import (
	"fmt"
)

const (
	COLOR_NONE = 0
	COLOR_1    = 1
	COLOR_2    = 2
	COLOR_3    = 3
	COLOR_4    = 4
	COLOR_5    = 5
	COLOR_6    = 6
	COLOR_7    = 7

	COLOR_MAX = 8
)

var stdTags [COLOR_MAX][3]int

const (
	colorBlack = (iota + 30)
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
	colorCyan
	colorWhite
)

//显示方式
const (
	opDefault   = 0
	opHighlight = 1
	opUnderline = 4
	opShine     = 5
	opInverse   = 7
	opUnvis     = 8
)

//前景
const (
	fBlack   = 30
	fRed     = 31
	fGreen   = 32
	fYellow  = 33
	fBlue    = 34
	fMagenta = 35
	fCyan    = 36
	fWhite   = 37
)

//背景
const (
	bBlack   = 40
	bRed     = 41
	bGreen   = 42
	bYellow  = 43
	bBlue    = 44
	bMagenta = 45
	bCyan    = 46
	bWhite   = 47
)

func colorInit() {
	stdTags = [COLOR_MAX][3]int{}
	stdTags[COLOR_1] = [3]int{opHighlight, fWhite, bBlack}
	stdTags[COLOR_2] = [3]int{opHighlight, fBlue, bBlack}
	stdTags[COLOR_3] = [3]int{opHighlight, fGreen, bBlack}
	stdTags[COLOR_4] = [3]int{opHighlight, fYellow, bBlack}
	stdTags[COLOR_5] = [3]int{opHighlight, fMagenta, bBlack}
	stdTags[COLOR_6] = [3]int{opHighlight, fCyan, bBlack}
	stdTags[COLOR_7] = [3]int{opHighlight, fRed, bBlack}

}

func colorString(level int, str string) string {
	if level < COLOR_1 || level >= COLOR_MAX {
		return str
	}

	colorStr := fmt.Sprintf(" %c[%d;%d;%dm%s%s%c[0m ", 0x1b, stdTags[level][0], stdTags[level][2], stdTags[level][1], "", str, 0x1b)
	return colorStr
}
