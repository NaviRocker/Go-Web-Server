// Get the message button element
const messageButton = document.getElementById('message-button');

// Get the message element
const messageElement = document.getElementById('message');

// Add a click event listener to the message button
messageButton.addEventListener('click', () => {
  // Make a GET request to the server
  fetch('/message')
    .then(response => response.text())
    .then(message => {
      // Update the message element with the response from the server
      messageElement.textContent = message;
    });
});
