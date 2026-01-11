# Security

## Related Docs
- `docs/INVARIANTS.md`
- `docs/PRD.md`

## Canonical Security Statement
Go 1.25.5 is the latest patch release of the Go 1.25 series at time of pinning. Known toolchain-level CVEs prior to 1.25.5 (notably in `crypto/x509`) are patched in this version. As of the pin date, there are no known unpatched CVEs in the Go standard library affecting typical CLI workloads.

Go’s security guarantees do not extend to third-party modules, application logic, or handling of untrusted input beyond stdlib contracts. Security posture is conditional: CLI tools that do not parse untrusted certificates or operate as network services are outside the known CVE impact envelope, while tools that do process untrusted certs, HTTP traffic, or structured attacker-controlled input must rely on correct application-level handling.

Pinning to Go 1.25.5 mitigates known toolchain vulnerabilities as of the pin date, but does not eliminate dependency- or application-level risk.
