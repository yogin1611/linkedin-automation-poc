# LinkedIn Automation — Technical Proof of Concept

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-Educational-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Proof%20of%20Concept-yellow.svg)](README.md)

> **⚠️ Educational & Evaluation Use Only**
> 
> This repository demonstrates clean architecture patterns for browser automation systems. It does **not** interact with live production systems and is designed for learning purposes only.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Demo Video](#demo-video)
- [Important Disclaimers](#important-disclaimers)
- [Project Goals](#project-goals)
- [Architecture](#architecture)
- [Features](#features)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Technical Deep Dive](#technical-deep-dive)
- [Safety & Ethics](#safety--ethics)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)

---

## 🎯 Overview

This project is a **technical proof-of-concept** showcasing enterprise-grade Go architecture for browser automation systems. It demonstrates:

- **Clean Architecture**: Interface-driven design with clear separation of concerns
- **Behavior Modeling**: Human-like interaction patterns and timing simulation
- **State Management**: Persistent, resumable execution flows
- **Anti-Detection Strategies**: Architectural patterns for stealth operations
- **Safety-First Design**: Mock implementations preventing unintended execution

**What this is NOT:**
- ❌ A production-ready automation tool
- ❌ A LinkedIn scraper or bot
- ❌ A system that bypasses platform safeguards

---

## 🎥 Demo Video

A short walkthrough demonstrating:
- Project setup and folder structure
- Architecture and design decisions
- Mock vs Rod browser implementation
- Human-like behavior simulation
- State persistence and safe execution

▶️ **Watch the demo here:**  
[VIDEO DEMO](https://youtu.be/ieyAdklGjZg)

---

## ⚠️ Important Disclaimers

### Legal & Ethical Notice

**Automating LinkedIn violates their Terms of Service** and may result in:
- Permanent account suspension
- Legal action from LinkedIn
- Violation of CFAA (Computer Fraud and Abuse Act)
- Civil liability

This project:
- ✅ Uses **mock implementations** by default
- ✅ Demonstrates **architecture and design patterns**
- ✅ Teaches **responsible engineering practices**
- ✅ Does **not** interact with live systems

### Platform Compliance

All browser automation code is:
- Abstracted behind interfaces
- Disabled by default using build tags
- Implemented for educational demonstration only
- Not intended for production use

---

## 📌 Project Goals

This proof-of-concept demonstrates proficiency in:

| Goal | Description |
|------|-------------|
| **Clean Architecture** | Modular, testable, maintainable Go code following SOLID principles |
| **Interface Design** | Strategy pattern for pluggable browser implementations |
| **Behavior Simulation** | Realistic human interaction modeling with randomized timing |
| **State Persistence** | JSON-based state management for resumable operations |
| **Anti-Detection Theory** | Conceptual understanding of fingerprinting and evasion |
| **Safety Engineering** | Build-time controls preventing accidental execution |

---

## 🏗️ Architecture

### High-Level Overview
```
┌─────────────────────────────────────────────────────────┐
│                     cmd/app/                            │
│                  Application Entry Point                │
└─────────────────────┬───────────────────────────────────┘
                      │
          ┌───────────┴───────────┐
          │                       │
┌─────────▼─────────┐   ┌────────▼─────────┐
│  Browser Layer    │   │  Behavior Engine │
│  (Interface)      │   │  (Human-like)    │
├───────────────────┤   ├──────────────────┤
│ • MockBrowser     │   │ • ThinkTime      │
│ • RodBrowser*     │   │ • Randomization  │
└─────────┬─────────┘   │ • Timing Curves  │
          │             └──────────────────┘
          │
┌─────────▼─────────────────────────────────┐
│           Storage & State Layer           │
│  • Persistent JSON state                  │
│  • Visit tracking                         │
│  • Resume capability                      │
└───────────────────────────────────────────┘
```

*Rod implementation disabled by default via build tags

### Directory Structure
```
linkedin-automation-poc/
├── cmd/
│   └── app/
│       └── main.go              # Application entry point
├── internal/
│   ├── behavior/
│   │   ├── behavior.go          # Human-like behavior engine
│   │   ├── fingerprint.go       # Fingerprint masking strategies
│   │   ├── human.go             # Human interaction patterns
│   │   ├── mouse.go             # Mouse movement simulation
│   │   ├── schedule.go          # Activity scheduling
│   │   ├── scroll.go            # Scrolling behavior
│   │   └── typing.go            # Typing rhythm simulation
│   ├── browser/
│   │   ├── interface.go         # Browser interface definition
│   │   ├── mock_browser.go      # Safe mock implementation (default)
│   │   └── rod_browser.go       # Rod implementation (build-tagged)
│   ├── config/
│   │   ├── config.go            # Configuration management
│   │   └── loader.go            # Environment loader
│   ├── logger/
│   │   └── logger.go            # Structured logging
│   └── storage/
│       ├── json_store.go        # JSON storage implementation
│       └── state.go             # State structure definitions
├── .env.example                 # Example environment variables
├── go.mod                       # Go module dependencies
├── go.sum                       # Dependency checksums
├── README.md                    # Project documentation
├── state.json                   # Persistent state (generated at runtime, gitignored)
└── LICENSE                      # License information
```

---

## ✨ Features

### Core Capabilities

#### 🎭 Human Behavior Simulation
- **Randomized Think Time**: Cognitive delay modeling (2-5 seconds)
- **Long Pauses**: Natural breaks (10-30 seconds)
- **Non-Deterministic Timing**: Jittered execution patterns
- **Activity Scheduling**: Business hours awareness

#### 🕵️ Anti-Detection Strategies (Conceptual)
- Browser fingerprint masking (abstracted)
- Mouse movement path modeling (Bézier curves)
- Human typing rhythm simulation
- Random scrolling patterns
- Rate limiting and throttling
- User-agent rotation (simulated)

#### 💾 State Management
- Persistent JSON storage
- Visit history tracking
- Last execution timestamp
- Safe resume after interruption
- Idempotent operations

#### 🔒 Safety Controls
- Mock browser by default
- Build-tag guarded Rod execution
- No credential storage
- No live API interactions
- Simulation-only fingerprinting

---

## 🚀 Getting Started

### Prerequisites

- **Go 1.25+** ([Download](https://golang.org/dl/))
- Basic understanding of Go modules
- Familiarity with browser automation concepts

### Installation
```bash
# Clone the repository
git clone https://github.com/yogin1611/linkedin-automation-poc.git
cd linkedin-automation-poc

# Install dependencies
go mod download

# Verify installation
go build ./cmd/app
```

### Running the Demo (Safe Mode)
```bash
# Execute with mock browser (default)
go run ./cmd/app
```

**Sample Output (abridged):**
```
[INFO] Initializing browser automation system (Mock Mode)
[INFO] Applying fingerprint masking strategy...
[INFO] Simulating mouse movement (Bézier curve)
[INFO] Simulating human scrolling behavior
[INFO] Simulating typing with human rhythm
[INFO] Mock navigation to: https://example.com
[INFO] State persisted to state.json
[INFO] Execution complete
```

### Optional: Rod Build (Advanced)

⚠️ **Not Recommended for Evaluation**
```bash
# Enable Rod browser (requires Chrome/Chromium)
go run -tags=rod ./cmd/app
```

This mode:
- Launches a real browser instance
- May trigger antivirus alerts
- Should only be used in isolated test environments
- Is **not** necessary to demonstrate the project

---

## 📁 Project Structure

### Key Components

#### 1. Browser Abstraction (`internal/browser/`)
```go
// Browser interface allows pluggable implementations
type Browser interface {
    type Browser interface {
    Open(url string) error
    Close() error
}
```

**Implementations:**
- `MockBrowser`: Default safe implementation
- `RodBrowser`: Real browser control (build-tagged)

#### 2. Behavior Engine (`internal/behavior/`)

Provides realistic human interaction patterns:
```go
// Think simulates human cognitive delay
func Think() {
    time.Sleep(randomDuration(2, 5) * time.Second)
}

// LongPause simulates natural breaks
func LongPause() {
    time.Sleep(randomDuration(10, 30) * time.Second)
}
```

#### 3. State Persistence (`internal/storage/`)

Manages execution state across runs:
```go
type State struct {
    VisitedURLs  []string  `json:"visited_urls"`
    LastExecution time.Time `json:"last_execution"`
}
```

#### 4. Configuration (`internal/config/`)

Environment-based configuration:
```go
type Config struct {
    BrowserMode    string
    HeadlessMode   bool
    StateFilePath  string
    LogLevel       string
}
```

---

## 🔬 Technical Deep Dive

### Human Behavior Modeling

The behavior engine implements several strategies to mimic human interaction:

#### Timing Randomization
```go
// Base timing with jitter
baseDelay := 2 * time.Second
jitter := time.Duration(rand.Intn(3000)) * time.Millisecond
actualDelay := baseDelay + jitter
```

#### Mouse Movement Simulation

- Bézier curve path generation
- Variable speed along path
- Occasional overshooting/correction
- Natural acceleration/deceleration

#### Typing Rhythm

- Character-by-character delays (50-150ms)
- Occasional typos and corrections
- Word boundary pauses
- Variable typing speed

### Anti-Detection Architecture

All anti-detection features are **simulated** for safety:

| Technique | Implementation | Status |
|-----------|----------------|--------|
| Fingerprint Masking | Interface-based | Simulated |
| Canvas Fingerprinting | Strategy pattern | Abstracted |
| WebGL Fingerprinting | Mock implementation | Abstracted |
| User-Agent Rotation | Configuration-based | Simulated |
| Mouse Movement | Bézier curves | Modeled |
| Timing Randomization | Behavior engine | Functional |

### State Management Flow
```
┌──────────┐
│  Start   │
└────┬─────┘
     │
     ▼
┌──────────────┐      ┌──────────────┐
│ Load State   │─────▶│ State Exists?│
└──────────────┘      └─────┬────────┘
                            │
                    ┌───────┴────────┐
                    │ Yes        No  │
                    ▼                ▼
            ┌──────────────┐  ┌──────────────┐
            │ Resume From  │  │ Create New   │
            │ Last Point   │  │ State        │
            └──────┬───────┘  └──────┬───────┘
                   │                 │
                   └────────┬────────┘
                            ▼
                    ┌──────────────┐
                    │ Execute Task │
                    └──────┬───────┘
                           ▼
                    ┌──────────────┐
                    │ Persist State│
                    └──────────────┘
```

---

## 🔐 Safety & Ethics

### Design Principles

1. **Safe by Default**: Mock implementations prevent accidental execution
2. **Explicit Opt-In**: Real browser requires build tags
3. **No Credentials**: System never stores or transmits auth data
4. **Transparent Logging**: All actions logged for audit trail
5. **Educational Focus**: Code comments explain "why not" alongside "how"

### Responsible Use Guidelines

This project is intended for:
- ✅ Learning Go architecture patterns
- ✅ Understanding browser automation concepts
- ✅ Studying anti-detection strategies theoretically
- ✅ Demonstrating engineering skills

This project should **never** be used for:
- ❌ Automating LinkedIn or similar platforms
- ❌ Scraping data without authorization
- ❌ Bypassing platform security measures
- ❌ Violating Terms of Service

### Platform Compliance Statement

The maintainers of this project:
- Respect platform Terms of Service
- Discourage unauthorized automation
- Provide this code for educational purposes only
- Accept no liability for misuse

---

## 📚 Documentation

### Additional Resources

- [Go Browser Automation Best Practices](docs/best-practices.md)
- [Anti-Detection Strategy Overview](docs/anti-detection.md)
- [Architecture Decision Records](docs/adr/)
- [API Reference](docs/api-reference.md)

### Related Technologies

- [Rod](https://go-rod.github.io/) - Browser automation library
- [Chromium DevTools Protocol](https://chromedevtools.github.io/devtools-protocol/)
- [Playwright](https://playwright.dev/) - Similar automation framework

---

## 🤝 Contributing

### Guidelines

Contributions are welcome for:
- Architecture improvements
- Better mock implementations
- Documentation enhancements
- Test coverage expansion

**We do NOT accept:**
- Code that enables platform ToS violations
- Anti-detection bypasses for live systems
- Credential handling implementations
- Production automation features

### Development Setup
```bash
# Fork and clone the repository
git clone https://github.com/yogin1611/linkedin-automation-poc.git

# Create a feature branch
git checkout -b feature/your-feature-name

# Make changes and test
go test ./...

# Submit pull request
```

---

## 📄 License

This project is licensed under the **Educational Use License**.

**Key Terms:**
- ✅ Study and learn from the code
- ✅ Reference in educational contexts
- ✅ Use as architecture example
- ❌ Deploy to production
- ❌ Use for platform automation
- ❌ Commercial exploitation

See [LICENSE](LICENSE) file for full terms.

---

## 👨‍💻 Author

**Heerath Bhat**
- GitHub: [@yogin1611](https://github.com/yogin1611)
- LinkedIn: [Heerath Bhat](https://linkedin.com/in/heerathbhat)
- Email: heerath.bhat@gmail.com

---

## 🙏 Acknowledgments

- Rod library authors for excellent browser automation tools
- Go community for architectural patterns and best practices
- Open-source contributors who value ethical software development

---

## 📞 Support

### Questions?

- 📧 Email: heerath.bhat@gmail.com
- 💬 Discussions: [GitHub Discussions](https://github.com/yogin1611/linkedin-automation-poc/discussions)
- 🐛 Issues: [GitHub Issues](https://github.com/yogin1611/linkedin-automation-poc/issues)

### Disclaimer

This project is provided "as is" without warranty. The authors are not responsible for misuse or violations of platform Terms of Service. Always respect website policies and legal boundaries.

---

<div align="center">

**⭐ If you found this educational, please star the repository! ⭐**

Made with ❤️ for learning and ethical engineering

</div>