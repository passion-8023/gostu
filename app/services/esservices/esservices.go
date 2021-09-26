package esservices

const Mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 1,
		"max_result_window": 1000000
	},
	"mappings":{
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
				"type":"date"
			},
			"tags":{
				"type":"keyword"
			},
			"location":{
				"type":"geo_point"
			},
			"suggest_field":{
				"type":"completion"
			}
		}
	}
}`


