FROM golang:1.16 as builder

COPY api.go .

# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
ENV CGO_ENABLED=0
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /api api.go

###

FROM scratch

# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
COPY --from=builder /api .
CMD ["./api"]
