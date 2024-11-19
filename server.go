package main

import (
	"fmt"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func setimg(input string) {
	lines := convertToMarkdown(input)

	outputPath := "/img/output.png"
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

	// 保存为 PNG
	if err := dc.SavePNG(outputPath); err != nil {
		log.Fatalf("failed to save PNG: %v", err)
	}

	fmt.Printf("Markdown image saved to %s\n", outputPath)
}
