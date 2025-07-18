package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// The following context keys can be used to modify the server URL used by the client.
	// However, we only support a single server configuration at the moment.
	// This needs to be exposed to the user in the future when we support multiple server configurations.

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	ClientName       string            `json:"clientName,omitempty"`
	NIOSHostURL      string            `json:"niosHostURL,omitempty"`
	NIOSUsername     string            `json:"niosUsername,omitempty"`
	NIOSPassword     string            `json:"niosPassword,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
	DefaultExtAttrs  map[string]struct{ Value string }
	ClientCert       []byte
	ClientKey        []byte
	SslVerify        bool
}

// NewConfiguration returns a new Configuration object.
// The following default values are set:
// - ClientName: "nios-go-client"
// - UserAgent: "nios-go-client/version"
// - Debug: false
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		ClientName:       "nios-go-client",
		NIOSHostURL:      lookupEnv(envNiosHostURL, ""),
		NIOSUsername:     lookupEnv(envNiosUsername, ""),
		NIOSPassword:     lookupEnv(envNiosPassword, ""),
		DefaultHeader:    make(map[string]string),
		Debug:            lookupEnvBool(envIBLogLevel, true),
		UserAgent:        fmt.Sprintf("nios-%s/%s", sdkIdentifier, version),
		Servers:          ServerConfigurations{},
		OperationServers: map[string]ServerConfigurations{},
		DefaultExtAttrs:  make(map[string]struct{ Value string }),
		ClientCert:       readFile(envClientCertPath),
		ClientKey:        readFile(envClientKeyPath),
		SslVerify:        false,
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, ReportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, ReportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, ReportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, ReportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

// lookupEnv is a function that returns the value of the environment variable named by the key
// or the default value if the environment variable is not set.
func lookupEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func lookupEnvBool(key string, def bool) bool {
	if logLvlStr, ok := os.LookupEnv(key); ok {
		if logLvl, err := strconv.ParseBool(logLvlStr); err == nil {
			return logLvl
		}
	}
	return def
}

func readFile(filepath string) []byte {
	filepath = lookupEnv(filepath, "")
	if filepath == "" {
		return nil
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("Error reading client cert file '%s': %v", filepath, err)
		return nil
	}
	return data
}
