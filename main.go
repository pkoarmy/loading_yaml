package main

import (
	"fmt"
	"myapp/lib" // Module name 'myapp'
)

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
	CONFIG = "./bucket.yaml" // This file won't exist yet, but the code refers to it.
)

func main() {
	fmt.Println("hello")

	yamlConfig, err := lib.LoadYAML(CONFIG)
	if err != nil {
		fmt.Println(" error on Load YAML", CONFIG, err)
		// Consider os.Exit(1) here or other error handling
		return // Exit if YAML loading fails
	}

	for k, v_interface := range yamlConfig {
		// v_interface is of type interface{}. We need to assert it to map[string]interface{}
		// which is the type our helper functions expect for their first argument.
		vv, ok := v_interface.(map[string]interface{})
		if !ok {
			fmt.Printf("Warning: Invalid structure for key '%s' in YAML. Expected a map, got %T.\n", k, v_interface)
			continue // Skip this entry if it's not a map
		}

		ci := configItems{}
		ci.cname = lib.GetString(vv, "cname", "")
		ci.flatten_header = lib.GetString(vv, "flatten_header", "")
		ci.match_percent = lib.GetInt(vv, "match_percent", 0) // Defaulting to 0 if nil or wrong type
		ci.match_bucket = lib.GetString(vv, "match_bucket", "")
		ci.match_variable = lib.GetString(vv, "match_variable", "")
		ci.attach_header = lib.GetBool(vv, "attach_header", false) // Defaulting to false

		bucketConfig[k] = ci
	}

	// Optional: Print the loaded config to verify
	// for key, config := range bucketConfig {
	//  	fmt.Printf("Loaded config for %s: %+v\n", key, config)
	// }
}
