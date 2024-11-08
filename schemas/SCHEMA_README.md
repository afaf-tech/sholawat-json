# JSON Schema

  This guide provides a naming format for JSON Schema files to maintain consistency and ease of management.

## Overview of JSON Schema

    JSON Schema provides a clear framework for defining and validating the structure of JSON data. It allows developers to specify data types, rules, and constraints, ensuring consistency across systems. It also simplifies integration and testing by providing a shared standard for data exchange and supports both human-readable and machine-readable documentation. JSON Schema is widely supported by a range of development tools, making it easy to implement and use across various projects.

## Schema Specification Version

All JSON Schema files in this project adhere to the **2020-12** version of the JSON Schema specification. This version offers expanded capabilities and improvements over previous drafts (e.g., draft-04, draft-06, and draft-07). To learn more about this specification, refer to the [JSON Schema Specification 2020-12](https://json-schema.org/) and [Understanding JSON Schema](https://json-schema.org/understanding-json-schema/).

## Naming Format

Each JSON Schema file should follow this pattern:

### Components

- **`<app_name>`**: Name of the application or project.
- **`<schema_descriptor>`**: Descriptive schema name, e.g., `diba_rawi`, `burdah_albushiri`.
- **`<version>`**: Schema version in the format `vX`, where `X` is the version (e.g., `v1`, `v2`).

Examples
- initial version`burdah_fasl_v1.schema.json`

- subsequent versions:`burdah_fasl_v2.schema.json`

- semantic versioning (optional):`burdah_fasl_v1.1.schemajson`

## Versioning

- **Major Updates**: Increment version (e.g., `v1`, `v2`) for changes that are not backward-compatible.
- **Minor Updates**: Use minor versions (e.g., `v1.1`, `v1.2`) for backward-compatible changes.
