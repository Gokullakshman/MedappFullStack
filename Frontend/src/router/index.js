import Vue from "vue";
import VueRouter from "vue-router";


Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Login",
    component: () =>
    import("../views/login.vue")
  },
  {
      path: '/dashboard/:role/:login_id/:login_history_id/:user_id',
      name: "Dashboard",
      component: () =>
      import("../views/Dashboard.vue")
    },
  
];

const router = new VueRouter({
  routes,
});

export default router;
