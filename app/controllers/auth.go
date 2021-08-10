package controllers

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gostu/app/validates"
	"gostu/app/validates/rules"
	"gostu/pkg/response"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"math"
	"os"
)

func Captcha(c *gin.Context) {

}

func PasswordLogin(c *gin.Context) {
	var data validates.SignUpParam
	if err := c.ShouldBind(&data); err != nil {
		//获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator.ValidationErrors类型的错误直接返回
			response.ErrorResponse(c, response.ValidateCheckError, err.Error())
			return
		}

		//validator.ValidationErrors类型的错误进行翻译
		response.ErrorResponse(c, response.ValidateCheckError, rules.Translate(errs))
		return
	}
	response.SuccessResponse(c, "成功了", data)
}

//GBK转utf8的方法
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}

	//source, err := file.Open()
	//if err != nil {
	//	response.ErrorResponse(c, response.ValidateCheckError, err.Error())
	//	return
	//}
	//defer source.Close()
	//
	//var content []byte
	//var tmp = make([]byte, 128)
	//
	//for {
	//	n, err := source.Read(tmp)
	//	if err == io.EOF {
	//		break;
	//	}
	//	if err != nil {
	//		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
	//		return
	//	}
	//
	//	content = append(content, tmp[:n]...)
	//}
	//
	//c.Writer.WriteString(ConvertToString(string(content), "gbk", "utf-8"));

	dst := fmt.Sprintf("/upload/%s", file.Filename)
	if err := c.SaveUploadedFile(file, "."+dst); err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	response.SuccessResponse(c, "上传成功", dst)

}

// 显示图片接口
func ShowFile(c *gin.Context) {
	//imageName := c.Query("imageName")
	var data validates.ShowUploadFile
	if err := c.ShouldBind(&data); err != nil {
		//获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator.ValidationErrors类型的错误直接返回
			response.ErrorResponse(c, response.ValidateCheckError, err.Error())
			return
		}
		//validator.ValidationErrors类型的错误进行翻译
		response.ErrorResponse(c, response.ValidateCheckError, rules.Translate(errs))
		return
	}
	url := "." + data.Url
	c.File(url)
}

func FileTest(c *gin.Context) {
	file, err := os.Open("./upload/test.txt")
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	defer file.Close()

	var content []byte
	var tmp = make([]byte, 20)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			break
		}
		if err != nil {
			response.ErrorResponse(c, response.ValidateCheckError, err.Error())
			return
		}
		content = append(content, tmp[:n]...)
	}

	c.Writer.Write(content)
}

func FileTestBufio(c *gin.Context) {
	file, err := os.Open("./upload/test.txt")
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	defer file.Close()

	var content string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				content += line
			}
			fmt.Println("文件读完了")
			break
		}

		if err != nil {
			response.ErrorResponse(c, response.ValidateCheckError, err.Error())
			return
		}
		content += line
	}

	c.Writer.WriteString(content)
}

func FileTestIoutil(c *gin.Context) {
	content, err := ioutil.ReadFile("./upload/test.txt")
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	c.Writer.Write(content)
}

func FileTestWrite(c *gin.Context) {
	content := c.Query("content")
	file, err := os.OpenFile("./upload/tmp.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	defer file.Close()

	//file.Write([]byte(content))
	//file.WriteString("哈哈哈，就是要学go，不服你也来学习啊")

	write := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		write.WriteString("哈哈哈，就是要学go，不服你也来学习啊\n")
	}
	write.Flush()

	err = ioutil.WriteFile("./upload/tmp.txt", []byte(content), 0664)
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
}

func GetSinPng(c *gin.Context) {
	const size = 300
	//设置图片尺寸
	pic := image.NewGray(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//上色
			pic.SetGray(x, y, color.Gray{255})
		}
		pic.SetGray(x, size/2, color.Gray{0})
	}
	for x := 0; x < size; x++ {
		//计算定义域
		s := (float64(x) / size) * (2 * math.Pi)

		//计算值域，先缩小长度，然后下移
		y := -math.Sin(s)*size/2 + size/2
		//将图形填充到像素中
		pic.SetGray(x, int(y), color.Gray{0})
	}

	file, err := os.Create("./upload/sin.png")
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	defer file.Close()
	err = png.Encode(file, pic)
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	//response.SuccessResponse(c, "成功", "/upload/sin.png")
	c.File("./upload/sin.png")
}

func GetCosPng(c *gin.Context) {
	const (
		xSize = 1000
		ySize = 500
	)

	pic := image.NewGray(image.Rect(0, 0, xSize, ySize))
	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
		pic.SetGray(x, ySize/2, color.Gray{0})
	}

	for x := 0; x < xSize; x++ {
		s := (float64(x) / xSize) * (8 * math.Pi)
		y := -math.Cos(s)*ySize/2 + ySize/2
		pic.SetGray(x, int(y), color.Gray{0})
	}

	file, err := os.Create("./upload/cos.png")
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	err = png.Encode(file, pic)
	if err != nil {
		response.ErrorResponse(c, response.ValidateCheckError, err.Error())
		return
	}
	c.File("./upload/cos.png")
}
