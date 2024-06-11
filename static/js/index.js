/// <reference types="leaflet" />

let markers
let timeIndex
let data

const tooltip         = document.getElementById("tooltip")
const idText          = document.getElementById("idText")
const locationText    = document.getElementById("locationText")
const descriptionText = document.getElementById("descriptionText")
const tooltipImage    = document.getElementById("tooltipImage")
const timeIndexText   = document.getElementById("timeIndexText")

async function init() {
  data = await d3.json("/api/data")
  timeIndex = data.length

  let map = L.map('map')
    .setView([60.385, 5.34], 14.5)

  markers = L.layerGroup().addTo(map)

  map.on('mousedown', function() {
    tooltip.style.opacity = 0.0
  })
  map.on('zoom', function() {
    tooltip.style.opacity = 0.0
  })

  L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
  }).addTo(map);

  tooltipImage.onerror = function() {
    tooltipImage.src = "/static/images/placeholder.jpg"
  }

  populateCircles()
}


function populateCircles() {
  markers.clearLayers()
  timeIndexText.innerText = timeIndex

  data.slice(0, timeIndex).forEach(d => {
    const circle = L.circle([d.latitude, d.longitude], {
      color: 'red',
      fillColor: '#f03',
      fillOpacity: 0.5,
      radius: 7
    }).addTo(markers);

    circle.data = d

    circle.on('mouseover', function () {
      this.setStyle({
        color: "blue",
        fillColor: '#30f',
        fillOpacity: 0.8,
      })
    })

    circle.on('mouseout', function () {
      this.setStyle({
        color: 'red',
        fillColor: '#f03',
        fillOpacity: 0.5, 
      }) 
    })

    circle.on('click', function (e) {
      L.DomEvent.stopPropagation(e);
      tooltip.style.transform = `translate(${e.containerPoint.x}px, ${e.containerPoint.y}px)`

      idText.innerText = `Id: ${this.data.id}`
      locationText.innerText = `Location:\n\  Lat: ${this.data.latitude}\n  Lon: ${this.data.longitude}`
      descriptionText.innerText = `Description: ${this.data.description ? this.data.description : "No description"}`
      tooltipImage.src = `/static/images/${this.data.id}.jpg`

      tooltip.style.opacity = 1;
    })
  })
}


function goBackTimeline() {
  if (timeIndex > 0) {
    timeIndex--
  populateCircles()
  }
}

function goForwardTimeline() {
  if (timeIndex < data.length) {
    timeIndex++
    populateCircles()
  }
}


init()
