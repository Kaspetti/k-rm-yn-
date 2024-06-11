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

  populateCircles()
}


function populateCircles() {
  markers.clearLayers()
  timeIndexText.innerText = timeIndex

  data.slice(0, timeIndex).forEach(d => {
    const circle = L.circle([d.WKT[1], d.WKT[0]], {
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

      idText.innerText = `Id: ${this.data.navn}`
      locationText.innerText = `Location:\n\  Lat: ${this.data.WKT[1]}\n  Lon: ${this.data.WKT[0]}`
      descriptionText.innerText = `Description: ${this.data.beskrivelse ? this.data.beskrivelse : "No description"}`
      tooltipImage.src = `/api/images?id=${this.data.navn}`

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
