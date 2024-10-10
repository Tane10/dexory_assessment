function handlerFileUpload() {
  document
    .getElementById("uploadForm")
    .addEventListener("submit", async function (event) {
      event.preventDefault();

      const fileInput = document.getElementById("file");
      const file = fileInput.files[0]; // Get the selected file

      if (!file) {
        alert("Please select a file to upload.");
        return;
      }

      const type = file.name.split(".")[1];

      const formData = new FormData(); // Create a FormData object
      formData.append("file", file); // Append the file to the FormData object
      formData.append("type", type); // Append the file type (or read from the hidden input)

      try {
        const response = await fetch(`/upload?type=${type}`, {
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
    });
}

document.addEventListener("DOMContentLoaded", function () {
  const dropdown = document.querySelector(".dropdown");
  const dropbtn = document.querySelector(".dropbtn");

  // Toggle dropdown visibility on button click
  dropbtn.addEventListener("click", function () {
    dropdown.querySelector(".dropdown-content").classList.toggle("show");
  });

  // Close dropdown if clicked outside
  window.addEventListener("click", function (event) {
    if (!event.target.matches(".dropbtn")) {
      const dropdowns = document.getElementsByClassName("dropdown-content");
      for (let i = 0; i < dropdowns.length; i++) {
        const openDropdown = dropdowns[i];
        if (openDropdown.classList.contains("show")) {
          openDropdown.classList.remove("show");
        }
      }
    }
  });

  handlerFileUpload();
});
