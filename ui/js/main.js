let button = document.querySelector('#validate-form-button');
button.addEventListener('click', function (){
    let nameInput = document.querySelector('#name');
    let surNameInput = document.querySelector('#surname');
    let emailInput = document.querySelector('#email');
    let passwordInput = document.querySelector('#password');
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
        console.log(data);
        Validation(data);
    })
        .catch ((error) => {
        console.error('Error:', error);
    });

})

function Validation(data) {
    for (let i = 0; i < data.Errors.length; i++) {

        let error = data.Errors[i];

    }
}