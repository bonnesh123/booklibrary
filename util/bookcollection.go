package util

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type Book struct {
	ID       int    `mapstructure:"id"`
	Name     string `mapstructure:"name"`
	Author   string `mapstructure:"author"`
	Labels   string `mapstructure:"labels"`
	Quantity int    `mapstructure:"quantity"`
}

type Collection struct {
	BookList []Book `mapstructure:"booklist"`
}

var BookCollection *Collection

func (i Collection) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}

func LoadCollection(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("collection")
	viper.SetConfigType("json")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("Error in reading config", err)
		return
	}

	if err = viper.Unmarshal(&BookCollection); err != nil {
		fmt.Println("Unable to decode into struct", err)
		return
	}
	return
}
