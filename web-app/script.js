const wrapper = document.querySelector('.wrapper');
const signUpLink = document.querySelector('.signUp-link');
const signInLink = document.querySelector('.signIn-link');
const signInForm = document.querySelector('.sign-in');
const signUpForm = document.querySelector('.sign-up');

// const apiLink = "https://critical-tobe-student-iitropar-6d8618df.koyeb.app"
const apiLink = "http://localhost:8000"

signUpLink.addEventListener('click', () => {
    wrapper.classList.add('signup-clicking');
    wrapper.classList.remove('signin-clicking');
});

signInLink.addEventListener('click', () => {
    wrapper.classList.add('signin-clicking');
    wrapper.classList.remove('signup-clicking');
});

function onSignUp(e) {
    e.preventDefault();
    // Creating CORE request to my api
    const reqOptions = {
        mode: 'cors',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name: document.getElementById('Name').value,
            email: document.getElementById('Email').value,
            password: document.getElementById('Password').value,
        })
    }
    fetch(apiLink + '/signup', reqOptions)
    .then(response => response.json())
    .then((data)=> {
        alert("Account Created Successfully!!");
        document.getElementById('Name').innerHTML = "";
        document.getElementById('Email').innerHTML = "";
        document.getElementById('Password').innerHTML = "";
    })
}

async function onSignIn(e) {
    e.preventDefault();
    // Creating CORS request to my api
    const reqOptions = {
        mode: 'cors',
        method: 'POST',
        // To get the cookie from backend
        credentials: "include",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "email": document.getElementById('Email_Signin').value,
            "password": document.getElementById('Password_Signin').value,
        })
    }
    fetch(apiLink + '/login', reqOptions)
    .then((response) => response.json() )
    .then((data)=> {
        if (data.message === "Success") {
            alert("Logged In Successfully!!");
            document.getElementById('Email_Signin').innerHTML = "";
            document.getElementById('Password_Signin').innerHTML = "";
            window.location.href = "http://127.0.0.1:5500/JWT_Authentication/web-app/main-page.html";
        } else {
            alert(data.message);
        }
    });   
}