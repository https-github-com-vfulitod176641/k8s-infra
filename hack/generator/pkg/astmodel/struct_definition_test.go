/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_NewStructDefinition_GivenValues_InitializesFields(t *testing.T) {
	g := NewGomegaWithT(t)

	const name = "demo"
	const group = "group"
	const version = "2020-01-01"
	fullNameField := createStringField("fullName", "Full legal name")
	familyNameField := createStringField("familyName", "Shared family name")
	knownAsField := createStringField("knownAs", "Commonly known as")

	ref := NewTypeName(*NewLocalPackageReference(group, version), name)
	definition := NewStructDefinition(ref, NewStructType().WithFields(fullNameField, familyNameField, knownAsField))

	definitionGroup, definitionPackage, err := definition.Name().PackageReference.GroupAndPackage()
	g.Expect(err).To(BeNil())

	g.Expect(definition.Name().name).To(Equal(name))
	g.Expect(definitionGroup).To(Equal(group))
	g.Expect(definitionPackage).To(Equal(version))
	g.Expect(definition.structType.fields).To(HaveLen(3))
}

func Test_StructDefinitionAsAst_GivenValidStruct_ReturnsNonNilResult(t *testing.T) {
	g := NewGomegaWithT(t)

	ref := NewTypeName(*NewLocalPackageReference("group", "2020-01-01"), "name")
	field := NewStructDefinition(ref, NewStructType())
	node := field.AsDeclarations(nil)

	g.Expect(node).NotTo(BeNil())
}

func createStringField(name string, description string) *FieldDefinition {
	return NewFieldDefinition(FieldName(name), name, StringType).WithDescription(&description)
}

func createIntField(name string, description string) *FieldDefinition {
	return NewFieldDefinition(FieldName(name), name, IntType).WithDescription(&description)
}
