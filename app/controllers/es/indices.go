package es

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v6"
	"gostu/app/services/esservices"
	"gostu/pkg/config"
	"gostu/pkg/elasticsearch"
	"gostu/pkg/response"
	"net/http"
	"strconv"
	"time"
)

func CreateIndex(ctx *gin.Context)  {
	esIndex := config.AppConfig.GetString("es.test_index_one")
	esType := config.AppConfig.GetString("es.test_type_one")
	//第一种方案：使用结构体
	//employee := esservices.Employee{
	//	FirstName: "王",
	//	LastName:  "鑫",
	//	Age:       26,
	//	About:     "一个PHP开发工程师，集帅气和才华与一身的阳光大男孩",
	//	Interests: []string{"羽毛球", "听歌"},
	//}
	//res, err := elasticsearch.Client.Index().
	//	Index(esIndex).
	//	Type(esType).
	//	Id("1").
	//	BodyJson(employee).
	//	Do(context.Background())
	//if err != nil {
	//	response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	//response.SuccessResponse(ctx, "创建成功", res)
	//return

	//第二种方案：使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	//e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	res, err := elasticsearch.Client.Index().
		Index(esIndex).
		Type(esType).
		Id("3").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(ctx, "创建成功", res)
	return
}

func CheckIndex(ctx *gin.Context)  {
	var esIndex string =  config.AppConfig.GetString("es.test_index_two")
	exists, err := elasticsearch.Client.IndexExists(esIndex).Do(context.Background())
	if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
	}
	if exists {
		info, err := elasticsearch.Client.IndexGet(esIndex).Do(context.Background())
		if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		response.SuccessResponse(ctx, "存在", info)
		return
	} else {
		result, err := elasticsearch.Client.CreateIndex(esIndex).BodyString(esservices.Mapping).Do(context.Background())
		if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		if !result.Acknowledged {
			response.ErrorResponse(ctx, http.StatusBadRequest, "创建失败")
			return
		}
		response.SuccessResponse(ctx, "创建成功", result)
		return
	}
}

func DeleteIndex(ctx *gin.Context)  {
	esIndex := ctx.DefaultQuery("index", config.AppConfig.GetString("es.test_index_two"))
	do, err := elasticsearch.Client.DeleteIndex(esIndex).Do(context.Background())
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(ctx, "删除成功", do)
	return
}

func BulkData(ctx *gin.Context)  {
	esIndex :=  config.AppConfig.GetString("es.test_index_two")
	users := []esservices.User{
		esservices.User{
			Name:    "马云",
			Age:     57,
			Married: true,
			Created: time.Now().Format("2006-01-02 15:04:05"),
			Tags:    "阿里巴巴创始人",
		},
		esservices.User{
			Name:    "马化腾",
			Age:     50,
			Married: true,
			Created: time.Now().Format("2006-01-02 15:04:05"),
			Tags:    "腾讯集团创始人",
		},
		esservices.User{
			Name:    "李彦宏",
			Age:     40,
			Married: true,
			Created: time.Now().Format("2006-01-02 15:04:05"),
			Tags:    "百度创始人",
		},
	}

	bulkRequest := elasticsearch.Client.Bulk()
	for _, user := range users {
		doc := elastic.NewBulkIndexRequest().Index(esIndex).Type(esIndex).Doc(user)
		bulkRequest = bulkRequest.Add(doc)
	}
	res, err := bulkRequest.Do(context.Background())
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	failed := res.Failed()
	l := len(failed)
	if l > 0 {
		response.ErrorResponse(ctx, http.StatusBadRequest, "添加失败" + strconv.Itoa(l))
		return
	}
	response.SuccessResponse(ctx, "添加成功", nil)
	return
}

