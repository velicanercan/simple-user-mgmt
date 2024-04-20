<template>
  <layout-div>
    <h2 class="text-center mt-5 mb-3">User Details</h2>
    <div class="card">
      <div class="card-header">
        <router-link class="btn btn-outline-info float-right" to="/"
          >View All Users
        </router-link>
      </div>
      <div class="card-body">
        <b className="text-muted">First Name:</b>
        <p>{{ user.first_name }}</p>
        <b className="text-muted">Last Name:</b>
        <p>{{ user.last_name }}</p>
        <b className="text-muted">Email:</b>
        <p>{{ user.email }}</p>
        <b className="text-muted">Birth Date:</b>
        <p>{{ user.birth_date }}</p>
      </div>
    </div>
  </layout-div>
</template>

<script>
import axios from "axios";
import LayoutDiv from "../LayoutDiv.vue";
import Swal from "sweetalert2";

export default {
  name: "UserShow",
  components: {
    LayoutDiv,
  },
  data() {
    return {
      user: {
        first_name: "",
        last_name: "",
        email: "",
        birth_date: "",
      },
      isSaving: false,
    };
  },
  created() {
    const id = this.$route.params.id;
    axios
      .get(`/users/${id}`)
      .then((response) => {
        let userInfo = response.data;
        this.user.first_name = userInfo.first_name;
        this.user.last_name = userInfo.last_name;
        this.user.email = userInfo.email;
        this.user.birth_date = userInfo.birth_date;
        return response;
      })
      .catch((error) => {
        Swal.fire({
          icon: "error",
          title: "An Error Occured!",
          showConfirmButton: false,
          timer: 1500,
        });
        return error;
      });
  },
  methods: {},
};
</script>
