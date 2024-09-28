package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	logTag = "config"
)

var (
	// errNilConfig is returned when a nil reference is passed in as Un/Marshaler reference
	errNilConfig = errors.New("Config object is empty.")

	logReplacementsOnce sync.Once
	verbose             bool
)

// Verbose enables or disables verbose logging. The default is false, which means this library will not print logs.
func Verbose(enabled bool) {
	verbose = enabled
}

// Loader represents a configuration loader delegate
type loader func(string) ([]byte, error)

// SaveJSONFile will save your struct to the given filename,
// this is a good way to create a json template when your
//
// note: for developer only, to generate json template
// usage: 1) modify TestSaveConfigToFile with your config struct
//  2. go test -run TestSaveConfigToFile
//  3. grab your_beautifully_formatted_config_file.json and comment the test case as it was
func SaveJSONFile(filename string, config interface{}) error {
	bytes, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0660)
}

// LoadJSONFile gets your config from the json file,
// and fills your struct with the option
func LoadJSONFile(filename string, config interface{}) error {
	if config == nil {
		return errNilConfig
	}

	// Default to loading from file, for safety
	loadConfig := loader(loadFromFile)
	
	// Load the configuration
	bytes, err := loadConfig(filename)
	if err != nil {
		return err
	}
	return parseConfig(bytes, config)
}

// parseConfig parses the configuration from JSON (or something else in future).
func parseConfig(b []byte, config interface{}) error {
	return json.Unmarshal(UpdateFromEnv(b), config)
}

// LoadJSONEnvPathOrPanic calls LoadJSONEnvPath but panics on error
func LoadJSONEnvPathOrPanic(envVar string, config interface{}) {
	if err := LoadJSONEnvPath(envVar, config); err != nil {
		panic(fmt.Errorf("failed to load config file with error %s", err))
	}
}

// LoadJSONEnvPath gets your config from the json file provided by env var,
// and fills your struct with the option
func LoadJSONEnvPath(envVar string, config interface{}) error {
	if config == nil {
		return errNilConfig
	}

	filename := os.Getenv(envVar)
	if filename == "" {
		return fmt.Errorf("Env var is empty: %s", envVar)
	}
	//logPrintf(logTag+" : loading config from envVar %s, file = %s", envVar, filename)
	return LoadJSONFile(filename, config)
}

// UpdateFromEnv is used during testing to update the config from environmental settings (needed for Go.CD and docker)
func UpdateFromEnv(bytes []byte) []byte {
	configAsString := string(bytes)

	configAsString, replacements := replaceConfigEnvVars(configAsString)
	if len(replacements) > 0 {
		logReplacementsOnce.Do(func() {
			//logPrintf(logTag + " : detected env vars, replacing " + strings.Join(replacements, ", "))
		})
	}

	return []byte(configAsString)
}

// ReplaceConfigEnvVars replaces $ENV_VAR strings with their equivalents defined in the environment.
// No substitutions are made if the corresponding environment variable is not defined.
//
// For a list of supported variables, see configEnvVars.
func ReplaceConfigEnvVars(input string) string {
	result, _ := replaceConfigEnvVars(input)
	return result
}

// This list is from scripts/env-vars.sh
var configEnvVars = []string{
	"REDIS_HOST",
	"REDIS_CLUSTER_HOST",
	"MYSQL_HOST",
	"POSTGRES_HOST",
	"GRABDDB_HOST",
	"ETCD_HOST",
	"ELASTICSEARCH_HOST",
	"SCYLLA_HOST",
	"KAFKA_HOST",
	"DD_AGENT_HOST",
}

func replaceConfigEnvVars(input string) (string, []string) {
	var replacements []string

	for _, envVar := range configEnvVars {
		if value := os.Getenv(envVar); value != "" {
			input = strings.Replace(input, "$"+envVar, value, -1)
			replacements = append(replacements, "$"+envVar+" with "+value)
		}
	}

	return input, replacements
}

// loads a file from HTTP
func loadFromHTTP(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// Write the body to file
	//logPrintf("%s : loading config from HTTP %s", logTag, uri)
	return ioutil.ReadAll(resp.Body)
}

// loads a file from OS File
func loadFromFile(uri string) ([]byte, error) {
	//logPrintf("%s : loading config from OS File %s", logTag, uri)
	return ioutil.ReadFile(uri)
}
