import "./assets/main.css";

import { createApp } from "vue";
import App from "@/App.vue";
import { VueQueryPlugin } from "@tanstack/vue-query";
import router from "@/router.ts";

const app = createApp(App);

app.use(VueQueryPlugin);
app.use(router);

app.mount("#app");
