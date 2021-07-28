FROM frolvlad/alpine-glibc
COPY ./build/web-teller /app/
WORKDIR /app
ENTRYPOINT ["/app/web-teller"]