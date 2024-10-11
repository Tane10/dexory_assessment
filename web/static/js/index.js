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

// document.addEventListener("DOMContentLoaded", () => {});
