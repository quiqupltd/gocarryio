# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.6.0]

- Add `manufacturer_id` to item-request
- Add upload-document, bulk-upload-documents and retry-document-upload endpoints
- Add `document_id`, `carrier_upload_status`, `carrier_upload_date` and `message` to document-object
- Pin oapi-codegen to v2.5.1 so regeneration does not churn the generated client

## [v0.5.0]

- Add registration_numbers field to FreeFormRequest and LocationObject for customs/VAT support
- Fix naming for the address codes

## [v0.4.1]

- Add support for location codes

## [v0.4.0]

- Update lib configuration to be compatible with golang 1.23

## [v0.3.0] - 2025-05-16

- Add oauth/token endpoint

## [v0.2.0] - 2025-05-15

- Add Authorization header and oauth/token endpoint

## [v0.1.0] - 2025-05-13

- Initial go sdk creation
