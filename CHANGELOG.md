# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](http://semver.org).

## [v2.0.0](https://github.com/ghacts/go/compare/v1.0.0...v2.0.0) (2026-08-09)

### Bug Fixes

- update actions' version

- avoid duplicate artifact names

- upgrade actions

- update syntax

- **action:** set build output directory for build command

- **test:** set artifact name

### Dev Bug Fixes

- **github-workflow:** add missing parameter to disable creating pr steps

### Documentations

- **readme:** add contact and license information

### Features

- **action:** support customize the artifact retention days

- **action:** support custom build command

- **action:** allow skipping upload the artifact

- **action:** add upload artifact step

Full Changelog: [v1.0.0...v2.0.0](https://github.com/ghacts/go/compare/v1.0.0...v2.0.0)

## v1.0.0 (2023-06-15)

### Bug Fixes

- **action:** change action name

- **action:** configure branding icon and color

- **action:** change default golangci-lint version to latest

- **action:** remove trailing slashes

- **action:** set cache dependency path for step setup-go

- **action:** set working directory for build and test step

- **action:** set working directory

- **action:** fix wrong project path

- **action:** correct conditional logics

### Dev Bug Fixes

- **taskfile:** remove version tag prefix from release branch name

### Dev Features

- **changelog:** configure changelog generator

- **github-workflow:** add workflow for automate releasing new version

- **taskfile:** add task files

### Documentations

- **readme:** add documentation for action

### Features

- **action:** add ability to skipping step(s), changing working directory

- **action:** add action.yml

Full Changelog: [v1.0.0](https://github.com/ghacts/go/commits/v1.0.0)
