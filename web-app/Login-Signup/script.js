const wrapper = document.querySelector('.wrapper');
const signUpLink = document.querySelector('.signUp-link');
const signInLink = document.querySelector('.signIn-link');
const signInForm = document.querySelector('.sign-in');
const signUpForm = document.querySelector('.sign-up');

signUpLink.addEventListener('click', () => {
    wrapper.classList.add('signup-clicking');
    wrapper.classList.remove('signin-clicking');
});

signInLink.addEventListener('click', () => {
    wrapper.classList.add('signin-clicking');
    wrapper.classList.remove('signup-clicking');
});