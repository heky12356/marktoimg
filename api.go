package main

import (
	"fmt"
	"image/color"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// 递归处理节点及其子节点
func processNode(node ast.Node, source []byte, textType string) {
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		processNodeByType(child, source, textType)
	}
}

func processNodeByType(node ast.Node, source []byte, textType string) {
	switch n := node.(type) {
	case *ast.Text:
		processText(n, source, textType)
	case *ast.Heading:
		fmt.Printf("\n")
		newLine("0")
		headingType := "h" + strconv.Itoa(int(n.Level))
		processNode(n, source, headingType)
	case *ast.Emphasis:
		if n.Level == 2 {
			fmt.Print("Bold Text: ")
			processNode(n, source, "bold")
		} else if n.Level == 1 {
			fmt.Print("Italic Text: ")
			processNode(n, source, "italic")
		}
	case *ast.ListItem:
		fmt.Printf("ListItem: \n")
		newLine("0")
		Listintend()
		processNode(n, source, "listItem")
	case *ast.Link:
		fmt.Print("Link: ")
		processNode(n, source, "link")
	case *ast.AutoLink:
		fmt.Printf("AutoLink: ")
		processNode(n, source, "link")
	case *ast.Blockquote:
		fmt.Printf("Blockquote: ")
		processNode(n, source, "quotep")
	case *ast.Paragraph:
		newLine("0")
		fmt.Printf("Paragraph: ")
		if textType == "quotep" {
			processNode(n, source, "quote")
		} else {
			processNode(n, source, "0")
		}
	default:
		fmt.Printf("Node Type: %T\n", n)
		if _, ok := n.(*ast.TextBlock); !ok {
			fmt.Printf("\n")
		}
		processNode(node, source, textType)
	}
}

func processText(n *ast.Text, source []byte, textType string) {
	text := n.Segment.Value(source)
	switch textType {
	case "bold":
		fmt.Printf("bold: %s", text)
		drawBold(text)
	case "italic":
		fmt.Printf("italic: %s", text)
		drawItalic(text)
	case "listItem":
		fmt.Printf("listItem: %s", text)
		drawList(text)
	case "quote":
		fmt.Printf("quote: %s", text)
		drawquote(text)
	case "h1", "h2", "h3", "h4", "h5", "h6":
		level, _ := strconv.Atoi(textType[1:])
		fmt.Printf("%s: %s", textType, text)
		drawHeading(text, level)
	default:
		fmt.Printf("%s: %s", textType, text)
		drawText(text, "0")
	}
}

// // 提取link节点的文本(包含链接本身)
// func extractText(node ast.Node, source []byte) string {
// 	var result string
// 	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
// 		if textNode, ok := n.(*ast.Text); ok && entering {
// 			result += string(textNode.Segment.Value(source))
// 		}
// 		return ast.WalkContinue, nil
// 	})
// 	return result
// }

// 初始化画布
func imginit() {
	x = 0.0
	y = 0.0
	canvaWidth = OriginCanvaWidth
	canvaHeight = OrifinCanvaHeight
}

func initCanvas() *gg.Context {

	// 设置画布大小
	width := canvaWidth
	height := canvaHeight

	x += fontIndentLeft
	y += fontIndentTop
	// 创建 gg context
	dc := gg.NewContext(width, height)

	// 设置白色背景
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// 加载字体
	fontPath := nomarlFont
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		log.Fatalf("failed to load font: %v", err)
	}

	// 设置默认颜色为黑色
	dc.SetColor(color.Black)
	return dc
}

// 将字符串转换为 Markdown 格式的代码块
func convertToMarkdown(input string) []string {
	// 分割字符串
	lines := strings.Split(input, "\n")

	// 检查是不是有尖括号括起来的链接

	for i, line := range lines {
		re := regexp.MustCompile(`<([^>]+)>`)
		// 定义正则表达式，匹配 - [ ] 和 - [X]
		re2 := regexp.MustCompile(`- \[( |(?i)x)\]`)
		lines[i] = re.ReplaceAllString(line, `$1`)
		fmt.Printf("line %d: %s\n", i, lines[i])
		// 将 - [ ] 和 - [x] 替换为 (未完成) 和 (已完成)
		lines[i] = re2.ReplaceAllStringFunc(lines[i], func(match string) string {
			if match == "- [ ]" {
				return "- checkn"
			} else if match == "- [X]" {
				return "- checkx"
			} else if match == "- [x]" {
				return "- checkx"
			}
			return match
		})
	}
	return lines
}

func setimg(input string) {
	lines := convertToMarkdown(input)

	outputPath := "img/output.png"
	logoPath := "logo/io_logo.png" // logo图片路径
	// 初始化 Goldmark 并解析 Markdown
	md := goldmark.New()

	// 初始化画布
	imginit()
	dc = initCanvas()
	// 递归处理节点
	for _, line := range lines {
		reader := text.NewReader([]byte(line))
		document := md.Parser().Parse(reader)
		processNode(document, []byte(line), "0")
	}

	// 加载logo图片
	logo, err := gg.LoadImage(logoPath)
	if err != nil {
		log.Fatalf("failed to load logo image: %v", err)
	}

	// 将logo图片绘制到画布上
	dc.DrawImage(logo, 0, canvaHeight-logoHeight-100) //

	// 保存为 PNG
	if err := dc.SavePNG(outputPath); err != nil {
		log.Fatalf("failed to save PNG: %v", err)
	}

	fmt.Printf("Markdown image saved to %s\n", outputPath)
}
