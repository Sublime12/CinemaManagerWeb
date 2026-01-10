<script setup lang="ts">
import { useGetMovieQuery } from "@/composables/movies/queries";
import { computed, toRefs, watchEffect } from "vue";
import ItemGroup from "@/components/ui/item/ItemGroup.vue";
import { CardTitle } from "@/components/ui/card";
import moment from "moment";

const props = defineProps<{
  id: string;
}>();

const { id } = toRefs(props);
const { data: movie, isFetching, isError, error } = useGetMovieQuery(id);
const formattedDate = computed(() => {
  if (!movie.value) return undefined;
  return moment(movie.value.published_at).format("MMMM D, YYYY");
});

const formattedDuration = computed(() => {
  if (!movie.value) return undefined;
  return movie.value.length.humanize();
});

watchEffect(() => {
  if (error.value) {
    console.error("Error: ", error.value);
  }
});
</script>

<template>
  <div v-if="isFetching">Loading movie...</div>
  <div v-else-if="isError">Error while loading movie...</div>
  <ItemGroup v-else-if="movie" class="space-y-7">
    <div>
      <CardTitle class="text-5xl">
        {{ movie.name }}
      </CardTitle>
      <span>
        {{ formattedDate }}
      </span>
    </div>
    <div class="flex space-x-6">
      <div>
        <img
          class="rounded-lg border-3 border-gray-300 w-70"
          src="@/assets/movie-img-1.webp"
          alt=""
        />
      </div>
      <!-- movie details -->
      <div class="flex flex-col space-y-7">
        <div class="flex space-x-13 justify-between">
          <div class="">
            <span class="font-semibold">LENGTH</span><br />
            <span> {{ formattedDuration }} </span>
          </div>
          <div>
            <span class="font-semibold">RATINGS</span><br />
            <span> 7.00 </span>
          </div>
          <div>
            <span class="font-semibold">LANGUAGE</span><br />
            <span> {{ movie.language }} </span>
          </div>
        </div>
        <div>
          <span class="font-semibold">GENRE</span><br />
          <div class="space-x-2">
            <span v-for="genre in movie.genres"> {{ genre }}, </span>
          </div>
        </div>
      </div>
    </div>
    <div class="text-base">
      {{ movie.description }}
    </div>
  </ItemGroup>
</template>
