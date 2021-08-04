# Opinionated template for Go + React + TypeScript web apps using single binary bundling

## The template contains the following:
- a Go project (/app)   
- a React/TypeScript project (/webapp) created with create-react-app
- some (again, opinionated) basic configuration for VS Code

### For the Go app, the following is used:
- [Gorilla Mux](github.com/gorilla/mux) for routing
- a sample API using a subrouter
- a file server that bundles content of a 'dist' folder into the binary using `//go:embed` and serves the content from there

### For the React app
- It's just a default CRA-generated app, except
- it also contains a .env file which ensures that running `npm run build` will copy the output files into the proper directory in the Go app tree for bundling purposes