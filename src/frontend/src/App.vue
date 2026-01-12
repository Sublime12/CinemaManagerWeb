<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import NavigationMenu from '@/components/ui/navigation-menu/NavigationMenu.vue';
import {
  NavigationMenuContent,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from 'reka-ui';
import NavigationMenuItem from '@/components/ui/navigation-menu/NavigationMenuItem.vue';
import { Clapperboard, Menu } from 'lucide-vue-next';
import { VueQueryDevtools } from '@tanstack/vue-query-devtools';
</script>

<template>
  <!-- <header> -->
  <!--   <nav> -->
  <!--     <RouterLink to="/">Home</RouterLink> -->
  <!--     <RouterLink to="/about">About</RouterLink> -->
  <!--     <RouterLink to="/movies">Movies</RouterLink> -->
  <!--   </nav> -->
  <!-- </header> -->

  <div class="px-8">
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
          <NavigationMenuItem>
            <NavigationMenuLink>
              <RouterLink to="/">Threatre</RouterLink>
            </NavigationMenuLink>
          </NavigationMenuItem>
          <NavigationMenuLink>
            <RouterLink to="/">Search</RouterLink>
          </NavigationMenuLink>

          <NavigationMenuItem> </NavigationMenuItem>
          <NavigationMenuItem>
            <NavigationMenuLink>
              <RouterLink to="/">Search</RouterLink>
            </NavigationMenuLink>
          </NavigationMenuItem>
          <NavigationMenuItem>
            <NavigationMenuLink>
              <RouterLink to="/">Home</RouterLink>
            </NavigationMenuLink>
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
    </div>
    <!-- body -->
    <div class="flex justify-center">
      <div class="container max-w-7xl">
        <RouterView />
      </div>
    </div>
    <!-- end body -->
    <VueQueryDevtools />
  </div>
</template>
