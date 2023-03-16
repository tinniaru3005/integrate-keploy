# integrate-keploy

## Go Application with Keploy Integration

This is a simple Go application that demonstrates how to integrate with Keploy for deploying and managing application releases. Keploy is a deployment and release management platform that helps teams ship code faster and more reliably.

## Getting Started

To get started with this application, you will need to have Go installed on your system. You will also need to sign up for a Keploy account and obtain your API key and secret.

### Installation

Clone this repository to your local machine:

```
git clone https://github.com/tinniaru3005/integrate-keploy
```

Navigate to the project directory:
```
cd go-keploy-integration
```

Install the dependencies:
```
go mod download
```
### Configuration
Set the following environment variables with your Keploy API key, secret, and environment name:
```
export KEPLOY_API_KEY=your_api_key_here
export KEPLOY_API_SECRET=your_api_secret_here
export KEPLOY_ENVIRONMENT=your_environment_name_here
```

### Usage
Build and run the application:
```
go run main.go
```

Visit http://localhost:8080/ in your web browser to see the application running.

## Contributing
Contributions are welcome! If you find any bugs or issues, please open an issue in this repository.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Code Sample Description

This code does the following:

Defines the name and version of the Go application.
Sets up the Keploy client using the Keploy API key and secret stored in environment variables.
Builds the Go binary.
Creates a new release in Keploy.
Uploads the binary to the release in Keploy.
Deploys the release to the environment specified in an environment variable.
Starts an HTTP server listening on port 8080.
Prints the deployment status.
