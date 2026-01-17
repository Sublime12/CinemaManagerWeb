<script setup lang="ts">
import { onMounted } from 'vue';
import { useLogoutMutation } from '@/composables/auth/queries.ts';
import router, { ROUTE_NAME } from '@/router';

const { mutateAsync } = useLogoutMutation()

onMounted(async () => {
  await mutateAsync();
  router.push({
    name: ROUTE_NAME.HOME,
  })
});
</script>

<template></template>
