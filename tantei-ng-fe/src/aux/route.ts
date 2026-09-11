import axios from 'axios';

// export const api = axios.create({
//     baseURL: "https://tantei-ng-be.vercel.app",
//     withCredentials: true
// })

export const api = axios.create({
	baseURL: 'http://localhost:28080',
	withCredentials: true
});
