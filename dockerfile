FROM golang:1.24-alpine as go_builder

WORKDIR /usr/src

COPY go.* ./
RUN go mod download 

COPY . . 

RUN go build -v -o fullstack_app

FROM alpine:latest as app

WORKDIR /app
COPY --from=go_builder /usr/src/fullstack_app .
## copy other things here like any sort of frontend/dist
# COPY --from=bun_builder /usr/src/frontend/dist ./dist

CMD [ "fullstack_app", "serve" ]
