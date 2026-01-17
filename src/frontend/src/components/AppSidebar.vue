<script setup lang="ts">
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarInset,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarProvider,
  SidebarRail,
  SidebarTrigger,
} from '@/components/ui/sidebar';
import { GalleryVerticalEnd } from 'lucide-vue-next';
import { computed, ref } from 'vue';
import MovieCard from './movies/MovieCard.vue';
import { useGetMoviesQuery } from '@/composables/movies/queries';
import { Home } from 'lucide-vue-next';
import CreateMovieDialog from '@/components/movies/CreateMovieDialog.vue';

enum ADMIN_PANEL {
  MOVIES = 'MOVIES',
  THEATHRES = 'THEATRES',
}
const selectedPanel = ref<ADMIN_PANEL>(ADMIN_PANEL.MOVIES);

const { data: movies } = useGetMoviesQuery();

const selectedClass = 'bg-gray-200 hover:bg-gray-200 border-2';
const moviesSelected = computed(() => selectedPanel.value == ADMIN_PANEL.MOVIES);
const theatresSelected = computed(() => selectedPanel.value == ADMIN_PANEL.THEATHRES);
</script>

<template>
  <SidebarProvider>
    <Sidebar>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg">
              <div
                class="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg"
              >
                <GalleryVerticalEnd class="size-4" />
              </div>
              <div class="grid flex-1 text-left text-sm leading-tight">
                <span class="truncate font-semibold">CineSub</span>
              </div>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Admin</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              <SidebarMenuItem>
                <SidebarMenuButton as-child>
                  <div
                    @click="selectedPanel = ADMIN_PANEL.MOVIES"
                    :class="[moviesSelected ? selectedClass : '']"
                  >
                    <Home />
                    <span>Movies</span>
                  </div>
                </SidebarMenuButton>
                <SidebarMenuButton as-child>
                  <div
                    @click="selectedPanel = ADMIN_PANEL.THEATHRES"
                    :class="[theatresSelected ? selectedClass : '']"
                  >
                    <Home />
                    <span>Theatres</span>
                  </div>
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter />
      <SidebarRail />
    </Sidebar>
    <SidebarInset class="w-full max-w-none">
      <header
        class="flex h-16 shrink-0 items-center justify-between gap-2 transition-[width,height] ease-linear group-has-data-[collapsible=icon]/sidebar-wrapper:h-12"
      >
        <div class="flex items-center gap-2 px-4">
          <SidebarTrigger class="-ml-1" />
        </div>
        <div class="px-4">
          <CreateMovieDialog />
        </div>
      </header>
      <div class="flex flex-1 flex-col gap-4 p-4 pt-0">
        <div class="bg-muted/50 min-h-screen flex-1 rounded-xl md:min-h-min">
          <div
            v-if="selectedPanel == ADMIN_PANEL.MOVIES"
            class="flex flex-row flex-wrap justify-around space-y-2 space-x-4"
          >
            <div v-for="movie in movies">
              <MovieCard :key="movie.id" :movie="movie" />
            </div>
          </div>
        </div>
      </div>
    </SidebarInset>
  </SidebarProvider>
</template>
