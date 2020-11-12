(function() {
    document.getElementById("start").addEventListener('click', (e) => {
        e.stopPropagation()
        document.getElementById("start").classList.add('hide')
        document.getElementById("buzz").classList.remove("hide")

        this.ws = new WebSocket('ws://' + window.location.host + '/ws')
        this.ws.addEventListener('message', function(e) {
            console.log(e)
            new Audio('/Buzzer-sound.mp3').play()
        });

        document.getElementById("buzz").addEventListener("click", (e) => {
            e.stopPropagation()
            this.ws.send(JSON.stringify({ "user": "default" }))
        })
    })
})()
