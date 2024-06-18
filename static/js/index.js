/// <reference types="leaflet" />


let markers
let data

const tooltip         = document.getElementById("tooltip")
const idText          = document.getElementById("idText")
const locationText    = document.getElementById("locationText")
const descriptionText = document.getElementById("descriptionText")
const tooltipImage    = document.getElementById("tooltipImage")


async function init() {
  data = await d3.json("/api/data")

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
      const tooltipRect = tooltip.getBoundingClientRect();
      const screenWidth = window.innerWidth;
      const screenHeight = window.innerHeight;

      let translateX = e.containerPoint.x;
      let translateY = e.containerPoint.y;

      if (translateX + tooltipRect.width > screenWidth) {
        translateX = screenWidth - tooltipRect.width;
      }

      if (translateY + tooltipRect.height > screenHeight) {
        translateY = screenHeight - tooltipRect.height;
      }

      tooltip.style.transform = `translate(${translateX}px, ${translateY}px)`;

      idText.innerText = `Id: ${this.data.id}`
      locationText.innerText = `Koordinater: ${this.data.latitude}, ${this.data.longitude}`
      descriptionText.innerText = `Beskrivelse: ${this.data.description ? this.data.description : "No description"}`

      tooltipImage.src = `/static/images/${this.data.id}.jpg`;
      tooltip.style.opacity = 1
    })
  })
}


function openImg() {
  window.open(tooltipImage.src, "_blank").focus()
}


init()
