import { createRouter, createWebHistory } from "vue-router";
import Index from "../views/Index.vue";
import Search from "../views/Search.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "index",
      component: Search,
    },
    {
      path: "/search",
      name: "search",
      component: Search,
    },
  ],
});

export default router;
