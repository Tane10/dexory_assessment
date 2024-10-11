const fileUploadHandler = async (event) => {
  event.preventDefault();

  //TODO: Fix index.js:19 => POST http://localhost:8080/upload?type=json 415 (Unsupported Media Type)

  const fileInput = document.getElementById("file");
  const file = fileInput.files[0]; // Get the selected file

  if (!file) {
    alert("Please select a file to upload.");
    return;
  }

  const fileType = file.name.split(".").pop(); // Get the file extension (e.g., csv, json)

  const formData = new FormData(); // Create a FormData object
  formData.append("file", file); // Append the file to the FormData object
  formData.append("type", fileType); // Append the file type (or read from the hidden input)

  try {
    const response = await fetch(`/upload?type=${fileType}`, {
      method: "POST",
      body: formData, // Set the body to the FormData object
    });

    if (response.ok) {
      const result = await response.text(); // Get the response as text (or JSON if you prefer)

      console.log(result);
      document.getElementById(
        "message"
      ).innerText = `Upload successful: ${result}`;
    }

    // else {
    //   const errorText = await response.text();
    //   document.getElementById(
    //     "message"
    //   ).innerText = `Upload failed: ${errorText}`;
    // }
  } catch (error) {
    console.error(error);
  }
};

const handleReportCheckboxChange = (selectedCheckbox) => {
  // Get all checkboxes
  const checkboxes = document.querySelectorAll(".report-item-input");
  // Loop through each checkbox and uncheck if it's not the selected one
  checkboxes.forEach((checkbox) => {
    if (checkbox !== selectedCheckbox) {
      checkbox.checked = false; // Uncheck the checkbox
    }
  });
};

const generateReportHandler = async (event) => {
  if (event) event.preventDefault();
  const checkedBoxes = document.querySelectorAll(".item-checkbox:checked");
  const selectedFiles = Array.from(checkedBoxes).map((cb) => cb.value);

  console.log(selectedFiles);

  if (selectedFiles.length != 2) {
    alert("Please select 2 files");
    return;
  }

  try {
    const generateReportResp = await fetch("/report", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        files: selectedFiles,
      }),
    });
    if (!generateReportResp.ok) {
      throw new Error("Failed to generate the report");
    }
    const reportData = await generateReportResp.json();
    console.log(reportData);
  } catch (error) {
    console.error(error);
    return;
  }
};

const reportDisplayHandler = async () => {
  const selectedReport = document.querySelectorAll(".report-item-input");

  let report = "";
  selectedReport.forEach((r) => {
    if (r.checked) {
      report = r.value;
    }
  });

  if (selectedReport.length == 0) {
    alert("Please select a report to view.");
    return;
  }

  try {
    const response = await fetch(`/view?file=reports/${report}`);
    if (!response.ok) {
      throw new Error(response);
    }

    const data = await response.json();

    const table = document.createElement("table");
    table.className = "table table-bordered"; // Use Bootstrap classes for styling

    const thead = document.createElement("thead");

    const headerRow = document.createElement("tr");
    headerRow.innerHTML = `
        <th>Location</th>
        <th>Scanned</th>
        <th>Occupied</th>
        <th>Expected Barcodes</th>
        <th>Detected Barcodes</th>
        <th>Outcome</th>
    `;
    thead.appendChild(headerRow);
    table.appendChild(thead);

    // Create table body
    const tbody = document.createElement("tbody");
    data.forEach((item) => {
      const row = document.createElement("tr");
      row.innerHTML = `
            <td>${item.Location}</td>
            <td>${item.Scanned ? "Yes" : "No"}</td>
            <td>${item.Occupied ? "Yes" : "No"}</td>
            <td>${
              item.ExpectedItems
                ? item.ExpectedItems.split(", ")
                    .map((barcode) => barcode.trim())
                    .join(", ")
                : ""
            }</td>
            <td>${
              item.DetectedBarcodes
                ? item.DetectedBarcodes.split(", ")
                    .map((barcode) => barcode.trim())
                    .join(", ")
                : ""
            }</td>
            <td>${item.Outcome}</td>
        `;
      tbody.appendChild(row);
    });
    table.appendChild(tbody);

    // Clear previous table and append the new table to the container
    const reportTableContainer = document.getElementById(
      "reportTableContainer"
    );

    const reportContainerParent = document.getElementById("reportTableParent");
    reportContainerParent.hidden = true;
    reportTableContainer.innerHTML = ""; // Clear previous content
    reportTableContainer.appendChild(table);
    reportContainerParent.hidden = false;
  } catch (err) {
    console.error(err);
    return;
  }
};

const downloadHandler = async () => {
  try {
    const response = await fetch(
      `/view?file=reports/report_10-10-2024_20:17:19_kPm0bHOaQdOkGOFHkDfmyQ.json&action=download`
    );
    if (!response.ok) {
      throw new Error(response);
    }
  } catch (err) {
    throw err;
  }
};
