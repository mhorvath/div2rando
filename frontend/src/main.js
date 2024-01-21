import { createApp } from "vue";
import App from "./App.vue";
import PrimeVue from "primevue/config";
import FormatterPlugin from './plugins/formatter'
import "primevue/resources/themes/lara-dark-amber/theme.css";
import 'primeicons/primeicons.css'
import ToastService from 'primevue/toastservice';

const app = createApp(App);
app.use(PrimeVue);
app.use(FormatterPlugin);
app.use(ToastService);

app.mount("#app");
