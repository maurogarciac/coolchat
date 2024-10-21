
htmx.defineExtension("ws-html-response", {
    transformResponse : function(text, xhr, elt) {

      
        var response = JSON.parse(JSON.parse(text).message)
        console.log(response)
        var chatbox = document.getElementById("chatbox")
    
        // Full message div
        var messageDiv= document.createElement("div")
        // if presponse.user === me, div class="message-mine" with align right and bgcolor dark blue
        // else div class="other"
        messageDiv.id = "message"

        // Header for user and timestamp
        var msgHeaderDiv = document.createElement("div")
        msgHeaderDiv.id = "msg-headers"

        var sentBy = document.createElement("div")
        sentBy.id = "sent-by"
        sentBy.textContent = `${response.user}:`
        msgHeaderDiv.appendChild(sentBy) 

        var timeSent = document.createElement("div")
        timeSent.id = "time-sent" 
        timeSent.textContent = `${response.sent}`
        msgHeaderDiv.appendChild(timeSent)

        // Actual message text
        
        var msgContentDiv = document.createElement("div")
        msgContentDiv.textContent = `${response.message}`
        msgContentDiv.id = "msg-content"

        chatbox.appendChild(messageDiv)
        messageDiv.appendChild(msgHeaderDiv)
        messageDiv.appendChild(msgContentDiv)
    }
});

    // function getCookieByName(name) {
    //     const cookies = document.cookie.split(";");
    //     for (let cookie of cookies) {
    //         cookie = cookie.trim();
    //         if (cookie.startsWith(name + "=")) {
    //             return cookie.substring(name.length + 1);
    //         }
    //     }
    //     return null;
    // }
    
    // console.log(getCookieByName("user"));