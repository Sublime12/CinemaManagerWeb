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
import router, { ROUTE_NAME as ROUTENAME } from '@/router';

const formSchema = toTypedSchema(LoginFormSchema);

const { handleSubmit, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues: {
    username: '',
    password: '',
  },
});

const { mutateAsync } = useLoginMutation();

const onSubmit = handleSubmit(async (data) => {
  const response = await mutateAsync(data);
  // toast(response.message);

  toast('', {
    description: h(
      'pre',
      { class: 'bg-code text-code-foreground mt-2 overflow-x-auto rounded-md p-4' },
      h('code', response.message),
    ),
    position: 'bottom-right',
    class: 'flex flex-col gap-2',
    style: {
      '--border-radius': 'calc(var(--radius)  + 4px)',
    },
  });

  router.push({
    name: ROUTENAME.HOME,
  });
});
</script>

<template>
  <div class="flex justify-center pt-20">
    <Card class="w-full sm:max-w-md">
      <CardHeader class="flex justify-center">
        <CardTitle class="text-2xl">Login</CardTitle>
      </CardHeader>
      <CardContent>
        <form id="login-form" @submit="onSubmit">
          <FieldGroup>
            <VeeField v-slot="{ field, errors }" name="username">
              <Field :data-invalid="!!errors.length">
                <FieldLabel for="username"> Username: </FieldLabel>
                <Input
                  id="username"
                  v-bind="field"
                  placeholder="Enter your username"
                  autocomplete="off"
                  :aria-invalid="!!errors.length"
                />
                <FieldError v-if="errors.length" :errors="errors" />
              </Field>
            </VeeField>
            <VeeField v-slot="{ field, errors }" name="password">
              <Field :data-invalid="!!errors.length">
                <FieldLabel for="password"> Password: </FieldLabel>
                <Input
                  id="password"
                  v-bind="field"
                  placeholder="Enter your password"
                  autocomplete="off"
                  :aria-invalid="!!errors.length"
                />
                <FieldError v-if="errors.length" :errors="errors" />
              </Field>
            </VeeField>
            <Item v-if="isError"> Error while login </Item>
          </FieldGroup>
        </form>
      </CardContent>
      <CardFooter>
        <Field orientation="horizontal">
          <Button type="button" variant="outline" @click="resetForm"> Reset </Button>
          <Button type="submit" form="login-form"> Submit </Button>
        </Field>
      </CardFooter>
    </Card>
  </div>
</template>
