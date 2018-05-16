FROM alpine AS build

ENV IPCHECK_VERSION="0.2"

RUN apk --no-cache add curl \
 && curl -o ipcheck https://github.com/ahammond/ipcheck/releases/download/v$IPCHECK_VERSION/ipcheck \
 && chmod a+x ipcheck

FROM scratch
COPY --from=build ipcheck /

ENTRYPOINT ["/ipcheck"]
