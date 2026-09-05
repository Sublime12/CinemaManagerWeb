import { api } from '@/api';
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query';
import z from 'zod';

export const LoginFormSchema = z.object({
  username: z
    .string()
    .min(5, 'username must be at least 5 characters.')
    .max(32, 'username must be at most 32 characters.'),
  password: z.string(),
});

const MessageSchema = z.object({
  message: z.string(),
});

const MeSchema = z.object({
  user: z.number(),
  is_admin: z.boolean(),
  username: z.string().optional(),
  name: z.string().optional(),
});

type Message = z.infer<typeof MessageSchema>;

export type MeResponse = z.infer<typeof MeSchema>;

export type LoginForm = z.infer<typeof LoginFormSchema>;

export function useGetMeQuery() {
  return useQuery({
    queryKey: ['me'],
    queryFn: async () => {
      try {
        const response = await api.get<MeResponse>(`/me`);
        return MeSchema.parse(response.data);
      } catch (err) {
        return null;
      }
    },
    retry: false,
    staleTime: 1000 * 60 * 5,
  });
}

export function useLoginMutation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationKey: ['login'],
    mutationFn: async (form: LoginForm) => {
      const message = await api.post<Message>(`/login`, form);
      return message.data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['me'] });
    },
  });
}

export function useLogoutMutation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationKey: ['logout'],
    mutationFn: async () => {
      await api.post(`/logout`);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['me'] });
    },
  });
}
