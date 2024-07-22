FROM alpine:edge AS builder

WORKDIR /app

COPY . ./

RUN apk add --no-cache go git ca-certificates build-base \
 && go build -buildmode=pie -ldflags="-s -w" -trimpath -o summaly main.go

FROM alpine:edge

WORKDIR /app

COPY --from=builder /app/backdance /app/backdance

RUN apk add --no-cache ca-certificates tini \
 && addgroup -g 743 app \
 && adduser -u 743 -G app -D -h /app app \
 && chmod +x /app/backdance \
 && chown -R app:app /app

USER app
CMD [ "tini", "--", "/app/backdance" ]