FROM docker.io/golang:1.21.5-bullseye as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/tftools

CMD ["tftools"]

# Distroless image
# FROM gcr.io/distroless/static-debian12:debug
# COPY --from=build /go/bin/tftools /app/tftools

# ENV plan=/tmp/plan.json

# Pending to add entrypoint
# CMD ["/app/tftools", "summarize", "<", $plan]