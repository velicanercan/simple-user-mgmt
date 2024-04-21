<template>
  <layout-div>
    <h2 class="text-center mt-5 mb-3">Edit User {{ `${form.first_name}` }}</h2>
    <div class="card">
      <div class="card-header">
        <router-link class="btn btn-outline-info float-right" to="/"
          >Back</router-link
        >
      </div>
      <div class="card-body">
        <el-form
          :model="form"
          :rules="rules"
          ref="userForm"
          label-width="120px"
        >
          <el-form-item label="First Name" prop="first_name">
            <el-input v-model="form.first_name"></el-input>
          </el-form-item>
          <el-form-item label="Last Name" prop="last_name">
            <el-input v-model="form.last_name"></el-input>
          </el-form-item>
          <el-form-item prop="email" label="Email">
            <el-input v-model="form.email"></el-input>
          </el-form-item>
          <el-form-item label="Birth Date" prop="birth_date">
            <el-date-picker
              v-model="form.birth_date"
              type="date"
              value-format="YYYY-MM-DD"
              placeholder="Select date"
            ></el-date-picker>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleUpdate">Update</el-button>
            <el-button @click="resetForm">Reset</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </layout-div>
</template>

<script>
import { reactive, getCurrentInstance } from "vue";
import { useVuelidate } from "@vuelidate/core";
import { required, email } from "@vuelidate/validators";
import axios from "axios";
import LayoutDiv from "../LayoutDiv.vue";
import Swal from "sweetalert2";

export default {
  name: "UserEdit",
  props: {
    user: Object,
  },
  components: {
    LayoutDiv,
  },
  setup(props) {
    const currentInstance = getCurrentInstance();
    const router = currentInstance.appContext.config.globalProperties.$router;
    const userID = router.currentRoute.value.params.id;
    const form = reactive({
      first_name: props.user ? props.user.first_name : "",
      last_name: props.user ? props.user.last_name : "",
      email: props.user ? props.user.email : "",
      birth_date: props.user ? props.user.birth_date : "",
    });

    const rules = {
      first_name: { required },
      last_name: { required },
      email: { required, email },
      birth_date: { required },
    };

    const v$ = useVuelidate(rules, form);
    const fetchUserData = async () => {
      try {
        const response = await axios.get(`/users/${userID}`);
        const userData = response.data;
        form.first_name = userData.first_name;
        form.last_name = userData.last_name;
        form.email = userData.email;
        form.birth_date = userData.birth_date;
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };
    
    const handleUpdate = () => {
      v$.value.$validate();
      if (!v$.value.$error) {
        axios
          .put(`/users/${userID}`, form)
          .then((response) => {
            Swal.fire({
              icon: "success",
              title: "User updated successfully!",
              showConfirmButton: false,
              timer: 1500,
            });
            setTimeout(() => {
              router.push("/");
            }, 1500);
            return response;
          })
          .catch((error) => {
            let err = error.response.data.error;
            Swal.fire({
              icon: "error",
              title: err,
              showConfirmButton: false,
              timer: 1500,
            });
            return error;
          });
      } else {
          let errString = 'Please fill in all required fields.';
          for (let i in v$.value.$errors) {
            if (v$.value.$errors[i].$validator === 'email') {
              errString = 'Please enter a valid email address.';
              break;
            }
          }
          Swal.fire({
            icon: "error",
            title: errString,
            showConfirmButton: false,
            timer: 1000,
          });
        }
    };

    const resetForm = () => {
      v$.value.$reset();
      form.first_name = props.user ? props.user.first_name : "";
      form.last_name = props.user ? props.user.last_name : "";
      form.email = props.user ? props.user.email : "";
      form.birth_date = props.user ? props.user.birth_date : "";
    };

    fetchUserData();

    return {
      form,
      rules,
      handleUpdate,
      resetForm,
    };
  },
};
</script>
