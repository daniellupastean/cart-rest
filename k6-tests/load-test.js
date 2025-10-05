import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "2m", target: 50 },
    { duration: "5m", target: 50 },
    { duration: "2m", target: 100 },
    { duration: "5m", target: 100 },
    { duration: "2m", target: 0 },
  ],
  thresholds: {
    http_req_duration: ["p(95)<1000", "p(99)<2000"],
    http_req_failed: ["rate<0.05"],
  },
};

const BASE_URL = "http://localhost:8080";

export default function () {
  const cartId = `user-${__VU}`;
  const productId = `prod-${Math.floor(Math.random() * 1000)}`;

  if (Math.random() < 0.6) {
    const addPayload = JSON.stringify({
      productId: productId,
      quantity: Math.floor(Math.random() * 5) + 1,
    });
    const res = http.post(`${BASE_URL}/cart/${cartId}/items`, addPayload, {
      headers: { "Content-Type": "application/json" },
    });
    check(res, { "add ok": (r) => r.status === 200 });
  } else if (Math.random() < 0.9) {
    const res = http.get(`${BASE_URL}/cart/${cartId}`);
    check(res, { "get ok": (r) => r.status === 200 || r.status === 404 });
  } else {
    if (Math.random() < 0.5) {
      const updatePayload = JSON.stringify({
        quantity: Math.floor(Math.random() * 10) + 1,
      });
      const res = http.put(
        `${BASE_URL}/cart/${cartId}/items/${productId}`,
        updatePayload,
        {
          headers: { "Content-Type": "application/json" },
        }
      );
      check(res, { "update ok": (r) => r.status === 200 || r.status === 404 });
    } else {
      const res = http.del(`${BASE_URL}/cart/${cartId}/items/${productId}`);
      check(res, { "delete ok": (r) => r.status === 200 || r.status === 404 });
    }
  }

  sleep(Math.random() * 2 + 0.5);
}
