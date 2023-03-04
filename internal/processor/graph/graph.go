package graph

import (
   "InstantiateInterfacesFromYamlConfig/internal/processor"
   "fmt"
)

func init() {
   processor.Register(processor.MSGRAPH, MakeOne, Cast)
}

func init() {
   fmt.Println("graph.init")
}

type GraphStruct struct {
   Kind    processor.Kind `yaml:"Kind"`
   MSGraph string         `yaml:"MSGraph"`
}

type Graph GraphStruct

func (g *Graph) Send() {

}

func (g *Graph) Read() {

}

func MakeOne() interface{} {
   return &Graph{}
}

func Cast(i interface{}) processor.MailProcessor {
   g := i.(*Graph)
   return g
}
