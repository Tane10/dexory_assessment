# File Upload and Comparison Web Application

**NOTE: Requires Go version 1.23 or higher**

## Overview

This web application allows users to upload CSV or JSON files and compare the contents of the uploaded files. It provides an intuitive interface for selecting files, viewing comparison reports, and downloading the results.

## Features

- **File Upload**: Upload CSV or JSON files for comparison.
- **File Comparison**: Select files to compare after uploading.
- **Report Generation**: Generates a detailed comparison report based on the file contents.
- **User Interface**: Built with HTML, CSS, and Bootstrap.
- **Report Interface**: Reports are displayed in an easy-to-read table format.

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS, Bootstrap
- **JavaScript**: For asynchronous requests
- **Data Formats**: JSON and CSV

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/tane10/dexory_assignment.git
   cd dexory_assignment
   ```

2. **Install Dependencies**:

   Ensure you have Go installed. Download the latest version from [Go's official website](https://golang.org/dl/).

   Once Go is installed, run the following command to install necessary Go modules:

   ```bash
   go mod tidy
   ```

3. **Run the Application**:

   To start the web application, navigate to the project directory and run:

   ```bash
   go run cmd/main.go
   ```

   The server will start at `http://localhost:8080`.

## Usage

1. Open your web browser and navigate to `http://localhost:8080`.
2. Use the _Choose a file_ button to upload either a CSV or JSON file from your local machine.
3. After uploading, you can view the uploaded files by clicking _Select Files_.
4. Select one CSV file and one JSON file, then click the _Click to Generate Reports_ button to generate a report.
5. To view generated reports, click _Select a Report_ and then click _View_.
6. The report will be displayed in a table format.
7. If needed, download the generated report by selecting a report and clicking _Download_.

## How to Run Tests

### Running Unit Tests:

You can run all unit tests in the project by executing:

```bash
go test ./... -v
```

This will run the tests in verbose mode, providing detailed information about each test.

### HTTP API Testing:

In addition to unit tests, you can test API endpoints using the `test.http` file provided. This file contains predefined HTTP requests for testing the API.

To run these tests:

1. Open `test.http` in an IDE that supports HTTP requests (such as VS Code or IntelliJ).
2. Execute the individual HTTP requests directly from the IDE, or use a tool like `curl` or [HTTPie](https://httpie.io/) to manually send requests.

### Bash Script for File Upload:

You can also test file uploads via the provided bash script `robotFileUpload.sh`. This script allows uploading files to the server via the terminal.

1. Ensure the script has executable permissions:

   ```bash
   chmod +x scripts/robotFileUpload.sh
   ```

2. Run the script by specifying the file path and file type (either `csv` or `json`):

   ```bash
   ./scripts/robotFileUpload.sh path/to/file.csv csv
   ```
