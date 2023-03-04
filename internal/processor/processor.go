package processor

import (
   "fmt"
   "gopkg.in/yaml.v3"
)

type MailProcessor interface {
   Read()
   Send()
}

type Processor struct {
   MailProcessor MailProcessor
}

type Kind string

const (
   SMTPIMAP Kind = "SMTPIMAP"
   MSGRAPH  Kind = "MSGRAPH"
)

var constructors = make(map[Kind]func() interface{}, 10)
var casts = make(map[Kind]func(interface{}) MailProcessor, 10)

func Register(k Kind, f func() interface{}, c func(interface{}) MailProcessor) {
   constructors[k] = f
   casts[k] = c
}

type TypeDef struct {
   Kind Kind `yaml:"Kind"`
}

func (p *Processor) UnmarshalYAML(value *yaml.Node) (err error) {
   //  Decode into a temp struct to get the Kind
   td := &TypeDef{}
   err = value.Decode(td)
   if err != nil {
      return
   }

   if td.Kind == "" {
      err = fmt.Errorf("missing Kind")
      return
   }

   // Get the constructor for the Kind
   constructor, ok := constructors[td.Kind]
   if !ok {
      err = fmt.Errorf("missing constructor: %s", td.Kind)
      return
   }

   // Get the cast for the Kind
   cast, ok := casts[td.Kind]
   if !ok {
      err = fmt.Errorf("missing cast: %s", td.Kind)
      return
   }

   // Get an instance of the Kind
   tp := constructor()
   // Decode into the Kind's instance
   err = value.Decode(tp)
   if err != nil {
      return
   }

   // Cast the decoded Kind instance to a Processor
   mp := cast(tp)
   p.MailProcessor = mp

   return
}
