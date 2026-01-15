package strictpkg

import "example.com/strictpkg/internal/deep"

// This uses a type from 'example.com/strictpkg/internal/deep'.
// Current pkg is 'example.com/strictpkg'.
// Target pkg is 'example.com/strictpkg/internal/deep'.
// HasPrefix(target, current + "/internal") is TRUE.
// So this should fail.

func ExposeDeep(d deep.DeepType) {} // want "exported func ExposeDeep uses internal type example.com/strictpkg/internal/deep.DeepType"
