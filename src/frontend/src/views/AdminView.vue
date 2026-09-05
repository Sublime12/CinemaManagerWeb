<script setup lang="ts">
import AppSidebar from '@/components/AppSidebar.vue';
import { useGetMeQuery } from '@/composables/auth/queries';
import { computed } from 'vue';
import { ROUTE_NAME } from '@/router';
import { ShieldAlert, LogIn, Film } from 'lucide-vue-next';

const { data: meData, isFetching } = useGetMeQuery();

const isLoggedIn = computed(() => !!meData.value?.user);
const isAdmin = computed(() => !!meData.value?.user && !!meData.value?.is_admin);
</script>

<template>
  <div class="min-h-screen">
    <!-- Loading auth -->
    <div
      v-if="isFetching"
      class="flex min-h-[60vh] flex-col items-center justify-center space-y-3 text-slate-400"
    >
      <Film class="h-8 w-8 animate-spin text-rose-500" />
      <p class="text-xs font-semibold">Verifying administrator privileges...</p>
    </div>

    <!-- Render Admin -->
    <AppSidebar v-else-if="isAdmin" />

    <!-- Unauthorized -->
    <div v-else class="flex min-h-[65vh] flex-col items-center justify-center px-4">
      <div
        class="w-full max-w-md space-y-5 rounded-3xl border border-slate-800 bg-slate-900 p-8 text-center shadow-2xl"
      >
        <div
          class="inline-flex rounded-2xl border border-rose-500/20 bg-rose-500/10 p-4 text-rose-500"
        >
          <ShieldAlert class="h-10 w-10" />
        </div>

        <div class="space-y-2">
          <h2 class="text-2xl font-black text-white">403 Admin Access Required</h2>
          <p class="text-xs leading-relaxed text-slate-400">
            <span v-if="isLoggedIn"
              >Your account does not have administrator privileges. Only admin accounts (such as
              <code class="rounded bg-slate-950 px-1 text-rose-400">user01</code>) can access this
              page.</span
            >
            <span v-else
              >The Cinema Management Console requires administrator authentication. Please sign in
              to access movie scheduling and theatre controls.</span
            >
          </p>
        </div>

        <RouterLink
          :to="{ name: ROUTE_NAME.LOGIN }"
          class="inline-flex w-full items-center justify-center gap-2 rounded-xl bg-rose-600 px-5 py-3 text-xs font-bold text-white shadow-lg shadow-rose-600/25 transition-all hover:bg-rose-500"
        >
          <LogIn class="h-4 w-4" />
          <span>{{ isLoggedIn ? 'Switch Account' : 'Sign In to Admin Portal' }}</span>
        </RouterLink>
      </div>
    </div>
  </div>
</template>
