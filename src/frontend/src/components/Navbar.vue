<script setup lang="ts">
import { ROUTE_NAME } from '@/router';
import { useRoute, useRouter } from 'vue-router';
import {
  Clapperboard,
  Film,
  LayoutDashboard,
  LogIn,
  LogOut,
  ShieldCheck,
  User,
} from 'lucide-vue-next';
import { useGetMeQuery, useLogoutMutation } from '@/composables/auth/queries';
import { computed } from 'vue';
import { toast } from 'vue-sonner';

const route = useRoute();
const router = useRouter();

const { data: meData } = useGetMeQuery();
const { mutateAsync: logoutMutate } = useLogoutMutation();

const isLoggedIn = computed(() => !!meData.value?.user);
const isAdmin = computed(() => !!meData.value?.is_admin);

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
          :to="{ name: ROUTE_NAME.MOVIES }"
          :class="[
            'flex items-center gap-2 rounded-full px-4 py-1.5 text-xs font-semibold transition-all duration-200',
            isActive(ROUTE_NAME.MOVIES) || isActive(ROUTE_NAME.HOME)
              ? 'bg-rose-600 text-white shadow-md shadow-rose-600/20'
              : 'text-slate-400 hover:bg-slate-800 hover:text-white',
          ]"
        >
          <Film class="h-3.5 w-3.5" />
          <span>Movies</span>
        </RouterLink>

        <!-- Admin Dashboard Link (Only visible when connected AND is_admin) -->
        <RouterLink
          v-if="isAdmin"
          :to="{ name: ROUTE_NAME.ADMIN }"
          :class="[
            'flex items-center gap-2 rounded-full px-4 py-1.5 text-xs font-semibold transition-all duration-200',
            isActive(ROUTE_NAME.ADMIN)
              ? 'bg-rose-600 text-white shadow-md shadow-rose-600/20'
              : 'text-slate-400 hover:bg-slate-800 hover:text-white',
          ]"
        >
          <LayoutDashboard class="h-3.5 w-3.5 text-amber-400" />
          <span>Admin Dashboard</span>
        </RouterLink>
      </nav>

      <!-- Right Action / Auth Button -->
      <div class="flex items-center gap-3">
        <!-- Connected Status Badge & Logout Button -->
        <div v-if="isLoggedIn" class="flex items-center gap-2">
          <span
            v-if="isAdmin"
            class="hidden items-center gap-1.5 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-3 py-1 text-xs font-semibold text-emerald-400 sm:inline-flex"
          >
            <ShieldCheck class="h-3.5 w-3.5" />
            <span>Admin Active</span>
          </span>
          <span
            v-else
            class="hidden items-center gap-1.5 rounded-full border border-blue-500/20 bg-blue-500/10 px-3 py-1 text-xs font-semibold text-blue-400 sm:inline-flex"
          >
            <User class="h-3.5 w-3.5" />
            <span>Connected</span>
          </span>

          <button
            @click="handleLogout"
            class="inline-flex items-center gap-2 rounded-xl border border-slate-700 bg-slate-800 px-3.5 py-2 text-xs font-semibold text-slate-300 transition-all hover:border-rose-500/40 hover:bg-rose-600/20 hover:text-rose-400"
          >
            <LogOut class="h-3.5 w-3.5" />
            <span>Logout</span>
          </button>
        </div>

        <!-- Login Button (Visible when NOT connected) -->
        <RouterLink
          v-else
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
