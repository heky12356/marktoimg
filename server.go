package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getimg(c *gin.Context) {
	var input ImageInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 假设 setimg 函数会处理图片并生成 output.png
	setimg(input.Input)

	// 读取生成的图片文件
	imagePath := "img/output.png"
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read image"})
		return
	}

	// 设置响应头并返回图片数据
	c.Header("Content-Type", "img/png")
	c.Header("Content-Disposition", "inline; filename=output.png")
	c.Data(http.StatusOK, "img/png", imageData)
}
