# Event Markup Language Compliance Test

Tests whether an API created from EML by a builder such as les-node complies with the EML specification.

* Are all the command processor end points working?
* Are the read models working?
* Are all the business rules implemented?
* Are the required validation errors returned when executing invalid commands?

## Test EMD

Tests all features which can be used frm Event Markdown

```cd emd && make setup && sleep 1 && make test```


## Test EML

Tests all features which are supported in Event Markup Language 0.10.1-alpha, but are not part of EMD.

```cd eml && make setup && sleep 1 && make test```
