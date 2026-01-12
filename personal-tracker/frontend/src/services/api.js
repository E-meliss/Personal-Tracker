import axios from 'axios';

export const api = axios.create({
  // Configure via Vite env var when needed.
  // Dev default is "/api" so Vite can proxy to the Go backend (see vite.config.js).
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: 10000
});
