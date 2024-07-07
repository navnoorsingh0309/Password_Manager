// const apiLink = "https://critical-tobe-student-iitropar-6d8618df.koyeb.app"
const apiLink = "http://localhost:8000"

window.onload = function() {
    // Setting user id
    document.getElementById("user_id").innerHTML = "User Id: " + getCookieValue("user_id")
}

// Getting coookie
function getCookieValue(name) 
{
    const regex = new RegExp(`(^| )${name}=([^;]+)`)
    const match = document.cookie.match(regex)
    if (match) {
        return match[2]
    }
}