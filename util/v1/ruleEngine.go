package v1

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type myFact interface {
	*MyFactKnowledge | *Fact
}

func RuleEngineClient[T myFact](drls string, fact T) (*ast.KnowledgeBase, ast.IDataContext) {
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("MF", fact)
	if err != nil {
		panic(err)
	}

	// creating knowledge library and adding rules into it
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	// use the builder to add the definition to the knowledgeLibrary from a declared resource
	bs := pkg.NewBytesResource([]byte(drls))
	err = ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", bs)
	if err != nil {
		panic(err)
	}

	// to execute a KnowledgeBase
	knowledgeBase := knowledgeLibrary.NewKnowledgeBaseInstance("TutorialRules", "0.0.1")

	return knowledgeBase, dataCtx
}
