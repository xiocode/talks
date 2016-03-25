viper.SetConfigName("config")
viper.SetConfigType("yaml")
viper.AddConfigPath("config/")
viper.AddConfigPath(".")
err := viper.ReadInConfig() // Find and read the config file
if err != nil {             // Handle errors reading the config file
    panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("Config file changed:", e.Name)
})
