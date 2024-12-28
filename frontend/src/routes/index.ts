import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomePage from "../pages/HomePage.vue";
import ViewPage from "../pages/ViewPage.vue";
import LoginPage from "../pages/LoginPage.vue";
import NewPage from "../pages/NewPage.vue";

const routes: RouteRecordRaw[] = [
    { path: "/", component: HomePage },
    { path: "/view/:id", component: ViewPage },
    { path: "/new", component: NewPage },
    { path: "/login", component: LoginPage },
]

export const router = createRouter({
    history: createWebHistory(),
    routes
})