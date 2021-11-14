
const psw = document.getElementById("password")
const block = document.getElementById("indicator")

psw.addEventListener("input", checkPassword)

function checkPassword() {
    let password = psw.value;

    let s_letters = "qwertyuiopasdfghjklzxcvbnm";
    let b_letters = "QWERTYUIOPLKJHGFDSAZXCVBNM";
    let digits = "0123456789";
    let specials = "!@#$%^&*()_-+=\|/.,:;[]{}";

    let is_s = false;
    let is_b = false;
    let is_d = false;
    let is_sp = false;

    for (let i = 0; i < password.length; i++) {

        if (!is_s && s_letters.indexOf(password[i]) !== -1) is_s = true;
        else if (!is_b && b_letters.indexOf(password[i]) !== -1) is_b = true;
        else if (!is_d && digits.indexOf(password[i]) !== -1) is_d = true;
        else if (!is_sp && specials.indexOf(password[i]) !== -1) is_sp = true;
    }
    let rating = 0;
    if (is_s) rating++;
    if (is_b) rating++;
    if (is_d) rating++;
    if (is_sp) rating++;

    if (password.length < 6 && rating < 3) easyPassword(block);
    else if (password.length < 6 && rating >= 3) mediumPassword(block);
    else if (password.length >=9 && rating < 3) mediumPassword(block);
    else if (password.length >= 8 && rating >= 3) hardPassword(block);
    else if (password.length >= 6 && rating === 1) easyPassword(block);
    else if (password.length >= 6 && rating > 1 && rating < 4) mediumPassword(block);
    else if (password.length >= 6 && rating === 4) hardPassword(block);
}

function easyPassword(elem) {
    elem.style.width = "15%";
    elem.style.backgroundColor = "Crimson";
}

function mediumPassword(elem) {
    elem.style.width = "50%";
    elem.style.backgroundColor = "DarkOrange";
}

function hardPassword(elem) {
    elem.style.width = "100%"
    elem.style.backgroundColor = "SpringGreen";
}