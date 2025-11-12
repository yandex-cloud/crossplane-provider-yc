//go:build generate
// +build generate

/*
Copyright 2025 YANDEX LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Post-generation tool to inject custom validations into CRDs.
// This addresses known issues where controller-gen removes custom x-kubernetes-validations.
// See: https://github.com/crossplane/upjet/issues/78

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	indent20 = "                    "
	indent22 = "                      "
	indent24 = "                        "
)

// injectRoleImmutability injects immutability validation for the role field in IAM member CRDs.
func injectRoleImmutability(crdFile, resourceName string) error {
	fmt.Printf("Processing %s...\n", resourceName)

	// Read the file
	file, err := os.Open(crdFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Check if validation already exists
	for _, line := range lines {
		if strings.Contains(line, "self == oldSelf") {
			fmt.Printf("  ✓ Immutability validation already present in %s\n", resourceName)
			return nil
		}
	}

	// Find the location to inject the validation
	// We're looking for the role field in initProvider section
	inInitProvider := false
	inRoleField := false
	injectionLine := -1

	for i, line := range lines {
		// Track when we enter initProvider
		if strings.Contains(line, "initProvider:") {
			inInitProvider = true
			continue
		}

		// Track when we exit initProvider (next major section at same indentation level)
		if inInitProvider && strings.TrimSpace(line) != "" {
			// Check if we've exited initProvider by looking at indentation
			if !strings.HasPrefix(line, "                ") || // Less than 16 spaces
				(strings.HasPrefix(line, "              ") && !strings.HasPrefix(line, "                ")) {
				if !strings.Contains(line, "properties:") {
					inInitProvider = false
				}
			}
		}

		// Find role field in initProvider
		if inInitProvider && strings.TrimSpace(line) == "role:" {
			inRoleField = true
			continue
		}

		// Find the "type: string" line after role description in initProvider
		if inRoleField && strings.Contains(line, "type: string") && strings.HasPrefix(line, indent20) {
			injectionLine = i
			break
		}
	}

	if injectionLine == -1 {
		return fmt.Errorf("pattern not found - could not locate role field in initProvider")
	}

	// Prepare the validation lines to inject
	validationLines := []string{
		indent20 + "# WARNING: This will be deleted upon generation, see https://github.com/crossplane/upjet/issues/78 -",
		indent20 + "# WARNING: restore this manually!",
		indent20 + "x-kubernetes-validations:",
		indent22 + fmt.Sprintf("- message: 'Role can not be changed after creation. If changes are needed, delete this %s and create a new one.'", resourceName),
		indent24 + "rule: self == oldSelf",
	}

	// Insert the validation lines before the "type: string" line
	newLines := make([]string, 0, len(lines)+len(validationLines))
	newLines = append(newLines, lines[:injectionLine]...)
	newLines = append(newLines, validationLines...)
	newLines = append(newLines, lines[injectionLine:]...)

	// Write the modified content back
	output, err := os.Create(crdFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, line := range newLines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write line: %w", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	fmt.Printf("  ✓ Successfully injected immutability validation into %s\n", resourceName)
	return nil
}

func main() {
	fmt.Println("Post-processing generated CRDs...")

	// Get the current working directory (will be apis/ when run via go:generate)
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting working directory: %v\n", err)
		os.Exit(1)
	}

	// Determine project root (go up one level from apis/)
	projectRoot := filepath.Dir(wd)

	// Determine CRD directory
	crdDir := filepath.Join(projectRoot, "package", "crds")

	// Process FolderIAMMember CRD
	if err := injectRoleImmutability(
		filepath.Join(crdDir, "iam.yandex-cloud.jet.crossplane.io_folderiammembers.yaml"),
		"FolderIAMMember",
	); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing FolderIAMMember: %v\n", err)
		os.Exit(1)
	}

	if err := injectRoleImmutability(
		filepath.Join(crdDir, "iam.yandex-cloud.m.jet.crossplane.io_folderiammembers.yaml"),
		"FolderIAMMember",
	); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing FolderIAMMember: %v\n", err)
		os.Exit(1)
	}

	// Add more IAM member CRDs here if needed in the future
	// if err := injectRoleImmutability(
	// 	filepath.Join(crdDir, "iam.yandex-cloud.jet.crossplane.io_cloudiammembers.yaml"),
	// 	"CloudIAMMember",
	// ); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error processing CloudIAMMember: %v\n", err)
	// 	os.Exit(1)
	// }

	fmt.Println("Post-processing complete!")
}
