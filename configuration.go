package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"net/http"
	"time"
	"path/filepath"
)

//MinerConfig struct for json parse
type MinerConfig struct {
	Name string // NAME
	Pin  string // PIN-NUMBER OF GPIO
	IP   string // IP ADDRESS
	Info string // ADDITIONAL INFO
}

//ConfigurationFile struct to parse config.json
type ConfigurationFile struct {
	WaitSeconds     int           // Period of the timer checks in seconds
	StartupCheck    bool          // Check miners on startup
	Log             bool          //Enable or disable logging
	RemoteNotify    bool          //Remote notification telegram,pushover & etc
	TgBotActivate   bool          //Enable or disable Telegram bot
	TgAPIKey        string        //Telegram Api key for bot communicationg
	TgAdminUserName string        //Telegram Username which will control the bot
	Pushover        bool          //Enable or disable Pushover notifications
	PushoverToken   string        //Pushover access token
	PushoverUser    string        //Pushover user token
	Miners          []MinerConfig // An array of the
}

type NicehasherFile struct {
    Result struct {
    Error  string	  `json:"error"`
    	   Current []struct {
	   	   Profitability	string       `json:"profitability"`
		   Data  []interface{} 	`json:"data"`
		   Name        		string       `json:"name"`
		   Suffix       	string       `json:"suffix"`
		   Algo       		int         `json:"algo"`
	   } `json:"current"`
	   NhWallet		bool `json:"nh_wallet"`
	   AttackWrittenOff 	int  `json:"attack_written_off"`
	   Past           []struct {
	   	Data 	  	[][]interface{} `json:"data"`
		Algo 		int           `json:"algo"`
	   } `json:"past"`
	   Payments []struct {
		Amount string `json:"amount"`
		Fee    string `json:"fee"`
		TXID   string `json:"TXID"`
		Time   int    `json:"time"`
		Type   int    `json:"type"`
	   } `json:"payments"`
		   AttackAmount string `json:"attack_amount"`
		   Addr        string `json:"addr"`
		   AttackRepaid string `json:"attack_repaid"`
	   } `json:"result"`
	   Method string `json:"method"`
																														    }

//Config is the global Config variable
var Config ConfigurationFile
var Nicehashing NicehasherFile
//ReadConfig - read and parse the config file
func ReadConfig() (configFile ConfigurationFile) {
	//get binary dir
	//os.Args doesn't work the way we want with "go run". You can use next line
	//for local dev, but use the original for production.
	dir, err := filepath.Abs("./")
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	log.Notice("Reading file config.json...")
	file := dir + "/config.json"

	configFileContent, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error("Trying to read file config.json, but:", err)
		os.Exit(1)
	}

	log.Notice("Parsing configuration file...")
	err = json.Unmarshal(configFileContent, &configFile)
	if err != nil {
		log.Error("Parsing JSON content, but:", err)
		os.Exit(2)
	}

	log.Notice("Timer (time in seconds):", configFile.WaitSeconds)
	log.Notice("Found miner configurations:", len(configFile.Miners))

	if configFile.Pushover == true {
		log.Notice("Pushover notification is ENABLED")
		log.Notice("Pushover Token:", configFile.PushoverToken)
		log.Notice("Pushover User:", configFile.PushoverUser)
	}

	return
}
func ReadNicehash() (nicehashFile NicehasherFile) {
     url := "https://api.nicehash.com/api?method=stats.provider.ex&addr=3M5xEw7mszPj5G9Hs2M5HwfArvJiU4em5q"

     Client := http.Client{
     	    Timeout: time.Second * 2, // Maximum of 2 secs
	    }

     req, err := http.NewRequest("GET", url, nil)
     	  if err != nil {
	     log.Notice(err)
	     }
     req.Header.Set("User-Agent", "raspberry-autorrestarter")

     res, getErr := Client.Do(req)
          if getErr != nil {
	     log.Notice(getErr)
	     }
     body, readErr := ioutil.ReadAll(res.Body)
           if readErr != nil {
	     log.Notice(readErr)
	     }
     nicehasher := NicehasherFile{}
        jsonErr := json.Unmarshal(body, &nicehasher)
	    if jsonErr != nil {
	            log.Notice(jsonErr)
		        }
     if len(nicehasher.Result.Error)==0{			

     	log.Notice("Nicehash method:", nicehasher.Result.Current[0].Profitability)
	}else{
	log.Notice("Error", nicehasher.Result.Error)
	}
	
return


}
