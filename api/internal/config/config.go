package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/storyprotocol/protocol-api/pkg/logger"
	validator "gopkg.in/validator.v2"
	yaml "gopkg.in/yaml.v2"
)

type AppConfig struct {
	AppID                                  string   `yaml:"app_id"`
	Port                                   int64    `yaml:"port"`
	Env                                    string   `yaml:"env"`
	TheGraphAlphaEndpoint                  string   `yaml:"the_graph_alpha_endpoint"`
	TheGraphSepoliaBetaEndpoint            string   `yaml:"the_graph_beta_endpoint"`
	TheGraphSepolia0xSplitBetaEndpoint     string   `yaml:"the_graph_0x_split_beta_endpoint"`
	TheGraphRenaissanceBetaEndpoint        string   `yaml:"the_graph_renaissance_endpoint"`
	TheGraphRenaissance0xSplitBetaEndpoint string   `yaml:"the_graph_renaissance_0x_split_beta_endpoint"`
	ApiKeys                                []string `yaml:"api_keys"`
	OpenChainLookupEndpoint                string   `yaml:"openchain_lookup_endpoint"`
	ZettablockAPIKey                       string   `yaml:"zettablock_api_key"`
	ZettablockTransactionEndpoint          string   `yaml:"zettablock_trx_endpoint"`
	ZettablockACLEndpoint                  string   `yaml:"zettablock_acl_endpoint"`
	ZettablockCollectionEndpoint           string   `yaml:"zettablock_collection_endpoint"`
	ZettablockDisputeEndpoint              string   `yaml:"zettablock_dispute_endpoint"`
	ZettablockIPAssetEndpoint              string   `yaml:"zettablock_ipasset_endpoint"`
	ZettablockIPAPolicyEndpoint            string   `yaml:"zettablock_ipapolicy_endpoint"`
	ZettablockLicenseEndpoint              string   `yaml:"zettablock_license_endpoint"`
	ZettablockLicenseMfpEndpoint           string   `yaml:"zettablock_lmfp_endpoint"`
	ZettablockLicenseOwnerEndpoint         string   `yaml:"zettablock_license_owner_endpoint"`
	ZettablockModuleEndpoint               string   `yaml:"zettablock_module_endpoint"`
	ZettablockPermissionsEndpoint          string   `yaml:"zettablock_permissions_endpoint"`
	ZettablockPolicyEndpoint               string   `yaml:"zettablock_policy_endpoint"`
	ZettablockPolicyFwmEndpoint            string   `yaml:"zettablock_policyfwm_endpoint"`
	ZettablockRoyaltyPayEndpoint           string   `yaml:"zettablock_royalty_pay_endpoint"`
	ZettablockRoyaltyPolicyEndpoint        string   `yaml:"zettablock_royalty_policy_endpoint"`
	Zettablock0xSplitsEndpoint             string   `yaml:"zettablock_splits_endpoint"`
}

var (
	cfgFlag        = flag.String("config", "config.yaml", "config file")
	configInstance *AppConfig
)

// GetConfig loads the config and return cached instance once loaded
func GetConfig() (*AppConfig, error) {
	if configInstance != nil {
		return configInstance, nil
	}
	var cfg AppConfig

	cfgFiles := strings.Split(*cfgFlag, ",")
	logger.Info(cfgFiles)

	if err := LoadFiles(&cfg, cfgFiles...); err != nil {
		logger.Fatalf("Failed to load config file: %v", err)
	}

	configInstance = &cfg
	return &cfg, nil
}

// LoadFiles combines configs in the files and validate the configs
func LoadFiles(config interface{}, fileNames ...string) error {
	if len(fileNames) == 0 {
		return errors.New("no config files to load")
	}

	for _, file := range fileNames {
		fmt.Println("Load config file: " + file)
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(data, config); err != nil {
			return err
		}
	}

	if err := validator.Validate(config); err != nil {
		return err.(validator.ErrorMap)
	}
	return nil
}
