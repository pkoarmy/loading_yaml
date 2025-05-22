package main

import "fmt"

type configItems struct {
	cname          string
	flatten_header string
	match_percent  int
	match_bucket   string
	match_variable string
	attach_header  bool
}

var (
	bucketConfig map[string]configItems = make(map[string]configItems)
	err          error
)

const (
	CONFIG = "./bucket.yaml"
)

func main() {
	fmt.Println("hello")

	/*
		mydomain.com:
		  cname: mydomain.com
		  flatten_header: myheader
		  match_percent: <nil> # 50 vs 50 ( 100% )
		  match_bucket: <nil>
		  match_variable: <nil>
		  attach_header: true
	*/
	yamlConfig, err := lib.LoadYAML(CONFIG)
	if err != nil {
		fmt.Println(" error on Load YAML", CONFIG, err)
	}
	for k, v := range yamlConfig {
		vv := v.(map[string]any)
		ci := configItems{}
		if vv["cname"] != nil {
			ci.cname = vv["cname"].(string)
		}
		if vv["flatten_header"] != nil {
			ci.flatten_header = vv["flatten_header"].(string)
		}
		if vv["match_percent"] != nil {
			ci.match_percent = vv["match_percent"].(int)
		}
		if vv["match_bucket"] != nil {
			ci.match_bucket = vv["match_bucket"].(string)
		}
		if vv["match_variable"] != nil {
			ci.match_variable = vv["match_variable"].(string)
		}
		if vv["attach_header"] != nil {
			ci.attach_header = vv["attach_header"].(bool)
		}

		bucketConfig[k] = ci
	}

}
