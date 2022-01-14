package postgres

//import (
//	"database/sql"
//	"encoding/json"
//	"fmt"
//	_ "github.com/lib/pq"
//	
//)
//
//type Campaign struct {
//	UserID               int     `json:"user_id"`
//	CompanyURL           string  `json:"company_url"`
//	CompanyPrice         float64 `json:"company_price"`
//	CompanyID            int     `json:"company_id"`
//	CompanyCountry       string  `json:"company_country"`
//	AdTitle              string  `json:"ad_title"`
//	AdText               string  `json:"ad_text"`
//	AdIcon               string  `json:"ad_icon"`
//	AdImage              string  `json:"ad_image"`
//	CompanyBlackList     string  `json:"blacklist"`
//	CompanyWhiteList     string  `json:"whitelist"`
//	CompanyBlackListFeed string  `json:"blacklist_feed"`
//	CompanyWhiteListFeed string  `json:"whitelist_feed"`
//	Freshness            string  `json:"freshness"`
//}
//
//var Campaigns []Campaign
//
//func E(country string, platform string, browser string, category string, push_type string) string {
//
//	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//		config.Config["host"].(string),
//		config.Config["port"].(int),
//		config.Config["user"].(string),
//		config.Config["password"].(string),
//		config.Config["dbname"].(string))
//	db, err := sql.Open("postgres", psqlconn)
//
//	if err != nil {
//		return "error"
//	}
//
//	defer db.Close()
//
//	var query = `SELECT
//                    comapny->>'user_id' user_id,
//                    comapny->>'url' company_url,
//                    ((comapny->>'countries')::jsonb->>'%s')::numeric company_price,
//                    comapny->>'id' company_id,
//                    '%s' company_country,
//                    comapny->>'ad_title' ad_title,
//                    comapny->>'ad_text' ad_text,
//                    comapny->>'ad_icon' ad_icon,
//                    comapny->>'ad_image' ad_image,
//                    comapny->>'blacklist' blacklist,
//                    comapny->>'whitelist' whitelist,
//                    comapny->>'blacklist_feed' blacklist_feed,
//                    comapny->>'whitelist_feed' whitelist_feed,
//                    comapny->>'freshness' freshness
//                FROM campaigns
//
//                CROSS JOIN LATERAL jsonb_array_elements(campaigns.payload) comapny
//
//                WHERE ((comapny->>'countries')::jsonb->>'%s')::numeric > 0
//                AND ((comapny->>'type')::jsonb->>'%s')::boolean = true
//                AND (jsonb_array_length(comapny->'os') = 0 OR (comapny->>'os')::jsonb @> '"%s"')
//                AND (jsonb_array_length(comapny->'browser') = 0 OR (comapny->>'browser')::jsonb @> '"%s"')
//
//                ORDER BY ((comapny->>'countries')::jsonb->>'%s')::numeric DESC
//                LIMIT 250`
//
//	//AND (comapny->>'category')::integer IN (%s)
//
//	query = fmt.Sprintf(query, country, country, country, push_type, platform, browser, country)
//	//query = fmt.Sprintf(query, country, country, country, platform, browser, category, country)
//
//	rows, err := db.Query(query)
//
//	if err != nil {
//		return "error"
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		c := Campaign{}
//		rows.Scan(&c.UserID,
//			&c.CompanyURL,
//			&c.CompanyPrice,
//			&c.CompanyID,
//			&c.CompanyCountry,
//			&c.AdTitle,
//			&c.AdText,
//			&c.AdIcon,
//			&c.AdImage,
//			&c.CompanyBlackList,
//			&c.CompanyWhiteList,
//			&c.CompanyBlackListFeed,
//			&c.CompanyWhiteListFeed,
//			&c.Freshness)
//		Campaigns = append(Campaigns, c)
//	}
//
//	json, _ := json.Marshal(Campaigns)
//
//	Campaigns = nil
//
//	return string(json)
//}
