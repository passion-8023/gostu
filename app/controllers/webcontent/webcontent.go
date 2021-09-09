package webcontent

import (
	"github.com/gin-gonic/gin"
	"gostu/app/services/webservices"
	"net/http"
)

func GetWebContent(ctx *gin.Context)  {

	//webservices.StartChromeDriverService()
	cookieStr := webservices.StartChrome()
	ctx.JSON(http.StatusOK, gin.H{
		"cookie": cookieStr,
	})
	//curURL := "https://time.geekbang.org/column/article/357322";
	//
	////启动 chrome 及简单配置
	//var opts []selenium.ServiceOption
	//caps := selenium.Capabilities{
	//	"browserName": "chrome",
	//}
	//
	//// 禁止加载图片，加快渲染速度
	//imgCaps := map[string]interface{}{
	//	"profile.managed_default_content_settings.images": 2,
	//}
	//
	//chromeCaps := chrome.Capabilities{
	//	Prefs: imgCaps,
	//	Path:  "",
	//	Args: []string{
	//		"--headless",
	//		"--start-maximized",
	//		//"--window-size=1200x600",
	//		"--no-sandbox",
	//		"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	//		"--disable-gpu",
	//		"--disable-impl-side-painting",
	//		"--disable-gpu-sandbox",
	//		"--disable-accelerated-2d-canvas",
	//		"--disable-accelerated-jpeg-decoding",
	//		"--test-type=ui",
	//	},
	//}
	//caps.AddChrome(chromeCaps)
	//
	//// 启动 chromedriver server
	//logger.Logger.Infoln(config.AppConfig.GetInt("chrome.port"))
	//service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", config.AppConfig.GetInt("chrome.port"), opts...)
	//if err != nil {
	//	log.Printf("Error starting the ChromeDriver server: %v", err)
	//	return
	//}
	//defer service.Stop()
	//
	//// 打开 chrome 浏览器
	//wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", config.AppConfig.GetInt("chrome.port")))
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//defer wd.Quit()
	//
	////加载URL
	//err = wd.Get(curURL)
	//if err != nil {
	//	log.Println(fmt.Sprintf("Failed to load page: %s\n", err))
	//}
	//
	////判断加载完成
	//jsRt, err := wd.ExecuteScript("return document.readyState", nil)
	//if err != nil {
	//	log.Println("exe js err", err)
	//}
	//fmt.Println("jsRt", jsRt)
	//if jsRt != "complete" {
	//	log.Println("网页加载未完成")
	//	return
	//}
	//
	////获取网站内容
	//var frameHtml string
	//
	//time.Sleep(1 * time.Second)
	//frameHtml, err = wd.PageSource()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  200,
	//	"message": "获取成功",
	//	"data":    frameHtml,
	//})

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


//func getCookieStr(usernameText string, passwordText string) (cookieStr string) {
//	var (
//		driverPath = "/usr/bin/chromedriver"
//		port       = config.AppConfig.GetInt("chrome.port")
//	)
//
//	service, err := selenium.NewChromeDriverService(driverPath, port, []selenium.ServiceOption{}...)
//	if nil != err {
//		fmt.Println("start a chromedriver service failed", err.Error())
//		return
//	}
//	defer func() {
//		_ = service.Stop()
//	}()
//
//	// Connect to the WebDriver instance running locally.
//	caps := selenium.Capabilities{"browserName": "chrome"}
//	//禁止图片加载，加快渲染速度
//	imagCaps := map[string]interface{}{
//		"profile.managed_default_content_settings.images": 2,
//	}
//	chromeCaps := chrome.Capabilities{
//		Prefs: imagCaps,
//		Path:  "",
//		Args: []string{
//			"--headless",
//			"--no-sandbox",
//			"--disable-gpu-sandbox",
//			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36",
//		},
//	}
//	//以上是设置浏览器参数
//	caps.AddChrome(chromeCaps)
//	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
//	if err != nil {
//		fmt.Println("connect to the webDriver failed", err.Error())
//		return
//	}
//	defer func() {
//		_ = wd.Quit()
//	}()
//
//
//	err = wd.Get("https://passport.weibo.cn/signin/login?entry=mweibo&r=https://weibo.cn/")
//	if err != nil {
//		fmt.Println("get page failed", err.Error())
//		return
//	}
//	err = wd.Wait(Displayed(selenium.ByCSSSelector, "#loginName"))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	err = wd.Wait(Displayed(selenium.ByCSSSelector, "#loginPassword"))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	err = wd.Wait(Displayed(selenium.ByCSSSelector, "#loginAction"))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	username, err := wd.FindElement(selenium.ByCSSSelector, "#loginName")
//	if err != nil {
//		fmt.Println("get username failed", err.Error())
//		return
//	}
//	password, err := wd.FindElement(selenium.ByCSSSelector, "#loginPassword")
//	if err != nil {
//		fmt.Println("get username failed", err.Error())
//		return
//	}
//	submit, err := wd.FindElement(selenium.ByCSSSelector, "#loginAction")
//	if err != nil {
//		fmt.Println("get username failed", err.Error())
//		return
//	}
//	err = username.SendKeys(usernameText)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	err = password.SendKeys(passwordText)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	err = submit.Click()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	err = wd.Wait(func(wdTemp selenium.WebDriver) (b bool, e error) {
//		tit, err := wdTemp.Title()
//		if err != nil {
//			return false, nil
//		}
//		if tit != "我的首页" {
//			return false, nil
//		}
//		return true, nil
//	})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	cookieLst, err := wd.GetCookies()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	var cookieArr []string
//	for _, c := range cookieLst {
//		cookieArr = append(cookieArr, c.Name+"="+c.Value)
//	}
//	cookieStr = strings.Join(cookieArr, "; ")
//	return
//}
//
//
//func Displayed(by, elementName string) func(selenium.WebDriver) (bool, error) {
//	return func(wd selenium.WebDriver) (ok bool, err error) {
//		var el selenium.WebElement
//		el, err = wd.FindElement(by, elementName)
//		if err != nil {
//			return
//		}
//		ok, err = el.IsDisplayed()
//		return
//	}
//}


