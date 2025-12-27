# LinkedIn Automation — Technical Proof of Concept

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-Educational-blue.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Proof%20of%20Concept-yellow.svg)](README.md)

> **⚠️ Educational & Evaluation Use Only**
>
> This repository demonstrates clean architecture patterns for browser automation systems.  
> It does **not** interact with live production systems and is designed strictly for learning and technical evaluation.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Demo Video](#demo-video)
- [Important Disclaimers](#important-disclaimers)
- [Project Goals](#project-goals)
- [Architecture](#architecture)
- [Features](#features)
- [Messaging Workflow](#messaging-workflow)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Technical Deep Dive](#technical-deep-dive)
- [Safety & Ethics](#safety--ethics)
- [License](#license)

---

## 🎯 Overview

This project is a **technical proof-of-concept** showcasing enterprise-grade Go architecture for browser automation systems.

It demonstrates:

- **Clean Architecture** using interfaces and separation of concerns
- **Human-like behavior modeling** (mouse, typing, scrolling, timing)
- **Persistent state management** for resumable execution
- **Automated messaging workflow** (mocked, safe, idempotent)
- **Anti-detection concepts** implemented architecturally
- **Safety-first design** preventing real platform interaction

### What this is NOT
- ❌ A production automation tool  
- ❌ A real LinkedIn bot or scraper  
- ❌ A ToS-bypassing system  

---

## 🎥 Demo Video

A walkthrough demonstrating:

- Project structure and clean architecture
- Mock vs Rod browser design
- Human-like behavior simulation
- **Scraping + automated messaging workflow**
- Persistent message state handling

▶️ **Watch the demo here:**  
https://youtu.be/2iuYVnu-vMY

> Note: All browser actions in the demo are executed in **mock mode** for safety and compliance.

---

## ⚠️ Important Disclaimers

### Legal & Ethical Notice

Automating LinkedIn **violates their Terms of Service** and may result in:
- Permanent account suspension
- Legal consequences
- Civil liability

This project:
- ✅ Uses **mock browser implementations by default**
- ✅ Demonstrates **architecture and automation concepts only**
- ✅ Does **not** interact with live LinkedIn systems

### Platform Compliance

- Real browser automation is disabled by default
- Rod integration is guarded via build tags
- No credentials are stored or used
- No network calls to LinkedIn are performed

---

## 📌 Project Goals

This proof-of-concept demonstrates:

| Goal | Description |
|-----|------------|
| Clean Architecture | Modular, testable Go code following SOLID principles |
| Interface Design | Pluggable browser implementations |
| Behavior Simulation | Human-like timing, movement, and interaction |
| Messaging Workflow | Template-based, state-aware messaging |
| State Persistence | JSON-backed idempotent execution |
| Safety Engineering | Mock-first, ToS-compliant design |

---

## 🏗️ Architecture
```
┌───────────────────────────────┐
│      cmd/app/main.go          │
│      Application Entry        │
└───────────────┬───────────────┘
                │
    ┌──────────┴──────────┐
    │                     │
┌────▼────┐      ┌──────▼──────┐
│ Browser │      │  Behavior   │
│  Layer  │      │   Engine    │
│         │      │             │
│  Mock   │      │   Mouse     │
│  Rod*   │      │   Typing    │
└────┬────┘      │  Scrolling  │
     │           │   Timing    │
┌────▼──────────────────────────┐
│     State Store (JSON)        │
│     • Visited Profiles        │
│     • Messages Sent           │
│     • Idempotency             │
└───────────────────────────────┘
```

\* Rod browser implementation exists but is disabled by default.

---

## ✨ Features

### 🎭 Human-Like Behavior Simulation
- Randomized think time
- Natural mouse movement paths
- Human typing rhythm
- Scroll acceleration/deceleration
- Short & long pauses

### 🕵️ Anti-Detection Strategies (Conceptual)
- Browser fingerprint abstraction
- Timing randomization
- Rate limiting
- Mocked UA and viewport behavior

### 💾 State Management
- Persistent JSON storage
- Resume-safe execution
- Duplicate-action prevention

### 🔒 Safety Controls
- Mock browser by default
- No real platform interaction
- No credential handling
- Explicit opt-in for real browser builds

---

## 💬 Messaging Workflow

The messaging flow was added in response to assignment feedback and is fully implemented **architecturally**.

### Messaging Pipeline
```
Profile Visit (Scraping)
         ↓
Check if Already Messaged
         ↓
Generate Message Template
         ↓
Send Message (Mock Browser)
         ↓
Persist Message State
```

### Key Properties
- Template-based message generation
- Idempotent messaging (no duplicates)
- State-aware execution
- Safe mock execution
- Fully integrated with behavior engine

---

## 🚀 Getting Started

### Prerequisites
- Go **1.25+**
- Basic Go module knowledge

### Installation
```bash
git clone https://github.com/yogin1611/linkedin-automation-poc.git
cd linkedin-automation-poc
go mod download
go build ./cmd/app
```

### Run (Mock Mode)
```bash
go run ./cmd/app
```

Sample output:
```
[MOCK BROWSER] Opening URL
[SCROLL] Scrolling viewport
[TYPE] Reviewing profile...
[MOCK MESSAGE] Sending message
Workflow completed successfully (Mock Mode)
```

### Optional: Rod Build
```bash
go run -tags=rod ./cmd/app
```

⚠️ Not required for evaluation.

---

## 📁 Project Structure
```
linkedin-automation-poc/
├── cmd/app/main.go
├── internal/
│   ├── behavior/        # Human-like behavior simulation
│   ├── browser/         # Browser interface + mock/rod impl
│   ├── messaging/       # Messaging service + templates
│   ├── storage/         # Persistent state store
│   ├── config/          # Configuration loading
│   └── logger/          # Structured logging
├── state.json           # Generated at runtime (gitignored)
├── go.mod
├── go.sum
└── README.md
```

### Key Interfaces
```go
type Browser interface {
	Open(url string) error
	SendMessage(profileURL string, message string) error
	Close() error
}
```

---

## 🔬 Technical Deep Dive

### Messaging State Model
```go
type State struct {
	VisitedProfiles []string
	MessagesSent    map[string]string
}
```

Ensures:
- No duplicate messages
- Resume-safe execution
- Deterministic behavior

---

## 🔐 Safety & Ethics

### Design Principles
- Mock-first execution
- Explicit opt-in for real browser
- No credential usage
- Transparent logging
- Educational intent only

### Intended Use
- ✅ Learning automation architecture
- ✅ Internship evaluation
- ❌ Real platform automation
- ❌ ToS circumvention

---

## 📄 License

**Educational Use License**

- **Allowed:** Learning, reference, evaluation
- **Not allowed:** Production use, scraping, automation against real platforms

---

## 👨‍💻 Author

**Heerath Bhat**

- GitHub: https://github.com/yogin1611
- LinkedIn: https://linkedin.com/in/heerathbhat
- Email: heerath.bhat@gmail.com

---

⭐ Built with a focus on clean engineering, safety, and responsible automation.