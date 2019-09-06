package paykassasci

// VersionSCI is version sci
const VersionSCI = "0.4"

// PayKassaURL is actual url for request
const PayKassaURL = "https://paykassa.pro/sci/" + VersionSCI + "/index.php"

// Available func
const (
	ConfirmOrder       = "sci_confirm_order"
	CreateOrder        = "sci_create_order"
	CreateOrderGetData = "sci_create_order_get_data"
)

// Available currency
const (
	Payeer          = 1  // USD|RUB
	PerfectMoney    = 2  // USD
	AdvCash         = 4  // USD|RUB
	Berty           = 7  // USD|RUB
	BitCoin         = 11 // BTC
	Ethereum        = 12 // ETH
	Litecoin        = 14 // LTC
	Dogecoin        = 15 // DOGE
	Dash            = 16 // DASH
	BitcoinCash     = 18 // BCH
	Zcash           = 19 // ZEC
	Monero          = 20 // XMR
	EthereumClassic = 21 // ETC
	Ripple          = 22 // XRP
	NEO             = 23 // NEO
	Gas             = 24 // GAS
	BitcoinSV       = 25 // BSV
	Waves           = 26 // WAVES
	TRON            = 27 // TRX
	Stellar         = 28 // XLM
)

var currencyCode = map[int]string{
	BitCoin:         "BTC",
	Ethereum:        "ETH",
	Litecoin:        "LTC",
	Dogecoin:        "DOGE",
	Dash:            "DASH",
	BitcoinCash:     "BCH",
	Zcash:           "ZEC",
	Monero:          "XMR",
	EthereumClassic: "ETC",
	Ripple:          "XRP",
	NEO:             "NEO",
	Gas:             "GAS",
	BitcoinSV:       "BSV",
	Waves:           "WAVES",
	TRON:            "TRX",
	Stellar:         "XLM",
}

// CurrencyCode returned code by currency id
func CurrencyCode(id int) string {
	return currencyCode[id]
}
