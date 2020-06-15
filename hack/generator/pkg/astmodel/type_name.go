/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/ast"
)

// TypeName is a name associated with another Type (it also is usable as a Type)
type TypeName struct {
	PackageReference PackageReference // Note: This has to be a value and not a ptr because this type is used as the key in a map
	name             string
}

// NewTypeName is a factory method for creating a TypeName
func NewTypeName(pr PackageReference, name string) *TypeName {
	return &TypeName{pr, name}
}

// Name returns the package-local name of the type
func (typeName *TypeName) Name() string {
	return typeName.name
}

// AsType implements Type for TypeName
func (typeName *TypeName) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	// If our package is being referenced, we need to ensure we include a selector for that reference
	if imp, ok := codeGenerationContext.PackageImports()[typeName.PackageReference]; ok {
		return &ast.SelectorExpr{
			X:   ast.NewIdent(imp.PackageName()),
			Sel: ast.NewIdent(typeName.Name()),
		}
	}

	// Sanity assertion that the type we're generating is in the same package that the context is for
	if !codeGenerationContext.currentPackage.Equals(&typeName.PackageReference) {
		panic(fmt.Sprintf(
			"no reference for %v included in package %v",
			typeName.name,
			codeGenerationContext.currentPackage))
	}

	return ast.NewIdent(typeName.name)
}

// References indicates whether this Type includes any direct references to the given Type
func (typeName *TypeName) References(d *TypeName) bool {
	return typeName.Equals(d)
}

// RequiredImports returns all the imports required for this definition
func (typeName *TypeName) RequiredImports() []*PackageReference {
	return []*PackageReference{&typeName.PackageReference}
}

// Equals returns true if the passed type is the same TypeName, false otherwise
func (typeName *TypeName) Equals(t *TypeName) bool {
	return typeName.name == t.name &&
		typeName.PackageReference.Equals(&t.PackageReference)
}
