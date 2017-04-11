"use strict";

const http = require('http')

setInterval(() => {
    const dt = new Date()
    const body = JSON.stringify({
        ts: dt
    });
    const request = new http.ClientRequest({
        hostname: "127.0.0.1",
        port: 8082,
        path: "/get_stuff",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Content-Length": Buffer.byteLength(body)
        }
    });
    request.end(body);
}, 1000);