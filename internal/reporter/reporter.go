package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ar-ease/performance-monitor/internal/analyzer"
)

// GenerateReport creates a report in the specified format
func GenerateReport(result *analyzer.AnalysisResult, format string) error {
	switch format {
	case "console":
		return generateConsoleReport(result)
	case "json":
		return generateJSONReport(result)
	case "html":
		return generateHTMLReport(result)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}

func generateConsoleReport(result *analyzer.AnalysisResult) error {
	fmt.Println("\n📊 Performance Analysis Report")
	fmt.Println("==============================")
	fmt.Printf("Project: %s (%s)\n", result.ProjectInfo.Type, result.ProjectInfo.Framework)
	fmt.Printf("Analyzed at: %s\n\n", result.Timestamp.Format(time.RFC3339))

	if result.BuildMetrics != nil {
		fmt.Println("🏗️  Build Metrics:")
		fmt.Printf("  Build Time: %v\n", result.BuildMetrics.BuildTime)
		fmt.Printf("  Bundle Size: %.2f MB\n", float64(result.BuildMetrics.BundleSize)/(1024*1024))
		fmt.Printf("  Dependencies: %d\n", result.BuildMetrics.Dependencies)
		if len(result.BuildMetrics.Warnings) > 0 {
			fmt.Printf("  Warnings: %d\n", len(result.BuildMetrics.Warnings))
		}
		fmt.Println()
	}

	if result.RuntimeMetrics != nil {
		fmt.Println("⚡ Runtime Metrics:")
		fmt.Printf("  Startup Time: %v\n", result.RuntimeMetrics.StartupTime)
		fmt.Printf("  Memory Usage: %.2f MB\n", float64(result.RuntimeMetrics.MemoryUsage)/(1024*1024))
		fmt.Printf("  CPU Usage: %.2f%%\n", result.RuntimeMetrics.CPUUsage)
		fmt.Println()
	}

	if result.StaticMetrics != nil {
		fmt.Println("🔍 Static Analysis:")
		fmt.Printf("  Lines of Code: %d\n", result.StaticMetrics.LinesOfCode)
		fmt.Printf("  Complexity: %d\n", result.StaticMetrics.Complexity)
		fmt.Printf("  Test Coverage: %.2f%%\n", result.StaticMetrics.TestCoverage)
		fmt.Println()
	}

	return nil
}

func generateJSONReport(result *analyzer.AnalysisResult) error {
	filename := fmt.Sprintf("performance-report-%s.json",
		result.Timestamp.Format("2006-01-02-15-04-05"))

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create JSON report file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(result); err != nil {
		return fmt.Errorf("failed to encode JSON report: %w", err)
	}

	fmt.Printf("📄 JSON report saved to: %s\n", filename)
	return nil
}

func generateHTMLReport(result *analyzer.AnalysisResult) error {
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
    <title>Performance Report - %s</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background-color: #f5f5f5; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        .header { background: #2c3e50; color: white; padding: 20px; border-radius: 8px; margin-bottom: 20px; }
        .metric-section { margin: 20px 0; }
        .metric-card { background: #f8f9fa; padding: 15px; border-radius: 6px; margin: 10px 0; }
        .metric-value { font-size: 1.5em; font-weight: bold; color: #3498db; }
        .metric-label { color: #7f8c8d; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Performance Analysis Report</h1>
            <p>Project: %s (%s)</p>
            <p>Generated: %s</p>
        </div>
        
        %s
    </div>
</body>
</html>
`

	var sections string

	if result.BuildMetrics != nil {
		sections += fmt.Sprintf(`
        <div class="metric-section">
            <h3>🏗️ Build Metrics</h3>
            <div class="metric-card">
                <div class="metric-label">Build Time</div>
                <div class="metric-value">%v</div>
            </div>
            <div class="metric-card">
                <div class="metric-label">Bundle Size</div>
                <div class="metric-value">%.2f MB</div>
            </div>
            <div class="metric-card">
                <div class="metric-label">Dependencies</div>
                <div class="metric-value">%d</div>
            </div>
        </div>
		`, result.BuildMetrics.BuildTime,
			float64(result.BuildMetrics.BundleSize)/(1024*1024),
			result.BuildMetrics.Dependencies)
	}

	filename := fmt.Sprintf("performance-report-%s.html",
		result.Timestamp.Format("2006-01-02-15-04-05"))

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create HTML report file: %w", err)
	}
	defer file.Close()

	content := fmt.Sprintf(htmlTemplate,
		result.ProjectInfo.Type,
		result.ProjectInfo.Type,
		result.ProjectInfo.Framework,
		result.Timestamp.Format(time.RFC3339),
		sections)

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write HTML report: %w", err)
	}

	fmt.Printf("📄 HTML report saved to: %s\n", filename)
	return nil
}
