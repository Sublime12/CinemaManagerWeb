import { api } from '@/api';
import { useQuery } from '@tanstack/vue-query';
import z from 'zod';

const MovieSchema = z.object({
  id: z.number(),
  name: z.string(),
  description: z.string(),
});

const MoviesSchema = z.array(MovieSchema)

export type Movie = z.infer<typeof MovieSchema>;

export function useGetMoviesQuery() {
  return useQuery({
    queryKey: ['get-movies'],
    queryFn: async () => {
      const movies = await api.get<Movie[]>(`/movies`);

      return MoviesSchema.parse(movies.data);
    },
  });
}
