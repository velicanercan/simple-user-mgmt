<template>
  <div>
    <div class="table-header">
      <input type="text" v-model="search" placeholder="Search users">
      <el-button type="primary" @click="openCreateForm">Create User</el-button>
    </div>
    <el-table :data="filteredUsers" style="width: 100%">
      <el-table-column prop="id" label="ID" width="80"></el-table-column>
      <el-table-column prop="first_name" label="First Name"></el-table-column>
      <el-table-column prop="last_name" label="Last Name"></el-table-column>
      <el-table-column prop="email" label="Email"></el-table-column>
      <el-table-column prop="birth_date" label="Birth Date"></el-table-column>
      <el-table-column label="Actions" width="200">
        <template #default="{ row }">
          <el-button @click="editUser(row)">Edit</el-button>
          <el-button type="danger" @click="deleteUser(row)">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 20, 30, 40]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="totalUsers"
    ></el-pagination>
    <el-dialog v-model:visible="createFormVisible" title="Create User" :before-close="handleCreateFormClose">
      <UserForm v-model:user="selectedUser" />
    </el-dialog>
  </div>
</template>

<script>
import axios from 'axios';
import { ref, computed, onMounted } from 'vue';

export default {
  name: 'UserList',
  setup() {
    const currentPage = ref(1);
    const pageSize = ref(10);
    const search = ref('');
    const users = ref([]);
    const createFormVisible = ref(false);
    const selectedUser = ref(null);
    
    const filteredUsers = computed(() => {
      return users.value.filter(user => {
        return user.first_name.toLowerCase().includes(search.value.toLowerCase()) ||
               user.last_name.toLowerCase().includes(search.value.toLowerCase()) ||
               user.email.toLowerCase().includes(search.value.toLowerCase());
      });
    });

    const getUsers = async () => {
      try {
        const apiUrl = process.env.VUE_APP_API_URL;
        const response = await axios.get(apiUrl + '/users');
        users.value = response.data;
      } catch (error) {
        console.error('Error fetching users:', error);
      }
    };

    onMounted(() => {
      getUsers();
    });

    const handleSizeChange = (val) => {
      pageSize.value = val;
    };

    const handleCurrentChange = (val) => {
      currentPage.value = val;
    };

    const editUser = () => {
      // Implement edit user functionality
    };

    const deleteUser = () => {
      // Implement delete user functionality
    };

    const openCreateForm = () => {
      createFormVisible.value = true;
    };

    const handleCreateFormClose = () => {
      createFormVisible.value = false;
    };

    return {
      search,
      filteredUsers,
      currentPage,
      pageSize,
      totalUsers: computed(() => users.value.length),
      handleSizeChange,
      handleCurrentChange,
      editUser,
      deleteUser,
      createFormVisible,
      selectedUser,
      openCreateForm,
      handleCreateFormClose
    };
  }
};
</script>

<style scoped>
.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}
</style>
