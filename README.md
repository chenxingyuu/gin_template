# gin_template

## Features

V1.0.0

- [x] Config
- [x] Redis、MySQL
- [x] Response
- [x] Docker Deploy

## Usage

### 1. Clone

```shell
git clone https://github.com/chenxingyuu/gin_template.git
```

### 2. Config

```shell
cd gin_template
cp config_template.yaml config.yaml
```

### 3. Run

```shell
go run cmd/server/main.go
```

## Deploy

### 1. Build

```shell
docker build -t gin_template .
```

### 2. Run

```shell
docker run -d -p 8080:8080 --name gin_template gin_template
# or 
docker-compose up -d
```

## Directory

```
├── cmd                     # Directory containing the application's command-line interface.
│   └── server              # Executable for running the server. 
├── config                  # Directory containing the application's configuration files.
│   ├── config.go           # Configuration file parsing.
│   ├── constant.go         # Constant.
│   └── env.go              # Environment variable parsing.
├── deployment              # Directory containing the application's deployment files.
│   ├── Dockerfile          # Dockerfile.
│   └── docker-compose.yaml # Docker-compose.
├── internal                # Directory containing the application's private application and library code.
│   └── app                 # Application code.
├── pkg                     # Directory containing the application's public library code.
│   ├── datastore           # Datastore.
│   └── response            # Response.
├── LICENSE                 # License.    
├── README.md               # README.
├── config_template.yaml    # Configuration template.
├── go.mod                  # Go module file specifying the project's dependencies.
├── go.sum                  # File containing checksums for the project's dependencies.
└── main.go                 # Application entry point.

```



