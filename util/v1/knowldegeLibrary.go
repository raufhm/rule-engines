package v1

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"time"
)

type MyFactKnowledge struct {
	IntAttribute     int64
	StringAttribute  string
	BooleanAttribute bool
	FloatAttribute   float64
	TimeAttribute    time.Time
	WhatToSay        string
}

func (mf *MyFactKnowledge) GetWhatToSay(sentence string) string {
	return fmt.Sprintf("Let say \"%s\"", sentence)
}
func KnowledgeLibrary() {
	myFact := &MyFactKnowledge{
		IntAttribute:     123,
		StringAttribute:  "Some string value",
		BooleanAttribute: true,
		FloatAttribute:   1.234,
		TimeAttribute:    time.Now(),
	}

	// basic rule as a raw string in DSL
	drls := `
	rule CheckValues "Check the default values" salience 10 {
    when 
        MF.IntAttribute == 123 && MF.StringAttribute == "Some string value"
    then
        MF.WhatToSay = MF.GetWhatToSay("Hello Grule");
        Retract("CheckValues");
	}
	`
	kb, dataCtx := RuleEngineClient(drls, myFact)
	e := engine.NewGruleEngine()
	err := e.Execute(dataCtx, kb)
	if err != nil {
		panic(err)
	}
	fmt.Println(myFact.WhatToSay)
}
