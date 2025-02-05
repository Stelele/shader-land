import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { router } from './routes'
import { clerkPlugin } from '@clerk/vue'

const PUBLISHABLE_KEY = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY

if (!PUBLISHABLE_KEY) {
    throw new Error('Add your Clerk Publishable Key to the .env.local file')
}

createApp(App)
    .use(router)
    .use(clerkPlugin, { publishableKey: PUBLISHABLE_KEY })
    .mount('#app')
