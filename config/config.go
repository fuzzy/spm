package main

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
Parse our config.
TODO: macro expansion after loading the JSON config
*/
type SpmConfig struct {
	Color       bool `json:"Color"`
	Directories struct {
		System struct {
			BinpkgRoot  string `json:"BinpkgRoot"`
			DistFiles   string `json:"DistFiles"`
			GlobalRoot  string `json:"GlobalRoot"`
			InstallRoot string `json:"InstallRoot"`
			Profiles    string `json:"Profiles"`
			SessionRoot bool   `json:"SessionRoot"`
		} `json:"System"`
		User struct {
			BinpkgRoot  string `json:"BinpkgRoot"`
			DistFiles   string `json:"DistFiles"`
			GlobalRoot  string `json:"GlobalRoot"`
			InstallRoot string `json:"InstallRoot"`
			Profiles    string `json:"Profiles"`
			SessionRoot string `json:"SessionRoot"`
		} `json:"User"`
	} `json:"Directories"`
	Logging struct {
		Facility bool   `json:"Facility"`
		Logfile  string `json:"Logfile"`
		Loglevel string `json:"Loglevel"`
		Output   string `json:"Output"`
	} `json:"Logging"`
	PerUser bool `json:"PerUser"`
}

var SpmCfg SpmConfig

func LoadConfig(cfg string) error {
	data, e := ioutil.ReadFile(cfg)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	e = json.Unmarshal([]byte(data), &SpmCfg)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
