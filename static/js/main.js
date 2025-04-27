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

async function vote(jokeId, isFunny) {
    const userCookieID = getCookie("user_cookie_id");
    const voteData = {
        user_cookie_id: userCookieID,
        joke_id: Number(jokeId),
        is_like: isFunny
    };

    // console.log("Sending vote data:", voteData);

    try {
        const response = await fetch(`${window.apiUrl}/vote`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(voteData)
        });

        const result = await response.json();
        
        if (response.ok) {
            console.log(result.message);
            clearNotification();

            // Nếu muốn sau khi vote thì gọi lại get joke để lấy joke mới.
            // getJoke();
        } else {
            console.error('Error voting:', result.error);
            if (result.error === "you have already voted for this joke") {
                showNotification("you have already voted for this joke", "warning");
            } else {
                showNotification("An error occurred while voting!", "error");
            }
        }
    } catch (error) {
        console.error('Error submitting vote:', error);
        showNotification("Error submitting vote!", "error");
    }
}

function showNotification(message, type = 'success') {
    const notification = document.querySelector('.notification');
    notification.classList.remove('hidden');
    notification.innerHTML = `
        <div class="flex items-center max-w-xs p-4 mb-4 text-sm text-white rounded-lg shadow ${
            type === 'success' ? 'bg-green-500' : type === 'warning' ? 'bg-yellow-500' : 'bg-red-500'
        }" role="alert">
            <svg class="inline w-5 h-5 mr-3" fill="currentColor" viewBox="0 0 20 20">
                <path d="${
                    type === 'success'
                        ? 'M16.707 5.293a1 1 0 00-1.414 0L8 12.586 4.707 9.293a1 1 0 10-1.414 1.414l4 4a1 1 0 001.414 0l8-8a1 1 0 000-1.414z'
                        : type === 'warning'
                        ? 'M8.257 3.099c.765-1.36 2.722-1.36 3.487 0l6.516 11.591c.75 1.337-.213 3.01-1.742 3.01H3.483c-1.528 0-2.492-1.673-1.741-3.01L8.257 3.1zM11 14a1 1 0 10-2 0 1 1 0 002 0zm-1-2a1 1 0 01-1-1V8a1 1 0 012 0v3a1 1 0 01-1 1z'
                        : 'M8.257 3.099c.765-1.36 2.722-1.36 3.487 0l6.516 11.591c.75 1.337-.213 3.01-1.742 3.01H3.483c-1.528 0-2.492-1.673-1.741-3.01L8.257 3.1zM11 14a1 1 0 10-2 0 1 1 0 002 0zm-1-2a1 1 0 01-1-1V8a1 1 0 012 0v3a1 1 0 01-1 1z'
                }"></path>
            </svg>
            <div>${message}</div>
        </div>
    `;

    setTimeout(() => {
        notification.classList.add('hidden');
        notification.innerHTML = '';
    }, 2000);
}

function clearNotification() {
    const notification = document.querySelector('.notification');
    notification.classList.add('hidden');
    notification.innerHTML = '';
}

document.addEventListener('DOMContentLoaded', function () {
    getJoke();
    
    document.querySelector('.vote-funny').addEventListener('click', function () {
        const jokeId = document.querySelector('.joke-id').value;
        vote(jokeId, true);
    });
    
    document.querySelector('.vote-not-funny').addEventListener('click', function () {
        const jokeId = document.querySelector('.joke-id').value;
        vote(jokeId, false);
    });
});
