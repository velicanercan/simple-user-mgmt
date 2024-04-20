<template>
  <layout-div>
    <div class="container">
      <h2 class="text-center mt-5 mb-3">User Management</h2>
      <div class="search-bar"></div>
      <div class="card">
        <div class="card-header">
          <router-link to="/create" class="btn btn-outline-primary"
            >Create New User
          </router-link>
          <div style="float: right">
            <input
              v-model="search"
              type="text"
              class="form-control"
              placeholder="Search..."
            />
          </div>
        </div>
        <template v-if="isLoading">
          <div class="text-center">Loading users...</div>
        </template>
        <template v-else>
          <div class="card-body">
            <el-table :data="filteredUsers" :row-key="users.id" stripe>
              <el-table-column
                prop="first_name"
                label="First Name"
                width="150"
              ></el-table-column>
              <el-table-column
                prop="last_name"
                label="Last Name"
                width="150"
              ></el-table-column>
              <el-table-column prop="email" label="Email"></el-table-column>
              <el-table-column
                prop="birth_date"
                label="Birth Date"
                width="180"
              ></el-table-column>
              <el-table-column label="Action" width="240">
                <template #default="scope">
                  <router-link
                    :to="{ path: `/show/${scope.row.id}` }"
                    class="el-button el-button--primary el-button--small"
                  >
                    Show
                  </router-link>
                  <router-link
                    :to="{ path: `/edit/${scope.row.id}` }"
                    class="el-button el-button--success el-button--small mx-2"
                  >
                    Edit
                  </router-link>
                  <el-button
                    type="danger"
                    size="small"
                    @click="handleDelete(scope.row.id)"
                  >
                    Delete
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="currentPage"
              :page-sizes="[1, 10, 20, 30, 40]"
              :page-size="pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="totalUsers"
            ></el-pagination>
          </div>
        </template>
      </div>
    </div>
  </layout-div>
</template>

<script>
import axios from "axios";
import LayoutDiv from "../LayoutDiv.vue";
import Swal from "sweetalert2";
import { ref, watchEffect } from "vue";
import { ElTable, ElTableColumn } from "element-plus";

export default {
  name: "UserList",
  components: {
    LayoutDiv,
    ElTable,
    ElTableColumn,
  },
  setup() {
    const currentPage = ref(1);
    const users = ref([]);
    const pageSize = ref(10);
    const search = ref("");
    const isLoading = ref(true);
    const filteredUsers = ref([]);
    const totalUsers = ref(0);

    watchEffect(() => {
      const startIndex = (currentPage.value - 1) * pageSize.value;
      const endIndex = startIndex + pageSize.value;

      const filtered = users.value.filter((user) => {
        return (
          user.first_name.toLowerCase().includes(search.value.toLowerCase()) ||
          user.last_name.toLowerCase().includes(search.value.toLowerCase()) ||
          user.email.toLowerCase().includes(search.value.toLowerCase())
        );
      });

      filteredUsers.value = filtered.slice(startIndex, endIndex);
      totalUsers.value = filtered.length;
    });

    const handleSizeChange = (val) => {
      pageSize.value = val;
    };
    const handleCurrentChange = (val) => {
      currentPage.value = val;
    };

    return {
      users,
      isLoading,
      search,
      pageSize,
      filteredUsers,
      currentPage,
      totalUsers,
      handleCurrentChange,
      handleSizeChange,
    };
  },
  created() {
    this.fetchUserList();
  },

  methods: {
    fetchUserList() {
      this.isLoading = true;
      axios
        .get("/users")
        .then((response) => {
          this.users = response.data;
          this.isLoading = false;
          this.filteredUsers = this.users;
          return response;
        })
        .catch((error) => {
          console.error("Error fetching users:", error);
          this.isLoading = false;
          return error;
        });
    },
    handleDelete(id) {
      Swal.fire({
        title: "Are you sure?",
        text: "You won't be able to revert this!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#3085d6",
        cancelButtonColor: "#d33",
        confirmButtonText: "Yes, delete it!",
      }).then((result) => {
        if (result.isConfirmed) {
          axios
            .delete(`/users/${id}`)
            .then((response) => {
              Swal.fire({
                icon: "success",
                title: "User deleted successfully!",
                showConfirmButton: false,
                timer: 1500,
              });
              this.fetchUserList();
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
        }
      });
    },
  },
};
</script>
