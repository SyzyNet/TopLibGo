package shared

import toplib "toplib/pkg"

// All global variables are defined here and can be accessed from any package.
var (
	AccountRegex     = "^[a-zA-Z0-9]{3,20}:[a-zA-Z0-9]{3,20}$"
	ProxyRegex       = `((http|https|socks4|socks5):\/\/)?([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?`
	HitsNumber   int = 0
	FailsNumber  int = 0
	CPM          int = 0
	Retries      int = 0
	AccountList  []toplib.Account
	HitList      []*toplib.Account
	ProxyList    []*toplib.Proxy
	OnCheck      = make(chan string)
)

//Make a hits flag channel
