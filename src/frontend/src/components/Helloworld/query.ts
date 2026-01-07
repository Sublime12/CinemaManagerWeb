import { api } from "@/api";
import { useQuery } from "@tanstack/vue-query";

type Message = {
  message: string
}

export function useHelloworld() {
  return useQuery({
    queryKey: ['helloworld'],
    queryFn: async () => {
      const res = await api.get<Message>('/')
      return res.data.message
    }
  })
}
