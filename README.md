fronted [![Coverage Status](https://coveralls.io/repos/getlantern/fronted/badge.png)](https://coveralls.io/r/getlantern/fronted)&nbsp;[![GoDoc](https://godoc.org/github.com/getlantern/fronted?status.png)](http://godoc.org/github.com/getlantern/fronted)
==========
To install:

`go get github.com/getlantern/fronted`

For docs:

`godoc github.com/getlantern/fronted`

See [ddftool](https://github.com/getlantern/ddftool) for more details on how to generate and tests fronting domains for the supported CDNs.

[!NOTE]
Since the masquerade domains and IP addresses can change, tests might fail and they need to be updated. You can basically ping some of the masquerade domains (from `default_masquerade.go`) and update the IPs accordingly.

To generate an updated domain fronting configuration file, just run:

```
./updateFrontedConfig.bash
```