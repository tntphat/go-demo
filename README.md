# gateway

### Ideas:
- All income APIs are handled by API gateway.
- The authentication and authorization are included as well (use JWT).

### Services port:
- api-gateway: 1000
- user-service: 1001

### How to run:
- copy conf.yaml.default to conf.yaml
- go run server.go