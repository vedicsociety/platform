// Implementation of the Configuration interface
// The DefaultConfig struct implements the Configuration interface using a map.
// Nested configuration sections are also expressed as maps.
// An individual configuration setting can be requested by separating a section name from a setting name,
// such as logging:level, or a map containing all of the settings can be requested using a section name, such as logging.
// To define the methods that accept a default value, see a file named config_default_fallback.go to the config folder
package config

import "strings"

type DefaultConfig struct {
    configData map[string]interface{}
}

func (c *DefaultConfig) get(name string) (result interface{}, found bool) {
    data := c.configData
    for _, key := range strings.Split(name, ":") {
        result, found = data[key]
        if newSection, ok := result.(map[string]interface{}); ok && found {
            data = newSection
        } else {
            return 
        }
    }
    return
}

func (c *DefaultConfig) GetSection(name string) (section Configuration, found bool) {
    value, found := c.get(name)
    if (found) {
        if sectionData, ok := value.(map[string]interface{}) ; ok {
            section = &DefaultConfig { configData: sectionData }
        }
    }
    return
}

func (c *DefaultConfig) GetString(name string) (result string, found bool) {
    value, found := c.get(name)
    if (found) { result = value.(string) }
    return
}

func (c *DefaultConfig) GetInt(name string) (result int, found bool) {
    value, found := c.get(name)
    if (found) { result =  int(value.(float64)) }
    return
}

func (c *DefaultConfig) GetBool(name string) (result bool, found bool) {
    value, found := c.get(name)
    if (found) { result = value.(bool) }
    return
}

func (c *DefaultConfig) GetFloat(name string) (result float64, found bool) {
    value, found := c.get(name)
    if (found) { result = value.(float64) }
    return
}
