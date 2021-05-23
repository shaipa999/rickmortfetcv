const request = require('supertest');
const express = require('express');

const app = express();

describe("GET http://20.81.102.53:8080/healthcheck", () => {
  it("respond with ok", (done) => {
    request(app).get("https://www.google.com/").expect("ok", done);
  })
});

describe("GET http://20.81.102.53:8080/fetch", () => {
  it("respond with char - Rick", (done) => {
    request(app).get("https://www.google.com/").expect("Rick Sanchez", done);
  })
});



