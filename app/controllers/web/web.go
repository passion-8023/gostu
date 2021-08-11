package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gostu/pkg/logger"
	"gostu/pkg/response"
	"io/ioutil"
	"net/http"
	"strings"
)

//const cookie = "LF_ID=1626694902749-9551268-770204; GCID=de3deb9-120f8ef-4d06adc-1d6be51; GRID=de3deb9-120f8ef-4d06adc-1d6be51; _ga=GA1.2.809667571.1627012879; _gcl_au=1.1.558334945.1627012881; gksskpitn=c884f9f5-7364-4b95-8b74-2613b3792139; _gid=GA1.2.1663183633.1628648705; GCESS=BggBAwYEma6pxwUEAAAAAAQEAC8NAAIEa4ITYQcE6Sl6MAkBAQwBAQMEa4ITYQsCBgAKBAAAAAANAQEBCJ31HgAAAAAA; Hm_lvt_59c4ff31a9ee6263811b23eb921a5083=1628664667,1628665244,1628666256,1628668501; Hm_lvt_022f847c4e3acd44d4a2481d9187f1e6=1628665244,1628666256,1628668489,1628668501; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%222028957%22%2C%22first_id%22%3A%2217b3434ed17cc8-01566b11789309-434a0b2e-2073600-17b3434ed18cdd%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E5%BC%95%E8%8D%90%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC%22%2C%22%24latest_referrer%22%3A%22https%3A%2F%2Faccount.infoq.cn%2F%22%2C%22%24latest_landing_page%22%3A%22https%3A%2F%2Ftime.geekbang.org%2F%3Fm%3D0%26d%3D3%26c%3D11%26sort%3D0%26order%3Dsort%22%2C%22%24latest_utm_source%22%3A%22geektimeWeb%22%2C%22%24latest_utm_medium%22%3A%22menu%22%2C%22%24latest_utm_campaign%22%3A%22newregister%22%2C%22%24latest_utm_term%22%3A%22pc_interstitial_1206%22%7D%2C%22%24device_id%22%3A%2217abe92632f1de-0cfb29a316e87f-434a0b2e-2073600-17abe9263308f1%22%7D; ERID=bf3bf5e-6ddf76a-f64fa95-d46b9bd; SERVERID=1fa1f330efedec1559b3abbcb6e30f50|1628668563|1628662098; _gat=1; Hm_lpvt_59c4ff31a9ee6263811b23eb921a5083=1628668540; Hm_lpvt_022f847c4e3acd44d4a2481d9187f1e6=1628668540; gk_process_ev={%22count%22:37%2C%22utime%22:1628668494311%2C%22referrer%22:%22https://time.geekbang.org/?m=0&d=3&c=11&sort=0&order=sort%22%2C%22target%22:%22%22%2C%22referrerTarget%22:%22page_geektime_login%22}"
const cookie = ""


func GetWebList(ctx *gin.Context)  {
	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		"https://time.geekbang.org/serv/v1/column/articles",
		strings.NewReader("{\"cid\":112,\"size\":100,\"prev\":0,\"order\":\"earliest\",\"sample\":false,\"chapter_ids\":[\"578\",\"579\",\"580\",\"581\",\"582\"]}"),
	)
	if err != nil {
		logger.Logger.Errorln("http read failed", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Set("Host", "time.geekbang.org")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "60")
	req.Header.Set("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Origin", "https://time.geekbang.org")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://time.geekbang.org/column/article/13159")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	logger.Logger.Infoln(resp, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Errorln("http read failed", err)
	}

	result := make(map[string]interface{})
	_ = json.Unmarshal(body, &result)
	response.SuccessResponse(ctx, "", result)
}

func GetWebContent(ctx *gin.Context)  {
	client := &http.Client{}

	req, err := http.NewRequest(
		"POST",
		"https://time.geekbang.org/serv/v1/article",
		strings.NewReader("{\"id\":13176,\"include_neighbors\":true,\"is_freelyread\":true}"),
	)
	if err != nil {
		logger.Logger.Errorln("http read failed", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Set("Host", "time.geekbang.org")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "60")
	req.Header.Set("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Origin", "https://time.geekbang.org")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://time.geekbang.org/column/article/13176")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	resp, err := client.Do(req)
	logger.Logger.Infoln(resp, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Errorln("http read failed", err)
	}

	result := make(map[string]interface{})
	_ = json.Unmarshal(body, &result)
	response.SuccessResponse(ctx, "", result)
}

//
//func GetWebContent(ctx *gin.Context)  {
//	client := &http.Client{}
//
//	var Body io.Reader
//	req, err := http.NewRequest(
//		"GET",
//		"https://time.geekbang.org/column/article/12655",
//		Body,
//	)
//	if err != nil {
//		logger.Logger.Errorln("http read failed", err)
//	}
//
//	//req.Header.Set("Content-Type", "application/json")
//	//req.Header.Set("Cookie", cookie)
//	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
//	//req.Header.Set("Host", "time.geekbang.org")
//	//req.Header.Set("Connection", "keep-alive")
//	//req.Header.Set("Content-Length", "60")
//	//req.Header.Set("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
//	//req.Header.Set("Accept", "application/json, text/plain, */*")
//	//req.Header.Set("sec-ch-ua-mobile", "?0")
//	//req.Header.Set("Origin", "https://time.geekbang.org")
//	//req.Header.Set("Sec-Fetch-Site", "same-origin")
//	//req.Header.Set("Sec-Fetch-Mode", "cors")
//	//req.Header.Set("Sec-Fetch-Dest", "empty")
//	//req.Header.Set("Referer", "https://time.geekbang.org/column/article/12655")
//	//req.Header.Set("Accept-Encoding", "gzip, deflate, br")
//	//req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
//
//	resp, err := client.Do(req)
//	logger.Logger.Infoln(resp, err)
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		logger.Logger.Errorln("http read failed", err)
//	}
//	//
//	//result := make(map[string]interface{})
//	//_ = json.Unmarshal(body, &result)
//	response.SuccessResponse(ctx, "", string(body))
//}
