FROM apline

ENV IPCHECK_VERSION="0.2"
RUN curl --output /usr/local/sbin/ipcheck https://github.com/ahammond/ipcheck/releases/download/v$IPCHECK_VERSION/ipcheck
