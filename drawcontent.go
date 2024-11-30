package main

import (
	"image/color"
	"log"
	"regexp"

	"github.com/fogleman/gg"
)

// 不同节点类型的处理函数

// 一般文字
func drawText(text []byte, texttype string) {
	for len(text) > 0 {
		width, _ := dc.MeasureString(string(text))
		if x+width > float64(dc.Width()) {
			// 计算可以放下的字符数
			for i := 1; i <= len(text); i++ {
				subWidth, _ := dc.MeasureString(string(text[:i]))
				if x+subWidth > float64(dc.Width()) {
					// 绘制当前行可以放下的部分
					dc.DrawString(string(text[:i-1]), x, y)
					x += subWidth
					text = text[i-1:]
					newLine(texttype)
					addCanvasHeight()
					break
				}
			}
		} else {
			dc.DrawString(string(text), x, y)
			addCanvasHeight()
			x += width
			break
		}
	}
}

// 粗体
func drawBold(text []byte) {
	// 通过切换字体文件来绘制粗体
	if err := dc.LoadFontFace(boldFont, fontSize); err != nil {
		log.Fatalf("failed to load bold font: %v", err)
	}
	drawText(text, "0")
	//切换回正常字体
	if err := dc.LoadFontFace(nomarlFont, fontSize); err != nil {
		log.Fatalf("failed to load font: %v", err)
	}
}

// 斜体
func drawItalic(text []byte) {
	// 通过切换字体文件来绘制斜体
	if err := dc.LoadFontFace(italicFont, fontSize); err != nil {
		log.Fatalf("failed to load bold font: %v", err)
	}
	drawText(text, "0")
	//切换回正常字体
	if err := dc.LoadFontFace(nomarlFont, fontSize); err != nil {
		log.Fatalf("failed to load font: %v", err)
	}
}

// 换行
func newLine(texttype string) {
	y += lineHeight
	// x = fontIndentLeft
	if texttype == "quote" {
		x = fontIndentLeft * 2
	} else if texttype == "list" {
		x = fontIndentLeft * 3.5
	} else {
		x = fontIndentLeft
	}
}

// 引用
func drawquote(text []byte) {
	// 通过切换大小和颜色来绘制引用
	if err := dc.LoadFontFace(italicFont, quoteFontSize); err != nil {
		log.Fatalf("failed to load bold font: %v", err)
	}
	dc.SetColor(color.RGBA{R: 211, G: 211, B: 211, A: 255}) // light grey color
	x += fontIndentLeft
	drawText(text, "quote")
	//切换回正常字体
	if err := dc.LoadFontFace(nomarlFont, fontSize); err != nil {
		log.Fatalf("failed to load font: %v", err)
	}
	dc.SetColor(color.Black)
	//计算宽带然后让x+宽度
	width, _ := dc.MeasureString(string(text))
	x += width
}

// 列表
func Listintend() {
	if err := dc.LoadFontFace(boldFont, listFontSize); err != nil {
		log.Fatalf("failed to load bold font: %v", err)
	}
	dc.DrawString("·", x+fontIndentLeft, y)
	x += fontIndentLeft * 2
}
func drawList(text []byte) {
	// 通过切换大小来绘制列表
	re2 := regexp.MustCompile(`checkx|checkn`)
	if re2.Match(text) {
		// 绘制复选框
		coverText(x-fontSize, y-fontSize)
		x -= fontIndentLeft
		drawCheckbox(dc, x, y, 30, re2.FindString(string(text)) == "checkx")
		text = []byte(re2.ReplaceAllString(string(text), ""))
		x += fontIndentLeft
	}
	if err := dc.LoadFontFace(nomarlFont, listFontSize); err != nil {
		log.Fatalf("failed to load bold font: %v", err)
	}
	//x += fontIndentLeft
	drawText(text, "list")
	//切换回正常字体
	if err := dc.LoadFontFace(nomarlFont, fontSize); err != nil {
		log.Fatalf("failed to load font: %v", err)
	}
}

// 标题
func drawHeading(text []byte, level int) {
	// 通过切换大小来绘制标题
	if err := dc.LoadFontFace(boldFont, headingFontSizemap[level]); err != nil {
		log.Fatalf("failed to load bold font: %v", err)
	}
	newLine("0")
	dc.DrawString(string(text), x, y)
	addCanvasHeight()
	drawLine(dc, x, y+float64(margin), float64(headLineLength)-x, y+float64(margin))
	newLine("0")
	//切换回正常字体
	if err := dc.LoadFontFace(nomarlFont, fontSize); err != nil {
		log.Fatalf("failed to load font: %v", err)
	}
	//计算宽带然后让x+宽度
	width, _ := dc.MeasureString(string(text))
	x += width
	// 根据不同字体大小来换行
	//y += headingFontSizemap[level] * 1.2
}

// 绘制线条
func drawLine(dc *gg.Context, x1, y1, x2, y2 float64) {
	if setHeadLine {
		dc.SetLineWidth(headLineWidth)
		dc.SetColor(headLineColor)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
		dc.SetColor(color.Black)
	}
}

// 绘制复选框
func drawCheckbox(dc *gg.Context, x, y, size float64, checked bool) {
	t_x := x
	t_y := y - size
	// 绘制复选框边框
	dc.DrawRectangle(t_x, t_y, size, size)
	dc.Stroke()

	if checked {
		// 绘制打勾符号
		dc.DrawLine(t_x+size*0.2, t_y+size*0.5, t_x+size*0.4, t_y+size*0.7)
		dc.DrawLine(t_x+size*0.4, t_y+size*0.7, t_x+size*0.8, t_y+size*0.3)
		dc.Stroke()
	}
}

// 画布长度自增
func addCanvasHeight() {
	// 检查画布高度是否足够
	if y+lineHeight+float64(logoHeight) > float64(dc.Height()) {
		// 增加画布高度
		canvaHeight += int(0.2*float64(canvaHeight)) + logoHeight
		newDc := initCanvas()

		// 将旧画布的内容复制到新画布上
		newDc.DrawImage(dc.Image(), 0, 0)

		// 替换旧画布
		dc = newDc
	}
}

func coverText(x, y float64) {
	// 设置颜色为白色
	dc.SetColor(color.White)
	// 绘制白色长方形
	dc.DrawRectangle(x, y, fontSize, fontSize)
	// 填充长方形
	dc.Fill()
	// 恢复颜色
	dc.SetColor(color.Black)
}
