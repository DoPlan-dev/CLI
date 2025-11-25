import http from "k6/http";
import { check, sleep } from "k6";
import { Trend } from "k6/metrics";

const BASE_URL = __ENV.K6_BASE_URL || "http://localhost:3000";

const qrDuration = new Trend("http_req_duration", true, {
  endpoint: "qr",
});
const analyticsDuration = new Trend("http_req_duration", true, {
  endpoint: "analytics",
});

export const options = {
  scenarios: {
    qr_load: {
      executor: "ramping-arrival-rate",
      startRate: 5,
      timeUnit: "1s",
      preAllocatedVUs: 20,
      stages: [
        { target: 20, duration: "1m" },
        { target: 40, duration: "2m" },
        { target: 10, duration: "1m" },
      ],
      exec: "qrScenario",
    },
    analytics_poll: {
      executor: "constant-arrival-rate",
      rate: 5,
      timeUnit: "1s",
      duration: "4m",
      preAllocatedVUs: 5,
      exec: "analyticsScenario",
      startTime: "30s",
    },
  },
  thresholds: {
    'http_req_duration{endpoint:"qr"}': ["p(95)<100", "avg<75"],
    'http_req_duration{endpoint:"analytics"}': ["p(95)<80", "avg<60"],
    http_req_failed: ["rate<0.01"],
  },
};

const qrPayloads = [
  { text: "https://example.com", format: "png", size: 200, errorCorrection: "M" },
  { text: "https://doplan.dev/docs", format: "svg", size: 300, errorCorrection: "H" },
  { text: "Edge cases are fun!", format: "png", size: 150, errorCorrection: "L" },
  {
    text: "https://example.com/product?id=123&utm_source=k6",
    format: "svg",
    size: 400,
    errorCorrection: "Q",
  },
];

export function qrScenario() {
  const payload = qrPayloads[Math.floor(Math.random() * qrPayloads.length)];
  const res = http.post(`${BASE_URL}/api/qr`, JSON.stringify(payload), {
    headers: { "Content-Type": "application/json" },
  });

  qrDuration.add(res.timings.duration);
  check(res, {
    "qr status is 200": (r) => r.status === 200,
    "qr response has base64": (r) => r.json("qrCode")?.includes("data:image"),
  });

  sleep(0.5);
}

export function analyticsScenario() {
  const res = http.get(`${BASE_URL}/api/analytics`);
  analyticsDuration.add(res.timings.duration);

  check(res, {
    "analytics status is 200": (r) => r.status === 200,
    "analytics has stats": (r) => !!r.json("totalGenerations"),
  });

  sleep(1);
}

