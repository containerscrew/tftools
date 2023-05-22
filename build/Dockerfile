# Start by building the application.
FROM golang:1.19 as build

ARG TOOL_NAME

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/$TOOL_NAME

# Develop mode using :debug tag in distroless (allow shell)
FROM gcr.io/distroless/static-debian11:debug as development
COPY --from=build /go/bin/$TOOL_NAME /
CMD ["/$TOOL_NAME"]

# Distroless for production deployments
FROM gcr.io/distroless/static-debian11 as production
COPY --from=build /go/bin/$TOOL_NAME /
CMD ["/$TOOL_NAME"]
