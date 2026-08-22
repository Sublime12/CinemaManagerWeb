<script setup lang="ts">
import { RouterView } from 'vue-router';
import { VueQueryDevtools } from '@tanstack/vue-query-devtools';
import { Toaster } from 'vue-sonner';
import Navbar from '@/components/Navbar.vue';
import { Film, Heart, ShieldCheck, Ticket } from 'lucide-vue-next';
</script>

<template>
  <div
    class="dark bg-background text-foreground selection:bg-primary/30 flex min-h-screen flex-col antialiased"
  >
    <!-- Main Top Navigation -->
    <Navbar />

    <!-- Sonner Notifications -->
    <Toaster position="top-right" theme="dark" richColors class="z-50 w-full" />

    <!-- Main Body Container -->
    <main class="container mx-auto max-w-7xl flex-1 px-4 py-6 md:px-8">
      <RouterView />
    </main>

    <!-- Footer -->
    <footer
      class="border-border/40 bg-card/40 text-muted-foreground mt-16 border-t py-10 text-sm backdrop-blur-md"
    >
      <div class="container mx-auto max-w-7xl px-4 md:px-8">
        <div class="mb-8 grid grid-cols-1 gap-8 md:grid-cols-4">
          <div class="space-y-3">
            <div class="text-foreground flex items-center gap-2 text-lg font-bold">
              <div class="from-primary to-accent rounded-lg bg-gradient-to-tr p-1.5 text-white">
                <Film class="h-5 w-5" />
              </div>
              <span
                class="bg-gradient-to-r from-white via-slate-200 to-slate-400 bg-clip-text text-transparent"
                >SubCine Manager</span
              >
            </div>
            <p class="text-muted-foreground text-xs leading-relaxed">
              Elevating the moviegoing experience with seamless screening schedules, interactive
              ticket booking, and real-time box office management.
            </p>
          </div>

          <div>
            <h4 class="text-foreground mb-3 flex items-center gap-1.5 text-sm font-semibold">
              <Ticket class="text-primary h-4 w-4" /> Cinema Experience
            </h4>
            <ul class="space-y-2 text-xs">
              <li class="hover:text-primary cursor-pointer transition-colors">Now Showing</li>
              <li class="hover:text-primary cursor-pointer transition-colors">
                IMAX & Dolby Cinema
              </li>
              <li class="hover:text-primary cursor-pointer transition-colors">Upcoming Releases</li>
              <li class="hover:text-primary cursor-pointer transition-colors">Concessions & F&B</li>
            </ul>
          </div>

          <div>
            <h4 class="text-foreground mb-3 flex items-center gap-1.5 text-sm font-semibold">
              <ShieldCheck class="text-accent h-4 w-4" /> Management
            </h4>
            <ul class="space-y-2 text-xs">
              <li class="hover:text-primary cursor-pointer transition-colors">Admin Dashboard</li>
              <li class="hover:text-primary cursor-pointer transition-colors">
                Screenings & Showtimes
              </li>
              <li class="hover:text-primary cursor-pointer transition-colors">
                Theatres & Screens
              </li>
              <li class="hover:text-primary cursor-pointer transition-colors">
                Box Office Analytics
              </li>
            </ul>
          </div>

          <div>
            <h4 class="text-foreground mb-3 text-sm font-semibold">System Status</h4>
            <div
              class="inline-flex items-center gap-2 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-3 py-1.5 text-xs font-medium text-emerald-400"
            >
              <span class="h-2 w-2 animate-pulse rounded-full bg-emerald-500"></span>
              All Cinema Systems Operational
            </div>
          </div>
        </div>

        <div
          class="border-border/30 flex flex-col items-center justify-between gap-4 border-t pt-6 text-xs md:flex-row"
        >
          <p>© 2026 SubCine Manager. All rights reserved.</p>
          <p class="flex items-center gap-1">
            Crafted with <Heart class="text-primary fill-primary h-3.5 w-3.5" /> for cinema
            enthusiasts
          </p>
        </div>
      </div>
    </footer>

    <VueQueryDevtools />
  </div>
</template>
