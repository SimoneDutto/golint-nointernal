package analyzer_test

import (
	"testing"

	"example.com/custom-linter/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	testdata := analysistest.TestData()
	// Run on both packages
	analysistest.Run(t, testdata, analyzer.NoInternalTypesAnalyzer, "example.com/mylib/public", "example.com/strictpkg")
}
