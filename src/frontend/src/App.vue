<script setup lang="ts">
import { RouterView } from 'vue-router';

import { VueQueryDevtools } from '@tanstack/vue-query-devtools';
import { Toaster } from 'vue-sonner';
import Navbar from '@/components/Navbar.vue';
</script>

<template>
  <div class="px-8">
    <Navbar />
    <Toaster class="w-full"/>
    <!-- body -->
    <div class="flex justify-center">
      <div class="container max-w-7xl">
        <RouterView />
      </div>
    </div>
    <!-- end body -->
    <VueQueryDevtools />
  </div>
</template>
