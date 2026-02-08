# Tiny-Go-Shell

A simple interactive shell written in Go, inspired by Linux terminals.  
It supports executing system commands, changing directories (`cd`), viewing command history, and has a minimal custom prompt.

---

## Features

- Execute system commands
- Change directories using `cd`
- Command history (`history` command)
- Exit shell using `exit`
- Minimal, easy-to-read prompt: `username@hostname: current_dir$`

---

## Requirements

- **Go** 1.21 or higher  
  Install from [golang.org](https://golang.org/dl/)

- **Operating System**  
  Recommended: Linux (Ubuntu)  
  If on Windows, **WSL (Windows Subsystem for Linux) with Ubuntu** is required  
  [WSL installation guide](https://learn.microsoft.com/en-us/windows/wsl/install)

- **Dependencies**  
  Uses only Go’s standard library — no external packages required

---

## Setup & Run

1. Clone the repository:

```bash
git clone https://github.com/SinghVanshraj/Tiny-Go-Shell.git
cd Tiny-Go-Shell
