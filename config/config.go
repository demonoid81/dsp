package config

var Config = map[string]interface{}{
	"Crypto":           "5368316e676520742879732370617371",
	"Click_Url":        "https://api.labyrinthads.com",
	"Media_Url":        "http://i.labyrinthads.com",
	"Url_Redirect":     "https://labyrinthads.com",
	"host":             "127.0.0.1",
	"port":             5432,
	"user":             "postgres",
	"password":         "9mcP7ZxEuMNcFU",
	"dbname":           "dsp",
	"revshare":         0.00,
	"mongo_database":   "dsp",
	"mongo_collection": "requests",
	"mongo_url":        "mongodb://dspadmin:2XqUnHQbjGhw2pM68TkgGKe@localhost:27017",
	"Kafka": map[string]interface{}{
		"subscribe": map[string]interface{}{
			"kafkaURL": "localhost:9092",
			"topic":    "subscribeTopic",
		},
		"click": map[string]interface{}{
			"kafkaURL": "localhost:9092",
			"topic":    "clickTopic",
		},
		"clickdsp": map[string]interface{}{
			"kafkaURL": "localhost:9092",
			"topic":    "clickdspTopic",
		},
	},
	"SSP": map[string]interface{}{
		"YxXs1HSVgDRRC9T": map[string]interface{}{
			"ssp_name": "Mgid",
			"ssp_id":   "102",
			"dsp": []map[string]interface{}{
				{
					"dsp_id":              "102",
					"profit":              0.00,
					"source_id_blacklist": []string{},
					"country_blacklist":   []string{},
					"country_whitelist":   []string{},
				},
			},
		},
		"xBXhzjjhhYDFU7r": map[string]interface{}{
			"ssp_name": "test",
			"ssp_id":   "101",
			"dsp": []map[string]interface{}{
				{
					"dsp_id":              "102",
					"profit":              0.00,
					"source_id_blacklist": []string{"102101back_block_lp_2087", "102101ramos_w10_1304_us_edge_new", "102101lux_w10_0303-dagger-2033", "102101ramos_w10_0605_us_edge", "102101terame_w10_1609_edge_486231", "102101ramos_w10_1608_topiced", "102101ramos_w10_0907_us_edge", "102101ceca_w10_1608_kasa", "102101yue_w10_1208_manunited"},
					"country_blacklist":   []string{},
					"country_whitelist":   []string{},
				},
			},
		},
	},
}
