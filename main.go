package main

import (
	"fmt"
	"strings"

	c "nishanth.io/vipersss/config"

	"github.com/spf13/viper"
)

func main() {
	// Set the file name of the configurations file
	vipe := viper.NewWithOptions(viper.KeyDelimiter("_"))
	vipe.SetConfigName("config")
	vipe.SetConfigType("yml")

	// Set the path to look for the configurations file
	vipe.AddConfigPath(".")

	vipe.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Enable VIPER to read Environment Variables
	vipe.AutomaticEnv()

	var configuration c.Configurations

	if err := vipe.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	vipe.SetDefault("database.dbname", "test_db")

	err := vipe.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println(vipe.AllKeys())
	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	//this should have the same value as SERVER_PORT
	fmt.Println("Port is\t\t", configuration.Server.Port)
	fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
	fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Database is\t", vipe.GetString("database.dbname"))
	fmt.Println("server Port is\t\t", vipe.GetInt("server.port"))
	fmt.Println("server Port is\t\t", vipe.GetInt("SERVER_PORT"))
	fmt.Println("EXAMPLE_PATH is\t", vipe.GetString("EXAMPLE_PATH"))
	fmt.Println("EXAMPLE_VAR is\t", vipe.GetString("EXAMPLE_VAR"))

}
