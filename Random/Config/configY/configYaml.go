package yaml

import (
    // "fmt"
    // "log"
    // "ioutil"
)

type ConfigYaml struct {
    Server struct {
        Port string `yaml:"port"`
        Host string `yaml:"host"`
    } `yaml:"server"`
    Database struct {
        Username string `yaml:"user"`
        Password string `yaml:"pass"`
    } `yaml:"database"`
}

/*
func (c *ConfigYaml) GetConf() *ConfigYaml {

    yamlFile, err := ioutil.ReadFile("conf.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    return c
}
func ConfigYaml(){
	f, err := os.Open("config.yml")
	if err != nil {
   	 	processError(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
    	processError(err)
	}

}
*/
