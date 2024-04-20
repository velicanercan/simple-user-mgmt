<template>
  <div>
    <el-form :model="form" :rules="rules" ref="userForm" label-width="120px">
      <el-form-item label="First Name" prop="first_name">
        <el-input v-model="form.first_name"></el-input>
      </el-form-item>
      <el-form-item label="Last Name" prop="last_name">
        <el-input v-model="form.last_name"></el-input>
      </el-form-item>
      <el-form-item prop="email" label="Email" ><el-input v-model="form.email"></el-input>
      </el-form-item>
      <el-form-item label="Birth Date" prop="birth_date">
        <el-date-picker v-model="form.birth_date" type="date" placeholder="Select date"></el-date-picker>
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
      first_name: props.user ? props.user.first_name : '',
      last_name: props.user ? props.user.last_name : '',
      email: props.user ? props.user.email : '',
      birth_date: props.user ? props.user.birth_date : ''
    });

    const rules = {
      first_name: { required},
      last_name: { required},
      email: { required, email },
      birth_date: { required }

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
      form.first_name = '';
      form.last_name = '';
      form.email = '';
      form.birth_date = '';
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
