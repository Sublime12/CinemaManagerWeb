<script setup="ts">
import { useHelloworld } from '@/components/Helloworld/query';

const { data } = useHelloworld();
</script>

<template>
  <div class="about">
    <h1>This is an about pagexx {{ data }}</h1>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
