package main

import (
   "github.com/davecgh/go-spew/spew"
   "reflect"
   "testing"
)

func Test_parseTraceConfig(t *testing.T) {
   tests := []struct {
      name       string
      configYaml string
      want       *TraceConfig
      wantErr    bool
   }{
      {
         name: "Test",
         configYaml: `
Tests:
  - MTAgent:
      Smtp: "SMTP Processor"
      Kind: "SMTPIMAP"
    TestAgent:
      MSGraph: "MSGraph Processor"
      Kind: "MSGRAPH"
`,
         want:    nil,
         wantErr: false,
      },
   }
   for _, tt := range tests {
      t.Run(tt.name, func(t *testing.T) {
         got, err := parseTraceConfig([]byte(tt.configYaml))
         if (err != nil) != tt.wantErr {
            t.Errorf("parseTraceConfig() error = %v, wantErr %v", err, tt.wantErr)
            return
         }
         if !reflect.DeepEqual(got, tt.want) {
            t.Errorf("parseTraceConfig() got: %+v, want: %v", *got, tt.want)
            spew.Dump(got)
         }
      })
   }
}
