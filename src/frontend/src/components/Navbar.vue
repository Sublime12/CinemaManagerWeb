<script setup lang="ts">
import { ROUTE_NAME } from '@/router';
import { reactive } from 'vue';
import { NavigationMenuList } from './ui/navigation-menu';
import { NavigationMenu } from './ui/navigation-menu';
import { NavigationMenuItem } from './ui/navigation-menu';
import { NavigationMenuLink } from './ui/navigation-menu';

const links = reactive([
  { name: 'Login', route: ROUTE_NAME.LOGIN, sortOrder: 1 },
  { name: 'Home', route: ROUTE_NAME.HOME, sortOrder: 1 },
]);
</script>

<template>
<div class="sticky top-0 flex items-center justify-between bg-white py-6">
  <div class="flex space-x-2">
    <Menu />
    <span> Menu</span>
  </div>
  <div class="flex flex-col items-center">
    <Clapperboard />
    <div class="font-bold">SubCine</div>
  </div>
  <NavigationMenu class="w-full">
    <NavigationMenuList class="flex max-w-none justify-between space-x-2 px-8">
      <NavigationMenuItem
        v-for="link in links"
        class="hover:text-white rounded hover:font-semibold hover:border hover:border-gray-300"
      >
        <NavigationMenuLink>
          <RouterLink :to="{ name: link.route }">{{ link.name }}</RouterLink>
        </NavigationMenuLink>
      </NavigationMenuItem>
    </NavigationMenuList>
  </NavigationMenu>
</div>
</template>
