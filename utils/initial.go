package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"tapi-controller/models"
	"tapi-controller/token"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const (
	dbDataPath = "/db_data/context.json"
	configPath = "."
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type Config struct {
	Address            string `mapstructure:"SERVER_ADDRESS"`
	Port               int    `mapstructure:"SERVER_PORT"`
	Usewrname          string `mapstructure:"USERNAME"`
	Password           string `mapstructure:"PASSWORD"`
	PrivateKeyFilename string `mapstructure:"PRIVATE_KEY_FILENAME"`
	PublicKeyFilename  string `mapstructure:"PUBLIC_KEY_FILENAME"`
}

type InMemoryDb struct {
	TapiCommonOperationalState     models.TapiCommonOperationalState
	TapiCommonLifecycleState       models.TapiCommonLifecycleState
	TapiCommonAdministrativeState  models.TapiCommonAdministrativeState
	TapiCommonLayerProtocolName    models.TapiCommonLayerProtocolName
	TapiCommonTerminationDirection models.TapiCommonTerminationDirection
	TapiCommonTerminationState     models.TapiCommonTerminationState
	TapiCommonPortRole             models.TapiCommonPortRole
	TapiCommonPortDirection        models.TapiCommonPortDirection
	TapiCommonForwardingDirection  models.TapiCommonForwardingDirection
	TapiCommonContext              models.TapiCommonContext
}

func InitMemoryDb() InMemoryDb {
	dbDataPathRelative := fmt.Sprintf("%s%s", basepath, dbDataPath)
	log.Println("Start to init in memory DB")
	jsonFile, err := os.Open(dbDataPathRelative)
	if err != nil {
		log.Panicf("Cannot open file: %s, with error: %x", dbDataPathRelative, err)
	}
	byteJson, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Panicf("Cannot read file: %s, with error: %x", dbDataPathRelative, err)
	}
	var context models.TapiCommonContext
	json.Unmarshal(byteJson, &context)
	res := InMemoryDb{
		TapiCommonOperationalState:     models.InitTapiCommonOperationalState(),
		TapiCommonLifecycleState:       models.InitTapiCommonLifecycleState(),
		TapiCommonAdministrativeState:  models.InitTapiCommonAdministrativeState(),
		TapiCommonLayerProtocolName:    models.InitTapiCommonLayerProtocolName(),
		TapiCommonTerminationDirection: models.InitTapiCommonTerminationDirection(),
		TapiCommonTerminationState:     models.InitTapiCommonTerminationState(),
		TapiCommonPortRole:             models.InitTapiCommonPortRole(),
		TapiCommonPortDirection:        models.InitTapiCommonPortDirection(),
		TapiCommonForwardingDirection:  models.InitTapiCommonForwardingDirection(),
		TapiCommonContext:              context,
	}
	log.Println("In memroy DB created")
	return res

}

func LoadConfig() (config Config, err error) {
	log.Println("Start to load config")
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	log.Println("config loaded")
	err = viper.Unmarshal(&config)
	return
}

func LoadConfigForTest() (config Config, err error) {
	log.Println("Start to load config")
	viper.AddConfigPath(configPath)
	viper.SetConfigName("../.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	log.Println("config loaded")
	err = viper.Unmarshal(&config)
	return
}

func InitTokenMaker(config Config) (token.TokenMaker, error) {
	log.Println("Load RSA keys")
	privateKeyPath := fmt.Sprintf("%s/keys/%s", basepath, config.PrivateKeyFilename)
	publicKeyPath := fmt.Sprintf("%s/keys/%s", basepath, config.PublicKeyFilename)

	prKeyFile, err := os.Open(privateKeyPath)
	if err != nil {
		message := fmt.Sprintf("cannot open file: %s, with error: %x", privateKeyPath, err)
		log.Println(message)
		return nil, errors.New(message)
	}
	prKeybytes, err := io.ReadAll(prKeyFile)
	if err != nil {
		message := fmt.Sprintf("cannot read file: %s, with error: %x", privateKeyPath, err)
		log.Println(message)
		return nil, errors.New(message)
	}
	prKey, err := jwt.ParseRSAPrivateKeyFromPEM(prKeybytes)
	if err != nil {
		message := fmt.Sprintf("cannot parse file: %s, with error: %x", privateKeyPath, err)
		log.Println(message)
		return nil, errors.New(message)
	}
	log.Println("loaded private key")
	pubKeyFile, err := os.Open(publicKeyPath)
	if err != nil {
		message := fmt.Sprintf("cannot open file: %s, with error: %x", publicKeyPath, err)
		log.Println(message)
		return nil, errors.New(message)
	}
	pubKeybytes, err := io.ReadAll(pubKeyFile)
	if err != nil {
		message := fmt.Sprintf("cannot read file: %s, with error: %x", publicKeyPath, err)
		log.Println(message)
		return nil, errors.New(message)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeybytes)
	if err != nil {
		message := fmt.Sprintf("cannot parse file: %s, with error: %x", publicKeyPath, err)
		log.Println(message)
		return nil, errors.New(message)
	}
	log.Println("loaded public key")
	return &token.JwtMaker{
		PrivateKey: prKey,
		PublicKey:  pubKey,
	}, nil
}
