package main

import (
	//	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SplatSpec struct {
	Author  string `json:"Author" yaml:"Author"`
	Binpkg  bool   `json:"Binpkg" yaml:"Binpkg"`
	Compile []struct {
		Arguments   []string   `json:"Arguments" yaml:"Arguments"`
		Command     string     `json:"Command" yaml:"Command"`
		Directory   string     `json:"Directory" yaml:"Directory"`
		Environment [][]string `json:"Environment" yaml:"Environment"`
		User        string     `json:"User" yaml:"User"`
	} `json:"Compile" yaml:"Compile"`
	Configure []struct {
		Arguments   []string   `json:"Arguments" yaml:"Arguments"`
		Command     string     `json:"Command" yaml:"Command"`
		Directory   string     `json:"Directory" yaml:"Directory"`
		Environment [][]string `json:"Environment" yaml:"Environment"`
		User        string     `json:"User" yaml:"User"`
	} `json:"Configure" yaml:"Configure"`
	Email     string `json:"Email" yaml:"Email"`
	Homepage  string `json:"Homepage" yaml:"Homepage"`
	Inclusive bool   `json:"Inclusive" yaml:"Inclusive"`
	Install   []struct {
		Arguments   []string   `json:"Arguments" yaml:"Arguments"`
		Command     string     `json:"Command" yaml:"Command"`
		Directory   string     `json:"Directory" yaml:"Directory"`
		Environment [][]string `json:"Environment" yaml:"Environment"`
		User        string     `json:"User" yaml:"User"`
	} `json:"Install" yaml:"Install"`
	License  string   `json:"License" yaml:"License"`
	Package  string   `json:"Package" yaml:"Package"`
	Requires []string `json:"Requires" yaml:"Requires"`
	Setup    []struct {
		Arguments   []string   `json:"Arguments" yaml:"Arguments"`
		Command     string     `json:"Command" yaml:"Command"`
		Directory   string     `json:"Directory" yaml:"Directory"`
		Environment [][]string `json:"Environment" yaml:"Environment"`
		User        string     `json:"User" yaml:"User"`
	} `json:"Setup" yaml:"Setup"`
	Sources struct {
		Data    [][]string `json:"Data" yaml:"Data"`
		Patches [][]string `json:"Patches" yaml:"Patches"`
		Sha1    string     `json:"Sha1" yaml:"Sha1"`
		Uris    []string   `json:"Uris" yaml:"Uris"`
	} `json:"Sources" yaml:"Sources"`
	Teardown []struct {
		Arguments   []string   `json:"Arguments" yaml:"Arguments"`
		Command     string     `json:"Command" yaml:"Command"`
		Directory   string     `json:"Directory" yaml:"Directory"`
		Environment [][]string `json:"Environment" yaml:"Environment"`
		User        string     `json:"User" yaml:"User"`
	} `json:"Teardown" yaml:"Teardown"`
	Version string `json:"Version" yaml:"Version"`
}

func ReadSplat(f string) (SplatSpec, error) {
	retv := SplatSpec{}
	fptr, err := ioutil.ReadFile(f)

	if err != nil {
		fmt.Println("ERROR: ioutil.ReadFile(f)")
		panic(err)
	}

	if f[len(f)-4:] == "json" {
		// read in json file
		err := yaml.Unmarshal(fptr, &retv)
		if err != nil {
			fmt.Println("ERROR: json unmarshal failed.")
			panic(err)
		}
	} else if f[len(f)-4:] == "yaml" {
		// read in yaml file
		err := yaml.Unmarshal(fptr, &retv)
		if err != nil {
			fmt.Println("ERROR: yaml unmarshal failed.")
			panic(err)
		} else {
			fmt.Println(err)
		}
	}

	return retv, nil
}
