# New Release HOW-TO

* Add release notes for the new version to CHANGELOG.md.

* Search & replace the current version number with the new one. (Except in CHANGELOG.md)

* ```make dist```

* git commit

* (e.g.) ```git tag -m 'release-0.10.0-alpha' -a release-0.10.0-alpha && git push origin --tags```

