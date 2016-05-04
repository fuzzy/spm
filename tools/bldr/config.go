package main

/*
 Author: Mike 'Fuzzy' Partin
 Copyright: (c) 2016-2018
 Email: fuzzy@fumanchu.org
 License: See LICENSE.md for details
*/

import (
	"encoding/json"
	"os"
)

type DirConfig struct {
	System struct {
		BinPkgRoot  string `json:"binPkgRoot"`
		DistFiles   string `json:"distFiles"`
		GlobalRoot  string `json:"globalRoot"`
		InstRoot    string `json:"instRoot"`
		Profiles    string `json:"profiles"`
		SessionRoot bool   `json:"sessionRoot"`
	} `json:"System"`
	User struct {
		BinPkgRoot  string `json:"binPkgRoot"`
		DistFiles   string `json:"distFiles"`
		GlobalRoot  string `json:"globalRoot"`
		InstRoot    string `json:"instRoot"`
		Profiles    string `json:"profiles"`
		SessionRoot string `json:"sessionRoot"`
	} `json:"User"`
}
