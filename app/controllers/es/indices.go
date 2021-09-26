package es

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateIndex(ctx *gin.Context)  {
	//exists, err := elasticsearch.Client.IndexExists("wx_user").Do(context.Background())
	//if err != nil {
	//	response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	//	return
	//}
	//if !exists {
	//	createResult, err := elasticsearch.Client.CreateIndex("wx_user").BodyString(esservices.Mapping).Do(context.Background())
	//	if err != nil {
	//		response.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
	//		return
	//	}
	//	if !createResult.Acknowledged {
	//		response.ErrorResponse(ctx, http.StatusBadRequest, "创建失败")
	//		return
	//	}
	//}


	data := []byte("今天我来讲哈希算法")
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(cipherStr)
	//fmt.Printf("%x\n", md5.Sum(data))
	fmt.Printf("%x\n", cipherStr)
	//fmt.Println(hex.EncodeToString(cipherStr))
}
