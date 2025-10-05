import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "2m", target: 100 },
    { duration: "3m", target: 100 },
    { duration: "2m", target: 200 },
    { duration: "3m", target: 200 },
    { duration: "2m", target: 300 },
    { duration: "3m", target: 300 },
    { duration: "2m", target: 400 },
    { duration: "3m", target: 400 },
    { duration: "2m", target: 0 },
  ],
  thresholds: {
    http_req_duration: ["p(95)<3000", "p(99)<5000"],
    http_req_failed: ["rate<0.1"],
  },
};

const BASE_URL = "http://localhost:8080";

export default function () {
  const cartId = `user-${__VU}`;
  const productId = `prod-${Math.floor(Math.random() * 1000)}`;

  const addPayload = JSON.stringify({
    productId: productId,
    quantity: Math.floor(Math.random() * 5) + 1,
  });

  const res = http.post(`${BASE_URL}/cart/${cartId}/items`, addPayload, {
    headers: { "Content-Type": "application/json" },
  });

  check(res, { completed: (r) => r.status !== 0 });
  sleep(Math.random() * 0.5);
}
