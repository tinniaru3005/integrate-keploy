package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Define the app's name and version
	appName := "my-go-app"
	appVersion := "1.0.0"

	// Set up Keploy client
	keployClient := NewKeployClient(os.Getenv("KEPLOY_API_KEY"), os.Getenv("KEPLOY_API_SECRET"))

	// Build the Go binary
	err := exec.Command("go", "build", "-o", appName).Run()
	if err != nil {
		fmt.Println("Error building Go binary:", err)
		return
	}

	// Create a new release
	release, err := keployClient.CreateRelease(appName, appVersion)
	if err != nil {
		fmt.Println("Error creating Keploy release:", err)
		return
	}

	// Upload the binary to the release
	err = keployClient.UploadBinary(release.ID, appName, appName)
	if err != nil {
		fmt.Println("Error uploading binary to Keploy release:", err)
		return
	}

	// Deploy the release to the environment
	deployment, err := keployClient.DeployRelease(release.ID, os.Getenv("KEPLOY_ENVIRONMENT"))
	if err != nil {
		fmt.Println("Error deploying Keploy release:", err)
		return
	}

	// Print the deployment status
	fmt.Println("Deployment status:", deployment.Status)

	// Start the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
