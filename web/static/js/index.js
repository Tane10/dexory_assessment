document.addEventListener("DOMContentLoaded", function () {
    const dropdown = document.querySelector('.dropdown');
    const dropbtn = document.querySelector('.dropbtn');

    // Toggle dropdown visibility on button click
    dropbtn.addEventListener('click', function () {
        dropdown.querySelector('.dropdown-content').classList.toggle('show');
    });

    // Close dropdown if clicked outside
    window.addEventListener('click', function (event) {
        if (!event.target.matches('.dropbtn')) {
            const dropdowns = document.getElementsByClassName("dropdown-content");
            for (let i = 0; i < dropdowns.length; i++) {
                const openDropdown = dropdowns[i];
                if (openDropdown.classList.contains('show')) {
                    openDropdown.classList.remove('show');
                }
            }
        }
    });
});
