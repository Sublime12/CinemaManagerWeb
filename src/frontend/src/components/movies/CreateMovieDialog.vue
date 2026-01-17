<script setup lang="ts">
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
import { Label } from '@/components/ui/label';
import FormInput from '../cinema_ui/FormInput.vue';
import { toTypedSchema } from '@vee-validate/zod';
import { useForm, Field as VeeField } from 'vee-validate';
import { toast } from 'vue-sonner';
import { z } from 'zod';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Field, FieldDescription, FieldError, FieldGroup, FieldLabel } from '@/components/ui/field';
import {
  InputGroup,
  InputGroupAddon,
  InputGroupText,
  InputGroupTextarea,
} from '@/components/ui/input-group';

const formSchema = toTypedSchema(
  z.object({
    title: z
      .string()
      .min(5, 'Bug title must be at least 5 characters.')
      .max(32, 'Bug title must be at most 32 characters.'),
    description: z
      .string()
      .min(20, 'Description must be at least 20 characters.')
      .max(100, 'Description must be at most 100 characters.'),
    duration: z
      .string()
      .min(20, 'Description must be at least 20 characters.')
      .max(100, 'Description must be at most 100 characters.'),
  }),
);
const { handleSubmit, resetForm } = useForm({
  validationSchema: formSchema,
  initialValues: {
    title: '',
    description: '',
  },
});
const onSubmit = handleSubmit((data) => {
  toast('You submitted the following values:', data);
});
</script>

<template>
  <Dialog>
    <form>
      <DialogTrigger as-child>
        <Button variant="outline"> New Movie </Button>
      </DialogTrigger>
      <DialogContent class="sm:max-w-106.25">
        <DialogHeader>
          <DialogTitle>Edit profile</DialogTitle>
          <DialogDescription>
            Make changes to your profile here. Click save when you're done.
          </DialogDescription>
        </DialogHeader>
        <Card class="w-full sm:max-w-md">
          <CardHeader>
            <CardTitle>Bug Report</CardTitle>
            <CardDescription> Help us improve by reporting bugs you encounter. </CardDescription>
          </CardHeader>
          <CardContent>
            <form id="form-vee-demo" @submit="onSubmit">
              <FieldGroup>
                <FormInput
                  name="title"
                  label="Movie Title"
                  placeholder="Enter the name of the movie"
                />
                <VeeField v-slot="{ field, errors }" name="description">
                  <Field :data-invalid="!!errors.length">
                    <FieldLabel for="form-vee-demo-description"> Description </FieldLabel>
                    <InputGroup>
                      <InputGroupTextarea
                        id="form-vee-demo-description"
                        v-bind="field"
                        placeholder="I'm having an issue with the login button on mobile."
                        :rows="6"
                        class="min-h-24 resize-none"
                        :aria-invalid="!!errors.length"
                      />
                      <InputGroupAddon align="block-end">
                        <InputGroupText class="tabular-nums">
                          {{ field.value?.length || 0 }}/100 characters
                        </InputGroupText>
                      </InputGroupAddon>
                    </InputGroup>
                    <FieldDescription>
                      Include steps to reproduce, expected behavior, and what actually happened.
                    </FieldDescription>
                    <FieldError v-if="errors.length" :errors="errors" />
                  </Field>
                </VeeField>
              </FieldGroup>
            </form>
          </CardContent>
          <CardFooter>
            <Field orientation="horizontal"> </Field>
          </CardFooter>
        </Card>
        <DialogFooter>
          <DialogClose as-child>
            <Button variant="outline"> Cancel </Button>

            <Button type="submit" form="form-vee-demo"> Submit </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </form>
  </Dialog>
</template>
