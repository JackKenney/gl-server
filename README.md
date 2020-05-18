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

## Resources

The main source for containerizing the programs is from here:
https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e 

The main standards for structuring the project came from here:
https://github.com/golang-standards/project-layout
