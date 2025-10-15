import http from 'k6/http';

import {check, sleep} from 'k6';

export const options = {
  iterations: 1000,
  vus: 100
}

export default function () {
  const payload = JSON.stringify({
    name: "Movie",
    year: 2000
  });
  const params = {
    headers: {
      "Content-Type": "application/json"
    }
  }

  const res = http.post("http://localhost:7746/api/v1/movies", payload, params)

  check(res, { 'status was 201': r => r.status === 201})

  sleep(1)
}
