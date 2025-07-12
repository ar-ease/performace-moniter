package detector

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ProjectInfo struct {
	Type         string            `json:"type"`
	Framework    string            `json:"framework"`
	BuildTool    string            `json:"buildTool"`
	Language     string            `json:"language"`
	Dependencies map[string]string `json:"dependencies"`
	Scripts      map[string]string `json:"scripts"`
	RootPath     string            `json:"rootPath"`
}

type PackageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Scripts         map[string]string `json:"scripts"`
}

func DetectProject(rootPath string) (*ProjectInfo, error) {
	info := &ProjectInfo{
		RootPath:     rootPath,
		Dependencies: make(map[string]string),
		Scripts:      make(map[string]string),
	}

	// Check for different project types
	if hasFile(rootPath, "package.json") {
		return detectJSProject(rootPath, info)
	} else if hasFile(rootPath, "go.mod") {
		return detectGoProject(rootPath, info)
	} else if hasFile(rootPath, "requirements.txt") || hasFile(rootPath, "pyproject.toml") {
		return detectPythonProject(rootPath, info)
	} else if hasFile(rootPath, "pom.xml") || hasFile(rootPath, "build.gradle") {
		return detectJavaProject(rootPath, info)
	}

	return nil, fmt.Errorf("unknown project type - no recognized project files found")
}

func detectJSProject(rootPath string, info *ProjectInfo) (*ProjectInfo, error) {
	info.Type = "JavaScript"
	info.Language = "JavaScript"

	// Read package.json
	packagePath := filepath.Join(rootPath, "package.json")
	data, err := os.ReadFile(packagePath)
	if err != nil {
		return nil, fmt.Errorf("error reading package.json: %v", err)
	}

	var pkg PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("error parsing package.json: %v", err)
	}

	// Merge dependencies
	for k, v := range pkg.Dependencies {
		info.Dependencies[k] = v
	}
	for k, v := range pkg.DevDependencies {
		info.Dependencies[k] = v
	}
	info.Scripts = pkg.Scripts

	// Detect framework
	info.Framework = detectFramework(info.Dependencies)

	// Detect build tool
	info.BuildTool = detectBuildTool(rootPath, info.Dependencies)

	return info, nil
}

func detectFramework(deps map[string]string) string {
	if _, exists := deps["next"]; exists {
		return "Next.js"
	}
	if _, exists := deps["react"]; exists {
		return "React"
	}
	if _, exists := deps["vue"]; exists {
		return "Vue.js"
	}
	if _, exists := deps["angular"]; exists {
		return "Angular"
	}
	if _, exists := deps["svelte"]; exists {
		return "Svelte"
	}
	return "Vanilla JavaScript"
}

func detectBuildTool(rootPath string, deps map[string]string) string {
	if hasFile(rootPath, "vite.config.js") || hasFile(rootPath, "vite.config.ts") {
		return "Vite"
	}
	if hasFile(rootPath, "webpack.config.js") || hasFile(rootPath, "webpack.config.ts") {
		return "Webpack"
	}
	if _, exists := deps["parcel"]; exists {
		return "Parcel"
	}
	if hasFile(rootPath, "rollup.config.js") {
		return "Rollup"
	}
	return "Unknown"
}

func detectGoProject(rootPath string, info *ProjectInfo) (*ProjectInfo, error) {
	info.Type = "Go"
	info.Language = "Go"
	info.Framework = "Standard Library"
	info.BuildTool = "Go Build"
	return info, nil
}

func detectPythonProject(rootPath string, info *ProjectInfo) (*ProjectInfo, error) {
	info.Type = "Python"
	info.Language = "Python"
	// Could detect Django, Flask, FastAPI, etc.
	return info, nil
}

func detectJavaProject(rootPath string, info *ProjectInfo) (*ProjectInfo, error) {
	info.Type = "Java"
	info.Language = "Java"
	// Could detect Spring, Maven, Gradle, etc.
	return info, nil
}

func hasFile(rootPath, filename string) bool {
	_, err := os.Stat(filepath.Join(rootPath, filename))
	return err == nil
}
