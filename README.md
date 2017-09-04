ipcheck
===

[![Build Status](https://travis-ci.org/ahammond/ipcheck.svg?branch=master)](https://travis-ci.org/ahammond/ipcheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/ahammond/ipcheck)](https://goreportcard.com/report/github.com/ahammond/ipcheck)
[![codecov](https://codecov.io/gh/ahammond/ipcheck/branch/master/graph/badge.svg)](https://codecov.io/gh/ahammond/ipcheck)

A statically linked binary which, given a string decides if it is

- malformed: not an IPv4 address
- public: a routable IPv4 address
- private: a non-routable IPv4 per [rfc1918](https://tools.ietf.org/html/rfc1918)
- loopback: [rfc1700](https://tools.ietf.org/html/rfc1700)
- linklocal: [rfc3927](https://tools.ietf.org/html/rfc3927)
- 6to4: [rfc3068](https://tools.ietf.org/html/rfc3068)
- documentation: [rfc5737](https://tools.ietf.org/html/rfc5737)

See wikipedia for [all the gory details](https://en.wikipedia.org/wiki/Reserved_IP_addresses).

Usage
---

In your dockerfile, add something like this:
```dockerfile
FROM alpine

ENV IPCHECK_VERSION="0.1"

RUN apk --no-cache add curl  \
 && curl -o /usr/local/sbin/ipcheck https://github.com/ahammond/ipcheck/releases/download/v$IPCHECK_VERSION/ipcheck

CMD [ "entrypoint.sh" ]
```

And then in your entrypoint.sh or wherever you're doing validation:

```bash
IP_TYPE=$(ipcheck "$MY_IP")
if [ "public" != "$IP_TYPE" ]; then
  echo "I need a public IP, but $MY_IP is $IP_TYPE"
  exit 1
fi
```
