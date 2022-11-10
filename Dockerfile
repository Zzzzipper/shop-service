# Build stage
FROM golang AS build-env
ADD . /src/merchant-service
ENV CGO_ENABLED=0
RUN cd /src/merchant-service/cmd && go build -o /app

# Production stage
FROM scratch
COPY --from=build-env /app /
ENV POSTGRES_URL=postgresql://postgres:password@172.17.0.2:5432/merchant?sslmode=disable
ENTRYPOINT ["/app"]
