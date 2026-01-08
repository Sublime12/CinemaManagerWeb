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
import { VueQueryDevtools } from '@tanstack/vue-query-devtools'
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
    <div class="flex justify-between items-center py-6 sticky top-0 bg-white">
      <div class="flex space-x-2">
        <Menu />
        <span> Menu</span>
      </div>
      <div class="flex flex-col items-center">
        <Clapperboard />
        <div class="font-bold">SubCine</div>
      </div>
      <NavigationMenu class="w-full">
        <NavigationMenuList class="flex space-x-2 justify-between px-8 max-w-none">
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
    <RouterView />
    <VueQueryDevtools />
  </div>
</template>
