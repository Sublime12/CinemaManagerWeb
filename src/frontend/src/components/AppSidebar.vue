<script setup lang="ts">
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarInset,
  SidebarMenu,
  SidebarMenuItem,
  SidebarProvider,
  SidebarRail,
  SidebarTrigger,
} from '@/components/ui/sidebar';
import {
  GalleryVerticalEnd,
  Film,
  Building2,
  Calendar,
  Users,
  DollarSign,
  TrendingUp,
  Search,
} from 'lucide-vue-next';
import { computed, ref } from 'vue';
import MovieCard from './movies/MovieCard.vue';
import { useGetMoviesQuery } from '@/composables/movies/queries';
import MovieFormDialog from '@/components/movies/MovieFormDialog.vue';
import { ROUTE_NAME } from '@/router';

enum ADMIN_PANEL {
  MOVIES = 'MOVIES',
  THEATRES = 'THEATRES',
  SHOWTIMES = 'SHOWTIMES',
}
const selectedPanel = ref<ADMIN_PANEL>(ADMIN_PANEL.MOVIES);
const searchQuery = ref('');

const { data: movies, isFetching } = useGetMoviesQuery();

const moviesSelected = computed(() => selectedPanel.value === ADMIN_PANEL.MOVIES);
const theatresSelected = computed(() => selectedPanel.value === ADMIN_PANEL.THEATRES);
const showtimesSelected = computed(() => selectedPanel.value === ADMIN_PANEL.SHOWTIMES);

const filteredMovies = computed(() => {
  if (!movies.value) return [];
  if (!searchQuery.value) return movies.value;
  return movies.value.filter(
    (m) =>
      m.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
      m.description.toLowerCase().includes(searchQuery.value.toLowerCase()),
  );
});

// Mock Data for Theatres Management
const theatres = [
  {
    id: 1,
    name: 'SubCine Grand Multiplex',
    city: 'New York',
    screens: 8,
    address: '742 Broadway Ave',
  },
  {
    id: 2,
    name: 'SubCine Downtown IMAX',
    city: 'Los Angeles',
    screens: 5,
    address: '120 Sunset Blvd',
  },
  {
    id: 3,
    name: 'SubCine Cinema Deluxe',
    city: 'Chicago',
    screens: 6,
    address: '450 Michigan Ave',
  },
];
</script>

<template>
  <SidebarProvider class="dark text-foreground min-h-screen">
    <Sidebar class="border-border/60 bg-card border-r">
      <SidebarHeader class="border-border/40 border-b p-4">
        <SidebarMenu>
          <SidebarMenuItem>
            <div class="flex items-center gap-3 px-2 py-1">
              <div
                class="from-primary shadow-primary/20 rounded-xl bg-linear-to-tr via-rose-600 to-amber-500 p-2 text-white shadow-md"
              >
                <GalleryVerticalEnd class="h-5 w-5" />
              </div>
              <div class="grid flex-1 text-left leading-tight">
                <span class="truncate text-sm font-extrabold text-white">SubCine Admin</span>
                <span
                  class="text-muted-foreground truncate text-[10px] font-semibold tracking-wider uppercase"
                  >Cinema Manager</span
                >
              </div>
            </div>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>

      <SidebarContent class="p-3">
        <SidebarGroup>
          <SidebarGroupLabel
            class="text-muted-foreground mb-2 px-2 text-[11px] font-bold tracking-wider uppercase"
          >
            Navigation & Controls
          </SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu class="space-y-1">
              <SidebarMenuItem>
                <button
                  @click="selectedPanel = ADMIN_PANEL.MOVIES"
                  :class="[
                    'flex w-full items-center gap-3 rounded-xl px-3 py-2.5 text-xs font-semibold transition-all',
                    moviesSelected
                      ? 'bg-primary shadow-primary/20 text-white shadow-md'
                      : 'text-muted-foreground hover:bg-secondary hover:text-foreground',
                  ]"
                >
                  <Film class="h-4 w-4" />
                  <span>Movie Listings</span>
                </button>
              </SidebarMenuItem>

              <SidebarMenuItem>
                <button
                  @click="selectedPanel = ADMIN_PANEL.THEATRES"
                  :class="[
                    'flex w-full items-center gap-3 rounded-xl px-3 py-2.5 text-xs font-semibold transition-all',
                    theatresSelected
                      ? 'bg-primary shadow-primary/20 text-white shadow-md'
                      : 'text-muted-foreground hover:bg-secondary hover:text-foreground',
                  ]"
                >
                  <Building2 class="h-4 w-4" />
                  <span>Theatres & Auditoriums</span>
                </button>
              </SidebarMenuItem>

              <SidebarMenuItem>
                <button
                  @click="selectedPanel = ADMIN_PANEL.SHOWTIMES"
                  :class="[
                    'flex w-full items-center gap-3 rounded-xl px-3 py-2.5 text-xs font-semibold transition-all',
                    showtimesSelected
                      ? 'bg-primary shadow-primary/20 text-white shadow-md'
                      : 'text-muted-foreground hover:bg-secondary hover:text-foreground',
                  ]"
                >
                  <Calendar class="h-4 w-4" />
                  <span>Screening Schedules</span>
                </button>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter class="border-border/40 border-t p-4">
        <RouterLink
          :to="{ name: ROUTE_NAME.HOME }"
          class="text-muted-foreground hover:text-primary flex items-center gap-2 text-xs transition-colors"
        >
          <span>← Return to Public Site</span>
        </RouterLink>
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>

    <SidebarInset class="bg-background w-full max-w-none space-y-8 p-6">
      <!-- Top Action Bar Header -->
      <div
        class="border-border/40 flex flex-col items-start justify-between gap-4 border-b pb-6 sm:flex-row sm:items-center"
      >
        <div class="flex items-center gap-3">
          <SidebarTrigger class="bg-card border-border rounded-xl border p-2" />
          <div>
            <h1 class="text-2xl font-black tracking-tight text-white">Admin Management Console</h1>
            <p class="text-muted-foreground text-xs">
              Manage movie catalogs, screening halls, and box office activity
            </p>
          </div>
        </div>

        <div class="flex items-center gap-3">
          <MovieFormDialog />
        </div>
      </div>

      <!-- Quick Metrics Analytics Overview -->
      <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
        <div class="bg-card border-border/70 space-y-2 rounded-2xl border p-5 shadow-lg">
          <div class="text-muted-foreground flex items-center justify-between">
            <span class="text-xs font-semibold">Active Movies</span>
            <Film class="text-primary h-4 w-4" />
          </div>
          <div class="text-2xl font-black text-white">{{ movies?.length || 0 }}</div>
          <p class="flex items-center gap-1 text-[11px] font-medium text-emerald-400">
            <TrendingUp class="h-3 w-3" /> +2 added this week
          </p>
        </div>

        <div class="bg-card border-border/70 space-y-2 rounded-2xl border p-5 shadow-lg">
          <div class="text-muted-foreground flex items-center justify-between">
            <span class="text-xs font-semibold">Theatres & Complexes</span>
            <Building2 class="text-accent h-4 w-4" />
          </div>
          <div class="text-2xl font-black text-white">{{ theatres.length }}</div>
          <p class="text-muted-foreground text-[11px]">19 total screens active</p>
        </div>

        <div class="bg-card border-border/70 space-y-2 rounded-2xl border p-5 shadow-lg">
          <div class="text-muted-foreground flex items-center justify-between">
            <span class="text-xs font-semibold">Today's Box Office</span>
            <DollarSign class="h-4 w-4 text-emerald-400" />
          </div>
          <div class="text-2xl font-black text-white">$14,850</div>
          <p class="flex items-center gap-1 text-[11px] font-medium text-emerald-400">
            <TrendingUp class="h-3 w-3" /> +14.2% vs yesterday
          </p>
        </div>

        <div class="bg-card border-border/70 space-y-2 rounded-2xl border p-5 shadow-lg">
          <div class="text-muted-foreground flex items-center justify-between">
            <span class="text-xs font-semibold">Seat Occupancy</span>
            <Users class="h-4 w-4 text-blue-400" />
          </div>
          <div class="text-2xl font-black text-white">84%</div>
          <p class="text-[11px] text-slate-400">Peak hour 19:00 - 22:00</p>
        </div>
      </div>

      <!-- Main Panel View -->
      <div>
        <!-- Movies Panel -->
        <div v-if="moviesSelected" class="space-y-6">
          <div class="flex flex-col items-start justify-between gap-4 sm:flex-row sm:items-center">
            <h2 class="flex items-center gap-2 text-lg font-bold text-white">
              <Film class="text-primary h-5 w-5" />
              <span>Movies Catalog</span>
            </h2>

            <div class="relative w-full sm:w-64">
              <Search
                class="text-muted-foreground absolute top-1/2 left-3 h-4 w-4 -translate-y-1/2"
              />
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Filter movies..."
                class="bg-card border-border text-foreground placeholder:text-muted-foreground focus:border-primary w-full rounded-xl border py-2 pr-4 pl-9 text-xs focus:outline-none"
              />
            </div>
          </div>

          <div v-if="isFetching" class="text-muted-foreground p-8 text-center">
            Loading movie collection...
          </div>

          <div
            v-else-if="filteredMovies.length > 0"
            class="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4"
          >
            <div v-for="movie in filteredMovies" :key="movie.id">
              <MovieCard :movie="movie" :is-admin="true" />
            </div>
          </div>

          <div
            v-else
            class="bg-card/40 border-border/40 text-muted-foreground rounded-2xl border p-12 text-center text-xs"
          >
            No movies found matching criteria.
          </div>
        </div>

        <!-- Theatres Panel -->
        <div v-else-if="theatresSelected" class="space-y-6">
          <div class="flex items-center justify-between">
            <h2 class="flex items-center gap-2 text-lg font-bold text-white">
              <Building2 class="text-accent h-5 w-5" />
              <span>Cinema Complexes & Auditoriums</span>
            </h2>
          </div>

          <div class="grid grid-cols-1 gap-6 md:grid-cols-3">
            <div
              v-for="th in theatres"
              :key="th.id"
              class="bg-card border-border/70 space-y-3 rounded-2xl border p-6 shadow-lg"
            >
              <div class="flex items-center justify-between">
                <span
                  class="bg-accent/20 text-accent rounded-full px-2.5 py-0.5 text-[10px] font-bold uppercase"
                >
                  {{ th.city }}
                </span>
                <span class="text-muted-foreground text-xs">{{ th.screens }} Screens</span>
              </div>
              <h3 class="text-base font-bold text-white">{{ th.name }}</h3>
              <p class="text-muted-foreground text-xs">{{ th.address }}</p>
              <div class="border-border/40 flex items-center justify-between border-t pt-2 text-xs">
                <span class="font-semibold text-emerald-400">Active</span>
                <button class="text-primary font-semibold hover:underline">
                  Configure Screens →
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Showtimes Panel -->
        <div v-else-if="showtimesSelected" class="space-y-6">
          <div class="flex items-center justify-between">
            <h2 class="flex items-center gap-2 text-lg font-bold text-white">
              <Calendar class="text-primary h-5 w-5" />
              <span>Screening Schedules</span>
            </h2>
          </div>
          <div
            class="bg-card/40 border-border/40 text-muted-foreground rounded-2xl border p-12 text-center text-xs"
          >
            Showtime matrix configuration active. Select a theatre above to schedule screenings.
          </div>
        </div>
      </div>
    </SidebarInset>
  </SidebarProvider>
</template>
