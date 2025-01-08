# Sholawat JSON

All praise is due to Allah SWT for His countless blessings. May blessings and peace always be upon our beloved Prophet Muhammad SAW.

**sholawat-json** is a collection of sholawat in a simple JSON format, aimed at developers who wish to integrate them into websites or applications. Each JSON file represents a different sholawat and follows a predefined schema, ensuring ease of access and consistent structure. To maintain accuracy and structure, we implement `JSON schema` validation, making it easier to integrate the data confidently. This project is designed for lightweight applications, with a focus on speed and efficiency.

We strive for accuracy, but if you encounter any errors, please reach out. Your feedback is invaluable for improvement.

## CDN
Check out the [`/sholawat`](https://github.com/afaf-tech/sholawat-json/tree/master/sholawat) to see all available sholawat and source files. The JSON files are also available through [JSDELIVER](https://www.jsdelivr.com/package/npm/sholawat-json) CDN

### List Sholawat and Sources 
get the information with : 
```
https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/sholawat.json
```
### Get a Fasl

You can get a single chapter (fasl) by providing its faslNumber for each sholawat. Text, transliteration and translation are provided on each fasl. To get it you can provide: 
```
https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/{sholawatName}/{source}/fasl/{faslNumber}.json

```
For Example:

- Sholawat Diba from nu online fasl 1: `https://cdn.jsdelivr.net/npm/sholawat-json@latest/sholawat/diba/nu_online/fasl/10.json`

Get a Schema
You can access the JSON schema for each fasl in the schemas folder. These schemas ensure the data structure remains consistent and reliable. Use the following link to get the schema:

```bash
https://cdn.jsdelivr.net/npm/sholawat-json@latest/schemas/{schema_name}.schema.json
```
For example:

Burdah Fasl Schema (v1.0):
https://cdn.jsdelivr.net/npm/sholawat-json@latest/schemas/burdah_fasl_v1.0.schema.json

## Data Sources

The sholawat data in this project has been collected from various trusted sources, including reputable Islamic literature and applications. Each data entry has gone through a multi-stage validation process to ensure accuracy and authenticity.

## Contributions

We warmly welcome contributions from anyone who has ideas or suggestions for the development of this project. If you would like to contribute, please submit a Pull Request (PR) through GitHub.

## License

This project is licensed under the MIT License. For more details, please refer to the LICENSE file included in the repository.

## Author
This project was developed by afaf-tech. If you encounter any issues, please report them by creating a GitHub issue in the repository.