# Custom GolangCI-Lint Linter Scaffolding

This project contains scaffolding for a custom `golangci-lint` linter plugin.

## Structure

- `linter/`: Source code for the linter plugin.
  - `pkg/analyzer/`: Contains the actual analysis logic.
  - `plugin/`: Contains the main entry point for the plugin.
- `example/`: A sample project to test the linter against.
- `Makefile`: helper commands.

## Requirements

- Go 1.21+ (Must match the version used to compile `golangci-lint`)
- `golangci-lint` installed.

## Usage

1. **Build and Run**:
   ```bash
   make run-example
   ```
   This command handles:
   - Installing a local copy of `golangci-lint` to ensure exact Go version matching (required for plugins).
   - Building the plugin.
   - Running the linter against the `example/` directory.

## Troubleshooting

If you encounter "plugin was built with a different version of package ..." errors:
- Run `make clean` and `make run-example` again.
- The Makefile is set up to match the `golangci-lint` version with your local Go version.
# golint-nointernal
