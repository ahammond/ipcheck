FROM alpine AS build

ARG TAG="v0.3"

RUN apk --no-cache add curl \
 && curl --output ipcheck --silent --location https://github.com/ahammond/ipcheck/releases/download/${TAG}/ipcheck_linux_amd64 \
 && chmod a+x ipcheck

FROM scratch
COPY --from=build ipcheck /

ENTRYPOINT ["/ipcheck"]
