package smtp

import (
   "InstantiateInterfacesFromYamlConfig/internal/processor"
   "fmt"
)

func init() {
   fmt.Println("smtp.init")
   processor.Register(processor.SMTPIMAP, MakeOne, Cast)
}

type SMTPStruct struct {
   Kind processor.Kind `yaml:"Kind"`
   Smtp string         `yaml:"Smtp"`
}
type SMTP SMTPStruct

func (g *SMTP) Send() {

}

func (g *SMTP) Read() {

}

func MakeOne() interface{} {
   return &SMTP{}
}

func Cast(i interface{}) processor.MailProcessor {
   s := i.(*SMTP)
   return s
}
