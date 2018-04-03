// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/go-rotate"
	"github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// UnitTests runs the unit test suite.
func UnitTest() error { return mageextras.UnitTest() }

// IntegrationTest executes the integration test suite.
func IntegrationTest() error {
	mg.Deps(Install)

	exampleFilename := "example.txt"

	exampleFileRot13, err := os.Open(exampleFilename)

	if err != nil {
		return err
	}

	var rot13Out bytes.Buffer

	cmdRot13 := exec.Command("rot13")
	cmdRot13.Stdin = bufio.NewReader(exampleFileRot13)
	cmdRot13.Stdout = bufio.NewWriter(&rot13Out)
	cmdRot13.Stderr = os.Stderr

	if err := cmdRot13.Run(); err != nil {
		return err
	}

	rot13Expected := "Cebsrffbe ZpTbantnyy ghearq vagb n png.\n"

	rot13Observed := rot13Out.String()

	if rot13Observed != rot13Expected {
		return fmt.Errorf("Got %s, expected %s", rot13Observed, rot13Expected)
	}

	return nil
}

// Text runs unit and integration tests.
func Test() error { mg.Deps(UnitTest); mg.Deps(IntegrationTest); return nil }

// CoverHTML denotes the HTML formatted coverage filename.
var CoverHTML = "cover.html"

// CoverProfile denotes the raw coverage data filename.
var CoverProfile = "cover.out"

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML() error { mg.Deps(CoverageProfile); return mageextras.CoverageHTML(CoverHTML, CoverProfile) }

// CoverageProfile generates raw coverage data.
func CoverageProfile() error { return mageextras.CoverageProfile(CoverProfile) }

// GoVet runs go tool vet.
func GoVet() error { return mageextras.GoVet("-shadow") }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("rotate-%s", rotate.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/go-rotate"

// Goxcart cross-compiles Go binaries with additional targets enabled.
func Goxcart() error {
	return mageextras.Goxcart(
		artifactsPath,
		"-repo",
		repoNamespace,
		"-banner",
		portBasename,
	)
}

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Goxcart); return mageextras.Archive(portBasename, artifactsPath) }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("rot13") }

// CleanCoverage deletes coverage data.
func CleanCoverage() error {
	if err := os.RemoveAll(CoverHTML); err != nil {
		return err
	}

	return os.RemoveAll(CoverProfile)
}

// Clean deletes artifacts.
func Clean() error { mg.Deps(CleanCoverage); return os.RemoveAll(artifactsPath) }
