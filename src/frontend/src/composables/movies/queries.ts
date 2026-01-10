import { api } from "@/api";
import { useQuery } from "@tanstack/vue-query";
import moment from "moment";
import type { Ref } from "vue";
import z from "zod";

const MovieSchema = z.object({
  id: z.number(),
  name: z.string(),
  description: z.string(),
  published_at: z.coerce.date<string>(),
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

export function useGetMoviesQuery() {
  return useQuery({
    queryKey: ["get-movies"],
    queryFn: async () => {
      const response = await api.get<Movie[]>(`/movies`);

      return MoviesSchema.parse(response.data);
    },
  });
}

export function useGetMovieQuery(id: Ref<string>) {
  return useQuery({
    queryKey: ["get-movie", id.value],
    queryFn: async () => {
      const response = await api.get<Movie>(`/movies/${id.value}`);

      console.log("movie data: ", response.data);
      return MovieSchema.parse(response.data);
    },
  });
}
