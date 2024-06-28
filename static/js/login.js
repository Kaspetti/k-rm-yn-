const searchParams = new URLSearchParams(window.location.search)
const errorMessage = document.getElementById("error-message")


switch (searchParams.get("auth_status")) {
  case "invalid_credentials":
    errorMessage.innerText = "Incorrect username or password"
    break
  case "invalid_session":
    errorMessage.innerText = "Invalid session. Please login..."
    break
  case "no_session":
    errorMessage.innerText = "No session found. Please login..."
    break
  case "expired_session":
    errorMessage.innerText = "Your session has expired. Please login..."
    break
}
