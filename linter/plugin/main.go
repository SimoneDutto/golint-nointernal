package main

import (
"example.com/custom-linter/pkg/analyzer"
"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// AnalyzerPlugin is the exported variable that golangci-lint looks for.
var AnalyzerPlugin analyzerPlugin

// GetAnalyzers returns the analyzers to be run.
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.NoInternalTypesAnalyzer,
	}
}
