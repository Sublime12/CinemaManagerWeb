<script setup lang="ts">
import MovieCard from '@/components/movies/MovieCard.vue';
import { useGetMoviesQuery } from '@/composables/movies/queries';
import { watchEffect } from 'vue';

const { data: movies, isFetching, isError, error } = useGetMoviesQuery();
watchEffect(() => {
  if (isError.value) {
    console.log('Error: ', error.value)
  }
});
</script>
<template>
  <div>
    <div v-if="isFetching">Loading movies...</div>
    <div v-if="isError">Error while loading movies...</div>
    <div v-else class="flex flex-wrap justify-baseline space-x-4 space-y-10">
      <div v-for="m in movies">
        <MovieCard :movie="m"/>
      </div>
    </div>
  </div>
</template>
