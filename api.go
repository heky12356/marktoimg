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
		switch n := child.(type) {
		case *ast.Text: // 普通文本
			fmt.Printf("Text: ")
			text := n.Segment.Value(source)
			switch textType {
			case "bold":
				fmt.Printf("bold : %s", text)
				drawBold(text)
			case "italic":
				fmt.Printf("italic : %s", n.Segment.Value(source))
				drawItalic(text)
			case "listItem":
				fmt.Printf("listItem : %s", n.Segment.Value(source))
				drawList(text)
			case "quote":
				fmt.Printf("quote : %s", n.Segment.Value(source))
				drawquote(text)
			case "h1":
				fmt.Printf("h1 : %s", n.Segment.Value(source))
				drawHeading(text, 1)
			case "h2":
				fmt.Printf("h2 : %s", n.Segment.Value(source))
				drawHeading(text, 2)
			case "h3":
				fmt.Printf("h3 : %s", n.Segment.Value(source))
				drawHeading(text, 3)
			case "h4":
				fmt.Printf("h4 : %s", n.Segment.Value(source))
				drawHeading(text, 4)
			case "h5":
				fmt.Printf("h5 : %s", n.Segment.Value(source))
				drawHeading(text, 5)
			case "h6":
				fmt.Printf("h6 : %s", n.Segment.Value(source))
				drawHeading(text, 6)
			default:
				fmt.Printf(textType+"：%s", n.Segment.Value(source))
				drawText(text, "0")
			}
		case *ast.Heading:
			fmt.Printf("\n")
			newLine("0")
			htype := "h" + strconv.Itoa(int(n.Level))
			processNode(n, source, htype) // 递归处理子节点

		case *ast.Emphasis:
			if n.Level == 2 {
				fmt.Print("Bold Text: ")
				processNode(n, source, "bold") // 递归处理子节点
			} else if n.Level == 1 { // 斜体文本
				fmt.Print("Italic Text: ")
				processNode(n, source, "italic") // 递归处理子节点
			}
		case *ast.ListItem:
			fmt.Printf("ListItem: ")
			fmt.Printf("\n")
			newLine("0")
			processNode(n, source, "listItem") // 递归处理子节点
		case *ast.Link: // 链接
			fmt.Print("Link: ")
			processNode(n, source, "link") // 递归处理子节点
		case *ast.AutoLink:
			fmt.Printf("AutoLink:")
			processNode(n, source, "Link")
		case *ast.Blockquote: // 引用
			fmt.Printf("Blockquote: ")
			processNode(n, source, "quotep") // 递归处理子节点
		case *ast.Paragraph: // 段落
			newLine("0")
			fmt.Printf("Paragraph: ")
			if textType == "quotep" {
				processNode(n, source, "quote") // 递归处理子节点
			} else {
				processNode(n, source, "0") // 递归处理子节点
			}
		default:
			fmt.Printf("Node Type: %T\n", n)
			if _, ok := n.(*ast.TextBlock); !ok {
				fmt.Printf("\n")
				//newLine()
			}
			processNode(child, source, textType) // 递归处理子节点
		}
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
func initCanvas() *gg.Context {

	// 设置画布大小
	width := canvaWidth
	height := canvaHeight

	x, y = fontIndentLeft, fontIndentTop
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
		// 定义正则表达式，匹配 - [ ] 和 - [x]
		re2 := regexp.MustCompile(`- \[( |x)\]`)
		lines[i] = re.ReplaceAllString(line, `$1`)
		fmt.Printf("line %d: %s\n", i, lines[i])
		// 将 - [ ] 和 - [x] 替换为 (未完成) 和 (已完成)
		lines[i] = re2.ReplaceAllStringFunc(lines[i], func(match string) string {
			if match == "- [ ]" {
				return "- checkn"
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
	dc.DrawImage(logo, 0, 400) // 这里的10, 10是logo图片的绘制位置，可以根据需要调整

	// 保存为 PNG
	if err := dc.SavePNG(outputPath); err != nil {
		log.Fatalf("failed to save PNG: %v", err)
	}

	fmt.Printf("Markdown image saved to %s\n", outputPath)
}
