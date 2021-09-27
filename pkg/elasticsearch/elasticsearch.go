package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v6"
	"gostu/pkg/config"
	"log"
)

var Client *elastic.Client

func init() {
	var err error
	esUrl := fmt.Sprintf("%s://%s:%s", config.AppConfig.GetString("es.scheme"), config.AppConfig.GetString("es.host"), config.AppConfig.GetString("es.port"))
	fmt.Println(esUrl)
	//username := config.AppConfig.GetString("es.user")
	//password := config.AppConfig.GetString("es.passd")
	//连接客户端
	Client, err = elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	if err != nil {
		log.Fatal("连接客户端",err)
		return
	}
	info, code, err := Client.Ping(esUrl).Do(context.Background())
	if err != nil {
		log.Fatal("客户端Ping", err)
		return
	}
	fmt.Printf("Elasticsearch returned with code>: %d and version %s\n", code, info.Version.Number)

	version, err := Client.ElasticsearchVersion(esUrl)
	if err != nil {
		log.Fatal("版本号", err)
		return
	}
	fmt.Printf("es的版本为%s\n", version)
}


