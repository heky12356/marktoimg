package main

import (
	"fmt"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func main() {

	// 	source := []byte(`
	// # Heading 1

	// ## Heading 2

	// ### Heading 3

	// #### Heading 4

	// ##### Heading 5

	// ###### Heading 6

	// This is a paragraph with some **bold** text.

	// This is a paragraph with some *italic* text.

	// - List item 1
	// - List item 2

	// > This is a blockquote.
	// 	`)

	// 示例输入字符串
	input := "## 如何认领\r\n\r\n请直接在下方回复，我会把 issue 的 assignees 设置为你\r\n\r\n## 需求\r\n\r\n当有人创建 issue 报名时，自动将报名信息同步到 qq 群聊\r\n\r\n\r\n## 实现方式\r\n\r\n用户创建 issue，自动触发 github webhook 发送 post 请求; 这一步需要: \r\n\r\n- 在 github 设置配置 webhook url 等内容\r\n\r\n后端收到请求并处理; 这一步需要:\t\r\n\r\n- 解析 issue 内容, 提取出作者，标题等信息\r\n\r\n转发到 qq 群; 这一步需要:\r\n\r\n- 调用 qq 机器人发送消息\r\n\r\n## 文档资料\r\n\r\n- github webhook: https://docs.github.com/zh/webhooks\r\n- onebot（统一的聊天机器人应用接口标准）: https://onebot.dev/\r\n- napcat(qq 的 onebot 实现): https://napneko.com/\r\n- nonebot(python 机器人框架): https://nonebot.dev/docs/quick-start\r\n- fastapi(仅举例，似乎是 nonebot 默认推荐的框架): https://fastapi.tiangolo.com/zh/tutorial/first-steps/\r\n\r\n## 难度\r\n\r\n中等 (需要学习一定知识，但是功能体量不大，稍有挑战)\r\n\r\n## 能力要求\r\n\r\n- 需要能够正常访问 github，并注册账号\r\n- 知道 webhook 的概念\r\n- 了解后端(例如 fastapi)基础使用方法\r\n- 对编写 qq 机器人感兴趣\r\n\r\n## 关闭 Issue 前请确认以下内容\r\n\r\n- [x] 代码已经上传：<https://github.com/io-club/IOGAS-QQ>\r\n- [x] 准备参加茶话会分享：<https://github.com/io-club/share/issues/19>"
	//input := "<https://github.com/io-club/IOGAS-QQ>"
	// 字符串切割
	lines := convertToMarkdown(input)

	outputPath := "output.png"
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

// func drawMarkdownToImage(mdContent string, outputPath string) {
//

// 	md := goldmark.New()
// 	reader := text.NewReader([]byte(mdContent))
// 	document := md.Parser().Parse(reader)

//
// }
