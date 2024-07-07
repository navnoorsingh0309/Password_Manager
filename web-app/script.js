const wrapper = document.querySelector('.wrapper');
const signUpLink = document.querySelector('.signUp-link');
const signInLink = document.querySelector('.signIn-link');
const signInForm = document.querySelector('.sign-in');
const signUpForm = document.querySelector('.sign-up');

const apiLink = "https://critical-tobe-student-iitropar-6d8618df.koyeb.app"

signUpLink.addEventListener('click', () => {
    wrapper.classList.add('signup-clicking');
    wrapper.classList.remove('signin-clicking');
});

signInLink.addEventListener('click', () => {
    wrapper.classList.add('signin-clicking');
    wrapper.classList.remove('signup-clicking');
});

function onSignUp() {
    const reqOptions = {
        mode: 'no-cors',
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

function onSignIn() {
    const reqOptions = {
        mode: 'no-cors',
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            email: document.getElementById('Email_Signin').value,
            password: document.getElementById('Password_Signin').value,
        })
    }
    fetch(apiLink + '/login', reqOptions)
    .then(response => response.json())
    .then((data)=> {
        alert("SignIn Successful");
        document.getElementById('Email_Signin').innerHTML = "";
        document.getElementById('Password_Signin').innerHTML = "";
    })
}