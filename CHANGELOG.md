# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.5] - 2021-07-17
### Fixed
- Update version of github.com/aws/aws-sdk-go as existing version had a vulnerability

## [0.0.4] - 2021-04-26
### Added
- Scan() implemented
- ScanWithContext() implemented
- ScanPages() implemented
- ScanPagesWithContext() implemented

## [0.0.3] - 2021-04-07
### Added
- QueryPages() implemented
- QueryPagesWithContext() implemented

### Changed
- AddReturnPutItemOutput() to accept multiple PutItemOutputs
- AddReturnQueryOutput() to accept multiple QueryOutputs

## [0.0.2] - 2021-03-07
### Added
- PutItemWithContext() implemented
- QueryWithContext() implemented

## [0.0.1] - 2021-02-14
### Added
- Query() implemented

### Changed
- How inputs and outputs for PutItem() work

## [0.0.0] - 2021-02-04
### Added
- PutItem() implemented
