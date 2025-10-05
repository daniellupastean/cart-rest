import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  stages: [
    { duration: "1m", target: 5 },
    { duration: "2m", target: 5 },
    { duration: "1m", target: 0 },
  ],
  thresholds: {
    http_req_duration: ["p(95)<500"],
    http_req_failed: ["rate<0.01"],
  },
};

const BASE_URL = "http://localhost:8080";

export default function () {
  const cartId = `user-${__VU}`;
  const productId = `prod-${Math.floor(Math.random() * 100)}`;

  let res = http.get(`${BASE_URL}/health`);
  check(res, { "health check ok": (r) => r.status === 200 });
  sleep(1);

  const addPayload = JSON.stringify({
    productId: productId,
    quantity: Math.floor(Math.random() * 5) + 1,
  });

  res = http.post(`${BASE_URL}/cart/${cartId}/items`, addPayload, {
    headers: { "Content-Type": "application/json" },
  });
  check(res, { "add item ok": (r) => r.status === 200 });
  sleep(1);

  res = http.get(`${BASE_URL}/cart/${cartId}`);
  check(res, { "get cart ok": (r) => r.status === 200 });
  sleep(1);

  const updatePayload = JSON.stringify({
    quantity: Math.floor(Math.random() * 10) + 1,
  });

  res = http.put(
    `${BASE_URL}/cart/${cartId}/items/${productId}`,
    updatePayload,
    {
      headers: { "Content-Type": "application/json" },
    }
  );
  check(res, { "update item ok": (r) => r.status === 200 });
  sleep(1);

  res = http.del(`${BASE_URL}/cart/${cartId}/items/${productId}`);
  check(res, { "remove item ok": (r) => r.status === 200 });
  sleep(2);
}
