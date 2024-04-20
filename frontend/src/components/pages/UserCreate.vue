<template>
    <layout-div>
      <h2 class="text-center mt-5 mb-3">Create New User</h2>
      <div class="card">
        <div class="card-header">
          <router-link class="btn btn-outline-info float-right" to="/"
            >Back
          </router-link>
        </div>
        <div class="card-body">
        <el-form :model="form" :rules="rules" ref="userForm" label-width="120px">
        <el-form-item label="First Name" prop="first_name">
            <el-input v-model="form.first_name"></el-input>
        </el-form-item>
        <el-form-item label="Last Name" prop="last_name">
            <el-input v-model="form.last_name"></el-input>
        </el-form-item>
        <el-form-item prop="email" label="Email" ><el-input v-model="form.email"></el-input>
            <!-- email validation -->
            <el-form-item v-if="v$?.value?.email?.email" class="invalid-feedback">Please enter a valid email address.</el-form-item>
        </el-form-item>
        <el-form-item label="Birth Date" prop="birth_date">
            <el-date-picker v-model="form.birth_date" type="date" value-format="YYYY-MM-DD" placeholder="Select date"></el-date-picker>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="handleSave">Submit</el-button>
            <el-button @click="resetForm">Reset</el-button>
        </el-form-item>
        </el-form>
        </div>
      </div>
    </layout-div>
  </template>
  
  <script>
  import { reactive, getCurrentInstance } from 'vue';
  import { useVuelidate } from '@vuelidate/core';
  import { required, email } from '@vuelidate/validators';
  import axios from "axios";
  import LayoutDiv from "../LayoutDiv.vue";
  import Swal from "sweetalert2";


  export default {
    name: "UserCreate",
    props: {
      user: Object
    },
    components: {
      LayoutDiv,
    },
    setup(props) {
      const currentInstance = getCurrentInstance();
      const router = currentInstance.appContext.config.globalProperties.$router;
      const form = reactive({
        first_name: props.user ? props.user.first_name : '',
        last_name: props.user ? props.user.last_name : '',
        email: props.user ? props.user.email : '',
        birth_date: props.user ? props.user.birth_date : ''
      });
      
      const rules = {
        first_name: { required },
        last_name: { required },
        email: { required, email },
        birth_date: { required }
      };

      const v$ = useVuelidate(rules, form);

      const handleSave = () => {
        v$.value.$validate();
        if (!v$.value.$error) {
          // Submit form data
          axios
            .post("/users", form)
            .then((response) => {
              Swal.fire({
                icon: "success",
                title: "User saved successfully!",
                showConfirmButton: false,
                timer: 1500,
              });
              setTimeout(() => {
                router.push("/");
              }, 1500);
              return response;
            })
            .catch((error) => {
              console.log(form);
              let err = error.response.data.error;
              Swal.fire({
                icon: "error",
                title: err,
                showConfirmButton: false,
                timer: 1500,
              });
              return error;
            });
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
        handleSave,
        resetForm,
        v$ // return v$ for validation
      };
    },
  };
  </script>
