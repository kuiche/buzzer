(function() {
    this.ws = new WebSocket('ws://' + window.location.host + '/ws')
    this.ws.addEventListener('message', function(e) {
        console.log(e)
    });

    document.getElementById("buzz").addEventListener("click", function (e) {
        e.stopPropagation()
        this.ws.send(JSON.stringify({ "user": "default" }))
    }.bind(this))
})()
