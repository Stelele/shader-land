import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomePage from "../pages/HomePage.vue";
import ViewPage from "../pages/ViewPage.vue";

const routes: RouteRecordRaw[] = [
    { path: "/", component: HomePage },
    { path: "/view/:id", component: ViewPage },
    { path: "/view", component: ViewPage }
]

export const router = createRouter({
    history: createWebHistory(),
    routes
})