<script setup lang="ts">
import { ROUTE_NAME } from '@/router';
import { useRoute } from 'vue-router';
import { Clapperboard, Film, LayoutDashboard, LogIn } from 'lucide-vue-next';

const route = useRoute();

const links = [
  { name: 'Movies', route: ROUTE_NAME.MOVIES, icon: Film },
  { name: 'Admin Dashboard', route: ROUTE_NAME.ADMIN, icon: LayoutDashboard },
];

const isActive = (routeName: string) => route.name === routeName;
</script>

<template>
  <header
    class="sticky top-0 z-50 border-b border-slate-800 bg-slate-900/80 px-4 py-3.5 shadow-xl backdrop-blur-xl transition-all md:px-8"
  >
    <div class="container mx-auto flex max-w-7xl items-center justify-between gap-4">
      <!-- Brand Logo -->
      <RouterLink :to="{ name: ROUTE_NAME.HOME }" class="group flex items-center gap-3">
        <div
          class="relative rounded-xl bg-gradient-to-br from-rose-600 via-rose-500 to-amber-500 p-2.5 text-white shadow-lg shadow-rose-600/20 transition-transform duration-300 group-hover:scale-105"
        >
          <Clapperboard class="h-6 w-6" />
        </div>
        <div class="flex flex-col">
          <span
            class="bg-gradient-to-r from-white via-slate-100 to-slate-300 bg-clip-text text-xl font-extrabold tracking-tight text-transparent transition-colors group-hover:from-rose-400"
          >
            SubCine
          </span>
          <span class="text-[10px] font-semibold tracking-widest text-rose-400 uppercase"
            >Cinema Experience</span
          >
        </div>
      </RouterLink>

      <!-- Center Navigation Links -->
      <nav
        class="hidden items-center gap-1.5 rounded-full border border-slate-800 bg-slate-900/90 p-1.5 md:flex"
      >
        <RouterLink
          v-for="link in links"
          :key="link.name"
          :to="{ name: link.route }"
          :class="[
            'flex items-center gap-2 rounded-full px-4 py-1.5 text-xs font-semibold transition-all duration-200',
            isActive(link.route)
              ? 'bg-rose-600 text-white shadow-md shadow-rose-600/20'
              : 'text-slate-400 hover:bg-slate-800 hover:text-white',
          ]"
        >
          <component :is="link.icon" class="h-3.5 w-3.5" />
          <span>{{ link.name }}</span>
        </RouterLink>
      </nav>

      <!-- Right Action / User Menu -->
      <div class="flex items-center gap-3">
        <RouterLink
          :to="{ name: ROUTE_NAME.LOGIN }"
          :class="[
            'inline-flex items-center gap-2 rounded-xl border px-4 py-2 text-xs font-semibold shadow-sm transition-all',
            isActive(ROUTE_NAME.LOGIN)
              ? 'border-rose-500 bg-rose-600 text-white'
              : 'border-slate-700 bg-slate-800 text-slate-200 hover:border-rose-500/40 hover:bg-slate-700',
          ]"
        >
          <LogIn class="h-3.5 w-3.5 text-rose-500" />
          <span>Admin Login</span>
        </RouterLink>
      </div>
    </div>
  </header>
</template>
