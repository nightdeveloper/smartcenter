package settings

import (
	"encoding/json"
	"log"
	"path/filepath"
	"io/ioutil"
	"time"
)

type Config struct {
	TelegramKey		string		`json:"telegramKey"`
	TelegramOpId	int64		`json:"telegramOpId"`
	PushUserId		string		`json:"pushUserId"`
	PushToken		string		`json:"pushToken"`
	PushDeviceName	string		`json:"pushDeviceName"`
	LastAlive		*time.Time	`json:"lastAlive"`
	GetIPURL1		string		`json:"getIPURL1"`
	GetIPURL2		string		`json:"getIPURL2"`
	ProxyUrl		string		`json:"proxyUrl"`
	ProxyUser		string		`json:"proxyUser"`
	ProxyPass		string		`json:"proxyPass"`
}

func (c *Config) getFileName() string {
	absPath, _ := filepath.Abs("./");
	return absPath + "/config_sc.json";
}

func (c *Config) Load() {

	file, err := ioutil.ReadFile(c.getFileName())

	if err != nil {
		log.Fatal("Config reading error from " + c.getFileName() + " ", err);
		panic("config reading error");
	}

	err = json.Unmarshal(file, c);

	if err != nil || c == nil {
		log.Fatal("Config decoding error ", err);
		panic("config decoding error");
	}

	if c.TelegramKey == "" || c.TelegramOpId == 0 {
		log.Fatal("we need dropbox dir and podcasts list to go")
		panic("config content error")
	}

	log.Println("config read");
}

func (c *Config) Save() {

	out, err := json.MarshalIndent(c, "", "	")

	err = ioutil.WriteFile(c.getFileName(), out,0755)

	if err != nil {
		log.Fatal("Config write error ", err);
		panic("config write error");
	}

//	log.Println("Config saved");
}