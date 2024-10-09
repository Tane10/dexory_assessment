CURL Tests:

```bash
curl -X GET -H "Content-Type: application/json" http://localhost:8080/upload
```

Functional requirements
● A web application is available which can consume JSON from a robot and store it in an
appropriate data store.
● A user can select a previously-uploaded JSON payload, and upload a CSV report to
generate a comparison report.
● A user can view or export the comparison report in an appropriate format.
● The JSON and CSV upload endpoints accept the format supplied in the sample files
accompanying this coding exercise.

NOTE: there is a 10mb file upload limit on the upload route meaning the robot will have to chuck data in MAX 10mb files when doing the scans. this is because AWS API Gateway has a 10MB payload 1limit


NOTE: We are going under the assumption that the JSON files are being sent from the robot name like so: report-{MS timestamp}-{base64 encoded UUID}.json => report-1632931462-550e8400.json


TODO:

#### 1. Set Up the Server and Routes

```
FUNCTION main():
    INITIALIZE a web server
    DEFINE routes:
        GET /view/{id} -> viewHandler
        POST /upload -> uploadHandler
        GET /report -> reportHandler
        POST /report -> reportHandler
        POST /export -> exportHandler
    START the server on port 8080
```

#### 2. Create a GET Request for `/view/{id}`

```
FUNCTION viewHandler(w http.ResponseWriter, r *http.Request):
    EXTRACT the {id} from the URL parameters
    IF {id} is provided THEN:
        SEARCH for the file corresponding to {id}
        IF file found THEN:
            RETURN the file data as JSON
        ELSE:
            RETURN 404 Not Found
    ELSE:
        RETURN all files as JSON
```

#### 3. Create a POST Request for `/upload` [DONE]

```
FUNCTION uploadHandler(w http.ResponseWriter, r *http.Request):
    DETERMINE the Content-Type from the request header
    IF Content-Type is "application/json" THEN:
        PARSE the JSON data
        VALIDATE the data format
        SAVE the data to a file or database
        RETURN success response
    ELSE IF Content-Type is "text/csv" THEN:
        READ the CSV data
        VALIDATE the data format
        SAVE the data to a file or database
        RETURN success response
    ELSE:
        RETURN 400 Bad Request for unsupported Content-Type
```

#### 4. Create a Data Comparison Function

```
FUNCTION compareData(csvData, jsonData):
    CONVERT both datasets into a comparable format (e.g., maps, slices)
    COMPARE the two datasets:
        FOR each item in csvData:
            CHECK if it exists in jsonData
        RETURN comparison result (e.g., differences found)
```

#### 5. Create a GET & POST Request for `/report`

```
FUNCTION reportHandler(w http.ResponseWriter, r *http.Request):
    IF r.Method is GET THEN:
        FETCH existing reports from the storage
        RETURN reports as JSON
    ELSE IF r.Method is POST THEN:
        PARSE the request body for new report criteria
        VALIDATE the criteria
        GENERATE a new report based on the criteria
        SAVE the report
        RETURN success response
    ELSE:
        RETURN 405 Method Not Allowed
```

#### 6. Create a POST Request for `/export`

```
FUNCTION exportHandler(w http.ResponseWriter, r *http.Request):
    PARSE the request body for export criteria (e.g., report ID)
    VALIDATE the criteria
    GENERATE the export file (e.g., CSV or PDF)
    SET appropriate headers for file download (Content-Disposition)
    WRITE the file content to the response writer
```
