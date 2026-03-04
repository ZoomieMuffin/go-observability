# DEV-55: Add startup and verification steps to README

## Date
2026-03-04

## Context
Step 1 needed a single README-based flow so new contributors can run gateway and
worker, then verify normal and failure behavior without extra docs.

## Changes
- Added Step 1 quick-start instructions in README.md for worker and gateway
startup.
- Documented required gateway environment variables and default values:
WORKER_BASE_URL, HTTP_TIMEOUT_MS.
- Added reproducible curl verification examples for normal path (200) and failure
path (502 when worker is stopped).
- Added project links from README to portfolio and issue index.

## Validation
- Verified the README includes startup commands for both services.
- Verified the README includes environment variables with defaults.
- Verified the README includes both normal and failure curl checks with expected
responses.

## Evidence
- README.md
- doc/issues/INDEX.md

## Notes
- Keep README as the runtime entrypoint only; move per-issue implementation details to
doc/issues/.






##### Validation
- Human-reviewed on 2026-03-04(JST)
