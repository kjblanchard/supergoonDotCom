const navLinks = ['Home', 'About', 'Games', 'Family', 'Career']
navLinks.forEach((link) => {
    document.getElementById(`nav${link}Link`).addEventListener("click", function () {
        navigateTo(link);
    });
})

function navigateTo(section) {
    document.getElementById("mainSpaContent").innerHTML = section
    console.log("Navigating to", section);
}