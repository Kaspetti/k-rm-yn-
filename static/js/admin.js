/// <reference types="leaflet" />

const latInput = document.getElementById("latitude")
const lonInput = document.getElementById("longitude")

let map
let circle
let drag = false

function init() {
  map = L.map('map')
    .setView([60.385, 5.34], 14.5)

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

      latInput.value = e.latlng.lat.toFixed(4)
      lonInput.value = e.latlng.lng.toFixed(4)
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
