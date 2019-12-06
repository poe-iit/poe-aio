(function(){
  // Browser sanity check:
  if (!('querySelector' in document && 'addEventListener' in document)) {
    // Old, old browser. Say buh-bye
    // console.log('Old browser');
    return;
  }

  document.addEventListener('DOMContentLoaded', function(){
    // Get the modal
  var modal = document.getElementsByClassName("modal");
  var fireModal = document.getElementById("fire-modal");
  var weatherModal = document.getElementById("weather-modal");
  var shooterModal = document.getElementById("shooter-modal");
  var safetyModal = document.getElementById("safety-modal");

  // Get the button that opens the modal
  var fireButton = document.getElementById("fire");
  var weatherButton = document.getElementById("weather");
  var shooterButton = document.getElementById("shooter");
  var safetyButton = document.getElementById("safety");


  // Get the <span> element that closes the modal
  var span = document.getElementsByClassName("close")[0];

  // When the user clicks on the button, open the modal
  fireButton.onclick = function() {
    fireModal.style.display = "block";
    sendPostRequest("fire")
  }
  weatherButton.onclick = function() {
    weatherModal.style.display = "block";
    sendPostRequest("enviormental")
  }
  shooterButton.onclick = function() {
    shooterModal.style.display = "block";
    sendPostRequest("shooter")
  }
  safetyButton.onclick = function() {
    safetyModal.style.display = "block";
    sendPostRequest("safety")
  }

  // When the user clicks on <span> (x), close the modal
  span.onclick = function() {
    fireModal.style.display = "none";
    weatherModal.style.display = "none";
    shooterModal.style.display = "none";
    safetyModal.style.display = "none";
  }

  // When the user clicks anywhere outside of the modal, close it
  window.onclick = function(event) {
    if (event.target == fireModal || event.target == weatherModal || event.target == shooterModal || event.target == safetyModal) {
      fireModal.style.display = "none";
      weatherModal.style.display = "none";
      shooterModal.style.display = "none";
      safetyModal.style.display = "none";
    }
  }
  // End of DOMContentLoaded
  });




  function sendPostRequest(emergencyType) {
    var xhr = new XMLHttpRequest();
    xhr.open("POST", '/button', true);

    //Send the proper header information along with the request
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

    xhr.onreadystatechange = function() { // Call a function when the state changes.
        if (this.readyState === XMLHttpRequest.DONE && this.status === 200) {
            // Request finished. Do processing here.
        }
    }
    xhr.send("emergency=" +emergencyType);

  }

// End of IIFE
}());
