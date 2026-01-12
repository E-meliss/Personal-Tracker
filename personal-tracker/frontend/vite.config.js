import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    proxy: {
      // Frontend calls /api/* and Vite forwards to Go backend.
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  }
});
