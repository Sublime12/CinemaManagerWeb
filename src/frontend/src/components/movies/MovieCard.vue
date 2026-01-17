<script setup lang="ts">
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import type { Movie } from '@/composables/movies/queries';
import { toRefs } from 'vue';

const props = defineProps<{
  movie: Movie;
}>();

const { movie } = toRefs(props);
</script>

<template>
  <Card class="w-72">
    <RouterLink
      :to="{
        name: 'movie',
        params: {
          id: movie.id,
        },
      }"
    >
      <CardHeader>
        <CardTitle>{{ movie.name }} </CardTitle>
        <!-- <CardDescription>{{ movie.description }}</CardDescription> -->
      </CardHeader>
      <CardContent>
        <img class="rounded-lg" src="@/assets/movie-img-1.webp" alt="Movie a" />
      </CardContent>
      <CardFooter>
        <p>Watch trailer</p>
      </CardFooter>
    </RouterLink>
  </Card>
</template>
