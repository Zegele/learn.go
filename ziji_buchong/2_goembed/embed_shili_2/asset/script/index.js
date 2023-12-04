onload = () => {
    const txt = document.querySelector('h1');
    let hue = 0
    setInterval(() => {
        hue = (hue + 23) % 360;
        txt.style.color = `hsl(${hue}, 61%, 57%)`;
    }, 300);
};