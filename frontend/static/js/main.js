// Scroll to the bottom of chat on first load
document.addEventListener("DOMContentLoaded", function () {
    if (window.location.pathname.includes("/chat")) {
        scrollToBottom();
    }
});

htmx.defineExtension("ws-html-response", {
    transformResponse : function(text, xhr, elt) {

        var response = JSON.parse(JSON.parse(text).message)

        var chatbox = document.getElementById("chatbox")
    
        // Full message div

        var msgWrapperDiv = document.createElement("div")
        msgWrapperDiv.id = "message-wrapper"

        var messageDiv = document.createElement("div")

        if (response.user === getUsername()) {
            messageDiv.className = "right"
        } else {
            messageDiv.className = "left"
        }
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

        // Format time from message

        var date = new Date(response.ts);
        var dayom = String(date.getDate() + " " + getMonthString(date.getUTCMonth()))
        var hours = String(date.getUTCHours()).padStart(2, '0');
        var minutes = String(date.getUTCMinutes()).padStart(2, '0');
        
        timeSent.textContent = `${dayom} - ${hours}:${minutes}`
        msgHeaderDiv.appendChild(timeSent)

        // Actual message text
        var msgContentDiv = document.createElement("div")
        msgContentDiv.textContent = `${response.text}`
        msgContentDiv.id = "msg-content"

        chatbox.appendChild(msgWrapperDiv)
        msgWrapperDiv.appendChild(messageDiv)
        messageDiv.appendChild(msgHeaderDiv)
        messageDiv.appendChild(msgContentDiv)

        scrollToBottom()
    }
})

// Update ws message using new structure that includes username
document.addEventListener("htmx:wsConfigSend", function (event) {
    var newMessage = JSON.stringify({ text: event.detail.parameters.message, user: getUsername() })
    event.detail.parameters.message = newMessage
})

// Get username from placeholder value in form
function getUsername(){
    var formInput = document.querySelector('#form input[name="message"]')
    var placeholderText = formInput.getAttribute('placeholder')

    // Match text found after comma until question mark :)
    const regex = /,\s*([^?]+)\?/ 
    
    return placeholderText.match(regex)[1].trim() 
}

// Get short month string
function getMonthString(monthIndex){
    const monthNames = [
        "Jan", "Feb", "Mar", "Apr", "May", "Jun", 
        "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
    ]
    
    return monthNames[monthIndex]
}

function scrollToBottom(){
    var elem = document.getElementById('chatbox');
    elem.scrollTop = elem.scrollHeight;
}
