# Snippetbox by Go

## Project structure

```bash
.
├── README.md
├── cmd
│   └── web
│       ├── handlers.go
│       └── main.go
├── go.mod
├── internal
└── ui
    ├── html
    └── static
```
- `cmd` contain application specific
    - `cmd/web` for web application
    - can add more application such as CLI (`cmd/cli`) in the future
- `internal` 
    - contain non-application specific code (Ex. access SQL database)
    - code here can only be access to code in project
