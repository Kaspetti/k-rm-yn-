/// <reference types="leaflet" />


let markers
let data
let map

const tooltip         = document.getElementById("tooltip")
const idText          = document.getElementById("idText")
const locationText    = document.getElementById("locationText")
const descriptionText = document.getElementById("descriptionText")
const tooltipImage    = document.getElementById("tooltipImage")

let focusing = false

async function init() {
  data = await d3.json("/api/data")

  map = L.map('map')
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
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map);

  tooltipImage.onerror = function() {
    tooltipImage.src = `/static/images/0.jpg`;
  }

  populateCircles()
}


function populateCircles() {
  markers.clearLayers()

  data.forEach(d => {
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

    circle.on('click', async function (e) {
      L.DomEvent.stopPropagation(e);

      map.setView([this.data.latitude, this.data.longitude], 18)
      focusing = true
      map.on("zoomend moveend", function() {
        if (focusing) {
          tooltip.style.transform = `translate(${window.innerWidth/2}px, ${window.innerHeight/2}px)`;
          tooltip.style.opacity = 1
          focusing = false
        }
      })

      idText.innerText = `Id: ${this.data.id}`
      locationText.innerText = `Koordinater: ${this.data.latitude}, ${this.data.longitude}`
      descriptionText.innerText = `Beskrivelse: ${this.data.description ? this.data.description : "No description"}`

      tooltipImage.src = `/static/images/${this.data.id}.jpg`;
    })
  })
}


function openImg() {
  window.open(tooltipImage.src, "_blank").focus()
}


init()
