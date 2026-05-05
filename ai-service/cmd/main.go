package main

import (
	aimodel "ai-service/internal/ai-model"
	"ai-service/pkg/env"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)


func main(){
	env.LoadEnv()

	groqAi := aimodel.NewGroqClient()
	if err != nilfunc {
		
	}
}