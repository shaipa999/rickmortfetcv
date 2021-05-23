const request = require('supertest');
const express = require('express');

const app = express();

describe("GET https://www.google.com/", () => {
  it("respond with ok", (done) => {
    request(app).get("https://www.google.com/").expect("ok", done);
  })
});

describe("GET https://www.google.com/", () => {
  it("respond with char - Rick", (done) => {
    request(app).get("https://www.google.com/").expect("Rick Sanchez", done);
  })
});



