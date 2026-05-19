# Changelog

All notable changes to this project will be documented in this file.

## [1.7.2] - 2026-05-18

### Added
- **Project Management**: Added `POST /project` endpoint for administrators to append new projects to the portfolio.

### Changed
- **AI Agent**: Internal optimizations and chore updates for the AI chat integration.

## [1.7.1] - 2026-05-17

### Fixed
- **AI Agent Fixes**: Enhanced error reporting and validation for OpenRouter AI responses.

## [1.7.0] - 2026-05-14

### Added
- **NGL Proxy Integration**: Added `POST /ngl` endpoint to proxy anonymous messages to NGL.link.
- **Documentation Update**: Updated README with new features, architecture diagrams, and Wakatime badge.

## [1.6.3] - 2026-04-22

### Changed
- **AI Model Upgrade**: Migrated from Pollinations AI to `GPT-4o-mini` via OpenRouter for improved chat responses and reliability.
- **Dependency Update**: Updated project dependencies to their latest stable versions.

## [1.6.2] - 2026-04-02

### Added
- **New Resume Endpoint**: Added `GET /dev` endpoint to retrieve resume data from GitHub Gist.
- **Enhanced Route Architecture**: Formalized the development route integration within the unified endpoint matrix.

### Changed
- **Pagination Optimization**: Fine-tuned pagination limits for `/blog` (8 → 11) and `/certs` (8 → 5) to balance data density and performance.

## [1.6.1] - 2026-03-28

### Changed
- **Paginator Completion**: Successfully finalized the paginator logic for improved data navigation across paginated endpoints.
- **Request Limiting**: Implemented manual request limits and fine-tuned overall thresholds to enhance API stability and prevent abuse.
- **System Reliability**: Further optimized the cloud-based contact system for more reliable message delivery to administrators.

## [1.6.0] - 2026-03-28

### Added
- **Manual Request Limits**: Added manual request limits and reduced overall limits for better API stability.
- **Blog Editing**: Fully implemented blog edit capabilities for administrators via `PUT /blog`.

### Changed
- **Enhanced Reliability**: Optimized contact system for admin cloud-based message delivery.

## [1.5.2] - 2026-03-24

### Changed
- **Contact System Migration**: **Breaking Change**: Migrated contact system from email-based to admin cloud-base for improved reliability.
- **Singularity Responses**: Added `data` field on **experience** and **projects** for singular response formats.

## [1.5.1] - 2026-03-19

### Changed
- **Global Caching**: Implemented a global in-memory caching system to reduce GitHub Gist API dependency.
- **Experience Updates**: Updated experience endpoints (Changed from POST to PUT for updates).

## [1.5.0] - 2026-03-18

### Changed
- **Multimedia Update & Security**:
  - Renamed `/upload-image` to `/upload`.
  - Renamed `/images` to `/retrieve`.
  - Enhanced security by using `Abort` instead of simple JSON for failed admin requests.
  - Added support for all data types via `mimetype` identification.

## [1.4.1] - 2026-03-14

### Added
- **Certificates Endpoints**: Added listing and publishing endpoints for portfolio certifications.

### Changed
- **Pagination Tuning**: Adjusted list sizing from **10 → 15** for paginated responses.

### Fixed
- **Bug Fixes**: Fixed state interaction issues between clicked items and pagination.

## [1.4.0] - 2026-03-05

### Added
- **YouTube MP3 Helper**: Added `GET /yt` endpoint backed by RapidAPI.
- **Blog IDs**: Implemented auto-assignment of monotonically increasing IDs for blog posts.

### Changed
- **AI Chat Optimization**: Stripped Pollinations "Ad" footer and improved payload parsing.
- **Hot Reload Improvements**: Updated Air config to stop on build errors and clear screen on rebuild.

## [1.3.2] - 2026-02-14

### Changed
- **Documentation Overhaul**: Added architecture diagrams, component responsibilities matrix, and canonical reference guides.

## [1.3.0] - 2026-02-07

### Added
- **Manga Utility**: Added `GET /manga` supporting search, chapter listing, and reading workflows.
- **Telegram Storage Bridge**: Implemented `POST /upload` (admin) and `GET /retrieve` (public) for offloaded assets.

## [1.2.0] - 2026-01-19

### Changed
- **Enhanced AI Agent**: Improved response formatting and error handling for Pollinations AI.
- **Logging System**: Implemented comprehensive request logging with timestamp formatting.

## [1.1.0] - 2026-01-03

### Added
- **Three-tier Permission System**: Introduced `ALL`, `COOKIE`, and `ADMIN` access levels.
- **Auth Middleware**: Added cookie handler and admin-level protection.

## [1.0.0] - 2026-01-01

### Added
- **Baybayin Transliterator**: Initial release with complete Unicode character mapping and normalization.

## [0.8.0] - 2025-12-30

### Added
- **Initial Release**: Core RESTful API with Gin and GitHub Gist integration.
