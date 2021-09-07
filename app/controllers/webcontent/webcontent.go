package webcontent

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"gostu/pkg/config"
	"gostu/pkg/logger"
	"log"
	"net/http"
	"time"
)

func GetWebContent(ctx *gin.Context)  {
	curURL := "https://time.geekbang.org/column/article/357322";

	//启动 chrome 及简单配置
	var opts []selenium.ServiceOption
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imgCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imgCaps,
		Path:  "",
		Args: []string{
			"--headless",
			"--start-maximized",
			//"--window-size=1200x600",
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
			"--disable-gpu",
			"--disable-impl-side-painting",
			"--disable-gpu-sandbox",
			"--disable-accelerated-2d-canvas",
			"--disable-accelerated-jpeg-decoding",
			"--test-type=ui",
		},
	}
	caps.AddChrome(chromeCaps)

	// 启动 chromedriver server
	logger.Logger.Infoln(config.AppConfig.GetInt("chrome.port"))
	service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", config.AppConfig.GetInt("chrome.port"), opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
		return
	}
	defer service.Stop()

	// 打开 chrome 浏览器
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", config.AppConfig.GetInt("chrome.port")))
	if err != nil {
		log.Println(err)
		return
	}
	defer wd.Quit()

	//加载URL
	err = wd.Get(curURL)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	//判断加载完成
	jsRt, err := wd.ExecuteScript("return document.readyState", nil)
	if err != nil {
		log.Println("exe js err", err)
	}
	fmt.Println("jsRt", jsRt)
	if jsRt != "complete" {
		log.Println("网页加载未完成")
		return
	}

	//获取网站内容
	var frameHtml string

	time.Sleep(1 * time.Second)
	frameHtml, err = wd.PageSource()
	if err != nil {
		log.Println(err)
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "获取成功",
		"data":    frameHtml,
	})

	////解析 html 文件
	//var doc *goquery.Document
	//doc, err = goquery.NewDocumentFromReader(bytes.NewReader([]byte(frameHtml)))
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//logger.Logger.Infoln("文档内容", frameHtml)
	//doc.Find("li.s-result-item").Each(func(liIndex int, liItem *goquery.Selection) {
	//	// do something
	//})
}



