<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod';
import { useForm, Field as VeeField } from 'vee-validate';
import { toast } from 'vue-sonner';

import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Field, FieldError, FieldGroup, FieldLabel } from '@/components/ui/field';
import { Input } from '@/components/ui/input';

import { h } from 'vue';
import { LoginFormSchema, useLoginMutation } from '@/composables/auth/queries';
import router, { ROUTE_NAME } from '@/router';
import { Clapperboard, Lock, User, LogIn, KeyRound } from 'lucide-vue-next';

const formSchema = toTypedSchema(LoginFormSchema);

const { handleSubmit, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues: {
    username: '',
    password: '',
  },
});

const { mutateAsync, isError, isPending } = useLoginMutation();

const onSubmit = handleSubmit(async (data) => {
  try {
    const response = await mutateAsync(data);
    toast.success('Successfully logged in!', {
      description: response.message || 'Welcome back to SubCine Cinema Manager.',
    });

    router.push({
      name: ROUTE_NAME.ADMIN,
    });
  } catch (err: any) {
    toast.error('Authentication Failed', {
      description: err?.response?.data?.message || 'Invalid username or password.',
    });
  }
});
</script>

<template>
  <div class="relative flex min-h-[75vh] items-center justify-center px-4 py-12">
    <!-- Ambient Backdrop Light Spotlights -->
    <div
      class="bg-primary/15 pointer-events-none absolute -top-10 left-1/2 h-96 w-96 -translate-x-1/2 rounded-full blur-3xl"
    ></div>
    <div
      class="bg-accent/10 pointer-events-none absolute bottom-10 left-1/3 h-80 w-80 rounded-full blur-3xl"
    ></div>

    <div class="relative z-10 w-full max-w-md">
      <div class="bg-card border-border/80 space-y-6 rounded-3xl border p-6 shadow-2xl sm:p-8">
        <!-- Header Cinema Logo -->
        <div class="space-y-2 text-center">
          <div
            class="from-primary shadow-primary/20 mb-1 inline-flex rounded-2xl bg-gradient-to-tr via-rose-600 to-amber-500 p-3 text-white shadow-xl"
          >
            <Clapperboard class="h-8 w-8" />
          </div>
          <h2 class="text-2xl font-black tracking-tight text-white">Cinema Portal Login</h2>
          <p class="text-muted-foreground text-xs">
            Enter administrator credentials to manage movie schedules and theatres
          </p>
        </div>

        <form id="login-form" @submit="onSubmit" class="space-y-4">
          <VeeField v-slot="{ field, errors }" name="username">
            <div class="space-y-1.5">
              <label
                for="username"
                class="flex items-center gap-1.5 text-xs font-semibold text-slate-300"
              >
                <User class="text-primary h-3.5 w-3.5" /> Username
              </label>
              <div class="relative">
                <Input
                  id="username"
                  v-bind="field"
                  placeholder="admin"
                  autocomplete="off"
                  class="border-border text-foreground placeholder:text-muted-foreground/60 focus:border-primary focus:ring-primary w-full rounded-xl border bg-slate-900/80 px-4 py-2.5 text-sm transition-all focus:ring-1"
                  :aria-invalid="!!errors.length"
                />
              </div>
              <p v-if="errors.length" class="text-destructive text-[11px] font-medium">
                {{ errors[0] }}
              </p>
            </div>
          </VeeField>

          <VeeField v-slot="{ field, errors }" name="password">
            <div class="space-y-1.5">
              <label
                for="password"
                class="flex items-center gap-1.5 text-xs font-semibold text-slate-300"
              >
                <Lock class="text-primary h-3.5 w-3.5" /> Password
              </label>
              <div class="relative">
                <Input
                  id="password"
                  type="password"
                  v-bind="field"
                  placeholder="••••••••"
                  autocomplete="off"
                  class="border-border text-foreground placeholder:text-muted-foreground/60 focus:border-primary focus:ring-primary w-full rounded-xl border bg-slate-900/80 px-4 py-2.5 text-sm transition-all focus:ring-1"
                  :aria-invalid="!!errors.length"
                />
              </div>
              <p v-if="errors.length" class="text-destructive text-[11px] font-medium">
                {{ errors[0] }}
              </p>
            </div>
          </VeeField>

          <div
            v-if="isError"
            class="bg-destructive/10 border-destructive/20 text-destructive rounded-xl border p-3 text-center text-xs font-medium"
          >
            Invalid credentials. Please verify your username and password.
          </div>

          <div class="flex items-center gap-3 pt-2">
            <button
              type="button"
              @click="resetForm"
              class="bg-secondary/60 hover:bg-secondary border-border text-muted-foreground hover:text-foreground w-1/3 rounded-xl border px-4 py-2.5 text-xs font-semibold transition-all"
            >
              Reset
            </button>

            <button
              type="submit"
              :disabled="isPending"
              class="bg-primary shadow-primary/30 flex w-2/3 items-center justify-center gap-2 rounded-xl px-4 py-2.5 text-xs font-bold text-white shadow-lg transition-all hover:bg-rose-600 disabled:opacity-50"
            >
              <LogIn v-if="!isPending" class="h-4 w-4" />
              <span v-if="isPending">Authenticating...</span>
              <span v-else>Sign In to Dashboard</span>
            </button>
          </div>
        </form>

        <div class="border-border/30 border-t pt-2 text-center">
          <p class="text-muted-foreground text-[11px]">
            Demo Credentials:
            <code class="text-primary bg-card rounded px-1.5 py-0.5 font-mono">admin / admin</code>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
