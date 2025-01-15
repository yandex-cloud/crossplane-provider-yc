# Changelog

## 0.7.0 - January 15, 2025
### Added
* support for iam resources most of IAM resources

## v0.6.1 - December 18, 2024
### Fixed
* Allow ID references in bucket/grant.

## v0.6.0
### Added
* Add CDN resource.

### Fixed
* Revert terraform provider version to 0.123.0.

## v0.5.1
### Fixed
* Bump terraform provider version to 0.130.0.

## v0.5.0
### Added
* Support for opensearch cluster.

## v0.4.1

### Fixed
* External repository name and documentation.

## v0.4.0
### Added
* Support for MongodbUser and MongodbDatabase.
* "No-fork" handling for MongodbCluster and RedisCluster.

### Fixed
* Disallow changing FolderIAMMember role after creation.

## v0.3.2
### Fixed
* Export Compute instance FQDN and IP addresses to connection details secret.

## v0.3.1
### Fixed
* Export Kafka cluster FQDN to connection details secret.

## v0.3.0
### Added
* Support for Kafka resouces (clusters, topics, users, connectors).
