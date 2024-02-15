// NavigateTo located in nav.js which must be loaded first
document.addEventListener("DOMContentLoaded", function () {
    var fragment = window.location.hash.substring(1);
    if (fragment) {
        navigateTo(fragment)
    } else {
        navigateTo("Home")
    }
});