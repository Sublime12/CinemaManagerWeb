<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import moment from 'moment';
import defaultPoster from '@/assets/movie-img-1.webp';
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
import {
  CreateMovieFormSchema,
  useCreateMovieMutation,
  useUpdateMovieMutation,
  useUploadPosterMutation,
  type Movie,
} from '@/composables/movies/queries';
import {
  Film,
  Plus,
  Pencil,
  Clock,
  Globe,
  Calendar,
  Tag,
  FileText,
  Image as ImageIcon,
  Upload,
  Loader2,
} from 'lucide-vue-next';

const props = defineProps<{
  movie?: Movie;
}>();

const isEditing = computed(() => !!props.movie);

const isOpen = ref(false);
const fileInput = ref<HTMLInputElement | null>(null);
const previewUrl = ref<string | null>(null);

const formSchema = toTypedSchema(CreateMovieFormSchema);

const resolvePosterUrl = (url?: string | null) => {
  if (!url) return null;
  if (url.startsWith('http') || url.startsWith('blob:') || url.startsWith('data:')) {
    return url;
  }
  if (url.startsWith('/api')) {
    return url;
  }
  return `/api${url.startsWith('/') ? '' : '/'}${url}`;
};

const getInitialValues = () => {
  if (props.movie) {
    let lengthMinutes = 120;
    if (props.movie.length) {
      if (typeof props.movie.length.asMinutes === 'function') {
        lengthMinutes = Math.round(props.movie.length.asMinutes());
      } else if (typeof props.movie.length === 'number') {
        lengthMinutes = Math.round(props.movie.length);
      }
    }

    let publishedAt = new Date().toISOString().split('T')[0];
    if (props.movie.published_at) {
      publishedAt = moment(props.movie.published_at).isValid()
        ? moment(props.movie.published_at).format('YYYY-MM-DD')
        : String(props.movie.published_at).split('T')[0];
    }

    return {
      name: props.movie.name || '',
      description: props.movie.description || '',
      published_at: publishedAt,
      length_minutes: lengthMinutes,
      language: props.movie.language || 'English',
      genres: Array.isArray(props.movie.genres) ? props.movie.genres.join(', ') : '',
      image_url: props.movie.image_url || '',
    };
  }

  return {
    name: '',
    description: '',
    published_at: new Date().toISOString().split('T')[0],
    length_minutes: 120,
    language: 'English',
    genres: 'Action, Sci-Fi',
    image_url: '',
  };
};

const { handleSubmit, resetForm, setValues, setFieldValue } = useForm({
  validationSchema: formSchema,
  initialValues: getInitialValues(),
});

watch(
  () => [isOpen.value, props.movie],
  ([open]) => {
    if (open) {
      const initVals = getInitialValues();
      resetForm({ values: initVals });
      setValues(initVals);
      previewUrl.value = resolvePosterUrl(props.movie?.image_url);
    }
  },
  { immediate: true, deep: true },
);

const { mutateAsync: createMovie, isPending: isCreating } = useCreateMovieMutation();
const { mutateAsync: updateMovie, isPending: isUpdating } = useUpdateMovieMutation();
const { mutateAsync: uploadPoster, isPending: isUploading } = useUploadPosterMutation();

const isPending = computed(() => isCreating.value || isUpdating.value);

const dialogTitle = computed(() =>
  isEditing.value ? 'Edit Movie Listing' : 'Create New Movie Listing',
);

const dialogDescription = computed(() =>
  isEditing.value
    ? 'Modify the details below to update this movie in the database.'
    : 'Fill in the details below to add a new movie to the cinema management database.',
);

const submitButtonText = computed(() => {
  if (isPending.value) {
    return isEditing.value ? 'Saving Changes...' : 'Saving Movie...';
  }
  return isEditing.value ? 'Update Movie' : 'Publish Movie';
});

const handleFileSelect = async (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (!target.files || target.files.length === 0) return;

  const file = target.files[0];
  try {
    const res = await uploadPoster(file);
    setFieldValue('image_url', res.url);
    previewUrl.value = resolvePosterUrl(res.url);
    toast.success('Poster uploaded successfully!');
  } catch (err: any) {
    toast.error('Failed to upload poster image', {
      description: err?.response?.data?.message || err?.message || 'Please check file format.',
    });
  }
};

const triggerFileInput = () => {
  fileInput.value?.click();
};

const onSubmit = handleSubmit(
  async (values) => {
    try {
      if (isEditing.value && props.movie) {
        await updateMovie({ id: props.movie.id, form: values });
        toast.success('Movie updated successfully!', {
          description: `"${values.name}" details have been updated.`,
        });
      } else {
        await createMovie(values);
        toast.success('Movie created successfully!', {
          description: `"${values.name}" has been added to the cinema catalog.`,
        });
        resetForm();
        previewUrl.value = null;
      }
      isOpen.value = false;
    } catch (err: any) {
      const status = err?.response?.status;
      if (status === 401) {
        toast.error('Authentication Required', {
          description: 'You must log in as an Admin first.',
        });
      } else {
        toast.error(isEditing.value ? 'Failed to update movie' : 'Failed to create movie', {
          description: err?.response?.data?.message || err?.message || 'Server error occurred.',
        });
      }
    }
  },
  (validationErrors) => {
    console.warn('Form validation errors:', validationErrors);
    toast.error('Validation Error', {
      description: 'Please make sure all form fields are filled out correctly.',
    });
  },
);
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <slot>
        <Button
          v-if="isEditing"
          variant="outline"
          size="sm"
          class="flex items-center gap-1.5 rounded-xl border-slate-700 bg-slate-900/80 text-xs font-semibold text-slate-200 hover:border-rose-500/50 hover:bg-slate-800 hover:text-white"
        >
          <Pencil class="h-3.5 w-3.5 text-rose-500" />
          <span>Edit</span>
        </Button>
        <Button
          v-else
          class="flex items-center gap-2 rounded-xl bg-rose-600 px-4 py-2 text-xs font-bold text-white shadow-lg shadow-rose-600/25 hover:bg-rose-500"
        >
          <Plus class="h-4 w-4" />
          <span>Add New Movie</span>
        </Button>
      </slot>
    </DialogTrigger>

    <DialogContent
      class="max-h-[90vh] overflow-y-auto rounded-3xl border-slate-800 bg-slate-900 p-6 text-slate-100 shadow-2xl sm:max-w-lg"
    >
      <DialogHeader class="space-y-1 border-b border-slate-800 pb-4 text-left">
        <DialogTitle class="flex items-center gap-2 text-xl font-bold text-white">
          <Pencil v-if="isEditing" class="h-5 w-5 text-rose-500" />
          <Film v-else class="h-5 w-5 text-rose-500" />
          <span>{{ dialogTitle }}</span>
        </DialogTitle>
        <DialogDescription class="text-xs text-slate-400">
          {{ dialogDescription }}
        </DialogDescription>
      </DialogHeader>

      <form id="movie-form" @submit="onSubmit" class="space-y-4 py-2">
        <!-- Poster Image Upload Box -->
        <div class="space-y-1.5">
          <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
            <ImageIcon class="h-3.5 w-3.5 text-rose-500" /> Movie Poster Image
          </label>

          <input
            ref="fileInput"
            type="file"
            accept="image/*"
            class="hidden"
            @change="handleFileSelect"
          />

          <div
            @click="triggerFileInput"
            class="group relative flex cursor-pointer flex-col items-center justify-center overflow-hidden rounded-2xl border-2 border-dashed border-slate-700 bg-slate-950/60 p-4 transition-all hover:border-rose-500/60"
          >
            <div v-if="previewUrl" class="group relative h-44 w-full overflow-hidden rounded-xl">
              <img
                :src="previewUrl"
                alt="Poster preview"
                class="h-full w-full object-cover"
                @error="(e: Event) => ((e.target as HTMLImageElement).src = defaultPoster)"
              />
              <div
                class="absolute inset-0 flex items-center justify-center gap-2 bg-slate-950/60 text-xs font-bold text-white opacity-0 transition-opacity group-hover:opacity-100"
              >
                <Upload class="h-4 w-4" /> Change Image
              </div>
            </div>

            <div
              v-else-if="isUploading"
              class="flex flex-col items-center gap-2 py-6 text-slate-400"
            >
              <Loader2 class="h-6 w-6 animate-spin text-rose-500" />
              <span class="text-xs font-medium">Uploading poster to server...</span>
            </div>

            <div
              v-else
              class="flex flex-col items-center gap-2 py-5 text-slate-400 group-hover:text-slate-200"
            >
              <div class="rounded-full border border-slate-800 bg-slate-900 p-2.5 text-rose-500">
                <Upload class="h-5 w-5" />
              </div>
              <div class="text-center">
                <span class="text-xs font-semibold text-rose-400">
                  {{ isEditing ? 'Click to upload new poster image' : 'Click to upload poster image' }}
                </span>
                <p class="text-[10px] text-slate-500">PNG, JPG, WEBP up to 10MB</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Movie Title -->
        <VeeField v-slot="{ componentField, errors }" name="name">
          <div class="space-y-1">
            <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
              <Film class="h-3.5 w-3.5 text-rose-500" /> Movie Title
            </label>
            <Input
              v-bind="componentField"
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
          <VeeField v-slot="{ componentField, errors }" name="length_minutes">
            <div class="space-y-1">
              <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
                <Clock class="h-3.5 w-3.5 text-rose-500" /> Duration (Minutes)
              </label>
              <Input
                type="number"
                v-bind="componentField"
                placeholder="120"
                class="w-full rounded-xl border border-slate-700 bg-slate-950 px-3.5 py-2 text-sm text-slate-100 focus:border-rose-500"
              />
              <p v-if="errors.length" class="text-[11px] font-medium text-rose-500">
                {{ errors[0] }}
              </p>
            </div>
          </VeeField>

          <!-- Language -->
          <VeeField v-slot="{ componentField, errors }" name="language">
            <div class="space-y-1">
              <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
                <Globe class="h-3.5 w-3.5 text-rose-500" /> Audio Language
              </label>
              <Input
                v-bind="componentField"
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
          <VeeField v-slot="{ componentField, errors }" name="published_at">
            <div class="space-y-1">
              <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
                <Calendar class="h-3.5 w-3.5 text-rose-500" /> Release Date
              </label>
              <Input
                type="date"
                v-bind="componentField"
                class="w-full rounded-xl border border-slate-700 bg-slate-950 px-3.5 py-2 text-sm text-slate-100 focus:border-rose-500"
              />
              <p v-if="errors.length" class="text-[11px] font-medium text-rose-500">
                {{ errors[0] }}
              </p>
            </div>
          </VeeField>

          <!-- Genres -->
          <VeeField v-slot="{ componentField, errors }" name="genres">
            <div class="space-y-1">
              <label class="flex items-center gap-1.5 text-xs font-semibold text-slate-300">
                <Tag class="h-3.5 w-3.5 text-rose-500" /> Genres (Comma separated)
              </label>
              <Input
                v-bind="componentField"
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
            :disabled="isPending || isUploading"
            class="rounded-xl bg-rose-600 px-5 py-2 text-xs font-bold text-white shadow-lg shadow-rose-600/20 hover:bg-rose-500"
          >
            <span>{{ submitButtonText }}</span>
          </Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
