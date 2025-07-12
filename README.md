# Performance Monitor (pmon)

A CLI tool for monitoring and analyzing performance of local projects across different languages and frameworks.

## Features

- 🔍 **Auto-detect project type** (JavaScript, Go, Python, Java)
- 📊 **Framework-specific analysis** (React, Next.js, Vue.js, etc.)
- ⚡ **Multiple analysis types** (Build, Runtime, Static, Memory, Network)
- 📄 **Multiple output formats** (Console, JSON, HTML)
- 🎯 **Actionable insights** and optimization recommendations

## Installation

```bash
# Clone the repository
git clone https://github.com/ar-ease/performance-monitor.git
cd performance-monitor

# Build the CLI
go build -o pmon

# Make it globally available (optional)
sudo mv pmon /usr/local/bin/
```

## Usage

### Basic Commands

```bash
# Run full analysis (recommended)
pmon -a

# Run specific analysis
pmon -b          # Build performance only
pmon -r          # Runtime performance only
pmon -s          # Static code analysis
pmon -m          # Memory profiling
pmon -n          # Network analysis
```

### Output Formats

```bash
# Console output (default)
pmon -a

# JSON output
pmon -a --output json

# HTML report
pmon -a --output html
```

### Advanced Options

```bash
# Continuous monitoring
pmon -a --watch

# Set monitoring duration
pmon -r --duration 30s

# CI-friendly output
pmon -a --ci
```

## Supported Projects

### JavaScript/Node.js

- ✅ Vanilla JavaScript
- ✅ React
- ✅ Next.js
- ✅ Vue.js
- ✅ Angular
- ✅ Svelte

### Other Languages

- ✅ Go
- ✅ Python
- ✅ Java
- 🔄 More coming soon...

## What We Measure

### Build Performance

- Build time
- Bundle size
- Dependency analysis
- Build warnings/errors

### Runtime Performance

- Startup time
- Memory usage
- CPU utilization
- Performance bottlenecks

### Code Quality

- Lines of code
- Cyclomatic complexity
- Test coverage
- Static analysis issues

### Network Performance

- API response times
- Asset loading times
- Bandwidth usage

## Example Output

```
🚀 Performance Monitor CLI (pmon)
================================
🔍 Detected: JavaScript project
📁 Framework: React
🔧 Build Tool: Vite
📍 Root Path: /path/to/your/project
📦 Dependencies: 42 packages

🚀 Running full performance analysis...

📊 Performance Analysis Report
==============================
🏗️  Build Metrics:
  Build Time: 3.2s
  Bundle Size: 2.45 MB
  Dependencies: 42

⚡ Runtime Metrics:
  Startup Time: 150ms
  Memory Usage: 45.2 MB
  CPU Usage: 12.5%
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

JavaScript/Node.js:
├── Bundle size analysis
├── Dependency tree depth
├── Unused code detection
├── Performance bottlenecks
└── Runtime profiling

React Projects:
├── Component render times
├── State update performance
├── Virtual DOM operations
└── Hook optimization

Next.js Projects:
├── SSR/SSG performance
├── Route optimization
├── Image optimization
└── API route performance

pmon -a
├── 1. Project Detection
│ ├── Scan for package.json
│ ├── Identify framework (React/Next.js/Vue/etc.)
│ ├── Detect build tools (Webpack/Vite/etc.)
│ └── Check for config files
├── 2. Setup Monitoring
│ ├── Install temporary monitoring tools
│ ├── Inject performance hooks
│ └── Setup profiling environment
├── 3. Run Analysis
│ ├── Static analysis (bundle, dependencies)
│ ├── Build performance measurement
│ ├── Runtime performance testing
│ └── Resource usage monitoring
└── 4. Generate Report
├── Performance dashboard
├── Recommendations
└── Export options (JSON/HTML/PDF)
