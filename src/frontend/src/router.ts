import { createRouter, createWebHistory } from 'vue-router';
import AboutView from '@/views/AboutView.vue';
import MovieView from '@/views/MovieView.vue';
import MoviesView from '@/views/MoviesView.vue';
import LoginView from './views/LoginView.vue';
import LogoutView from './views/LogoutView.vue';

export enum ROUTE_NAME {
  HOME = 'home',
  ABOUT = 'about',
  MOVIES = 'movies',
  MOVIE = 'movie',
  LOGIN = 'login',
  LOGOUT = 'logout',
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: ROUTE_NAME.HOME,
      component: MoviesView,
    },
    {
      path: '/about',
      name: ROUTE_NAME.ABOUT,
      component: AboutView,
    },
    {
      path: '/movies',
      name: ROUTE_NAME.MOVIES,
      component: MoviesView,
    },
    {
      path: '/movies/:id',
      name: ROUTE_NAME.MOVIE,
      component: MovieView,
      props: true,
    },
    {
      path: '/login',
      name: ROUTE_NAME.LOGIN,
      component: LoginView,
    },
    {
      path: '/logout',
      name: ROUTE_NAME.LOGOUT,
      component: LogoutView,
    },
  ],
});

export default router;
