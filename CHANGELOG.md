# Releases

## 0.10.4-alpha (April 30, 2018)

### New Features

* les-node uses newly opensourced les-node-template (https://github.com/Adaptech/les-node-template) (version 20180430-237a5d5)

## 0.10.3-alpha (April 21, 2018)

### New Features

Command preconditions: Do not execute the command if the precondition isn't met:

* "UserRegistered MustHaveHappened"
* "UserDeleted MustNotHaveHappened"

Automated EML v0.10.x specification [compliance test suite for generated APIs](cmd/eml-compliance-test/README.md) to verify that command validation rules and read models behave as expected.

### Bug Fixes

* Undefined readmodel field causes ''Invalid payload for model' in read model.

## 0.10.2-alpha (April 18, 2018)

* Changed default EML and EMD file names so they say what's in the files.

* Running les and les-node without installation, via Docker.

* EMD parameter & properties validation bug fix.

* Automated multi-platform builds.

* Automated platform-specific installation instructions.

## 0.10.1-alpha (April 15, 2018)

Minor bug fixes

## 0.10.0-alpha (April 11, 2018)

Initial release of 'les' and 'les-node'.
