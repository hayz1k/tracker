import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      // Проксируем только API-запросы, остальные маршруты — отдаём index.html
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
      },
    },
    // Важное:
    // fallback для SPA включён по умолчанию, если не проксируются запросы,
    // поэтому здесь не нужно ничего дополнительно прописывать.
  },
})
