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

    var test_line = {
        type: "Feature",
        properties: {
            name: "test line",
            popupContent: "test popup content",
        },
        geometry: {
            type: "LineString",
            coordinates: [
                [-0.1, 51.5],
                [-0.05, 51.55],
                [0, 51.5],
                [0.05, 51.55],
                [0.1, 51.5],
            ],
        },
    };
    L.geoJSON(test_line).addTo(map);

    let loop = () => {
        lng += 0.0001;
        circle.setLatLng([lat, lng]);
        setTimeout(loop, 100);
    };
    loop();
};

document.addEventListener("DOMContentLoaded", main);
document.addEventListener("htmx:configRequest", (evt) => {
    const auth = "ligmaballs";
    evt.detail.headers["Authorization"] = auth;
});
