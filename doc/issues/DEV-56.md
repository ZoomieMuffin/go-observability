# DEV-56: Introduce CI with GitHub Actions

## Date
2026-03-02

## Context
The project needed a minimum quality gate for pull requests.

## Changes
- Added GitHub Actions workflow at .github/workflows/ci.yml.
- Configured CI to run on pull_request.
- Added go test ./...
- Added go vet ./...
- Added go build ./...

## Validation
- Confirmed CI starts automatically on pull_request.
- Confirmed checks fail when a step fails.
- Confirmed checks pass when all steps succeed.

## Evidence
- .github/workflows/ci.yml
- PR #29

## Notes
- This issue covers CI only.
- CD is out of scope and handled separately.

Reviewer note: Formatting and final wording were explicitly requested and reviewed by a human.
