<script setup lang="ts">
import { useGetMovieQuery } from '@/composables/movies/queries';
import { toRefs } from 'vue';
import ItemGroup from '@/components/ui/item/ItemGroup.vue';

const props = defineProps<{
  id: string,
}>();

const { id } = toRefs(props)
const { data: movie, isFetching, isError } = useGetMovieQuery(id);
</script>

<template>
  <div v-if="isFetching"> Loading movie...</div>
  <div v-if="isError"> Error while loading movie...</div>
  <ItemGroup v-else>
    {{ movie?.description }}
  </ItemGroup>
</template>
