# File Upload and Comparison Web Application

## Overview

This web application allows users to upload CSV or JSON files and compare the contents of the uploaded files. It provides a user-friendly interface for selecting files, viewing reports, and downloading results.

## Features

- **File Upload**: Users can upload CSV or JSON files.
- **File Comparison**: After uploading, users can select files to compare.
- **Report Generation**: The application generates a comparison report based on the uploaded files.
- **User Interface**: A clean and simple UI built with HTML, CSS, and Bootstrap for responsive design.
- **Scrollable Table**: Reports are displayed in a scrollable table format for easy viewing.

## Technologies Used

- Go (Golang) for the backend server
- HTML, CSS, Bootstrap for the frontend
- JavaScript for asynchronous requests
- JSON and CSV formats for data storage and exchange

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/tane10/dexory_assignment.git
   cd dexory_assignment
   ```

2. **Install Dependencies**:
   Make sure you have Go installed. You can download it from [the official Go website](https://golang.org/dl/).

3. **Run the Application**:
   Navigate to the project directory and run:
   ```bash
   go run main.go
   ```
   The application will start running on `http://localhost:8080`.

## Usage

1. Open your web browser and navigate to `http://localhost:8080`.
2. Use the "Choose a file to upload" button to select a CSV or JSON file from your local machine.
3. After uploading, you can select files from the dropdown to compare.
4. Click the "Click to Generate Reports" button to view the comparison report.
5. Reports will be displayed in a scrollable table format.
6. Optionally, you can download the report as needed.

## Directory Structure

```plaintext
/your_project
├── internal                  # Internal packages for private use
│   └── directory_utils.go    # Utility functions for handling directories
├── static                    # Static assets (CSS, JS)
│   ├── css
│   └── js
├── templates                 # HTML templates for rendering views
│   └── index.html
└── main.go                   # Main application entry point
```

## Functionality

### File Upload

- Users can upload files via a form that restricts uploads to CSV and JSON formats.

### File Comparison

- The application compares the contents of the uploaded files and generates a report displaying:
  - The name of the location
  - Whether or not the location was successfully scanned
  - Whether or not the location was occupied
  - The expected and detected barcodes

### Report Generation

- A report is generated and displayed in a scrollable HTML table, allowing for easy access to information.
