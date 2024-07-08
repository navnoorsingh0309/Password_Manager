// const apiLink = "https://critical-tobe-student-iitropar-6d8618df.koyeb.app"
const apiLink = "http://localhost:8000"

window.onload = function() {
    // Setting user id
    document.getElementById("user_id").innerHTML = "User Id: " + getCookieValue("_UR_SH_ID_")

    // Getting passwords
    fetch(apiLink + "/getpasses", {
        method: "cors",
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "x-jwt-token": getCookieValue("_PM_RPR_TK_AC_"),
        }
    })
    .then(respose => respose.json())
    .then(data => {
        data.forEach(element => {
            let table = document.getElementById("Table_Body");
            let table_row = document.createElement("tr");
            let table_col1 = document.createElement("td");
            let table_col2 = document.createElement("td");
            let table_col3 = document.createElement("td");
            table_col1.setAttribute("data-title", "Entity");
            table_col1.innerHTML = element["entity"];
            table_col2.setAttribute("data-title", "Email");
            table_col2.innerHTML = element["email"];
            table_col3.setAttribute("data-title", "Password");
            table_col3.innerHTML = element["password"];
            table_row.appendChild(table_col1);
            table_row.appendChild(table_col2);
            table_row.appendChild(table_col3);
            table.appendChild(table_row);
        });
    });
}

// Signing out
function signOut() {
    // Expiring the cookie
    document.cookie = "_PM_RPR_TK_AC_=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;"
    document.cookie = "_UR_SH_ID_=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;"
    window.location.href = "index.html";
}

// Creating new password
function newPassword() {
    document.getElementById("dialog").style.display = "block";
    document.getElementById("dialogOverlay").style.display = "block";
}
function closeNewPassModal() {
    document.getElementById("dialog").style.display = "none";
    document.getElementById("dialogOverlay").style.display = "none";
}
function addNewPassword(e) {
    e.preventDefault();
    // Creating new password
    fetch(apiLink + "/newpass", {
        method: "cors",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "x-jwt-token": getCookieValue("_PM_RPR_TK_AC_"),
        },
        body: JSON.stringify({
            entity: document.getElementById("Entity").value,
            email:  document.getElementById("Email").value,
            password: document.getElementById("Password").value
        })
    })
    .then(response => response.json())
    .then((data) => {
        if (data.message === "Success") {
            console.log("Password added successfully!!")
        } else {
            console.log("An unexpected error has occured!!")
        }
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