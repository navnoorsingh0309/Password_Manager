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
            "name": document.getElementById('Name_Signup').value,
            "email": document.getElementById('Email').value,
            "password": document.getElementById('Password').value,
        })
    }
    fetch(apiLink + '/signup', reqOptions)
    .then(response => response.json())
    .then((data)=> {
        if (data.message !== "Error") {
            alert("Account Created Successfully!!");
            document.getElementById('Name').innerHTML = "";
            document.getElementById('Email').innerHTML = "";
            document.getElementById('Password').innerHTML = "";
        } else {
            alert("An unexpected error occurred. Please try again");
        }
    })
}

async function onSignIn(e) {
    e.preventDefault();
    // Creating CORS request to my api
    const reqOptions = {
        mode: 'cors',
        method: 'POST',
        // To get the cookie from backend
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
        if (data.id !== -1) {
            // Storing JWT Token in cookies
            document.cookie = "jwt=" + data.token;
            document.cookie = "user_id=" + data.id;
            // Redirecting to main page after successful login with JWT Token in cookies.
            alert("Logged In Successfully!!");
            document.getElementById('Email_Signin').innerHTML = "";
            document.getElementById('Password_Signin').innerHTML = "";
            window.location.href = "http://127.0.0.1:5500/JWT_Authentication/web-app/main-page.html";
        } else {
            alert("An unexpected error occurred. Please try again");
        }
    });   
}