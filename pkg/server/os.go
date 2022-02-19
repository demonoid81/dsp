package server

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

//func (s *Server) addCountries(ctx context.Context) {
//	r := strings.NewReader("[{\"label\":\"Afghanistan\",\"value\":\"AF\"},{\"label\":\"Åland Islands\",\"value\":\"AX\"},{\"label\":\"Albania\",\"value\":\"AL\"},{\"label\":\"Algeria\",\"value\":\"DZ\"},{\"label\":\"American Samoa\",\"value\":\"AS\"},{\"label\":\"Andorra\",\"value\":\"AD\"},{\"label\":\"Angola\",\"value\":\"AO\"},{\"label\":\"Anguilla\",\"value\":\"AI\"},{\"label\":\"Antarctica\",\"value\":\"AQ\"},{\"label\":\"Antigua & Barbuda\",\"value\":\"AG\"},{\"label\":\"Argentina\",\"value\":\"AR\"},{\"label\":\"Armenia\",\"value\":\"AM\"},{\"label\":\"Aruba\",\"value\":\"AW\"},{\"label\":\"Ascension Island\",\"value\":\"AC\"},{\"label\":\"Australia\",\"value\":\"AU\"},{\"label\":\"Austria\",\"value\":\"AT\"},{\"label\":\"Azerbaijan\",\"value\":\"AZ\"},{\"label\":\"Bahamas\",\"value\":\"BS\"},{\"label\":\"Bahrain\",\"value\":\"BH\"},{\"label\":\"Bangladesh\",\"value\":\"BD\"},{\"label\":\"Barbados\",\"value\":\"BB\"},{\"label\":\"Belarus\",\"value\":\"BY\"},{\"label\":\"Belgium\",\"value\":\"BE\"},{\"label\":\"Belize\",\"value\":\"BZ\"},{\"label\":\"Benin\",\"value\":\"BJ\"},{\"label\":\"Bermuda\",\"value\":\"BM\"},{\"label\":\"Bhutan\",\"value\":\"BT\"},{\"label\":\"Bolivia\",\"value\":\"BO\"},{\"label\":\"Bosnia & Herzegovina\",\"value\":\"BA\"},{\"label\":\"Botswana\",\"value\":\"BW\"},{\"label\":\"Brazil\",\"value\":\"BR\"},{\"label\":\"British Indian Ocean Territory\",\"value\":\"IO\"},{\"label\":\"British Virgin Islands\",\"value\":\"VG\"},{\"label\":\"Brunei\",\"value\":\"BN\"},{\"label\":\"Bulgaria\",\"value\":\"BG\"},{\"label\":\"Burkina Faso\",\"value\":\"BF\"},{\"label\":\"Burundi\",\"value\":\"BI\"},{\"label\":\"Cambodia\",\"value\":\"KH\"},{\"label\":\"Cameroon\",\"value\":\"CM\"},{\"label\":\"Canada\",\"value\":\"CA\"},{\"label\":\"Canary Islands\",\"value\":\"IC\"},{\"label\":\"Cape Verde\",\"value\":\"CV\"},{\"label\":\"Caribbean Netherlands\",\"value\":\"BQ\"},{\"label\":\"Cayman Islands\",\"value\":\"KY\"},{\"label\":\"Central African Republic\",\"value\":\"CF\"},{\"label\":\"Ceuta & Melilla\",\"value\":\"EA\"},{\"label\":\"Chad\",\"value\":\"TD\"},{\"label\":\"Chile\",\"value\":\"CL\"},{\"label\":\"China\",\"value\":\"CN\"},{\"label\":\"Christmas Island\",\"value\":\"CX\"},{\"label\":\"Cocos (Keeling) Islands\",\"value\":\"CC\"},{\"label\":\"Colombia\",\"value\":\"CO\"},{\"label\":\"Comoros\",\"value\":\"KM\"},{\"label\":\"Congo - Brazzaville\",\"value\":\"CG\"},{\"label\":\"Congo - Kinshasa\",\"value\":\"CD\"},{\"label\":\"Cook Islands\",\"value\":\"CK\"},{\"label\":\"Costa Rica\",\"value\":\"CR\"},{\"label\":\"Côte d’Ivoire\",\"value\":\"CI\"},{\"label\":\"Croatia\",\"value\":\"HR\"},{\"label\":\"Cuba\",\"value\":\"CU\"},{\"label\":\"Curaçao\",\"value\":\"CW\"},{\"label\":\"Cyprus\",\"value\":\"CY\"},{\"label\":\"Czechia\",\"value\":\"CZ\"},{\"label\":\"Denmark\",\"value\":\"DK\"},{\"label\":\"Diego Garcia\",\"value\":\"DG\"},{\"label\":\"Djibouti\",\"value\":\"DJ\"},{\"label\":\"Dominica\",\"value\":\"DM\"},{\"label\":\"Dominican Republic\",\"value\":\"DO\"},{\"label\":\"Ecuador\",\"value\":\"EC\"},{\"label\":\"Egypt\",\"value\":\"EG\"},{\"label\":\"El Salvador\",\"value\":\"SV\"},{\"label\":\"Equatorial Guinea\",\"value\":\"GQ\"},{\"label\":\"Eritrea\",\"value\":\"ER\"},{\"label\":\"Estonia\",\"value\":\"EE\"},{\"label\":\"Eswatini\",\"value\":\"SZ\"},{\"label\":\"Ethiopia\",\"value\":\"ET\"},{\"label\":\"Falkland Islands\",\"value\":\"FK\"},{\"label\":\"Faroe Islands\",\"value\":\"FO\"},{\"label\":\"Fiji\",\"value\":\"FJ\"},{\"label\":\"Finland\",\"value\":\"FI\"},{\"label\":\"France\",\"value\":\"FR\"},{\"label\":\"French Guiana\",\"value\":\"GF\"},{\"label\":\"French Polynesia\",\"value\":\"PF\"},{\"label\":\"French Southern Territories\",\"value\":\"TF\"},{\"label\":\"Gabon\",\"value\":\"GA\"},{\"label\":\"Gambia\",\"value\":\"GM\"},{\"label\":\"Georgia\",\"value\":\"GE\"},{\"label\":\"Germany\",\"value\":\"DE\"},{\"label\":\"Ghana\",\"value\":\"GH\"},{\"label\":\"Gibraltar\",\"value\":\"GI\"},{\"label\":\"Greece\",\"value\":\"GR\"},{\"label\":\"Greenland\",\"value\":\"GL\"},{\"label\":\"Grenada\",\"value\":\"GD\"},{\"label\":\"Guadeloupe\",\"value\":\"GP\"},{\"label\":\"Guam\",\"value\":\"GU\"},{\"label\":\"Guatemala\",\"value\":\"GT\"},{\"label\":\"Guernsey\",\"value\":\"GG\"},{\"label\":\"Guinea\",\"value\":\"GN\"},{\"label\":\"Guinea-Bissau\",\"value\":\"GW\"},{\"label\":\"Guyana\",\"value\":\"GY\"},{\"label\":\"Haiti\",\"value\":\"HT\"},{\"label\":\"Honduras\",\"value\":\"HN\"},{\"label\":\"Hong Kong SAR China\",\"value\":\"HK\"},{\"label\":\"Hungary\",\"value\":\"HU\"},{\"label\":\"Iceland\",\"value\":\"IS\"},{\"label\":\"India\",\"value\":\"IN\"},{\"label\":\"Indonesia\",\"value\":\"ID\"},{\"label\":\"Iran\",\"value\":\"IR\"},{\"label\":\"Iraq\",\"value\":\"IQ\"},{\"label\":\"Ireland\",\"value\":\"IE\"},{\"label\":\"Isle of Man\",\"value\":\"IM\"},{\"label\":\"Israel\",\"value\":\"IL\"},{\"label\":\"Italy\",\"value\":\"IT\"},{\"label\":\"Jamaica\",\"value\":\"JM\"},{\"label\":\"Japan\",\"value\":\"JP\"},{\"label\":\"Jersey\",\"value\":\"JE\"},{\"label\":\"Jordan\",\"value\":\"JO\"},{\"label\":\"Kazakhstan\",\"value\":\"KZ\"},{\"label\":\"Kenya\",\"value\":\"KE\"},{\"label\":\"Kiribati\",\"value\":\"KI\"},{\"label\":\"Kosovo\",\"value\":\"XK\"},{\"label\":\"Kuwait\",\"value\":\"KW\"},{\"label\":\"Kyrgyzstan\",\"value\":\"KG\"},{\"label\":\"Laos\",\"value\":\"LA\"},{\"label\":\"Latvia\",\"value\":\"LV\"},{\"label\":\"Lebanon\",\"value\":\"LB\"},{\"label\":\"Lesotho\",\"value\":\"LS\"},{\"label\":\"Liberia\",\"value\":\"LR\"},{\"label\":\"Libya\",\"value\":\"LY\"},{\"label\":\"Liechtenstein\",\"value\":\"LI\"},{\"label\":\"Lithuania\",\"value\":\"LT\"},{\"label\":\"Luxembourg\",\"value\":\"LU\"},{\"label\":\"Macao SAR China\",\"value\":\"MO\"},{\"label\":\"Madagascar\",\"value\":\"MG\"},{\"label\":\"Malawi\",\"value\":\"MW\"},{\"label\":\"Malaysia\",\"value\":\"MY\"},{\"label\":\"Maldives\",\"value\":\"MV\"},{\"label\":\"Mali\",\"value\":\"ML\"},{\"label\":\"Malta\",\"value\":\"MT\"},{\"label\":\"Marshall Islands\",\"value\":\"MH\"},{\"label\":\"Martinique\",\"value\":\"MQ\"},{\"label\":\"Mauritania\",\"value\":\"MR\"},{\"label\":\"Mauritius\",\"value\":\"MU\"},{\"label\":\"Mayotte\",\"value\":\"YT\"},{\"label\":\"Mexico\",\"value\":\"MX\"},{\"label\":\"Micronesia\",\"value\":\"FM\"},{\"label\":\"Moldova\",\"value\":\"MD\"},{\"label\":\"Monaco\",\"value\":\"MC\"},{\"label\":\"Mongolia\",\"value\":\"MN\"},{\"label\":\"Montenegro\",\"value\":\"ME\"},{\"label\":\"Montserrat\",\"value\":\"MS\"},{\"label\":\"Morocco\",\"value\":\"MA\"},{\"label\":\"Mozambique\",\"value\":\"MZ\"},{\"label\":\"Myanmar (Burma)\",\"value\":\"MM\"},{\"label\":\"Namibia\",\"value\":\"NA\"},{\"label\":\"Nauru\",\"value\":\"NR\"},{\"label\":\"Nepal\",\"value\":\"NP\"},{\"label\":\"Netherlands\",\"value\":\"NL\"},{\"label\":\"New Caledonia\",\"value\":\"NC\"},{\"label\":\"New Zealand\",\"value\":\"NZ\"},{\"label\":\"Nicaragua\",\"value\":\"NI\"},{\"label\":\"Niger\",\"value\":\"NE\"},{\"label\":\"Nigeria\",\"value\":\"NG\"},{\"label\":\"Niue\",\"value\":\"NU\"},{\"label\":\"Norfolk Island\",\"value\":\"NF\"},{\"label\":\"North Korea\",\"value\":\"KP\"},{\"label\":\"North Macedonia\",\"value\":\"MK\"},{\"label\":\"Northern Mariana Islands\",\"value\":\"MP\"},{\"label\":\"Norway\",\"value\":\"NO\"},{\"label\":\"Oman\",\"value\":\"OM\"},{\"label\":\"Pakistan\",\"value\":\"PK\"},{\"label\":\"Palau\",\"value\":\"PW\"},{\"label\":\"Palestinian Territories\",\"value\":\"PS\"},{\"label\":\"Panama\",\"value\":\"PA\"},{\"label\":\"Papua New Guinea\",\"value\":\"PG\"},{\"label\":\"Paraguay\",\"value\":\"PY\"},{\"label\":\"Peru\",\"value\":\"PE\"},{\"label\":\"Philippines\",\"value\":\"PH\"},{\"label\":\"Pitcairn Islands\",\"value\":\"PN\"},{\"label\":\"Poland\",\"value\":\"PL\"},{\"label\":\"Portugal\",\"value\":\"PT\"},{\"label\":\"Pseudo-Accents\",\"value\":\"XA\"},{\"label\":\"Pseudo-Bidi\",\"value\":\"XB\"},{\"label\":\"Puerto Rico\",\"value\":\"PR\"},{\"label\":\"Qatar\",\"value\":\"QA\"},{\"label\":\"Réunion\",\"value\":\"RE\"},{\"label\":\"Romania\",\"value\":\"RO\"},{\"label\":\"Russia\",\"value\":\"RU\"},{\"label\":\"Rwanda\",\"value\":\"RW\"},{\"label\":\"Samoa\",\"value\":\"WS\"},{\"label\":\"San Marino\",\"value\":\"SM\"},{\"label\":\"São Tomé & Príncipe\",\"value\":\"ST\"},{\"label\":\"Saudi Arabia\",\"value\":\"SA\"},{\"label\":\"Senegal\",\"value\":\"SN\"},{\"label\":\"Serbia\",\"value\":\"RS\"},{\"label\":\"Seychelles\",\"value\":\"SC\"},{\"label\":\"Sierra Leone\",\"value\":\"SL\"},{\"label\":\"Singapore\",\"value\":\"SG\"},{\"label\":\"Sint Maarten\",\"value\":\"SX\"},{\"label\":\"Slovakia\",\"value\":\"SK\"},{\"label\":\"Slovenia\",\"value\":\"SI\"},{\"label\":\"Solomon Islands\",\"value\":\"SB\"},{\"label\":\"Somalia\",\"value\":\"SO\"},{\"label\":\"South Africa\",\"value\":\"ZA\"},{\"label\":\"South Georgia & South Sandwich Islands\",\"value\":\"GS\"},{\"label\":\"South Korea\",\"value\":\"KR\"},{\"label\":\"South Sudan\",\"value\":\"SS\"},{\"label\":\"Spain\",\"value\":\"ES\"},{\"label\":\"Sri Lanka\",\"value\":\"LK\"},{\"label\":\"St. Barthélemy\",\"value\":\"BL\"},{\"label\":\"St. Helena\",\"value\":\"SH\"},{\"label\":\"St. Kitts & Nevis\",\"value\":\"KN\"},{\"label\":\"St. Lucia\",\"value\":\"LC\"},{\"label\":\"St. Martin\",\"value\":\"MF\"},{\"label\":\"St. Pierre & Miquelon\",\"value\":\"PM\"},{\"label\":\"St. Vincent & Grenadines\",\"value\":\"VC\"},{\"label\":\"Sudan\",\"value\":\"SD\"},{\"label\":\"Suriname\",\"value\":\"SR\"},{\"label\":\"Svalbard & Jan Mayen\",\"value\":\"SJ\"},{\"label\":\"Sweden\",\"value\":\"SE\"},{\"label\":\"Switzerland\",\"value\":\"CH\"},{\"label\":\"Syria\",\"value\":\"SY\"},{\"label\":\"Taiwan\",\"value\":\"TW\"},{\"label\":\"Tajikistan\",\"value\":\"TJ\"},{\"label\":\"Tanzania\",\"value\":\"TZ\"},{\"label\":\"Thailand\",\"value\":\"TH\"},{\"label\":\"Timor-Leste\",\"value\":\"TL\"},{\"label\":\"Togo\",\"value\":\"TG\"},{\"label\":\"Tokelau\",\"value\":\"TK\"},{\"label\":\"Tonga\",\"value\":\"TO\"},{\"label\":\"Trinidad & Tobago\",\"value\":\"TT\"},{\"label\":\"Tristan da Cunha\",\"value\":\"TA\"},{\"label\":\"Tunisia\",\"value\":\"TN\"},{\"label\":\"Turkey\",\"value\":\"TR\"},{\"label\":\"Turkmenistan\",\"value\":\"TM\"},{\"label\":\"Turks & Caicos Islands\",\"value\":\"TC\"},{\"label\":\"Tuvalu\",\"value\":\"TV\"},{\"label\":\"U.S. Outlying Islands\",\"value\":\"UM\"},{\"label\":\"U.S. Virgin Islands\",\"value\":\"VI\"},{\"label\":\"Uganda\",\"value\":\"UG\"},{\"label\":\"Ukraine\",\"value\":\"UA\"},{\"label\":\"United Arab Emirates\",\"value\":\"AE\"},{\"label\":\"United Kingdom\",\"value\":\"GB\"},{\"label\":\"United States\",\"value\":\"US\"},{\"label\":\"Uruguay\",\"value\":\"UY\"},{\"label\":\"Uzbekistan\",\"value\":\"UZ\"},{\"label\":\"Vanuatu\",\"value\":\"VU\"},{\"label\":\"Vatican City\",\"value\":\"VA\"},{\"label\":\"Venezuela\",\"value\":\"VE\"},{\"label\":\"Vietnam\",\"value\":\"VN\"},{\"label\":\"Wallis & Futuna\",\"value\":\"WF\"},{\"label\":\"Western Sahara\",\"value\":\"EH\"},{\"label\":\"Yemen\",\"value\":\"YE\"},{\"label\":\"Zambia\",\"value\":\"ZM\"},{\"label\":\"Zimbabwe\",\"value\":\"ZW\"},{\"label\":\"Other\",\"value\":\"other\"}]")
//	var ss []countries
//	err := json.NewDecoder(r).Decode(&ss)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(ss)
//	collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("countries")
//	for _, c := range ss {
//		_, err = collection.InsertOne(ctx, c)
//		if err != nil {
//			panic(err)
//		}
//	}
//}

func (s *Server) getOS(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var countries []country

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("os")
		cur, err := collection.Find(ctx, bson.D{{}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			//Create a value into which the single document can be decoded
			var elem country
			err := cur.Decode(&elem)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			countries = append(countries, elem)

		}
		if err := cur.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res, err := json.Marshal(countries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}
