/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func TestPackageGroupVersion_IncludesGeneratorVersion(t *testing.T) {
	g := NewGomegaWithT(t)

	pkg := NewPackageDefinition("longest-johns", "santiana", "latest-version")

	destDir, err := ioutil.TempDir("", "package_definition_test")
	g.Expect(err).To(BeNil())
	defer os.RemoveAll(destDir)

	fileCount, err := pkg.EmitDefinitions(destDir)
	g.Expect(err).To(BeNil())
	g.Expect(fileCount).To(Equal(0))

	gvFile := filepath.Join(destDir, "groupversion_info_gen.go")
	data, err := ioutil.ReadFile(gvFile)
	g.Expect(err).To(BeNil())
	g.Expect(string(data)).To(ContainSubstring("// Generator version: latest-version"))
}
