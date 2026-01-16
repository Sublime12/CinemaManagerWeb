import { api } from '@/api';
import { useMutation } from '@tanstack/vue-query';
import z from 'zod';

export const LoginFormSchema = z.object({
  username: z
    .string()
    .min(5, 'username must be at least 5 characters.')
    .max(32, 'username must be at most 32 characters.'),
  password: z.string(),
  //       .min(20, 'Description must be at least 20 characters.')
  //       .max(100, 'Description must be at most 100 characters.'),
});

const MessageSchema = z.object({
  message: z.string(),
});

type Message = z.infer<typeof MessageSchema>;

export type LoginForm = z.infer<typeof LoginFormSchema>;

export function useLoginMutation() {
  return useMutation({
    mutationKey: ['login'],
    mutationFn: async (form: LoginForm) => {
      const message = await api.post<Message>(`/login`, form);
      return message.data;
    },
  });
}

export function useLogoutMutation() {
  return useMutation({
    mutationKey: ['logout'],
    mutationFn: async () => {
      await api.post(`/logout`);
    },
  });
}
