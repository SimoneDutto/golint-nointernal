package analyzer

import (
"go/ast"
"go/token"
"go/types"
"strings"

"golang.org/x/tools/go/analysis"
)

var NoInternalTypesAnalyzer = &analysis.Analyzer{
	Name: "nointernaltypes",
	Doc:  "checks for exported functions using types from internal packages",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	currentPkg := pass.Pkg.Path()

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
fn, ok := n.(*ast.FuncDecl)
if !ok {
return true
}

if !fn.Name.IsExported() {
				return true
			}

			// We need the function object to get the signature
			obj, ok := pass.TypesInfo.Defs[fn.Name]
			if !ok {
				return true
			}
			
			sig, ok := obj.Type().(*types.Signature)
			if !ok {
				return true
			}

			prefix := "exported func " + fn.Name.Name
			if fn.Recv != nil {
				prefix = "exported method " + fn.Name.Name
			}

			// Check params
			checkTuple(pass, fn.Pos(), sig.Params(), prefix, currentPkg)
			// Check results
			checkTuple(pass, fn.Pos(), sig.Results(), prefix, currentPkg)

			return true
		})
	}
	return nil, nil
}

func checkTuple(pass *analysis.Pass, pos token.Pos, tuple *types.Tuple, prefix string, currentPkg string) {
	if tuple == nil {
		return
	}
	for i := 0; i < tuple.Len(); i++ {
		v := tuple.At(i)
		if isInternal(v.Type(), currentPkg) {
			pass.Reportf(pos, "%s uses internal type %s", prefix, v.Type().String())
		}
	}
}

func isInternal(t types.Type, currentPkg string) bool {
	// Unpack pointers, slices, maps, channels, arrays
	switch split := t.(type) {
	case *types.Pointer:
		return isInternal(split.Elem(), currentPkg)
	case *types.Slice:
		return isInternal(split.Elem(), currentPkg)
	case *types.Array:
		return isInternal(split.Elem(), currentPkg)
	case *types.Map:
		return isInternal(split.Key(), currentPkg) || isInternal(split.Elem(), currentPkg)
	case *types.Chan:
		return isInternal(split.Elem(), currentPkg)
	case *types.Named:
		// Check the package of the named type
		obj := split.Obj()
		if obj != nil && obj.Pkg() != nil {
			path := obj.Pkg().Path()
			if isInternalPath(path, currentPkg) {
				return true
			}
		}
		// Also check type arguments (generics) if any
		if tArgs := split.TypeArgs(); tArgs != nil {
			for i := 0; i < tArgs.Len(); i++ {
				if isInternal(tArgs.At(i), currentPkg) {
					return true
				}
			}
		}
	}
	return false
}

func isInternalPath(path string, currentPkg string) bool {
	// Check if path is strictly a sub-internal package of currentPkg
	// Pattern 1: currentPkg + "/internal"
	// Pattern 2: currentPkg + "/internal/"...
	
	internalRoot := currentPkg + "/internal"
	if path == internalRoot {
		return true
	}
	if strings.HasPrefix(path, internalRoot+"/") {
		return true
	}
	
	return false
}
