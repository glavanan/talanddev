package dbconnector

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Used to prevent multiple connection to the DB
var userDB *gorm.DB

var ExchangeRate map[string]float64

//InitDB init connection to the postgres DB
func InitDB(dsn string) {
	var err error
	userDB, err = gorm.Open(postgres.New(postgres.Config{DSN: dsn}))
	if err != nil {
		log.WithField("InitDB", "OpenGorm").Error(err)
	}
	// resp, err := http.Get("http://data.fixer.io/api/latest?access_key=bcc171454946d10e6d8c119a1c9528fb")
	// if err != nil {
	// 	log.WithField("InitDB", "GetExchangeRate").Error(err)
	// }
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.WithField("InitDB", "GetExchangeRate").Error(err)
	// }
	// var exchange model.Exchange
	// json.Unmarshal(body, &exchange)
	// ExchangeRate = exchange.Rates
	ExchangeRate = map[string]float64{"AED": 4.204469, "AFN": 104.681747, "ALL": 122.371207, "AMD": 544.410639, "ANG": 2.063407, "AOA": 683.262633, "ARS": 114.722566, "AUD": 1.561652, "AWG": 2.060716, "AZN": 1.950506, "BAM": 1.955725, "BBD": 2.311696, "BDT": 98.205804, "BGN": 1.955628, "BHD": 0.431584, "BIF": 2285.932481, "BMD": 1.144683, "BND": 1.549971, "BOB": 7.894274, "BRL": 6.245511, "BSD": 1.144948, "BTC": 1.7905344e-05, "BTN": 85.257057, "BWP": 13.130007, "BYN": 2.801581, "BYR": 22435.792002, "BZD": 2.307796, "CAD": 1.436389, "CDF": 2300.813778, "CHF": 1.054094, "CLF": 0.033224, "CLP": 916.754387, "CNY": 7.302741, "COP": 4446.86555, "CRC": 735.716749, "CUC": 1.144683, "CUP": 30.334107, "CVE": 110.662301, "CZK": 25.24718, "DJF": 203.433561, "DKK": 7.437626, "DOP": 64.793521, "DZD": 158.272488, "EGP": 17.990534, "ERN": 17.171667, "ETB": 54.147955, "EUR": 1, "FJD": 2.408266, "FKP": 0.839307, "GBP": 0.853127, "GEL": 3.6062, "GGP": 0.839307, "GHS": 6.999784, "GIP": 0.839307, "GMD": 59.642409, "GNF": 10874.491428, "GTQ": 8.856022, "GYD": 239.537081, "HKD": 8.91666, "HNL": 27.68421, "HRK": 7.51279, "HTG": 113.651871, "HUF": 367.21878, "IDR": 16252.098534, "ILS": 3.558024, "IMP": 0.839307, "INR": 85.140683, "IQD": 1671.237568, "IRR": 48391.485449, "ISK": 150.42327, "JEP": 0.839307, "JMD": 178.231945, "JOD": 0.811626, "JPY": 130.423499, "KES": 128.2622, "KGS": 97.049228, "KHR": 4664.584708, "KMF": 492.271478, "KPW": 1030.214596, "KRW": 1350.303159, "KWD": 0.345981, "KYD": 0.95414, "KZT": 493.629065, "LAK": 12156.536678, "LBP": 1754.799845, "LKR": 231.272507, "LRD": 165.464403, "LSL": 17.519422, "LTL": 3.379953, "LVL": 0.692408, "LYD": 5.231638, "MAD": 10.451394, "MDL": 20.198687, "MGA": 4564.428822, "MKD": 61.598695, "MMK": 2035.405309, "MNT": 3263.405505, "MOP": 9.189824, "MRO": 408.651729, "MUR": 49.454602, "MVR": 17.68579, "MWK": 934.06194, "MXN": 23.522171, "MYR": 4.754446, "MZN": 73.065565, "NAD": 17.525533, "NGN": 469.984487, "NIO": 40.316176, "NOK": 9.953124, "NPR": 136.413835, "NZD": 1.624724, "OMR": 0.440642, "PAB": 1.144938, "PEN": 4.602816, "PGK": 4.035053, "PHP": 56.999547, "PKR": 201.125089, "PLN": 4.645801, "PYG": 7873.075393, "QAR": 4.167836, "RON": 4.949654, "RSD": 117.616633, "RUB": 83.398233, "RWF": 1150.406682, "SAR": 4.293162, "SBD": 9.201687, "SCR": 16.039243, "SDG": 501.371659, "SEK": 10.035581, "SGD": 1.547864, "SHP": 1.576691, "SLL": 12591.516307, "SOS": 668.495414, "SRD": 24.543731, "STD": 23692.632488, "SVC": 10.018372, "SYP": 1438.834841, "SZL": 17.525524, "THB": 37.465906, "TJS": 12.920334, "TMT": 4.006391, "TND": 3.24232, "TOP": 2.585501, "TRY": 11.471677, "TTD": 7.773913, "TWD": 31.830251, "TZS": 2633.916578, "UAH": 29.951446, "UGX": 4041.582031, "USD": 1.144683, "UYU": 50.015866, "UZS": 12271.004988, "VEF": 2.4476787715169238e+11, "VND": 25921.352545, "VUV": 128.543848, "WST": 2.949841, "XAF": 655.927355, "XAG": 0.04529, "XAU": 0.000614, "XCD": 3.093564, "XDR": 0.814963, "XOF": 655.335292, "XPF": 119.762529, "YER": 286.457404, "ZAR": 17.579134, "ZMK": 10303.527122, "ZMW": 20.014295, "ZWL": 368.587544}
}

//GetDB check connection and GetDB
func GetDB() (*gorm.DB, error) {
	db, err := userDB.DB()
	if err != nil {
		log.WithField("CheckDB", "Access DB").Error(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.WithField("CheckDB", "First Ping failed").Error(err)
		err = db.Ping()
		if err != nil {
			log.WithField("CheckDB", "Second Ping failed, connection to DB not possible").Error(err)
			//Then a possibilities to try to open a new connection but not implemented for the test
			return nil, err
		}
	}
	return userDB, nil
}
