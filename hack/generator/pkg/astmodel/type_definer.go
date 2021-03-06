/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// TypeDefiner represents a named type in the output files, and knows how to generate the Go AST
type TypeDefiner interface {

	// Name is the name that will be bound to the type
	Name() *TypeName

	// Type is the type that the name will be bound to
	Type() Type

	// WithDescription adds (or removes!) a description for the defined type
	WithDescription(description *string) TypeDefiner

	// AsDeclarations generates the actual Go declarations
	AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl
}

// FileNameHint returns what a file that contains this definition (if any) should be called
// this is not always used as we might combine multiple definitions into one file
func FileNameHint(def TypeDefiner) string {
	return transformToSnakeCase(def.Name().name)
}
