package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ar-ease/performance-monitor/internal/analyzer"
	"github.com/ar-ease/performance-monitor/internal/detector"
	"github.com/ar-ease/performance-monitor/internal/reporter"
)

type Config struct {
	All      bool
	Build    bool
	Runtime  bool
	Static   bool
	Memory   bool
	Network  bool
	Output   string
	Duration string
	Watch    bool
	CI       bool
}

func main() {
	config := parseFlags()

	// Get current working directory
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}

	fmt.Println("🚀 Performance Monitor CLI (pmon)")
	fmt.Println("================================")

	// Detect project type
	projectInfo, err := detector.DetectProject(workDir)
	if err != nil {
		log.Fatal("❌ Error detecting project:", err)
	}

	fmt.Printf("🔍 Detected: %s project\n", projectInfo.Type)
	fmt.Printf("📁 Framework: %s\n", projectInfo.Framework)
	fmt.Printf("🔧 Build Tool: %s\n", projectInfo.BuildTool)
	fmt.Printf("📍 Root Path: %s\n", projectInfo.RootPath)

	if len(projectInfo.Dependencies) > 0 {
		fmt.Printf("📦 Dependencies: %d packages\n", len(projectInfo.Dependencies))
	}

	if len(projectInfo.Scripts) > 0 {
		fmt.Printf("📝 Scripts: %d available\n", len(projectInfo.Scripts))
		if config.All {
			fmt.Println("\nAvailable scripts:")
			for name, script := range projectInfo.Scripts {
				fmt.Printf("  • %s: %s\n", name, script)
			}
		}
	}

	if config.All {
		fmt.Println("\n🚀 Running full performance analysis...")
		result, err := analyzer.RunAnalysis(projectInfo, config)
		if err != nil {
			log.Fatal("❌ Error running analysis:", err)
		}

		err = reporter.GenerateReport(result, config.Output)
		if err != nil {
			log.Fatal("❌ Error generating report:", err)
		}
	} else {
		if config.Build {
			fmt.Println("\n🏗️  Running build performance analysis...")
			// TODO: Implement build analysis
		}
		if config.Runtime {
			fmt.Println("\n⚡ Running runtime performance analysis...")
			// TODO: Implement runtime analysis
		}
		if config.Static {
			fmt.Println("\n🔍 Running static code analysis...")
			// TODO: Implement static analysis
		}
		if config.Memory {
			fmt.Println("\n🧠 Running memory profiling...")
			// TODO: Implement memory profiling
		}
		if config.Network {
			fmt.Println("\n🌐 Running network analysis...")
			// TODO: Implement network analysis
		}
	}
}

func runFullAnalysis(projectInfo *detector.ProjectInfo, config Config) {
	fmt.Println("📊 Analysis includes:")
	fmt.Println("  • Bundle size analysis")
	fmt.Println("  • Build performance")
	fmt.Println("  • Runtime metrics")
	fmt.Println("  • Memory usage")
	fmt.Println("  • Code quality")

	// TODO: Implement actual analysis based on project type
	switch projectInfo.Framework {
	case "Next.js":
		fmt.Println("\n🔥 Next.js specific optimizations:")
		fmt.Println("  • SSR/SSG performance")
		fmt.Println("  • Route optimization")
		fmt.Println("  • Image optimization")
	case "React":
		fmt.Println("\n⚛️  React specific optimizations:")
		fmt.Println("  • Component render times")
		fmt.Println("  • Virtual DOM operations")
		fmt.Println("  • Hook optimization")
	case "Vue.js":
		fmt.Println("\n💚 Vue.js specific optimizations:")
		fmt.Println("  • Component performance")
		fmt.Println("  • Reactivity system")
	}

	fmt.Println("\n✅ Analysis complete! (TODO: Implement actual analysis)")
}

func parseFlags() Config {
	var config Config

	flag.BoolVar(&config.All, "a", false, "Run all performance analysis")
	flag.BoolVar(&config.All, "all", false, "Run all performance analysis")
	flag.BoolVar(&config.Build, "b", false, "Build performance analysis")
	flag.BoolVar(&config.Build, "build", false, "Build performance analysis")
	flag.BoolVar(&config.Runtime, "r", false, "Runtime performance analysis")
	flag.BoolVar(&config.Runtime, "runtime", false, "Runtime performance analysis")
	flag.BoolVar(&config.Static, "s", false, "Static code analysis")
	flag.BoolVar(&config.Static, "static", false, "Static code analysis")
	flag.BoolVar(&config.Memory, "m", false, "Memory profiling")
	flag.BoolVar(&config.Memory, "memory", false, "Memory profiling")
	flag.BoolVar(&config.Network, "n", false, "Network analysis")
	flag.BoolVar(&config.Network, "network", false, "Network analysis")
	flag.StringVar(&config.Output, "output", "console", "Output format (console, json, html)")
	flag.StringVar(&config.Duration, "duration", "10s", "Monitoring duration")
	flag.BoolVar(&config.Watch, "watch", false, "Continuous monitoring")
	flag.BoolVar(&config.CI, "ci", false, "CI-friendly output")

	flag.Parse()

	// If no specific flags, default to all
	if !config.Build && !config.Runtime && !config.Static && !config.Memory && !config.Network {
		config.All = true
	}

	return config
}
