import { api } from '@/api';
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query';
import moment from 'moment';
import type { Ref } from 'vue';
import z from 'zod';

const MovieSchema = z.object({
  id: z.number(),
  name: z.string(),
  description: z.string(),
  published_at: z.coerce.date(),
  length: z
    .number()
    .int()
    .positive()
    .transform((ms) => moment.duration(ms / 1e6)),
  language: z.string(),
  genres: z.array(z.string()),
});

const MoviesSchema = z.array(MovieSchema);

export type Movie = z.infer<typeof MovieSchema>;

export const CreateMovieFormSchema = z.object({
  name: z.string().min(2, 'Name must be at least 2 characters.'),
  description: z.string().min(10, 'Description must be at least 10 characters.'),
  published_at: z.string().min(1, 'Release date is required.'),
  length_minutes: z.number().min(1, 'Duration must be greater than 0.'),
  language: z.string().min(2, 'Language is required.'),
  genres: z.string().min(2, 'Genres are required (comma separated).'),
});

export type CreateMovieForm = z.infer<typeof CreateMovieFormSchema>;

export function useGetMoviesQuery() {
  return useQuery({
    queryKey: ['get-movies'],
    queryFn: async () => {
      const response = await api.get<Movie[]>(`/movies`);
      return MoviesSchema.parse(response.data);
    },
  });
}

export function useGetMovieQuery(id: Ref<string>) {
  return useQuery({
    queryKey: ['get-movie', id.value],
    queryFn: async () => {
      const response = await api.get<Movie>(`/movies/${id.value}`);
      return MovieSchema.parse(response.data);
    },
  });
}

export function useCreateMovieMutation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationKey: ['create-movie'],
    mutationFn: async (form: CreateMovieForm) => {
      const payload = {
        name: form.name,
        description: form.description,
        published_at: new Date(form.published_at).toISOString(),
        // Convert minutes to nanoseconds for Go time.Duration (minutes * 60 * 1e9)
        length: form.length_minutes * 60 * 1e9,
        language: form.language,
        genres: form.genres
          .split(',')
          .map((g) => g.trim())
          .filter(Boolean),
      };
      const response = await api.post(`/movies`, payload);
      return response.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['get-movies'] });
    },
  });
}
