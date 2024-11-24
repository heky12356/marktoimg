package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

// 画布
var dc *gg.Context
var x, y float64

// 字体大小
var fontSize = 60.0
var quoteFontSize = 40.0
var listFontSize = 50.0
var headingFontSizemap = map[int]float64{
	1: fontSize * 2.0,
	2: fontSize * 1.0,
	3: fontSize * 0.83,
	4: fontSize * 0.67,
	5: fontSize * 0.5,
	6: fontSize * 0.33,
}

// 字体缩进
var fontIndentLeft = 30.0
var fontIndentTop = 100.0

// 画布大小
var canvaWidth = 1080 * 2
var canvaHeight = 720 * 3

// 行高
var lineHeight = 70.0

// 字体路径
var nomarlFont = "./ttf/STFANGSO.TTF"
var boldFont = "./ttf/msyhbd.ttc"
var italicFont = "./ttf/SitkaVF-Italic.ttf"

// 标题下的横线
var setHeadLine = true
var headLineLength = canvaWidth // 实际上会减去fontIndentLeft * 2,左右各留一点空白
var headLineWidth = 2.0         // 线粗细
var headLineColor = color.RGBA{R: 211, G: 211, B: 211, A: 255}

// server
var serverPort = ":8000"

type ImageInput struct {
	Input string `json:"input"`
}
