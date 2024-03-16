import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path, { resolve } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);


// https://vitejs.dev/config/
export default defineConfig({
  root: resolve(__dirname, "./"),
  plugins: [react()],
  build: {
    outDir: resolve(__dirname, "dist"),
    rollupOptions: {
      input: {
        main: resolve(__dirname, "index.html"),
        page1: resolve(__dirname, "page1/index.html"),
        page2: resolve(__dirname, "page2/index.html"),
      },
    },
  },
})
