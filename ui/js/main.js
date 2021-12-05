let button = document.querySelector('#validate-form-button');

let nameInput = document.querySelector('#FirstName');
let surNameInput = document.querySelector('#Surname');
let emailInput = document.querySelector('#Email');
let passwordInput = document.querySelector('#Password');

let nameErrorMessage = document.querySelector('#nameErrorMessage');
let surnameErrorMessage = document.querySelector('#surnameErrorMessage');
let emailErrorMessage = document.querySelector('#emailErrorMessage');
let passwordErrorMessage = document.querySelector('#passwordErrorMessage');

button.addEventListener('click', function (){

    let form = {
        name: nameInput.value,
        surName: surNameInput.value,
        email: emailInput.value,
        password: passwordInput.value,
    }

    fetch(window.location.origin + '/validate', {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(form)
    }).then(response => response.json())
        .then(data => {
            checkValidation(data);
    })
        .catch ((error) => {
        console.error('Error:', error);
    });

})

function checkValidation(data) {
    clearErrors()
    for (let i = 0; i < data.errors.length; i++) {

        let fieldName = data.errors[i].fieldName;
        let errorMessage = data.errors[i].errorMessage;

        if (fieldName == "FirstName"){
            nameInput.classList.add("_err");
            nameErrorMessage.innerHTML = errorMessage;
            nameErrorMessage.classList.add("_errMsg");
        }
        if (fieldName == "Surname"){
            surNameInput.classList.add("_err");
            surnameErrorMessage.innerHTML = errorMessage;
            surnameErrorMessage.classList.add("_errMsg");
        }
        if (fieldName == "Email"){
            emailInput.classList.add("_err");
            emailErrorMessage.innerHTML = errorMessage;
            emailErrorMessage.classList.add("_errMsg");
        }
        if (fieldName == "Password"){
            passwordInput.classList.add("_err");
            passwordErrorMessage.innerHTML = errorMessage;
            passwordErrorMessage.classList.add("_errMsg");
        }
    }
}

function clearErrors() {
    nameInput.classList.remove("_err");
    surNameInput.classList.remove("_err");
    emailInput.classList.remove("_err");
    passwordInput.classList.remove("_err");

    nameErrorMessage.classList.remove("_errMsg");
    surnameErrorMessage.classList.remove("_errMsg");
    emailErrorMessage.classList.remove("_errMsg");
    passwordErrorMessage.classList.remove("_errMsg");

    nameErrorMessage.innerHTML = "";
    surnameErrorMessage.innerHTML = "";
    emailErrorMessage.innerHTML = "";
    passwordErrorMessage.innerHTML = "";
}