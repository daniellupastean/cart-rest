import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "30s", target: 10 },
    { duration: "10s", target: 500 },
    { duration: "1m", target: 500 },
    { duration: "10s", target: 10 },
    { duration: "30s", target: 10 },
    { duration: "10s", target: 500 },
    { duration: "1m", target: 500 },
    { duration: "30s", target: 0 },
  ],
  thresholds: {
    http_req_duration: ["p(95)<5000"],
    http_req_failed: ["rate<0.2"],
  },
};

const BASE_URL = "http://localhost:8080";

export default function () {
  const cartId = `user-${__VU}`;
  const productId = `prod-${Math.floor(Math.random() * 100)}`;

  const addPayload = JSON.stringify({
    productId: productId,
    quantity: 1,
  });

  const res = http.post(`${BASE_URL}/cart/${cartId}/items`, addPayload, {
    headers: { "Content-Type": "application/json" },
  });

  check(res, { completed: (r) => r.status !== 0 });
  sleep(0.1);
}
