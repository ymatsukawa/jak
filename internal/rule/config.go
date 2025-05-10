package rule

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/ymatsukawa/jak/internal/file"
)

const (
	DefaultConfigFile = "jak.toml"
	DefaultTimeout    = 30 // seconds
)

type Request struct {
	Name      string            `toml:"name"`
	Method    string            `toml:"method"`
	Path      string            `toml:"path"`
	Headers   []string          `toml:"headers"`
	RawBody   *string           `toml:"raw_body"`
	FormBody  *string           `toml:"form_body"`
	JsonBody  *string           `toml:"json_body"`
	Extract   map[string]string `toml:"extract"`
	DependsOn string            `toml:"depends_on"`
}

type Config struct {
	BaseUrl     string    `toml:"base_url"`
	Timeout     uint8     `toml:"timeout"`
	Concurrency bool      `toml:"concurrency"`
	IgnoreFail  bool      `toml:"ignore_fail"`
	Request     []Request `toml:"request"`
}

func LoadConfig(path string) (*Config, error) {
	configPath, err := file.AbsPath(path)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve config path: %w", err)
	}

	var config Config
	_, err = toml.DecodeFile(configPath, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode TOML: %w", err)
	}

	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	return &config, nil
}

func (c *Config) Validate() error {
	if err := c.validateBaseConfig(); err != nil {
		return err
	}

	return c.validateRequests()
}

func (c *Config) validateBaseConfig() error {
	if c.BaseUrl == "" {
		return fmt.Errorf("base_url is required in config")
	}
	if len(c.Request) == 0 {
		return fmt.Errorf("at least one request must be defined")
	}
	return nil
}

func (c *Config) validateRequests() error {
	nameSet := make(map[string]bool)
	for i, req := range c.Request {
		if err := validateRequest(req, i, nameSet); err != nil {
			return err
		}
	}
	return nil
}

func validateRequest(req Request, index int, nameSet map[string]bool) error {
	if err := validateRequestName(req.Name, index, nameSet); err != nil {
		return err
	}
	if err := validateRequestBasics(req); err != nil {
		return err
	}
	return validateRequestBody(req)
}

func validateRequestName(name string, index int, nameSet map[string]bool) error {
	if name == "" {
		return fmt.Errorf("request at index %d has no name", index)
	}
	if nameSet[name] {
		return fmt.Errorf("duplicate request name: %s", name)
	}
	nameSet[name] = true
	return nil
}

func validateRequestBasics(req Request) error {
	if req.Method == "" {
		return fmt.Errorf("method is required for request '%s'", req.Name)
	}
	if req.Path == "" {
		return fmt.Errorf("path is required for request '%s'", req.Name)
	}
	return nil
}

func validateRequestBody(req Request) error {
	bodyCount := 0
	if req.JsonBody != nil && *req.JsonBody != "" {
		bodyCount++
	}
	if req.FormBody != nil && *req.FormBody != "" {
		bodyCount++
	}
	if req.RawBody != nil && *req.RawBody != "" {
		bodyCount++
	}
	if bodyCount > 1 {
		return fmt.Errorf("multiple body types specified for request '%s'", req.Name)
	}
	return nil
}

func (r *Request) GetBody() string {
	if r.JsonBody != nil {
		return *r.JsonBody
	}
	if r.FormBody != nil {
		return *r.FormBody
	}
	if r.RawBody != nil {
		return *r.RawBody
	}

	return ""
}
