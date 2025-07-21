# PokeDex

A feature-rich, modern Pokédex built with Go—currently a CLI tool, soon to be a full-stack web application with Postgres, file uploads, and Dockerized deployment.

## Overview

GoPokédex is a personal project designed to fetch, display, and manage Pokémon data using the [PokeAPI](https://pokeapi.co/).  
It currently features a modular command-line interface (CLI) with extensible commands, robust error handling, and a clean codebase, structured for real-world development.  
The goal: learn modern Go development, practice API integration, and progressively build a **production-grade app** ready for team projects, research, or internships.

---

## Features

- **Modular CLI:**  
  Type-safe, extensible commands (help, exit, map/locations, etc.)
- **PokéAPI Integration:**  
  Fetches and paginates Pokémon locations via external HTTP client.
- **Project Structure:**  
  Uses idiomatic Go project layout, with separation between CLI, internal logic, and API layer.
- **Testing:**  
  Includes table-driven tests for utility functions (with plans for more).
- **GitHub Workflow:**  
  All features developed on branches, merged via PRs, and tracked using Git.

---

## Usage

### Prerequisites

- Go 1.20+ (or latest)
- Internet access (for PokéAPI calls)

### Running

```bash
git clone https://github.com/ECHIDNATHEG/PokeDex.git
cd PokeDex
go run main.go

Commands

    help — List available commands

    exit — Exit the Pokédex

    map — List next batch of Pokémon locations (paginated)

    mapb — Go back to previous batch of locations

Example

Pokedex > help
Available commands:
  help     : Gives Help about app
  exit     : Exits the program
  map      : Lists all the locations in batches of 20
  mapb     : Lists previous batch of locations

Pokedex > map
canalave-city-area
oreburgh-mine
```

## Project Structure

```
internal/           # API logic, config, and data models
temp/               # Temporary files and data
main.go             # CLI app entrypoint
main_test.go        # Basic tests (expanding!)
README.md           # This file
```

## Future Roadmap

I plan to evolve GoPokédex into a real-world, full-stack application, including:

- *Web Server Mode*:
    Switch between CLI and web server (HTTP REST API or web dashboard)

- *PostgreSQL Database*:
    Store favorites, search history, and cache PokéAPI results for speed and offline access

- *User Accounts*:
    Optional: support for logins, multiple users, and personalized data

- *File Server*:
    Upload and serve files (custom sprites, notes, screenshots)

- *Docker Containerization*:
    One-click setup using Docker & Docker Compose (app + Postgres)

- *CI/CD Integration*:
    Automated tests and builds with GitHub Actions

- *Deployment*:
    Live demo hosted on Render or Fly.io

