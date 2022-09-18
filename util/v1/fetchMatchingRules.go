package v1

import (
	"fmt"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

type Fact struct {
	LastUpdate int
}

func MatchingRules() {
	fact := &Fact{
		LastUpdate: 6000,
	}

	drls := `
	rule ResetTime "Reset the lastUpdate time" {
    when
        Fact.LastUpdate < Now()
    then
        Fact.LastUpdate = Now();
	}
	`

	kb, dataCtx := RuleEngineClient(drls, fact)
	e := engine.NewGruleEngine()
	ruleEntries, err := e.FetchMatchingRules(dataCtx, kb)
	if err != nil {
		panic(err)
	}
	fmt.Println(ruleEntries)
}
