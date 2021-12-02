let button = document.querySelector('#validate-form-button');
button.addEventListener('click', function (){
    let nameInput = document.querySelector('#FirstName');
    let surNameInput = document.querySelector('#Surname');
    let emailInput = document.querySelector('#Email');
    let passwordInput = document.querySelector('#Password');
    nameInput.classList.remove("_err");
    surNameInput.classList.remove("_err");
    emailInput.classList.remove("_err");
    passwordInput.classList.remove("_err");

    let form = {
        name: nameInput.value,
        surName: surNameInput.value,
        email: emailInput.value,
        password: passwordInput.value,
    }

    nameInput.style.backgroundColor = 'red solid';

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
            for (let i = 0; i < data.errors.length; i++) {

                let fieldName = data.errors[i].fieldName;
                let errorMessage = data.errors[i].errorMessage;

                let input = document.getElementById(fieldName);
                input.classList.add("_err");
                alert(errorMessage);
            }
    })
        .catch ((error) => {
        console.error('Error:', error);
    });

})
