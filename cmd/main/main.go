package main

import (
   "InstantiateInterfacesFromYamlConfig/internal/processor"
   // Force the init function to run and avoid circular references
   _ "InstantiateInterfacesFromYamlConfig/internal/processor/graph"
   // Force the init function to run and avoid circular references
   _ "InstantiateInterfacesFromYamlConfig/internal/processor/smtp"
   "gopkg.in/yaml.v3"
   "os"
)

type TraceConfig struct {
   Tests []*Test `yaml:"Tests"`
}

type Test struct {
   MTAgent   *processor.Processor `yaml:"MTAgent"`
   TestAgent *processor.Processor `yaml:"TestAgent"`
}

func main() {
   // Load yaml file

}
func loadTraceConfig(configFile string) (*TraceConfig, error) {
   configYaml, err := os.ReadFile(configFile)
   if err != nil {
      return nil, err
   }

   return parseTraceConfig(configYaml)
}

func parseTraceConfig(configYaml []byte) (*TraceConfig, error) {
   var config TraceConfig
   err := yaml.Unmarshal(configYaml, &config)
   if err != nil {
      return nil, err
   }
   return &config, err
}
