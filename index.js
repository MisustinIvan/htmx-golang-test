let main = () => {
    var map = L.map("map").setView([51.505, -0.09], 13);
    L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
        maxZoom: 19,
        attribution:
            '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
    }).addTo(map);
    var lat = 51.508;
    var lng = -0.11;
    var circle = L.circle([lat, lng], {
        color: "red",
        fillColor: "#f03",
        fillOpacity: 0.5,
        radius: 500,
    }).addTo(map);

    let loop = () => {
        lng += 0.0001;
        circle.setLatLng([lat, lng]);
        setTimeout(loop, 100);
    };
    loop();
};

document.addEventListener("DOMContentLoaded", main);
