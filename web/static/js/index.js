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
      alert(
        `File: ${file.name} has been uploaded! \nFile can be found if you click Select Files.`
      );
      window.location.reload();
    }
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

  const viewReport = document.getElementById("view-report-btn");
  const downloadReport = document.getElementById("download-report-btn");

  const allUnchecked = Array.from(checkboxes).every(
    (checkbox) => !checkbox.checked
  );

  if (selectedCheckbox.checked) {
    viewReport.disabled = false;
    downloadReport.disabled = false;
  }

  if (allUnchecked) {
    viewReport.disabled = true;
    downloadReport.disabled = true;
  }
};

const generateReportHandler = async (event) => {
  if (event) event.preventDefault();
  const checkedBoxes = document.querySelectorAll(".item-checkbox:checked");
  const selectedFiles = Array.from(checkedBoxes).map((cb) => cb.value);

  if (selectedFiles.length !== 2) {
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

    if (generateReportResp.ok) {
      const reportData = await generateReportResp.json();
      alert(
        `Report: ${reportData.filename} has been generated! \nReport can be found if you click Select a Report.`
      );
      window.location.reload();
    } else {
      console.error("Failed to generate the report");
    }
  } catch (error) {
    console.error(error);
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

  if (selectedReport.length === 0) {
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
            <td>${item.location}</td>
            <td>${item.scanned ? "Yes" : "No"}</td>
            <td>${item.occupied ? "Yes" : "No"}</td>
            <td>${
              item.expectedItems
                ? item.expectedItems
                    .split(", ")
                    .map((barcode) => barcode.trim())
                    .join(", ")
                : ""
            }</td>
            <td>${
              item.detectedBarcodes
                ? item.detectedBarcodes
                    .split(", ")
                    .map((barcode) => barcode.trim())
                    .join(", ")
                : ""
            }</td>
            <td>${item.outcome}</td>
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
  const selectedReport = document.querySelectorAll(".report-item-input");

  let report = "";
  selectedReport.forEach((r) => {
    if (r.checked) {
      report = r.value;
    }
  });

  if (selectedReport.length === 0) {
    alert("Please select a report to view.");
    return;
  }

  try {
    const response = await fetch(
      `/view?file=reports/${report}&action=download`
    );
    if (!response.ok) {
      throw new Error(response);
    }

    const blob = await response.blob();

    const url = window.URL.createObjectURL(blob);
    const link = document.createElement("a");
    link.href = url;

    link.download = report.split(".json")[0] + ".csv";

    document.body.appendChild(link);

    link.click();

    document.body.removeChild(link);

    window.URL.revokeObjectURL(url);
  } catch (err) {
    throw err;
  }
};

let checkedBoxes = [];

const handleSelectFilesCheckboxChange = (selectedCheckbox) => {
  if (selectedCheckbox.checked) {
    if (checkedBoxes.length >= 2) {
      // Uncheck the first selected checkbox
      const firstChecked = checkedBoxes.shift();
      firstChecked.checked = false; // Uncheck it
    }
    // Add the newly checked checkbox to the array
    checkedBoxes.push(selectedCheckbox);
  } else {
    // If the checkbox is being unchecked, remove it from the array
    checkedBoxes = checkedBoxes.filter((cb) => cb !== selectedCheckbox);
  }

  let jsonCount = 0;
  let csvCount = 0;

  for (const cb of checkedBoxes) {
    if (cb.value.match(".json")) jsonCount++;
    if (cb.value.match(".csv")) csvCount++;
  }

  const genReportButton = document.getElementById("generate-report");

  genReportButton.disabled = !(
    checkedBoxes.length === 2 &&
    jsonCount === 1 &&
    csvCount === 1
  );
};
