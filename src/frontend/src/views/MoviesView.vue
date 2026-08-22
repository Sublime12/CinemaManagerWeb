<script setup lang="ts">
import MovieCard from '@/components/movies/MovieCard.vue';
import { useGetMoviesQuery, type Movie } from '@/composables/movies/queries';
import { ROUTE_NAME } from '@/router';
import { ref, computed } from 'vue';
import { Film, Search, Sparkles, Ticket, Play, Filter, Flame } from 'lucide-vue-next';

const { data: movies, isFetching, isError, error } = useGetMoviesQuery();

const searchQuery = ref('');
const selectedGenre = ref('All');

const availableGenres = computed(() => {
  if (!movies.value) return ['All'];
  const genresSet = new Set<string>();
  movies.value.forEach((m) => {
    if (m.genres) {
      m.genres.forEach((g) => genresSet.add(g));
    }
  });
  return ['All', ...Array.from(genresSet)];
});

const filteredMovies = computed(() => {
  if (!movies.value) return [];
  return movies.value.filter((m) => {
    const matchesSearch = searchQuery.value
      ? m.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
        m.description.toLowerCase().includes(searchQuery.value.toLowerCase())
      : true;

    const matchesGenre =
      selectedGenre.value === 'All' ? true : m.genres && m.genres.includes(selectedGenre.value);

    return matchesSearch && matchesGenre;
  });
});
</script>

<template>
  <div class="space-y-10 pb-12">
    <!-- Hero Spotlight Showcase -->
    <div
      class="bg-card border-border/80 relative overflow-hidden rounded-3xl border p-6 shadow-2xl md:p-10"
    >
      <div
        class="from-card via-card/90 absolute inset-0 z-10 bg-gradient-to-r to-transparent"
      ></div>
      <div
        class="absolute top-0 right-0 bottom-0 w-full bg-cover bg-center opacity-40 blur-xs md:w-2/3"
        style="background-image: url('/src/assets/movie-img-1.webp')"
      ></div>

      <div class="relative z-20 max-w-2xl space-y-4">
        <div
          class="bg-primary/20 border-primary/30 text-primary inline-flex items-center gap-2 rounded-full border px-3 py-1 text-xs font-semibold tracking-wider uppercase"
        >
          <Flame class="fill-primary h-3.5 w-3.5" />
          <span>Featured Blockbuster</span>
        </div>

        <h1 class="text-3xl leading-tight font-black tracking-tight text-white md:text-5xl">
          Experience Cinema Like Never Before
        </h1>

        <p class="text-sm leading-relaxed text-slate-300 md:text-base">
          Book your tickets now for the latest IMAX 3D blockbusters, Dolby Atmos sound auditoriums,
          and exclusive premiere events.
        </p>

        <div class="flex flex-wrap items-center gap-3 pt-2">
          <RouterLink
            v-if="movies && movies.length > 0"
            :to="{ name: ROUTE_NAME.MOVIE, params: { id: movies[0].id } }"
            class="bg-primary shadow-primary/25 flex items-center gap-2 rounded-xl px-6 py-3 text-sm font-bold text-white shadow-xl transition-all hover:scale-102 hover:bg-rose-600"
          >
            <Ticket class="h-4 w-4" />
            <span>Book Tickets Now</span>
          </RouterLink>

          <button
            class="bg-secondary/80 hover:bg-secondary border-border text-foreground flex items-center gap-2 rounded-xl border px-5 py-3 text-sm font-semibold transition-all"
          >
            <Play class="text-primary fill-primary h-4 w-4" />
            <span>Watch Trailers</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Filter & Search Controls Header -->
    <div
      class="border-border/40 flex flex-col items-start justify-between gap-4 border-b pb-6 md:flex-row md:items-center"
    >
      <div>
        <h2 class="flex items-center gap-2 text-2xl font-bold text-white">
          <Film class="text-primary h-6 w-6" />
          <span>Now Showing</span>
          <span
            v-if="filteredMovies.length"
            class="bg-secondary text-muted-foreground rounded-full px-2.5 py-0.5 text-xs font-medium"
          >
            {{ filteredMovies.length }} Movies
          </span>
        </h2>
        <p class="text-muted-foreground mt-1 text-xs">
          Explore showtimes and reserve your favorite seats
        </p>
      </div>

      <!-- Search Input -->
      <div class="relative w-full md:w-72">
        <Search class="text-muted-foreground absolute top-1/2 left-3.5 h-4 w-4 -translate-y-1/2" />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search by title or description..."
          class="bg-card border-border text-foreground placeholder:text-muted-foreground focus:border-primary focus:ring-primary w-full rounded-xl border py-2.5 pr-4 pl-10 text-sm transition-all focus:ring-1 focus:outline-none"
        />
      </div>
    </div>

    <!-- Genre Category Filters -->
    <div class="scrollbar-none flex items-center gap-2 overflow-x-auto pb-2">
      <span
        class="text-muted-foreground mr-1 flex shrink-0 items-center gap-1 text-xs font-semibold"
      >
        <Filter class="text-primary h-3.5 w-3.5" /> Genres:
      </span>
      <button
        v-for="genre in availableGenres"
        :key="genre"
        @click="selectedGenre = genre"
        :class="[
          'shrink-0 rounded-full border px-4 py-1.5 text-xs font-semibold transition-all',
          selectedGenre === genre
            ? 'bg-primary border-primary shadow-primary/20 text-white shadow-md'
            : 'bg-card hover:bg-secondary text-muted-foreground border-border/60 hover:text-foreground',
        ]"
      >
        {{ genre }}
      </button>
    </div>

    <!-- Loading State -->
    <div
      v-if="isFetching"
      class="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4"
    >
      <div
        v-for="n in 4"
        :key="n"
        class="bg-card border-border/40 animate-pulse space-y-4 rounded-2xl border p-4"
      >
        <div class="aspect-[2/3] w-full rounded-xl bg-slate-800/60"></div>
        <div class="h-4 w-3/4 rounded bg-slate-800/60"></div>
        <div class="h-3 w-1/2 rounded bg-slate-800/60"></div>
      </div>
    </div>

    <!-- Error State -->
    <div
      v-else-if="isError"
      class="bg-destructive/10 border-destructive/20 space-y-3 rounded-2xl border p-8 text-center"
    >
      <p class="text-destructive font-semibold">Failed to load movie catalog.</p>
      <p class="text-muted-foreground text-xs">
        {{ error?.message || 'Please check backend server connection.' }}
      </p>
    </div>

    <!-- Movies Grid View -->
    <div
      v-else-if="filteredMovies.length > 0"
      class="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4"
    >
      <div v-for="movie in filteredMovies" :key="movie.id">
        <RouterLink :to="{ name: ROUTE_NAME.MOVIE, params: { id: movie.id } }">
          <MovieCard :movie="movie" />
        </RouterLink>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else class="bg-card/40 border-border/40 space-y-3 rounded-3xl border py-16 text-center">
      <Film class="text-muted-foreground mx-auto h-12 w-12 opacity-40" />
      <h3 class="text-lg font-bold text-white">No movies match your query</h3>
      <p class="text-muted-foreground mx-auto max-w-sm text-xs">
        Try clearing search keywords or choosing a different genre category.
      </p>
      <button
        @click="
          searchQuery = '';
          selectedGenre = 'All';
        "
        class="bg-secondary hover:bg-secondary/80 text-foreground border-border rounded-xl border px-4 py-2 text-xs font-semibold"
      >
        Reset Filters
      </button>
    </div>
  </div>
</template>
