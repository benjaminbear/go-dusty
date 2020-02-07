package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Bundles       []string                     `yaml:"Bundles"`
	EnvOverrides  map[string]map[string]string `yaml:"dusty_env_overrides"`
	RepoOverrides map[string]string            `yaml:"repo_overrides"`
	Username      string                       `yaml:"username"`
	SSHKeyPath    string                       `yaml:"ssh_key_path"`
	SetupHasRun   bool                         `yaml:"setup_has_run"`
	SpecsRepo     string                       `yaml:"specs_repo"`
	VmMemorySize  int                          `yaml:"vm_memory_size"`
}

type Configurator interface {
	AddBundle(bundle string)
	RemoveBundle(bundle string) error
	SetEnvOverride(application, envName, envValue string)
	UnsetEnvOverride(application, envName string)
	SetRepoOverrides(repository, override string)
	UnsetRepoOverride(repository string)
	SetUsername(username string)
	GetUsername() string
	SetSSHKeyPath(key string)
	GetSSHKeyPath() string
	SetSetupHasRun()
	HasSetupRun() bool
	SetSpecsRepo(repository string)
	GetSpecsRepo() string
	SetVmMemorySize(size int)
	GetVmMemorySize() int
	SaveConfiguration() error
}

// NewConfiguration simply returns a new empty Configurator.
func NewConfiguration() Configurator {
	return &Config{}
}

// LoadConfiguration tries to load an existing config from config folder.
// If no config is found an empty config is created.
func LoadConfiguration() (Configurator, error) {
	config := &Config{}
	cfgPath, err := defaultConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		return config, nil
	}

	configBytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// SaveConfiguration saves the configuration to the config file.
// If it does not exist it will be created, if it already exists it will be overwritten.
func (c *Config) SaveConfiguration() error {
	configBytes, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	cfgPath, err := defaultConfigPath()
	if err != nil {
		return err
	}

	// Check if directory exists and create it if necessary
	cfgDir := filepath.Dir(cfgPath)
	if _, err := os.Stat(cfgDir); os.IsNotExist(err) {
		err = os.MkdirAll(cfgDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Write new file or overwrites the old one
	err = ioutil.WriteFile(cfgPath, configBytes, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) AddBundle(bundle string) {
	c.Bundles = append(c.Bundles, bundle)
}

func (c *Config) RemoveBundle(bundle string) error {
	for i, v := range c.Bundles {
		if v == bundle {
			c.Bundles = append(c.Bundles[:i], c.Bundles[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("element %v not found in Bundles", bundle)
}

func (c *Config) SetEnvOverride(application, envName, envValue string) {
	if len(c.EnvOverrides) == 0 {
		c.EnvOverrides = make(map[string]map[string]string)
	}

	if len(c.EnvOverrides[application]) == 0 {
		c.EnvOverrides[application] = make(map[string]string)
	}

	c.EnvOverrides[application][envName] = envValue
}

func (c *Config) UnsetEnvOverride(application, envName string) {
	delete(c.EnvOverrides[application], envName)

	if len(c.EnvOverrides[application]) == 0 {
		delete(c.EnvOverrides, application)
	}
}

func (c *Config) SetRepoOverrides(repository, override string) {
	if len(c.RepoOverrides) == 0 {
		c.RepoOverrides = make(map[string]string)
	}

	c.RepoOverrides[repository] = override
}

func (c *Config) UnsetRepoOverride(repository string) {
	delete(c.RepoOverrides, repository)
}

func (c *Config) SetUsername(username string) {
	c.Username = username
}

func (c *Config) GetUsername() string {
	return c.Username
}

func (c *Config) SetSetupHasRun() {
	c.SetupHasRun = true
}

func (c *Config) HasSetupRun() bool {
	return c.SetupHasRun
}

func (c *Config) SetSpecsRepo(repository string) {
	c.SpecsRepo = repository
}

func (c *Config) GetSpecsRepo() string {
	return c.SpecsRepo
}

func (c *Config) SetVmMemorySize(size int) {
	c.VmMemorySize = size
}

func (c *Config) GetVmMemorySize() int {
	return c.VmMemorySize
}

func (c *Config) SetSSHKeyPath(key string) {
	c.SSHKeyPath = key
}

func (c *Config) GetSSHKeyPath() string {
	return c.SSHKeyPath
}
