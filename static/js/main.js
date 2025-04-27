function getCookie(name) {
    let cookieArr = document.cookie.split(';');
    for (let i = 0; i < cookieArr.length; i++) {
        let cookie = cookieArr[i].trim();
        if (cookie.indexOf(name + "=") == 0) {
            return cookie.substring(name.length + 1);
        }
    }
    return "";
}

function setCookie(name, value, days) {
    const d = new Date();
    d.setTime(d.getTime() + (days * 24 * 60 * 60 * 1000));
    let expires = "expires=" + d.toUTCString();
    document.cookie = name + "=" + value + ";" + expires + ";path=/";
}

async function getJoke() {
    const userCookieID = getCookie("user_cookie_id") || `user_${Math.random().toString(36).slice(2, 11)}`;
    
    if (!getCookie("user_cookie_id")) {
        setCookie("user_cookie_id", userCookieID, 7);
    }

    try {
        const response = await fetch(`${window.apiUrl}/joke?user_cookie_id=${userCookieID}`);
        const data = await response.json();
        
        if (data.joke) {
            document.querySelector('.joke-content').innerHTML = data.joke.text;
            document.querySelector('.joke-id').value = data.joke.id;
            document.querySelector('.vote-buttons').classList.remove('hidden');
        } else {
            document.querySelector('.joke-content').innerText = "That's all the jokes for today! Come back another day!";
            // Ẩn nút vote khi hết joke
            document.querySelector('.vote-buttons').classList.add('hidden');
        }
    } catch (error) {
        console.error('Error fetching joke:', error);
    }
}

document.addEventListener('DOMContentLoaded', function () {
    getJoke();
});
