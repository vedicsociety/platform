// Define the function that will load the data from the configuration file.
// The configuration file config.json defines two configuration sections, named logging and main.
// The logging section contains a single string configuration setting, named level.
// The main section contains a single string configuration setting named message.
// When adding configuration settings, pay close attention to the quote marks and the commas,
// both of which are required by JSON, but which are easy to omit.
package config

import (
	"encoding/json"
	"os"
	"strings"
)

// The Load function reads the contents of a file, decodes the JSON it contains into a map, and uses the map to create a DefaultConfig value.
func Load(fileName string) (config Configuration, err error) {
	var data []byte
	data, err = os.ReadFile(fileName)
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		m := map[string]interface{}{}
		err = decoder.Decode(&m)
		if err == nil {
			config = &DefaultConfig{configData: m}
		}
		loadEnv(config)
	}
	return
}

func loadEnv(c Configuration) {
	// get prefix
	pref, f := c.GetString("system:prefix")
	if f {
		// retrive all of environ vars as slice [](name=value)
		for _, env := range os.Environ() {
			// retrive names only
			pair := strings.SplitN(env, "=", 2)
			// filter with prefix only
			if strings.HasPrefix(pair[0], pref) {
				// trim prefix
				oskey := strings.TrimPrefix(pair[0], pref)
				// replace _ to : (by agreement)
				oskey = strings.ReplaceAll(oskey, "_", ":")
				// set value to config
				c.SetString(oskey, pair[1])
			}
		}
	}
}
