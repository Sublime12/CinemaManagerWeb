<script setup lang="ts">
import type { Movie } from '@/composables/movies/queries';
import { Clock, Star, Ticket, Calendar, Globe } from 'lucide-vue-next';
import { computed } from 'vue';
import moment from 'moment';
import defaultPoster from '@/assets/movie-img-1.webp';

const props = defineProps<{
  movie: Movie;
  isAdmin?: boolean;
}>();

const posterUrl = computed(() => {
  if (props.movie?.image_url) {
    if (props.movie.image_url.startsWith('http')) return props.movie.image_url;
    return `/api${props.movie.image_url}`;
  }
  return defaultPoster;
});

const formattedDuration = computed(() => {
  if (!props.movie?.length) return '2h 0m';
  const hours = props.movie.length.hours();
  const mins = props.movie.length.minutes();
  if (hours > 0) {
    return `${hours}h ${mins}m`;
  }
  return `${props.movie.length.asMinutes()}m`;
});

const releaseYear = computed(() => {
  if (!props.movie?.published_at) return '';
  return moment(props.movie.published_at).format('YYYY');
});
</script>

<template>
  <div
    class="group bg-card border-border/60 hover:shadow-primary/10 hover:border-primary/50 relative flex h-full flex-col overflow-hidden rounded-2xl border shadow-lg transition-all duration-300 hover:-translate-y-1 hover:shadow-2xl"
  >
    <!-- Poster Image Container -->
    <div class="relative aspect-[2/3] w-full overflow-hidden bg-slate-900">
      <img
        :src="posterUrl"
        :alt="movie.name"
        class="h-full w-full object-cover object-center transition-transform duration-500 group-hover:scale-105"
        @error="(e: Event) => ((e.target as HTMLImageElement).src = defaultPoster)"
      />

      <!-- Top Overlay Badges -->
      <div
        class="pointer-events-none absolute top-3 right-3 left-3 flex items-center justify-between"
      >
        <span
          class="inline-flex items-center gap-1 rounded-md border border-amber-400/30 bg-black/70 px-2.5 py-1 text-[11px] font-bold text-amber-400 backdrop-blur-md"
        >
          <Star class="h-3 w-3 fill-amber-400" />
          <span>8.5</span>
        </span>

        <span
          class="inline-flex items-center gap-1 rounded-md border border-white/10 bg-black/70 px-2.5 py-1 text-[11px] font-semibold text-slate-200 uppercase backdrop-blur-md"
        >
          <Globe class="text-primary h-3 w-3" />
          <span>{{ movie.language }}</span>
        </span>
      </div>

      <!-- Bottom Gradient Overlay -->
      <div
        class="from-card via-card/20 absolute inset-0 bg-gradient-to-t to-transparent opacity-90 transition-opacity group-hover:opacity-75"
      ></div>

      <!-- Quick Action Overlay Button -->
      <div
        class="absolute right-4 bottom-4 left-4 translate-y-2 opacity-0 transition-all duration-300 group-hover:translate-y-0 group-hover:opacity-100"
      >
        <button
          class="bg-primary shadow-primary/30 flex w-full items-center justify-center gap-2 rounded-xl px-4 py-2.5 text-xs font-semibold text-white shadow-lg hover:bg-rose-600"
        >
          <Ticket class="h-4 w-4" />
          <span>Book Tickets</span>
        </button>
      </div>
    </div>

    <!-- Movie Details Content -->
    <div class="flex flex-1 flex-col justify-between space-y-3 p-4">
      <div>
        <div class="text-muted-foreground mb-1 flex items-center justify-between text-xs">
          <span class="inline-flex items-center gap-1 text-slate-400">
            <Clock class="text-primary/80 h-3.5 w-3.5" />
            <span>{{ formattedDuration }}</span>
          </span>
          <span v-if="releaseYear" class="inline-flex items-center gap-1 text-slate-400">
            <Calendar class="h-3.5 w-3.5" />
            <span>{{ releaseYear }}</span>
          </span>
        </div>

        <h3
          class="text-foreground group-hover:text-primary line-clamp-1 text-base font-bold transition-colors"
        >
          {{ movie.name }}
        </h3>

        <p class="text-muted-foreground mt-1.5 line-clamp-2 text-xs leading-relaxed">
          {{ movie.description }}
        </p>
      </div>

      <!-- Genres Tags -->
      <div class="flex flex-wrap gap-1.5 pt-1">
        <span
          v-for="genre in movie.genres"
          :key="genre"
          class="bg-primary/10 text-primary border-primary/20 rounded-md border px-2 py-0.5 text-[10px] font-semibold"
        >
          {{ genre }}
        </span>
      </div>
    </div>
  </div>
</template>
