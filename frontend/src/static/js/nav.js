const navLinks = ['Home', 'About', 'Games', 'Family', 'Career']
// const apiEndpoint = '/spa/';
const apiEndpoint = 'http://localhost:8080/spa'
navLinks.forEach((link) => {
    document.getElementById(`nav${link}Link`).addEventListener("click", function () {
        navigateTo(link);
    });
})

function navigateTo(section) {

    console.log("Navigating to", section);

    fetch(apiEndpoint, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "TemplateFile": section.toLowerCase(),
            "Location": "hello"
        })
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.text()

        })
        .then(text => {
            document.getElementById("mainSpaContent").innerHTML = text

        })
}