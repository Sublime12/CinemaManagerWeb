<script setup lang="ts">
import { ref } from 'vue';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { toTypedSchema } from '@vee-validate/zod';
import { useForm, Field as VeeField } from 'vee-validate';
import { toast } from 'vue-sonner';
import { CreateMovieFormSchema, useCreateMovieMutation } from '@/composables/movies/queries';
import { Film, Plus, Clock, Globe, Calendar, Tag, FileText } from 'lucide-vue-next';

const isOpen = ref(false);
const formSchema = toTypedSchema(CreateMovieFormSchema);

const { handleSubmit, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues: {
    name: '',
    description: '',
    published_at: new Date().toISOString().split('T')[0],
    length_minutes: 120,
    language: 'English',
    genres: 'Action, Sci-Fi',
  },
});

const { mutateAsync, isPending } = useCreateMovieMutation();

const onSubmit = handleSubmit(async (values) => {
  try {
    await mutateAsync(values);
    toast.success('Movie created successfully!', {
      description: `"${values.name}" has been added to the cinema catalog.`,
    });
    resetForm();
    isOpen.value = false;
  } catch (err: any) {
    toast.error('Failed to create movie', {
      description: err?.response?.data?.message || 'Please check admin authentication.',
    });
  }
});
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button
        class="flex items-center gap-2 rounded-xl bg-rose-600 px-4 py-2 text-xs font-bold text-white shadow-lg shadow-rose-600/25 hover:bg-rose-500"
      >
        <Plus class="h-4 w-4" />
        <span>Add New Movie</span>
      </Button>
    </DialogTrigger>

    <DialogContent
      class="rounded-3xl border-slate-800 bg-slate-900 p-6 text-slate-100 shadow-2xl sm:max-w-lg"
    >
      <DialogHeader class="space-y-1 border-b border-slate-800 pb-4 text-left">
        <DialogTitle class="flex items-center gap-2 text-xl font-bold text-white">
          <Film class="h-5 w-5 text-rose-500" />
          <span>Create New Movie Listing</span>
        </DialogTitle>
        <DialogDescription class="text-xs text-slate-400">
          Fill in the details below to add a new movie to the cinema management database.
        </DialogDescription>
      </DialogHeader>

      <form id="create-movie-form" @submit="onSubmit" class="space-y-4 py-2">
        <!-- Movie Title -->
        <VeeField v-slot="{ field, errors }" name="name">
          <div class="space-y-1">
            <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
              <Film class="h-3.5 w-3.5 text-rose-500" /> Movie Title
            </label>
            <Input
              v-bind="field"
              placeholder="e.g. Inception: Remastered"
              class="w-full rounded-xl border border-slate-700 bg-slate-950 px-3.5 py-2 text-sm text-slate-100 focus:border-rose-500"
            />
            <p v-if="errors.length" class="text-[11px] font-medium text-rose-500">
              {{ errors[0] }}
            </p>
          </div>
        </VeeField>

        <!-- Description -->
        <VeeField v-slot="{ field, errors }" name="description">
          <div class="space-y-1">
            <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
              <FileText class="h-3.5 w-3.5 text-rose-500" /> Plot Synopsis / Description
            </label>
            <textarea
              v-bind="field"
              rows="3"
              placeholder="Enter a compelling overview of the movie..."
              class="w-full rounded-xl border border-slate-700 bg-slate-950 px-3.5 py-2 text-sm text-slate-100 focus:border-rose-500 focus:ring-1 focus:ring-rose-500 focus:outline-none"
            ></textarea>
            <p v-if="errors.length" class="text-[11px] font-medium text-rose-500">
              {{ errors[0] }}
            </p>
          </div>
        </VeeField>

        <!-- Two Column Grid for Metadata -->
        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <!-- Duration (Minutes) -->
          <VeeField v-slot="{ field, errors }" name="length_minutes">
            <div class="space-y-1">
              <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
                <Clock class="h-3.5 w-3.5 text-rose-500" /> Duration (Minutes)
              </label>
              <Input
                type="number"
                v-bind="field"
                placeholder="120"
                class="w-full rounded-xl border border-slate-700 bg-slate-950 px-3.5 py-2 text-sm text-slate-100 focus:border-rose-500"
              />
              <p v-if="errors.length" class="text-[11px] font-medium text-rose-500">
                {{ errors[0] }}
              </p>
            </div>
          </VeeField>

          <!-- Language -->
          <VeeField v-slot="{ field, errors }" name="language">
            <div class="space-y-1">
              <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
                <Globe class="h-3.5 w-3.5 text-rose-500" /> Audio Language
              </label>
              <Input
                v-bind="field"
                placeholder="English"
                class="w-full rounded-xl border border-slate-700 bg-slate-950 px-3.5 py-2 text-sm text-slate-100 focus:border-rose-500"
              />
              <p v-if="errors.length" class="text-[11px] font-medium text-rose-500">
                {{ errors[0] }}
              </p>
            </div>
          </VeeField>
        </div>

        <div class="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <!-- Release Date -->
          <VeeField v-slot="{ field, errors }" name="published_at">
            <div class="space-y-1">
              <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
                <Calendar class="h-3.5 w-3.5 text-rose-500" /> Release Date
              </label>
              <Input
                type="date"
                v-bind="field"
                class="w-full rounded-xl border border-slate-700 bg-slate-950 px-3.5 py-2 text-sm text-slate-100 focus:border-rose-500"
              />
              <p v-if="errors.length" class="text-[11px] font-medium text-rose-500">
                {{ errors[0] }}
              </p>
            </div>
          </VeeField>

          <!-- Genres -->
          <VeeField v-slot="{ field, errors }" name="genres">
            <div class="space-y-1">
              <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
                <Tag class="h-3.5 w-3.5 text-rose-500" /> Genres (Comma separated)
              </label>
              <Input
                v-bind="field"
                placeholder="Action, Sci-Fi, Thriller"
                class="w-full rounded-xl border border-slate-700 bg-slate-950 px-3.5 py-2 text-sm text-slate-100 focus:border-rose-500"
              />
              <p v-if="errors.length" class="text-[11px] font-medium text-rose-500">
                {{ errors[0] }}
              </p>
            </div>
          </VeeField>
        </div>

        <DialogFooter class="flex items-center justify-end gap-2 border-t border-slate-800 pt-4">
          <DialogClose as-child>
            <Button
              type="button"
              variant="outline"
              class="rounded-xl border-slate-700 text-xs text-slate-300 hover:bg-slate-800"
            >
              Cancel
            </Button>
          </DialogClose>

          <Button
            type="submit"
            :disabled="isPending"
            class="rounded-xl bg-rose-600 px-5 py-2 text-xs font-bold text-white shadow-lg shadow-rose-600/20 hover:bg-rose-500"
          >
            <span v-if="isPending">Saving Movie...</span>
            <span v-else>Publish Movie</span>
          </Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
