<template>
  <div>
    <el-form :model="form" :rules="rules" ref="userForm" label-width="120px">
      <el-form-item label="Name" prop="name">
        <el-input v-model="form.name"></el-input>
      </el-form-item>
      <el-form-item label="Email" prop="email">
        <el-input v-model="form.email"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm">Submit</el-button>
        <el-button @click="resetForm">Reset</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { reactive } from 'vue';
import { useVuelidate } from '@vuelidate/core';
import { required, email } from '@vuelidate/validators';

export default {
  name: 'UserForm',
  props: {
    user: Object
  },
  setup(props) {
    const form = reactive({
      name: props.user ? props.user.name : '',
      email: props.user ? props.user.email : ''
    });

    const rules = {
      name: { required },
      email: { required, email }
    };

    const v$ = useVuelidate(rules, form);

    const submitForm = () => {
      v$.value.$validate();
      if (!v$.value.$error) {
        // Submit form data
      }
    };

    const resetForm = () => {
      v$.value.$reset();
      form.name = '';
      form.email = '';
    };

    return {
      form,
      rules,
      submitForm,
      resetForm
    };
  }
};
</script>
