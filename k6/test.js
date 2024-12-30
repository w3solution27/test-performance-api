import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    stages: [
        { duration: '30s', target: 50 },  // aumenta para 50 usuários
        { duration: '1m', target: 100 }, // aumenta para 100 usuários
        { duration: '30s', target: 0 },  // reduz para 0 usuários
    ],
};

export default function () {
    const BASE_URL = 'http://localhost:3000/items';

    // GET
    const res = http.get(BASE_URL);
    if (res.status !== 200) {
        console.error(`GET failed with status: ${res.status}`);
    }

    // POST
    const payload = JSON.stringify({ name: `Item-${__ITER}` });
    const headers = { 'Content-Type': 'application/json' };
    const postRes = http.post(BASE_URL, payload, { headers });

    if (postRes.status !== 201) {
        console.error(`POST failed with status: ${postRes.status}`);
    }

    sleep(1);
}
