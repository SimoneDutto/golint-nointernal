package public

import "example.com/mylib/internal/secret"

// These used to trigger errors, but now shouldn't because secret
// is 'example.com/mylib/internal/secret', and current pkg is 'example.com/mylib/public'.
// internal/secret is NOT a child of public.

func ExportedFunc(s secret.SecretType) {}

func ExportedFuncReturn() secret.SecretType {
return secret.SecretType{}
}

func internalFunc(s secret.SecretType) {}

type MyPublicType struct{}

func (m *MyPublicType) ExportedMethod(s secret.SecretType) {}

func ExportedSlice(s []secret.SecretType) {}

func ExportedMap(m map[string]secret.SecretType) {}
