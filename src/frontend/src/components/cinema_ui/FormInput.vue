<script setup lang="ts">
import { toRefs } from 'vue';
import { Input } from '@/components/ui/input';
import { Field as VeeField } from 'vee-validate';
import { Field, FieldError, FieldLabel } from '@/components/ui/field';

const props = defineProps<{
  name: string;
  label?: string;
  placeholder?: string;
}>();

const { name, label, placeholder } = toRefs(props);
</script>
<template>
  <VeeField v-slot="{ field, errors }" :name="name">
    <Field :data-invalid="!!errors.length">
      <FieldLabel :for="name"> {{ label ?? name }} </FieldLabel>
      <Input
        :id="name"
        v-bind="field"
        :placeholder="placeholder ?? ''"
        autocomplete="off"
        :aria-invalid="!!errors.length"
      />
      <FieldError v-if="errors.length" :errors="errors" />
    </Field>
  </VeeField>
</template>
