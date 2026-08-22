<script setup lang="ts">
import { useGetMovieQuery } from '@/composables/movies/queries';
import { computed, toRefs, ref, watchEffect } from 'vue';
import moment from 'moment';
import {
  Clock,
  Star,
  Film,
  Ticket,
  Calendar,
  Globe,
  ArrowLeft,
  Play,
  Sparkles,
  MapPin,
  CheckCircle2,
} from 'lucide-vue-next';
import { ROUTE_NAME } from '@/router';

const props = defineProps<{
  id: string;
}>();

const { id } = toRefs(props);
const { data: movie, isFetching, isError, error } = useGetMovieQuery(id);

const selectedDate = ref('Today');
const selectedShowtime = ref<string | null>('19:15');

const formattedDate = computed(() => {
  if (!movie.value) return undefined;
  return moment(movie.value.published_at).format('MMMM D, YYYY');
});

const formattedDuration = computed(() => {
  if (!movie.value?.length) return '2h 0m';
  const hours = movie.value.length.hours();
  const mins = movie.value.length.minutes();
  if (hours > 0) {
    return `${hours}h ${mins}m`;
  }
  return `${movie.value.length.asMinutes()}m`;
});

const dates = ['Today, Aug 22', 'Tomorrow, Aug 23', 'Sun, Aug 24', 'Mon, Aug 25'];
const showtimes = [
  { time: '13:00', hall: 'Screen 1 (IMAX 3D)', price: '$16.50' },
  { time: '16:15', hall: 'Screen 2 (Dolby Atmos)', price: '$14.00' },
  { time: '19:15', hall: 'Screen 1 (IMAX 3D)', price: '$18.00' },
  { time: '22:00', hall: 'Screen 3 (VIP Suite)', price: '$22.00' },
];

watchEffect(() => {
  if (error.value) {
    console.error('Error: ', error.value);
  }
});
</script>

<template>
  <div class="space-y-8 pb-12">
    <!-- Top Back Navigation -->
    <div>
      <RouterLink
        :to="{ name: ROUTE_NAME.MOVIES }"
        class="bg-card border-border/60 text-muted-foreground hover:text-foreground hover:border-primary/50 inline-flex items-center gap-2 rounded-xl border px-3 py-1.5 text-xs font-semibold transition-all"
      >
        <ArrowLeft class="h-4 w-4" />
        <span>Back to Movies Catalog</span>
      </RouterLink>
    </div>

    <!-- Loading State -->
    <div v-if="isFetching" class="text-muted-foreground p-12 text-center">
      <Film class="text-primary mx-auto mb-3 h-10 w-10 animate-spin" />
      <p class="text-sm font-semibold">Loading movie details...</p>
    </div>

    <!-- Error State -->
    <div
      v-else-if="isError"
      class="bg-destructive/10 border-destructive/20 space-y-2 rounded-2xl border p-8 text-center"
    >
      <p class="text-destructive font-semibold">Could not load movie information.</p>
      <p class="text-muted-foreground text-xs">{{ error?.message }}</p>
    </div>

    <!-- Movie Details Content -->
    <div v-else-if="movie" class="space-y-10">
      <!-- Hero Header Box -->
      <div
        class="bg-card border-border/80 relative overflow-hidden rounded-3xl border p-6 shadow-2xl md:p-10"
      >
        <div
          class="from-card via-card/90 absolute inset-0 z-10 bg-gradient-to-r to-transparent"
        ></div>
        <div
          class="absolute top-0 right-0 bottom-0 w-full bg-cover bg-center opacity-30 blur-xs md:w-2/3"
          style="background-image: url('/src/assets/movie-img-1.webp')"
        ></div>

        <div class="relative z-20 flex flex-col items-start gap-8 md:flex-row md:items-center">
          <!-- Movie Poster Cover -->
          <div
            class="group relative aspect-[2/3] w-48 shrink-0 overflow-hidden rounded-2xl border border-white/10 bg-slate-900 shadow-2xl md:w-64"
          >
            <img
              src="@/assets/movie-img-1.webp"
              :alt="movie.name"
              class="h-full w-full object-cover"
            />
            <div
              class="absolute top-3 left-3 flex items-center gap-1 rounded-md border border-amber-400/30 bg-black/70 px-2.5 py-1 text-xs font-bold text-amber-400 backdrop-blur-md"
            >
              <Star class="h-3.5 w-3.5 fill-amber-400" />
              <span>8.5 / 10</span>
            </div>
          </div>

          <!-- Movie Info Block -->
          <div class="flex-1 space-y-4">
            <div class="flex flex-wrap items-center gap-2">
              <span
                v-for="genre in movie.genres"
                :key="genre"
                class="bg-primary/20 text-primary border-primary/30 rounded-full border px-3 py-1 text-xs font-semibold"
              >
                {{ genre }}
              </span>
              <span
                class="rounded-full border border-white/10 bg-white/5 px-3 py-1 text-xs font-semibold text-slate-300 uppercase"
              >
                {{ movie.language }}
              </span>
            </div>

            <h1 class="text-3xl leading-tight font-black tracking-tight text-white md:text-5xl">
              {{ movie.name }}
            </h1>

            <!-- Quick Specs Row -->
            <div class="flex flex-wrap items-center gap-6 text-xs text-slate-300 md:text-sm">
              <div class="flex items-center gap-1.5">
                <Clock class="text-primary h-4 w-4" />
                <span>{{ formattedDuration }}</span>
              </div>
              <div class="flex items-center gap-1.5">
                <Calendar class="text-accent h-4 w-4" />
                <span>{{ formattedDate }}</span>
              </div>
              <div class="flex items-center gap-1.5">
                <Globe class="h-4 w-4 text-blue-400" />
                <span>Subtitles Available</span>
              </div>
            </div>

            <p class="pt-2 text-sm leading-relaxed text-slate-300 md:text-base">
              {{ movie.description }}
            </p>

            <!-- Action CTA Buttons -->
            <div class="flex flex-wrap items-center gap-3 pt-4">
              <a
                href="#showtimes"
                class="bg-primary shadow-primary/25 flex items-center gap-2 rounded-xl px-6 py-3 text-sm font-bold text-white shadow-xl transition-all hover:bg-rose-600"
              >
                <Ticket class="h-4 w-4" />
                <span>Select Showtime & Book</span>
              </a>

              <button
                class="bg-secondary/80 hover:bg-secondary border-border text-foreground flex items-center gap-2 rounded-xl border px-5 py-3 text-sm font-semibold transition-all"
              >
                <Play class="text-primary fill-primary h-4 w-4" />
                <span>Watch Trailer</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Showtime Selection Section -->
      <div
        id="showtimes"
        class="bg-card border-border/80 space-y-6 rounded-3xl border p-6 shadow-xl md:p-8"
      >
        <div
          class="border-border/40 flex flex-col items-start justify-between gap-4 border-b pb-4 md:flex-row md:items-center"
        >
          <div>
            <h2 class="flex items-center gap-2 text-xl font-bold text-white">
              <Ticket class="text-primary h-5 w-5" />
              <span>Available Showtimes & Screens</span>
            </h2>
            <p class="text-muted-foreground mt-0.5 text-xs">
              Select a date and screening time to reserve your seats
            </p>
          </div>

          <!-- Date Selector Pills -->
          <div class="flex w-full items-center gap-2 overflow-x-auto pb-1 md:w-auto">
            <button
              v-for="d in dates"
              :key="d"
              @click="selectedDate = d"
              :class="[
                'shrink-0 rounded-xl border px-4 py-2 text-xs font-semibold transition-all',
                selectedDate === d
                  ? 'bg-primary border-primary text-white shadow-md'
                  : 'bg-secondary/60 hover:bg-secondary text-muted-foreground border-border/60 hover:text-foreground',
              ]"
            >
              {{ d }}
            </button>
          </div>
        </div>

        <!-- Showtimes Slot Grid -->
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 md:grid-cols-4">
          <div
            v-for="slot in showtimes"
            :key="slot.time"
            @click="selectedShowtime = slot.time"
            :class="[
              'relative cursor-pointer space-y-2 rounded-2xl border p-4 transition-all',
              selectedShowtime === slot.time
                ? 'bg-primary/10 border-primary shadow-primary/10 ring-primary shadow-lg ring-1'
                : 'bg-card hover:bg-secondary/60 border-border/60',
            ]"
          >
            <div class="flex items-center justify-between">
              <span class="text-lg font-black tracking-wide text-white">{{ slot.time }}</span>
              <span class="text-accent text-xs font-bold">{{ slot.price }}</span>
            </div>

            <div class="text-muted-foreground flex items-center gap-1 text-xs">
              <MapPin class="text-primary h-3 w-3" />
              <span>{{ slot.hall }}</span>
            </div>

            <div v-if="selectedShowtime === slot.time" class="absolute top-2 right-2">
              <CheckCircle2 class="text-primary h-4 w-4" />
            </div>
          </div>
        </div>

        <!-- Checkout / Reserve Action Bar -->
        <div
          class="border-border/40 bg-secondary/30 flex flex-col items-center justify-between gap-4 rounded-2xl border-t p-4 pt-4 md:flex-row"
        >
          <div class="text-xs text-slate-300">
            <span class="font-bold text-white">Selected Showtime:</span>
            <span class="text-primary ml-2 font-semibold"
              >{{ selectedDate }} at {{ selectedShowtime || '19:15' }}</span
            >
          </div>

          <button
            class="bg-primary shadow-primary/25 flex w-full items-center justify-center gap-2 rounded-xl px-8 py-3 text-sm font-bold text-white shadow-lg transition-all hover:bg-rose-600 md:w-auto"
          >
            <Ticket class="h-4 w-4" />
            <span>Proceed to Seat Selection</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
