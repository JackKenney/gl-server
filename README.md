# Generic Learning Server

This server is a REST server that takes training data at REST endpoints and makes predictions from other endpoints.

## Building

To build image run following command:
```bash
docker build . -t gl-server
```

## Running

To run an image, type following:
```bash
docker run -p 3000:3000 gl-server
```

## Testing

```bash
go test ./...
```
