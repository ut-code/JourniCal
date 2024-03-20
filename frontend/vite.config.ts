import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import { VitePWA } from 'vite-plugin-pwa'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), VitePWA({
    registerType: 'autoUpdate',
    devOptions: {
      enabled: true
    },
    injectRegister: 'auto',
    manifest: {
      name: 'JourniCal',
      short_name: 'JourniCal',
      description: 'カレンダーとジャーナルを組み合わせたアプリ',
      start_url: '.',
      display: "standalone",
      orientation: "portrait",
      theme_color: '#00D372',
      background_color: "#aaaaaa",
      icons: [
        {
          src: '192.png',
          sizes: '192x192',
          type: 'image/png',
        },
        {
          src: '512.png',
          sizes: '512x512',
          type: 'image/png',
        },
      ],
    }
  })],
})