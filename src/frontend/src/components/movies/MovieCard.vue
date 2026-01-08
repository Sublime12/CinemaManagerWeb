<script setup lang="ts">
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import type { Movie } from '@/composables/movies/queries';
import { toRefs } from 'vue';

const props = defineProps<{
  movie: Movie,
}>();

const { movie } = toRefs(props);
</script>

<template>
  <Card class="w-72">
    <CardHeader>
      <CardTitle>{{ movie.name }} </CardTitle>
      <!-- <CardDescription>{{ movie.description }}</CardDescription> -->
    </CardHeader>
    <CardContent>
      <RouterLink :to="{
        name: 'movie',
        params: {
          id: movie.id,
        }
      }">
      <img class="rounded-lg" src="@/assets/movie-img-1.webp" alt="Movie a">
      </RouterLink>
    </CardContent>
    <CardFooter>
      <p>Watch trailer</p>
    </CardFooter>
  </Card>
</template>
