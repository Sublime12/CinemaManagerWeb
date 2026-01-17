<script setup lang="ts">
import MovieCard from '@/components/movies/MovieCard.vue';
import { useGetMoviesQuery, type Movie } from '@/composables/movies/queries';
import router, { ROUTE_NAME } from '@/router';
import { watchEffect } from 'vue';

const { data: movies, isFetching, isError, error } = useGetMoviesQuery();

function moveToPageMovie(m: Movie) {}
watchEffect(() => {
  if (isError.value) {
    console.log('Error: ', error.value);
  }
});
</script>
<template>
  <div>
    <div v-if="isFetching">Loading movies...</div>
    <div v-if="isError">Error while loading movies...</div>
    <div v-else class="flex flex-wrap space-y-10 space-x-4">
      <div v-for="m in movies" :key="m.id">
        <RouterLink
          :to="{
            name: ROUTE_NAME.MOVIE,
            params: {
              id: m.id,
            },
          }"
        >
          <MovieCard :movie="m" class="cursor-pointer" @click="moveToPageMovie(m)" />
        </RouterLink>
      </div>
    </div>
  </div>
</template>
