<script setup lang="ts">
import { ROUTE_NAME } from '@/router';
import { useRoute, useRouter } from 'vue-router';
import { Clapperboard, Film, LayoutDashboard, LogIn, LogOut, ShieldCheck } from 'lucide-vue-next';
import { useGetMeQuery, useLogoutMutation } from '@/composables/auth/queries';
import { computed } from 'vue';
import { toast } from 'vue-sonner';

const route = useRoute();
const router = useRouter();

const { data: meData } = useGetMeQuery();
const { mutateAsync: logoutMutate } = useLogoutMutation();

const isLoggedIn = computed(() => !!meData.value?.user);

const handleLogout = async () => {
  try {
    await logoutMutate();
    toast.success('Logged out successfully');
    router.push({ name: ROUTE_NAME.HOME });
  } catch (err) {
    toast.error('Logout failed');
  }
};

const isActive = (routeName: string) => route.name === routeName;
</script>

<template>
  <header class="sticky top-0 z-50 bg-slate-900/80 backdrop-blur-xl border-b border-slate-800 px-4 md:px-8 py-3.5 shadow-xl transition-all">
    <div class="container mx-auto max-w-7xl flex items-center justify-between gap-4">
      
      <!-- Brand Logo -->
      <RouterLink :to="{ name: ROUTE_NAME.HOME }" class="flex items-center gap-3 group">
        <div class="relative p-2.5 rounded-xl bg-gradient-to-br from-rose-600 via-rose-500 to-amber-500 text-white shadow-lg shadow-rose-600/20 group-hover:scale-105 transition-transform duration-300">
          <Clapperboard class="w-6 h-6" />
        </div>
        <div class="flex flex-col">
          <span class="font-extrabold text-xl tracking-tight bg-gradient-to-r from-white via-slate-100 to-slate-300 bg-clip-text text-transparent group-hover:from-rose-400 transition-colors">
            SubCine
          </span>
          <span class="text-[10px] tracking-widest uppercase font-semibold text-rose-400">Cinema Experience</span>
        </div>
      </RouterLink>

      <!-- Center Navigation Links -->
      <nav class="hidden md:flex items-center gap-1.5 bg-slate-900/90 p-1.5 rounded-full border border-slate-800">
        <RouterLink
          :to="{ name: ROUTE_NAME.MOVIES }"
          :class="[
            'flex items-center gap-2 px-4 py-1.5 rounded-full text-xs font-semibold transition-all duration-200',
            isActive(ROUTE_NAME.MOVIES) || isActive(ROUTE_NAME.HOME)
              ? 'bg-rose-600 text-white shadow-md shadow-rose-600/20'
              : 'text-slate-400 hover:text-white hover:bg-slate-800'
          ]"
        >
          <Film class="w-3.5 h-3.5" />
          <span>Movies</span>
        </RouterLink>

        <!-- Admin Dashboard Link (Only visible when connected) -->
        <RouterLink
          v-if="isLoggedIn"
          :to="{ name: ROUTE_NAME.ADMIN }"
          :class="[
            'flex items-center gap-2 px-4 py-1.5 rounded-full text-xs font-semibold transition-all duration-200',
            isActive(ROUTE_NAME.ADMIN)
              ? 'bg-rose-600 text-white shadow-md shadow-rose-600/20'
              : 'text-slate-400 hover:text-white hover:bg-slate-800'
          ]"
        >
          <LayoutDashboard class="w-3.5 h-3.5 text-amber-400" />
          <span>Admin Dashboard</span>
        </RouterLink>
      </nav>

      <!-- Right Action / Auth Button -->
      <div class="flex items-center gap-3">
        
        <!-- Connected Status Badge & Logout Button -->
        <div v-if="isLoggedIn" class="flex items-center gap-2">
          <span class="hidden sm:inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-emerald-500/10 border border-emerald-500/20 text-emerald-400 text-xs font-semibold">
            <ShieldCheck class="w-3.5 h-3.5" />
            <span>Admin Active</span>
          </span>

          <button
            @click="handleLogout"
            class="inline-flex items-center gap-2 px-3.5 py-2 rounded-xl text-xs font-semibold bg-slate-800 hover:bg-rose-600/20 text-slate-300 hover:text-rose-400 border border-slate-700 hover:border-rose-500/40 transition-all"
          >
            <LogOut class="w-3.5 h-3.5" />
            <span>Logout</span>
          </button>
        </div>

        <!-- Login Button (Visible when NOT connected) -->
        <RouterLink
          v-else
          :to="{ name: ROUTE_NAME.LOGIN }"
          :class="[
            'inline-flex items-center gap-2 px-4 py-2 rounded-xl text-xs font-semibold transition-all shadow-sm border',
            isActive(ROUTE_NAME.LOGIN)
              ? 'bg-rose-600 text-white border-rose-500'
              : 'bg-slate-800 hover:bg-slate-700 text-slate-200 border-slate-700 hover:border-rose-500/40'
          ]"
        >
          <LogIn class="w-3.5 h-3.5 text-rose-500" />
          <span>Admin Login</span>
        </RouterLink>

      </div>

    </div>
  </header>
</template>
