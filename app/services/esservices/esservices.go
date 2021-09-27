package esservices

const Mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 1,
		"max_result_window": 1000000
	},
	"mappings":{
		"wx_user":{
			"properties":{
				"name":{
					"type":"text"
				},
				"age":{
					"type":"long"
				},
				"married":{
					"type":"boolean"
				},
				"created":{
					"type":"date",
					"format": "yyyy-MM-dd HH:mm:ss"
				},
				"tags":{
					"type":"keyword"
				}
			}
		}
	}
}`

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Married bool `json:"married"`
	Created string `json:"created"`
	Tags string `json:"tags"`
}



