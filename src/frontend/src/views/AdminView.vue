<script setup lang="ts">
import AppSidebar from '@/components/AppSidebar.vue';
import { useGetMeQuery } from '@/composables/auth/queries';
import { computed } from 'vue';
import { ROUTE_NAME } from '@/router';
import { ShieldAlert, LogIn, Film } from 'lucide-vue-next';

const { data: meData, isFetching } = useGetMeQuery();

const isLoggedIn = computed(() => !!meData.value?.user);
</script>

<template>
  <div class="min-h-screen">
    <!-- Loading auth status -->
    <div v-if="isFetching" class="flex flex-col items-center justify-center min-h-[60vh] text-slate-400 space-y-3">
      <Film class="w-8 h-8 animate-spin text-rose-500" />
      <p class="text-xs font-semibold">Verifying administrator session...</p>
    </div>

    <!-- Render Admin Dashboard if connected -->
    <AppSidebar v-else-if="isLoggedIn" />

    <!-- Unauthorized Banner if NOT connected -->
    <div v-else class="flex flex-col items-center justify-center min-h-[65vh] px-4">
      <div class="max-w-md w-full rounded-3xl bg-slate-900 border border-slate-800 p-8 text-center space-y-5 shadow-2xl">
        <div class="inline-flex p-4 rounded-2xl bg-rose-500/10 border border-rose-500/20 text-rose-500">
          <ShieldAlert class="w-10 h-10" />
        </div>

        <div class="space-y-2">
          <h2 class="text-2xl font-black text-white">Admin Access Restricted</h2>
          <p class="text-xs text-slate-400 leading-relaxed">
            The Cinema Management Console requires administrator authentication. Please sign in to access movie scheduling and theatre controls.
          </p>
        </div>

        <RouterLink
          :to="{ name: ROUTE_NAME.LOGIN }"
          class="inline-flex items-center justify-center gap-2 w-full py-3 px-5 rounded-xl bg-rose-600 hover:bg-rose-500 text-white font-bold text-xs shadow-lg shadow-rose-600/25 transition-all"
        >
          <LogIn class="w-4 h-4" />
          <span>Sign In to Admin Portal</span>
        </RouterLink>
      </div>
    </div>
  </div>
</template>
