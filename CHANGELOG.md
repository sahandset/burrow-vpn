# Changelog

All notable changes to this project are documented here.

---

## Week 1 – Initial Infrastructure Setup (2025-06-22)

### Added
- Project scaffold with clean Go module structure:
  - `cmd/` for CLI binaries
  - `internal/` for core modules (vpn, auth, etc.)
- TUN device support using `songgao/water` to create virtual network interfaces
- AES-GCM packet encryption for confidentiality and integrity
- mTLS authentication logic with certificate-based client/server validation
- CLI entrypoints for VPN client and server setup
- GitHub Actions CI workflow:
  - Module tidy
  - Linting with `golint`
  - Unit tests
  - Build for both client and server
