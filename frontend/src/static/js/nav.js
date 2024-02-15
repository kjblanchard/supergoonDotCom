const navLinks = ['Home', 'About', 'Games', 'Family', 'Career']
const apiEndpoint = 'http://localhost:8080/spa'
navLinks.forEach((link) => {
    document.getElementById(`nav${link}Link`).addEventListener("click", function () {
        navigateTo(link);
    });
})

function navigateTo(section) {
    fetch(apiEndpoint, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "TemplateFile": section.toLowerCase()
        })
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response from spa was not ok');
            }
            return response.text()
        })
        .then(responseText => {
            document.getElementById("mainSpaContent").innerHTML = responseText
        })
}