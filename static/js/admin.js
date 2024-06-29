/// <reference types="leaflet" />


const searchParams = new URLSearchParams(window.location.search)
const statusMessage = document.getElementById("status-message")


switch (searchParams.get("upload_status")) {
  case "file_error":
    errorMessage.innerText = "Error occured while getting file from form"
    break
  case "invalid_filetype":
    errorMessage.innerText = "Invalid filetype. Supported filetype: '.jpg'"
    break
  case "count_error":
    errorMessage.innerText = "Error while counting images. Automatic naming of the file unavailable. Try again later..."
    break
  case "invalid_lat":
    errorMessage.innerText = "Invalid latitude"
    break
  case "invalid_lon":
    errorMessage.innerText = "Invalid longitude"
    break
  case "file_upload_error":
    errorMessage.innerText = "Error while uploading the file to the server"
    break
  case "data_write_error":
    errorMessage.innerText = "Error while writing new data to the server"
    break
  case "data_read_error":
    errorMessage.innerText = "Error while reading old data for appending"
    break
}


const latInput = document.getElementById("latitude")
const lonInput = document.getElementById("longitude")

let map
let circle
let drag = false

function init() {
  map = L.map('map')
    .setView([60.385, 5.34], 16)

  L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map);

  circle = L.circle([60.385, 5.34], {
      color: 'red',
      fillColor: '#f03',
      fillOpacity: 0.5,
      radius: 7
  }).addTo(map);

  circle.on("click", function() {
    drag = !drag
  })

  map.on("mousemove", function(e) {
    if (drag) {
      circle.setLatLng(e.latlng)

      latInput.value = e.latlng.lat.toFixed(5)
      lonInput.value = e.latlng.lng.toFixed(5)
    }
  })
}


latInput.addEventListener("change", function() {
  circle.setLatLng([latInput.value, lonInput.value])
})

lonInput.addEventListener("change", function() {
  circle.setLatLng([latInput.value, lonInput.value])
})


init()
