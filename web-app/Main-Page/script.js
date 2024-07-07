// const apiLink = "https://critical-tobe-student-iitropar-6d8618df.koyeb.app"
const apiLink = "http://localhost:8000"

window.onload = function() {
    // Setting user id
    document.getElementById("user_id").innerHTML = "User Id: " + getCookieValue("user_id")

    // Getting passwords
    fetch(apiLink + "/getpasses?", {
        method: "cors",
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "x-jwt-token": getCookieValue("jwt"),
        }
    })
    .then(respose => respose.json())
    .then(data => {
        data.forEach(element => {
            let table = document.getElementById("Table_Body");
            let table_row = document.createElement("tr");
            let table_col1 = document.createElement("td");
            let table_col2 = document.createElement("td");
            table_col1.setAttribute("data-title", "Entity");
            table_col1.innerHTML = element["entity"];
            table_col2.setAttribute("data-title", "Password");
            table_col2.innerHTML = element["password"];
            table_row.appendChild(table_col1);
            table_row.appendChild(table_col2);
            table.appendChild(table_row);
        });
    });
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