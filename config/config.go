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
		"aeW4bHWkaGIMl6N": map[string]interface{}{
			"ssp_name": "Clickadilla_In-Page",
			"ssp_id": "111",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"qjHLq1RZXU3tk8j": map[string]interface{}{
			"ssp_name": "Clickadilla_In-Page",
			"ssp_id": "112",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"wKyozxnWWgdHECi": map[string]interface{}{
			"ssp_name": "Clickadilla_In-Page",
			"ssp_id": "113",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"STx9xtDhW0825P8": map[string]interface{}{
			"ssp_name": "Clickadilla_In-Page",
			"ssp_id": "114",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"lz5ouNvlSnlZOIf": map[string]interface{}{
			"ssp_name": "DaoPush",
			"ssp_id": "115",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "102",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"UvEH4RJTsa2oOpx": map[string]interface{}{
			"ssp_name": "DaoPush_Classic",
			"ssp_id": "116",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "102",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"NaENhvugH5WdQ1y": map[string]interface{}{
			"ssp_name": "DaoPush_In-Page",
			"ssp_id": "117",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"ngGnozEJ3syxUEp": map[string]interface{}{
			"ssp_name": "DaoPush_In-Page",
			"ssp_id": "118",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"wQdvtTF2b6fs2p2": map[string]interface{}{
			"ssp_name": "RexRTB",
			"ssp_id": "103",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "102",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"PJtnhleyzNYAdGk": map[string]interface{}{
			"ssp_name": "RiverTraffic_Classic",
			"ssp_id": "104",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "102",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"z8GBO0tCQwTw2bt": map[string]interface{}{
			"ssp_name": "RiverTraffic_In-Page",
			"ssp_id": "105",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"wlFmvmEl5spoPxk": map[string]interface{}{
			"ssp_name": "RiverTraffic_In-Page",
			"ssp_id": "106",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"FqYgmeGAJG9nJz6": map[string]interface{}{
			"ssp_name": "Clickadilla_In-Page",
			"ssp_id": "109",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"vv8VZtulWS5KvvC": map[string]interface{}{
			"ssp_name": "Clickadilla_In-Page",
			"ssp_id": "110",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "103",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"s9u1H1aJuJgYWJc": map[string]interface{}{
			"ssp_name": "Admaven_Adult_push",
			"ssp_id": "119",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "102",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"O29B7nMFGE6fFNv": map[string]interface{}{
			"ssp_name": "Admaven_MNSTRM_push",
			"ssp_id": "120",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "102",
					"profit": 0.00,
					"source_id_blacklist": []string{},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"uiUIGofvdGxzmA4": map[string]interface{}{
			"ssp_name": "Clickadilla_Classic",
			"ssp_id": "107",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "102",
					"profit": 0.00,
					"source_id_blacklist": []string{"102108305275"},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
		"2S46Snsmc1IaISv": map[string]interface{}{
			"ssp_name": "Clickadilla_Classic",
			"ssp_id": "108",
			"dsp": []map[string]interface{}{
				{
					"dsp_id": "102",
					"profit": 0.00,
					"source_id_blacklist": []string{"102108305275"},
					"country_blacklist": []string{},
					"country_whitelist": []string{},
				},
			},
		},
	},

}
