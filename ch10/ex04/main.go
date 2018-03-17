package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Package struct {
	ImportPath string
	Name       string
	Deps       []string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "Usage: need to add package directory.\n")
		os.Exit(1)
	}
	pkgInfo, err := getPkgInfos(os.Args[1:])
	if err != nil {
		log.Fatalf("Failed to get packages names: %s", err)
	}

	fmt.Print("Originals: ")
	printImportPath(pkgInfo)

	pkgDependencies, err := getPkgDependencies(pkgInfo)
	if err != nil {
		log.Fatalf("Failed to get packages names: %s", err)
	}
	if len(pkgDependencies) < 1 {
		fmt.Println("No dependencies")
		return
	}
	fmt.Print("Dependencies: ")
	printImportPath(pkgDependencies)
}
func printImportPath(pkg []Package) {
	for _, p := range pkg {
		fmt.Printf("%s ", p.ImportPath)
	}
	fmt.Println()
}

func getPkgInfos(pkges []string) ([]Package, error) {
	params := append([]string{"list", "-e", "-json"}, pkges...)
	out, err := exec.Command("go", params...).Output()
	if err != nil {
		return nil, err
	}
	pkgInfos := []Package{}
	d := json.NewDecoder(bytes.NewReader(out))
	for {
		var p Package
		err := d.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		pkgInfos = append(pkgInfos, p)
	}
	return pkgInfos, nil
}

func getPkgDependencies(original []Package) ([]Package, error) {
	originalImportPath := []string{}
	for _, p := range original {
		originalImportPath = append(originalImportPath, p.ImportPath)
	}

	params := []string{"list", "-e", "-json", "..."}
	out, err := exec.Command("go", params...).Output()
	if err != nil {
		return nil, err
	}
	pkgDependencies := []Package{}
	d := json.NewDecoder(bytes.NewReader(out))

	for {
		var p Package
		err := d.Decode(&p)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if isDependent(&p, originalImportPath) {
			pkgDependencies = append(pkgDependencies, p)
		}
	}
	return pkgDependencies, nil
}

func isDependent(p *Package, dirs []string) bool {
loop:
	for _, dir := range dirs {
		for _, deps := range p.Deps {
			if deps == dir {
				continue loop
			}
		}
		return false
	}
	return true
}
