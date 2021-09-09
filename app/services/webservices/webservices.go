package webservices

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"gostu/pkg/config"
	"log"
	"time"
)

//func StartChromeDriverService() {
//
//}

func StartChrome() (cookieStr string) {
	var (
		driverPath = "/usr/bin/chromedriver" //chrome驱动在服务器上的位置
		port       = config.AppConfig.GetInt("chrome.port") //端口号
		opts 	   = []selenium.ServiceOption{}
	)
	service, err := selenium.NewChromeDriverService(driverPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service failed", err.Error())
		return
	}

	defer func() {
		_ = service.Stop()
	}()

	//链接本地的浏览器 chrome
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
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)

	//调起一个chrome浏览器
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver failed", err.Error())
		return
	}
	defer func() {
		_ = wd.Quit()
	}()

	//极客时间登录链接
	err = wd.Get("https://account.geekbang.org/login?redirect=https://time.geekbang.org/")
	if err != nil {
		fmt.Println("get page failed", err.Error())
		return
	}

	//手机号输入框
	err = wd.Wait(Displayed(selenium.ByName, "cellphone"))
	if err != nil {
		fmt.Println("手机号输入", err)
		return
	}
	//密码输入框
	err = wd.Wait(Displayed(selenium.ByName, "password"))
	if err != nil {
		fmt.Println("密码", err)
		return
	}
	//协议勾选框
	err = wd.Wait(Displayed(selenium.ByCSSSelector, ".ThirdParty_agree-checkbox_2pT55"))
	if err != nil {
		fmt.Println("协议勾选框", err)
		return
	}
	//登录按钮
	err = wd.Wait(Displayed(selenium.ByClassName, "Button_button_3onsJ"))
	if err != nil {
		fmt.Println("登录按钮", err)
		return
	}

	username, err := wd.FindElement(selenium.ByName, "cellphone")
	if err != nil {
		fmt.Println("get username failed", err.Error())
		return
	}
	password, err := wd.FindElement(selenium.ByName, "password")
	if err != nil {
		fmt.Println("get password failed", err.Error())
		return
	}
	protocol, err := wd.FindElement(selenium.ByCSSSelector, "#agree")
	if err != nil {
		fmt.Println("get protocol failed", err.Error())
		return
	}
	submit, err := wd.FindElement(selenium.ByClassName, "Button_button_3onsJ")
	if err != nil {
		fmt.Println("submit failed", err.Error())
		return
	}
	err = username.SendKeys("12122388800")
	if err != nil {
		fmt.Println("设置手机号", err)
		return
	}
	err = password.SendKeys("67c04b90")
	if err != nil {
		fmt.Println("设置密码", err)
		return
	}
	err = protocol.SendKeys("checked")
	if err != nil {
		fmt.Println("设置协议", err)
		return
	}
	err = submit.Click()
	if err != nil {
		fmt.Println("登录按钮提交", err)
		return
	}

	time.Sleep(5* time.Second)


	//被重定向后我们就重新写入要进入的URL
	//加载URL
	err = wd.Get("https://time.geekbang.org/")
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
	return frameHtml



	//err = wd.Wait(func(wdTemp selenium.WebDriver) (b bool, e error) {
	//	tit, err := wdTemp.Title()
	//	logger.Logger.Infoln("标题信息", tit)
	//	if err != nil {
	//		return false, nil
	//	}
	//	if tit != "极客时间-轻松学习，高效学习-极客邦" {
	//		return false, nil
	//	}
	//	return true, nil
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	////获取cookie
	//cookieLst, err := wd.GetCookies()
	//if err != nil {
	//	fmt.Println("获取cookie", err)
	//	return
	//}
	//var cookieArr []string
	//for _, c := range cookieLst {
	//	cookieArr = append(cookieArr, c.Name+"="+c.Value)
	//}
	//cookieStr = strings.Join(cookieArr, "; ")
	return
}

func Displayed(by, elementName string) func(selenium.WebDriver) (bool, error) {
	return func(wd selenium.WebDriver) (ok bool, err error) {
		var el selenium.WebElement
		el, err = wd.FindElement(by, elementName)
		if err != nil {
			return
		}
		ok, err = el.IsDisplayed()
		return
	}
}


